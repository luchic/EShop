import { useNavigate } from "react-router-dom";
import shopBg from "@/assets/shop_bg.jpg";

const Account = () => {
  const navigate = useNavigate();
  const username = localStorage.getItem("shopUser") || "Mysterious Stranger";

  const handleLogout = () => {
    localStorage.removeItem("shopUser");
    navigate("/login");
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

      <div className="pixel-corner top-0 left-0" />
      <div className="pixel-corner top-0 right-0" style={{ transform: "scaleX(-1)" }} />
      <div className="pixel-corner bottom-0 left-0" style={{ transform: "scaleY(-1)" }} />
      <div className="pixel-corner bottom-0 right-0" style={{ transform: "scale(-1,-1)" }} />

      {/* Account Card */}
      <div
        className="relative z-10 border-4 border-gold p-8 flex flex-col gap-6 w-[420px]"
        style={{
          background: "rgba(15,8,2,0.85)",
          boxShadow: "6px 6px 0 hsl(var(--brown-dark)), inset 0 0 0 2px hsl(var(--gold))",
        }}
      >
        <h1 className="text-gold text-base text-center tracking-wider"
          style={{ textShadow: "2px 2px 0 hsl(var(--brown-dark))" }}
        >
          ✦ Adventurer's Guild Card ✦
        </h1>

        <div className="flex flex-col gap-4 border-2 border-brown-mid p-4" style={{ background: "rgba(10,5,0,0.4)" }}>
          <div className="flex justify-between items-center">
            <span className="text-[9px] text-muted-foreground">NAME</span>
            <span className="text-[11px] text-parchment">{username}</span>
          </div>
          <div className="flex justify-between items-center">
            <span className="text-[9px] text-muted-foreground">RANK</span>
            <span className="text-[11px] text-gold">Novice</span>
          </div>
          <div className="flex justify-between items-center">
            <span className="text-[9px] text-muted-foreground">GOLD</span>
            <span className="text-[11px] text-parchment">500 🪙</span>
          </div>
          <div className="flex justify-between items-center">
            <span className="text-[9px] text-muted-foreground">ITEMS</span>
            <span className="text-[11px] text-parchment">0</span>
          </div>
        </div>

        <div className="flex gap-3">
          <button
            onClick={() => navigate("/")}
            className="flex-1 text-[10px] font-pixel text-gold border-2 border-gold px-3 py-2.5 hover:bg-gold hover:text-brown-dark transition-colors"
            style={{ background: "rgba(15,8,2,0.7)" }}
          >
            ← Shop
          </button>
          <button
            onClick={handleLogout}
            className="flex-1 text-[10px] font-pixel text-destructive border-2 border-destructive px-3 py-2.5 hover:bg-destructive hover:text-foreground transition-colors"
            style={{ background: "rgba(15,8,2,0.7)" }}
          >
            Logout ✕
          </button>
        </div>
      </div>
    </div>
  );
};

export default Account;
