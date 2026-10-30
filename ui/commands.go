package ui

import (
	tea "charm.land/bubbletea/v2"
	model "github.com/DarylvdBerg/text-based-dungeon-crawler/system"
	"github.com/DarylvdBerg/text-based-dungeon-crawler/system/core"
)

type slotsLoadedMsg struct {
	slots []*model.Slot
}

func loadSlots() tea.Cmd {
	return func() tea.Msg {
		system := model.System{}
		slots := make([]*model.Slot, system.AvailableSlots())
		for i := range slots {
			slots[i] = core.ReadSave(i)
		}
		return slotsLoadedMsg{slots: slots}
	}
}
