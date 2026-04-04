import { ShopItem } from "@/data/shopItems";
import theodorImg from "@/assets/agent_theodor.png";
import Typewriter from "./Typewriter";
import { useState } from "react";

interface ShopkeeperPanelProps {
  selectedItem: ShopItem | null;
}

const ShopkeeperPanel = ({ selectedItem }: ShopkeeperPanelProps) => {
  const [showPrice, setShowPrice] = useState(false);

  const name = selectedItem ? `✦ ${selectedItem.name} ✦` : "Theodor says:";
  const text = selectedItem
    ? selectedItem.desc
    : "Welcome, traveller! Browse my wares and click an item to learn more.";

  return (
    <div className="w-[320px] min-w-[220px] flex flex-col justify-end px-2.5 pt-3 shrink-0">
      {/* Speech Bubble */}
      <div className="relative bg-bubble-bg border-4 border-brown-dark p-4 mb-3.5 min-h-[260px] flex flex-col gap-3.5"
        style={{
          boxShadow: "4px 4px 0 hsl(var(--brown-dark)), inset 0 0 0 2px hsl(var(--gold))",
          imageRendering: "pixelated",
        }}
      >
        {/* Bubble tail */}
        <div className="absolute -bottom-5 left-10 w-0 h-0"
          style={{
            borderLeft: "8px solid transparent",
            borderRight: "8px solid transparent",
            borderTop: "16px solid hsl(var(--brown-dark))",
          }}
        />
        <div className="absolute -bottom-3.5 left-[43px] w-0 h-0 z-10"
          style={{
            borderLeft: "5px solid transparent",
            borderRight: "5px solid transparent",
            borderTop: "12px solid hsl(var(--bubble-bg))",
          }}
        />

        <span className="text-sm text-dark-gold uppercase tracking-wider"
          style={{ textShadow: "1px 1px 0 hsl(var(--brown-dark))" }}
        >
          {name}
        </span>
        <span className="text-sm text-brown-dark leading-[1.9]">
          <Typewriter
            key={text}
            text={text}
            onComplete={() => setShowPrice(!!selectedItem)}
          />
        </span>
        {showPrice && selectedItem && (
          <span className="text-sm text-destructive mt-2 flex items-center gap-2">
            {selectedItem.price} 🪙
          </span>
        )}
      </div>

      {/* Shopkeeper */}
      <div className="flex justify-center items-end">
        <img
          src={theodorImg}
          alt="Shopkeeper Theodor"
          className="w-[220px] h-auto"
          style={{
            imageRendering: "pixelated",
            filter: "drop-shadow(4px 4px 0 rgba(0,0,0,0.7))",
            animation: "keeper-idle 2.4s steps(1) infinite",
          }}
        />
      </div>
    </div>
  );
};

export default ShopkeeperPanel;
