package auth

import (
	authapi "backend/shop/internal/api/auth"
	"backend/shop/internal/config"
	"backend/shop/internal/repository"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	providerName      = "42-intra"
	authorizeEndpoint = "https://api.intra.42.fr/oauth/authorize"
	tokenEndpoint     = "https://api.intra.42.fr/oauth/token"
	meEndpoint        = "https://api.intra.42.fr/v2/me"
	defaultScope      = "public"
)

type Handlers struct {
	repo       repository.Repository
	cfg        config.Config
	httpClient *http.Client
}

func NewHandlers(repo repository.Repository, cfg config.Config) *Handlers {
	return &Handlers{
		repo: repo,
		cfg:  cfg,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func Routers(mux *http.ServeMux, repo repository.Repository, cfg config.Config) *http.ServeMux {
	if mux == nil {
		return mux
	}

	handlers := NewHandlers(repo, cfg)
	mux.HandleFunc("/auth/42/login", handlers.handleLogin)
	mux.HandleFunc("/auth/42/callback", handlers.handleCallback)
	mux.HandleFunc("/user/login", handlers.handleLogin)
	return mux
}

func (h *Handlers) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	state, err := generateState()
	if err != nil {
		slog.Error("generate oauth state failed", slog.Any("err", err))
		http.Error(w, "Couldn't create login state", http.StatusInternalServerError)
		return
	}

	if err := h.repo.SaveOAuthState(state); err != nil {
		http.Error(w, "Couldn't save login state", http.StatusInternalServerError)
		return
	}

	response := authapi.LoginResponse{
		AuthorizeURL: buildAuthorizeURL(h.cfg, state),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("encode login response failed", slog.Any("err", err))
		http.Error(w, "Couldn't encode response", http.StatusInternalServerError)
	}
}

func (h *Handlers) handleCallback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	if code == "" || state == "" {
		http.Error(w, "Missing code or state", http.StatusBadRequest)
		return
	}

	if !h.repo.HasOAuthState(state) {
		http.Error(w, "Invalid oauth state", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteOAuthState(state); err != nil {
		http.Error(w, "Couldn't finish login", http.StatusInternalServerError)
		return
	}

	tokenResponse, err := h.exchangeCode(code)
	if err != nil {
		slog.Error("exchange oauth code failed", slog.Any("err", err))
		http.Error(w, "Couldn't exchange oauth code", http.StatusBadGateway)
		return
	}

	fortyTwoUser, err := h.fetchFortyTwoUser(tokenResponse.AccessToken)
	if err != nil {
		slog.Error("fetch 42 user failed", slog.Any("err", err))
		http.Error(w, "Couldn't fetch 42 user", http.StatusBadGateway)
		return
	}

	appUser, err := h.repo.UpsertOAuthUser(toOAuthUser(fortyTwoUser))
	if err != nil {
		http.Error(w, "Couldn't save user", http.StatusInternalServerError)
		return
	}

	sessionToken, err := createSessionToken(appUser, h.cfg.TokenSecret)
	if err != nil {
		slog.Error("create session token failed", slog.Any("err", err))
		http.Error(w, "Couldn't create session", http.StatusInternalServerError)
		return
	}

	response := authapi.CallbackResponse{
		SessionToken: sessionToken,
		User:         appUser,
		Provider:     providerName,
		Scopes:       scopesFromString(tokenResponse.Scope),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("encode callback response failed", slog.Any("err", err))
		http.Error(w, "Couldn't encode response", http.StatusInternalServerError)
	}
}

func generateState() (string, error) {
	buffer := make([]byte, 16)
	if _, err := rand.Read(buffer); err != nil {
		return "", err
	}

	return hex.EncodeToString(buffer), nil
}

func buildAuthorizeURL(cfg config.Config, state string) string {
	values := url.Values{}
	values.Set("client_id", cfg.FortyTwoClientID)
	values.Set("redirect_uri", cfg.FortyTwoRedirect)
	values.Set("response_type", "code")
	values.Set("scope", defaultScope)
	values.Set("state", state)

	return authorizeEndpoint + "?" + values.Encode()
}

func (h *Handlers) exchangeCode(code string) (authapi.FortyTwoTokenResponse, error) {
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("client_id", h.cfg.FortyTwoClientID)
	form.Set("client_secret", h.cfg.FortyTwoSecret)
	form.Set("code", code)
	form.Set("redirect_uri", h.cfg.FortyTwoRedirect)

	request, err := http.NewRequest(http.MethodPost, tokenEndpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return authapi.FortyTwoTokenResponse{}, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := h.httpClient.Do(request)
	if err != nil {
		return authapi.FortyTwoTokenResponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(response.Body, 2048))
		return authapi.FortyTwoTokenResponse{}, fmt.Errorf("token endpoint returned %d: %s", response.StatusCode, strings.TrimSpace(string(body)))
	}

	var tokenResponse authapi.FortyTwoTokenResponse
	if err := json.NewDecoder(response.Body).Decode(&tokenResponse); err != nil {
		return authapi.FortyTwoTokenResponse{}, err
	}

	return tokenResponse, nil
}

func (h *Handlers) fetchFortyTwoUser(accessToken string) (authapi.FortyTwoUserResponse, error) {
	request, err := http.NewRequest(http.MethodGet, meEndpoint, nil)
	if err != nil {
		return authapi.FortyTwoUserResponse{}, err
	}
	request.Header.Set("Authorization", "Bearer "+accessToken)

	response, err := h.httpClient.Do(request)
	if err != nil {
		return authapi.FortyTwoUserResponse{}, err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(response.Body, 2048))
		return authapi.FortyTwoUserResponse{}, fmt.Errorf("me endpoint returned %d: %s", response.StatusCode, strings.TrimSpace(string(body)))
	}

	var fortyTwoUser authapi.FortyTwoUserResponse
	if err := json.NewDecoder(response.Body).Decode(&fortyTwoUser); err != nil {
		return authapi.FortyTwoUserResponse{}, err
	}

	return fortyTwoUser, nil
}

func toOAuthUser(user authapi.FortyTwoUserResponse) authapi.OAuthUser {
	return authapi.OAuthUser{
		ProviderID:  user.ID,
		Login:       user.Login,
		DisplayName: user.DisplayName,
		Email:       user.Email,
	}
}

func createSessionToken(user authapi.AppUser, secret string) (string, error) {
	payload := fmt.Sprintf("%d:%s:%d", user.ID, user.Login, time.Now().UTC().Unix())
	mac := hmac.New(sha256.New, []byte(secret))
	if _, err := mac.Write([]byte(payload)); err != nil {
		return "", err
	}

	signature := hex.EncodeToString(mac.Sum(nil))
	rawToken := payload + "." + signature
	return base64.RawURLEncoding.EncodeToString([]byte(rawToken)), nil
}

func scopesFromString(scope string) []string {
	if strings.TrimSpace(scope) == "" {
		return []string{}
	}

	return strings.Fields(scope)
}
