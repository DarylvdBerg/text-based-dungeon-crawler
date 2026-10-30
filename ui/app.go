package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/DarylvdBerg/text-based-dungeon-crawler/ui/screens/slots"

	"charm.land/lipgloss/v2"
)

type App struct {
	screens []tea.Model
	focus   int
	width   int
	height  int
}

func New() App {
	return App{screens: []tea.Model{slots.NewLoading()}}
}

func (a App) Init() tea.Cmd {
	return loadSlots()
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case slotsLoadedMsg:
		return a.show(slots.New(msg.slots))
	}

	var cmds []tea.Cmd
	for i := range a.screens {
		var cmd tea.Cmd
		a.screens[i], cmd = a.screens[i].Update(msg)
		cmds = append(cmds, cmd)
	}
	return a, tea.Batch(cmds...)
}

func (a App) View() tea.View {
	parts := make([]string, len(a.screens))
	for index, screen := range a.screens {
		parts[index] = screen.View().Content
	}
	view := tea.NewView(lipgloss.JoinVertical(lipgloss.Left, parts...))
	view.AltScreen = true
	return view
}

func (a App) show(screens ...tea.Model) (tea.Model, tea.Cmd) {
	a.screens = screens
	a.focus = 0
	cmds := []tea.Cmd{a.resize()}
	for _, screen := range a.screens {
		cmds = append(cmds, screen.Init())
	}

	return a, tea.Batch(cmds...)
}

func (a *App) resize() tea.Cmd {
	part := tea.WindowSizeMsg{Width: a.width / max(len(a.screens), 1), Height: a.height}
	var cmds []tea.Cmd
	for i := range a.screens {
		var cmd tea.Cmd
		a.screens[i], cmd = a.screens[i].Update(part)
		cmds = append(cmds, cmd)
	}
	return tea.Batch(cmds...)
}
