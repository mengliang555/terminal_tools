package eventdeal

import (
	"log"
	"terminal_tools/data_info"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var _KeyboardEventMap map[string]func()

func init() {
	_KeyboardEventMap = make(map[string]func(), 0)
}

// sigle key
func RegisterKeyboardEvent(behave func(), IdF string, Other ...string) {
	if _, ok := _KeyboardEventMap[IdF]; ok {
		panic("dublicate key for event")
	} else {
		_KeyboardEventMap[IdF] = behave
	}

	for _, v := range Other {
		if _, ok := _KeyboardEventMap[IdF]; ok {
			log.Printf("[ERROR] dublicate key %s", v)
		} else {
			_KeyboardEventMap[v] = behave
		}
	}
}

func DealWithAllReigsterEvent() {
	uiEvents := ui.PollEvents()

	ui.Clear()
	for {
		select {
		case <-uiEvents:
		}
	}
	for v := range uiEvents {
		if v.Type == ui.KeyboardEvent {
			if val, ok := _KeyboardEventMap[v.ID]; ok {
				val()
			} else {
				log.Printf("[WARN]No such key behave %s", v.ID)
			}
		}
	}
}

func DealEvent(grid *ui.Grid, slg *widgets.SparklineGroup, lc *widgets.Plot, ls *widgets.List) {
	uiEvents := ui.PollEvents()
	tickerCount := 0
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Up>", "<Left>":
				ls.ScrollHalfPageUp()
				ui.Clear()
				ui.Render(grid)
			case "<Down>", "<Right>":
				ls.ScrollHalfPageDown()
				ui.Clear()
				ui.Render(grid)

			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}
		case <-ticker:
			start, end := tickerCount%data_info.MaxCount, (tickerCount+100)%data_info.MaxCount
			if start > end {
				start, end = end, start
			}
			slg.Sparklines[0].Data = data_info.SinFloat64[start:end]
			lc.Data[0] = data_info.SinFloat64[2*tickerCount:]
			ui.Render(grid)
			tickerCount = (tickerCount + 1) % data_info.MaxCount
		}
	}
}
