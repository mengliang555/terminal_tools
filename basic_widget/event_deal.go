package basicwidget

import (
	ui "github.com/gizak/termui/v3"
)

func EventDeal() {
	for ev := range ui.PollEvents() {
		if ev.Type == ui.KeyboardEvent {
			switch ev.ID {
			case "<C-c>", "q":
				return
			}
		}
	}
}
