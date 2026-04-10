import { useState } from "react";
import { useNavigate } from "react-router-dom";
import shopBg from "@/assets/shop_bg.jpg";

const Login = () => {
  const navigate = useNavigate();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = (e: React.FormEvent) => {
    e.preventDefault();
    // Placeholder — no real auth
    if (username && password) {
      localStorage.setItem("shopUser", username);
      navigate("/account");
    }
  };

  return (
    <div className="w-full h-screen overflow-hidden relative font-pixel flex items-center justify-center">
      {/* BG */}
      <div className="fixed inset-0 z-0" style={{
        backgroundImage: `url(${shopBg})`,
        backgroundSize: "cover",
        backgroundPosition: "center",
        imageRendering: "pixelated",
      }}>
        <div className="absolute inset-0" style={{ background: "rgba(15,8,2,0.65)" }} />
      </div>
      <div className="crt-scanlines" />

      {/* Pixel corners */}
      <div className="pixel-corner top-0 left-0" />
      <div className="pixel-corner top-0 right-0" style={{ transform: "scaleX(-1)" }} />
      <div className="pixel-corner bottom-0 left-0" style={{ transform: "scaleY(-1)" }} />
      <div className="pixel-corner bottom-0 right-0" style={{ transform: "scale(-1,-1)" }} />

      {/* Form */}
      <form
        onSubmit={handleLogin}
        className="relative z-10 border-4 border-gold p-8 flex flex-col gap-6 w-[380px]"
        style={{
          background: "rgba(15,8,2,0.85)",
          boxShadow: "6px 6px 0 hsl(var(--brown-dark)), inset 0 0 0 2px hsl(var(--gold))",
        }}
      >
        <h1 className="text-gold text-base text-center tracking-wider"
          style={{ textShadow: "2px 2px 0 hsl(var(--brown-dark))" }}
        >
          ✦ Enter the Shop ✦
        </h1>

        <div className="flex flex-col gap-2">
          <label className="text-[10px] text-parchment">Adventurer Name</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="bg-input border-2 border-brown-mid text-parchment text-[10px] font-pixel px-3 py-2 focus:border-gold focus:outline-none"
          />
        </div>

        <div className="flex flex-col gap-2">
          <label className="text-[10px] text-parchment">Secret Word</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="bg-input border-2 border-brown-mid text-parchment text-[10px] font-pixel px-3 py-2 focus:border-gold focus:outline-none"
          />
        </div>

        <button
          type="submit"
          className="text-[11px] font-pixel text-brown-dark bg-gold border-2 border-dark-gold px-4 py-3 hover:brightness-110 transition-all"
          style={{ boxShadow: "3px 3px 0 hsl(var(--brown-dark))" }}
        >
          Enter ⚔
        </button>

        <button
          type="button"
          onClick={() => navigate("/")}
          className="text-[9px] text-muted-foreground hover:text-gold transition-colors text-center"
        >
          ← Back to shop
        </button>
      </form>
    </div>
  );
};

export default Login;
