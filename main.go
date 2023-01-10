package main

import (
	"log"
	"terminal_tools/data_info"
	eventdeal "terminal_tools/event_deal"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	sl := widgets.NewSparkline()
	sl.Data = data_info.SinFloat64[:100]
	sl.LineColor = ui.ColorCyan
	sl.TitleStyle.Fg = ui.ColorWhite

	slg := widgets.NewSparklineGroup(sl)
	slg.Title = "Sparkline"

	lc := widgets.NewPlot()
	lc.Title = "braille-mode Line Chart"
	lc.Data = append(lc.Data, data_info.SinFloat64)
	lc.AxesColor = ui.ColorWhite
	lc.LineColors[0] = ui.ColorYellow

	ls := widgets.NewList()
	ls.Rows = data_info.MacUsage
	ls.Border = false

	eventdeal.RegisterKeyboardEvent(func() {
		ls.ScrollHalfPageUp()
		ui.Clear()
		ui.Render(grid)
	}, "<Up>", "<Left>")

	eventdeal.RegisterKeyboardEvent(func() {
		ls.ScrollHalfPageDown()
		ui.Clear()
		ui.Render(grid)
	}, "<Down>", "<Right>")

	p := widgets.NewParagraph()
	p.Text = "<> This row has 3 columns\n<- Widgets can be stacked up like left side\n<- Stacked widgets are treated as a single widget"
	p.Title = "Demonstration"

	grid.Set(
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, slg),
			ui.NewCol(1.0/2, lc),
		),
		ui.NewRow(1.0/2,
			ui.NewCol(1.0/2, ls),
			ui.NewCol(1.0/2, p),
		),
	)

	ui.Render(grid)
	eventdeal.DealEvent(grid, slg, lc, ls)
}
