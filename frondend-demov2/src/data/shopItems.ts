export interface ShopItem {
  id: number;
  name: string;
  desc: string;
  price: string;
  icon: string;
}

export const SHOP_ITEMS: ShopItem[] = [
  { id: 1, name: "Healing Potion", desc: "A shimmering red liquid that restores vitality. Brewed from rare mountain herbs.", price: "25 Gold", icon: "⚗" },
  { id: 2, name: "Enchanted Key", desc: "Opens any lock forged by mortal hands. Handle with care — it bites.", price: "80 Gold", icon: "🗝" },
  { id: 3, name: "Soul Gem", desc: "A gem that pulses with captured starlight. Useful for enchanting weapons.", price: "150 Gold", icon: "💎" },
  { id: 4, name: "Crystal Ball", desc: "Peer into the mists of fate. Results may vary. No refunds on prophecies.", price: "200 Gold", icon: "🔮" },
  { id: 5, name: "Iron Sword", desc: "A sturdy blade forged in the fires of Mount Ember. Good for beginners.", price: "45 Gold", icon: "⚔" },
  { id: 6, name: "Ancient Scroll", desc: "Contains a forgotten spell. The ink still glows faintly in moonlight.", price: "120 Gold", icon: "📜" },
  { id: 7, name: "Tower Shield", desc: "A massive shield that can block dragon fire. Very heavy, very reliable.", price: "95 Gold", icon: "🛡" },
  { id: 8, name: "Lucky Coin", desc: "Flip it before any battle for a chance of fortune. Heads you win, tails... well.", price: "10 Gold", icon: "🪙" },
  { id: 9, name: "Herb Bundle", desc: "A fragrant collection of medicinal herbs. Essential for any adventurer's kit.", price: "15 Gold", icon: "🌿" },
  { id: 10, name: "Ancient Vase", desc: "A beautifully preserved vase from a lost civilization. Collectors pay handsomely.", price: "300 Gold", icon: "🏺" },
  { id: 11, name: "Warding Eye", desc: "Protects the wearer from curses and dark magic. Stares back at evil.", price: "175 Gold", icon: "🧿" },
  { id: 12, name: "Lightning Rod", desc: "Channels storm energy into devastating attacks. Smells like ozone.", price: "250 Gold", icon: "⚡" },
  { id: 13, name: "Volatile Elixir", desc: "An unstable mixture that could heal or harm. For the brave or foolish.", price: "60 Gold", icon: "🧪" },
  { id: 14, name: "Wizard's Wand", desc: "A finely crafted wand of elder oak. Resonates with magical potential.", price: "400 Gold", icon: "🪄" },
  { id: 15, name: "Elven Bow", desc: "Light as a feather, deadly as a viper. Carved from a single branch.", price: "350 Gold", icon: "🏹" },
];
