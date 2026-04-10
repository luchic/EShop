import { useState } from "react";
import { useNavigate } from "react-router-dom";
import shopBg from "@/assets/shop_bg.jpg";
import { SHOP_ITEMS, ShopItem } from "@/data/shopItems";
import ShelfPanel from "@/components/ShelfPanel";
import ShopkeeperPanel from "@/components/ShopkeeperPanel";

const Index = () => {
  const [selectedItem, setSelectedItem] = useState<ShopItem | null>(null);
  const navigate = useNavigate();

  return (
    <div className="w-full h-screen overflow-hidden relative font-pixel">
      {/* Background */}
      <div
        className="fixed inset-0 z-0"
        style={{
          backgroundImage: `url(${shopBg})`,
          backgroundPosition: "center center",
          backgroundSize: "cover",
          backgroundRepeat: "no-repeat",
          imageRendering: "pixelated",
        }}
      >
        <div className="absolute inset-0" style={{ background: "rgba(15,8,2,0.38)" }} />
      </div>

      {/* CRT Scanlines */}
      <div className="crt-scanlines" />

      {/* Pixel corners */}
      <div className="pixel-corner top-0 left-0" />
      <div className="pixel-corner top-0 right-0" style={{ transform: "scaleX(-1)" }} />
      <div className="pixel-corner bottom-0 left-0" style={{ transform: "scaleY(-1)" }} />
      <div className="pixel-corner bottom-0 right-0" style={{ transform: "scale(-1,-1)" }} />

      {/* Nav bar */}
      <div className="fixed top-0 right-0 z-50 flex gap-2 p-3">
        <button
          onClick={() => navigate("/login")}
          className="text-[10px] font-pixel text-gold border-2 border-gold px-3 py-2 hover:bg-gold hover:text-brown-dark transition-colors"
          style={{ background: "rgba(15,8,2,0.7)" }}
        >
          Login
        </button>
        <button
          onClick={() => navigate("/account")}
          className="text-[10px] font-pixel text-gold border-2 border-gold px-3 py-2 hover:bg-gold hover:text-brown-dark transition-colors"
          style={{ background: "rgba(15,8,2,0.7)" }}
        >
          Account
        </button>
      </div>

      {/* Layout */}
      <div className="relative z-10 flex w-full h-screen">
        <ShelfPanel
          items={SHOP_ITEMS}
          selectedItem={selectedItem}
          onSelectItem={setSelectedItem}
        />
        <ShopkeeperPanel selectedItem={selectedItem} />
      </div>
    </div>
  );
};

export default Index;
