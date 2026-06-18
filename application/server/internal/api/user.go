package api

import "time"

type RegisterUser struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	SessionId string `json:"session_id"`
	UserId    int64  `json:"user_id"`
	Email     string `json:"email"`
	ExpiresAt int64  `json:"expires_at"`
}

type SessionData struct {
	SessionId string    `json:"session_id"`
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type GetUserByIdRequest struct {
	Email string `json:"email"`
}

type User struct {
	Id         int64     `db:"id"`
	Role       string    `db:"role"`
	FirstName  string    `db:"first_name"`
	SecondName string    `db:"second_name"`
	Email      string    `db:"email"`
	Password   []byte    `db:"password"`
	CreatedAt  time.Time `db:"created_at"`
}
