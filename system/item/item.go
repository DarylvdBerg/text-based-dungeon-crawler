package item

import (
	"github.com/DarylvdBerg/text-based-dungeon-crawler/system/player"
	"github.com/DarylvdBerg/text-based-dungeon-crawler/system/stats"
)

type Item interface {
	Name() string
	Description() string
}

type Consumable interface {
	Item
	Use(p *player.Player)
}

type Equipment struct {
	Item
	stats stats.Stats
}

