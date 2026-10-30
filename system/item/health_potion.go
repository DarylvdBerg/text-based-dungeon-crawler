package item

import "github.com/DarylvdBerg/text-based-dungeon-crawler/system/player"

var _ Consumable = HealthPotion{}

type HealthPotion struct {
	Amount int
}

// Name implements [Consumable].
func (h HealthPotion) Name() string {
	return "Health Potion"
}

// Description implements [Consumable].
func (h HealthPotion) Description() string {
	return "A potion that restores health when consumed."
}

// Use implements [Consumable].
func (h HealthPotion) Use(p *player.Player) {
	p.IncreaseHealth(h.Amount)
}
