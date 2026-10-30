package player

import "github.com/DarylvdBerg/text-based-dungeon-crawler/system/stats"

type Player struct {
	Name  string
	health int
	Stats stats.Stats
}

func (p *Player) IsAlive() bool {
	return p.health > 0
}

func (p *Player) IncreaseHealth(amount int) {
	p.health += amount
	if p.health > p.Stats.MaxHealth {
		p.health = p.Stats.MaxHealth
	}
}