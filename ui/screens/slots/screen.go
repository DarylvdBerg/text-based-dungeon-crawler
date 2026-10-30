package slots

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/DarylvdBerg/text-based-dungeon-crawler/system/model"
)

type ChosenMsg struct {
	Index int
	Slot  *model.Slot
}

type Model struct {
	slots   []*model.Slot
	cursor  int
	loading bool
	width   int
	height  int
}

func NewLoading() Model {
	return Model{loading: true}
}

func New(slots []*model.Slot) Model {
	return Model{slots: slots}
}

func (m Model) Models() []tea.Model {
	return nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		if m.loading {
			return m, nil
		}
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.slots)-1 {
				m.cursor++
			}
		case "enter":
			i := m.cursor
			return m, func() tea.Msg { return ChosenMsg{Index: i, Slot: m.slots[i]} }
		}
	}

	return m, nil
}
func (m Model) View() tea.View {
	if m.loading {
		return tea.NewView("Loading saves..")
	}

	var b strings.Builder
	b.WriteString("Choose a save slot\n\n")
	for i, s := range m.slots {
		cursor := " "
		if i == m.cursor {
			cursor = "> "
		}
		name := "New Game"
		if s != nil {
			name = s.Name
		}
		_, err := fmt.Fprintf(&b, "%s%d. %s\n", cursor, i+1, name)
		if err != nil {
			return tea.NewView("")
		}
	}
	b.WriteString("\n up/down move * enter select")
	return tea.NewView(b.String())
}
