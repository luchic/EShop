import { ShopItem } from "@/data/shopItems";
import ItemCard from "./ItemCard";

interface ShelfPanelProps {
  items: ShopItem[];
  selectedItem: ShopItem | null;
  onSelectItem: (item: ShopItem) => void;
}

const ITEMS_PER_SHELF = 5;

const ShelfPanel = ({ items, selectedItem, onSelectItem }: ShelfPanelProps) => {
  const shelves: ShopItem[][] = [];
  for (let i = 0; i < items.length; i += ITEMS_PER_SHELF) {
    shelves.push(items.slice(i, i + ITEMS_PER_SHELF));
  }

  return (
    <div className="flex-1 flex flex-col py-3.5 pl-3.5 pr-1.5 overflow-hidden">
      {/* Title */}
      <div
        className="text-base text-gold text-center tracking-widest mb-2.5 py-3 border-y-4 border-gold"
        style={{
          textShadow: "2px 2px 0 hsl(var(--brown-dark)), -1px -1px 0 hsl(var(--brown-mid))",
          background: "rgba(15,8,2,0.5)",
        }}
      >
        ✦ Theodor's Emporium ✦
      </div>

      {/* Scrollable shelves */}
      <div className="flex-1 overflow-y-auto overflow-x-hidden pr-1">
        {shelves.map((shelf, i) => (
          <div key={i} className="mb-5">
            {/* Items */}
            <div className="flex flex-nowrap gap-4 min-h-[170px] items-end px-3.5 pt-4 pb-2.5"
              style={{ background: "rgba(10,5,0,0.35)" }}
            >
              {shelf.map((item) => (
                <ItemCard
                  key={item.id}
                  item={item}
                  isSelected={selectedItem?.id === item.id}
                  onClick={() => onSelectItem(item)}
                />
              ))}
            </div>

            {/* Wooden plank */}
            <div
              className="h-5"
              style={{
                background: "repeating-linear-gradient(90deg, #7a4a1a 0px, #5c3a1e 6px, #6b4020 6px, #7a4a1a 12px)",
                borderTop: "5px solid #a0622a",
                borderBottom: "5px solid #3a1a00",
                boxShadow: "0 5px 0 #1a0800",
                imageRendering: "pixelated",
              }}
            />
          </div>
        ))}
      </div>
    </div>
  );
};

export default ShelfPanel;
