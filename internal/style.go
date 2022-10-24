package internal

import (
	"github.com/charmbracelet/lipgloss"
)

type Style struct{}

func (s *Style) activeTabBorder() lipgloss.Border {
	return lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}
}

func (s *Style) activeTab() lipgloss.Style {
	return s.tab().Copy().Border(s.activeTabBorder(), true)
}

func (s *Style) checkMark() string {
	return lipgloss.NewStyle().SetString("✓").
		Foreground(s.special()).
		PaddingRight(1).
		String()
}

func (s *Style) descStyle() lipgloss.Style {
	return lipgloss.NewStyle().MarginTop(2)
}

func (s *Style) infoStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).
		BorderForeground(s.subtle())
}

func (s *Style) divider() string {
	return lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(s.subtle()).
		String()
}

func (s *Style) list() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(s.subtle()).
		MarginRight(2).
		Height(8).
		Width(columnWidth + 1)
}

func (s *Style) listEntry(item string) string {
	return s.divider() + lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).
		Render(item)
}

func (s *Style) listChecked(item string) string {
	return s.checkMark() + lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).
		Render(item)
}

func (s *Style) listHeader(header string) string {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(s.subtle()).
		MarginRight(2).
		Render(header)
}

func (s *Style) listItem(item string) string {
	return lipgloss.NewStyle().PaddingLeft(3).Render(item)
}

func (s *Style) statusBarStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})
}

func (s *Style) statusStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Inherit(s.statusBarStyle()).
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#FF5F87")).
		Padding(0, 1).
		MarginRight(1)
}

func (s *Style) statusText() lipgloss.Style {
	return lipgloss.NewStyle().Inherit(s.statusBarStyle())
}

func (s *Style) subtle() lipgloss.TerminalColor {
	return lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
}

func (s *Style) highlight() lipgloss.TerminalColor {
	return lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
}

func (s *Style) special() lipgloss.TerminalColor {
	return lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
}

func (s *Style) tab() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(s.tabBorder(), true).
		BorderForeground(s.highlight()).
		Padding(0, 1)
}

func (s *Style) tabBorder() lipgloss.Border {
	return lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}
}

func (s *Style) tabGap() lipgloss.Style {
	return s.tab().Copy().
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)
}

func (s *Style) url(address string) string {
	return lipgloss.NewStyle().Foreground(s.special()).Render(address)
}

func (s *Style) width() int {
	return 96
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
