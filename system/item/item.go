package item

import "github.com/DarylvdBerg/text-based-dungeon-crawler/system/stats"

type Item interface {
	Name() string
	Description() string
}

type Consumable interface {
	Item
	Use() string
}

type Equipment struct {
	Item
	stats stats.Stats
}

