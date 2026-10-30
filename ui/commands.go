package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/DarylvdBerg/text-based-dungeon-crawler/system/core"
	"github.com/DarylvdBerg/text-based-dungeon-crawler/system/model"
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
