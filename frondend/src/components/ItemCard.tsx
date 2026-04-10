import { ShopItem } from "@/data/shopItems";

interface ItemCardProps {
  item: ShopItem;
  isSelected: boolean;
  onClick: () => void;
}

const ItemCard = ({ item, isSelected, onClick }: ItemCardProps) => {
  return (
    <div
      className="flex-none w-[136px] flex flex-col items-center gap-2.5 cursor-pointer relative p-2 transition-transform duration-[50ms] hover:-translate-y-2 group"
      onClick={onClick}
    >
      {/* Glow */}
      <div className="absolute inset-0 opacity-0 group-hover:opacity-100 pointer-events-none transition-opacity duration-200"
        style={{
          background: "radial-gradient(ellipse at center, rgba(245,200,66,0.25) 0%, transparent 70%)",
        }}
      />

      {/* Frame */}
      <div
        className={`w-28 h-28 flex items-center justify-center text-[28px] leading-none transition-all duration-100 ${
          isSelected
            ? "border-4 border-gold"
            : "border-4 border-brown-mid"
        }`}
        style={{
          background: "rgba(20,10,2,0.7)",
          boxShadow: isSelected
            ? "0 0 0 4px hsl(var(--gold)), 6px 6px 0 hsl(var(--brown-dark))"
            : "5px 5px 0 hsl(var(--brown-dark))",
          imageRendering: "pixelated",
        }}
      >
        <span className={isSelected ? "animate-[sparkle_1s_ease-in-out_infinite]" : ""}>
          {item.icon}
        </span>
      </div>

      {/* Label */}
      <div className="text-[10px] text-parchment text-center leading-relaxed max-w-[128px] break-words"
        style={{ textShadow: "1px 1px 0 #000" }}
      >
        {item.name}
      </div>
    </div>
  );
};

export default ItemCard;
