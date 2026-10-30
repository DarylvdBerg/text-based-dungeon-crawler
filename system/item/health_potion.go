package item

var _ Consumable = HealthPotion{}

type HealthPotion struct {
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
func (h HealthPotion) Use() string {
	panic("unimplemented")
}
