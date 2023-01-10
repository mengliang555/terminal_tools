package basicwidget

import (
	ui "github.com/gizak/termui/v3"
)

var _defaultGrid *ui.Grid

func init() {
	grid := ui.NewGrid()
	// termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, 500, 500)
	_defaultGrid = grid
}

func GetDefaultGrid() *ui.Grid {
	return _defaultGrid
}

// 2 column or 3 column

func InitGrid(grid *ui.Grid, val ...interface{}) []ui.GridItem {
	totalLength := len(val)
	var Col []ui.GridItem
	if totalLength == 1 {
		Col = []ui.GridItem{
			ui.NewCol(1, val),
		}
	} else if totalLength%3 == 0 {
		f := getDD(3, val...)
		Col = []ui.GridItem{
			ui.NewCol(1.0/3, f[0]),
			ui.NewCol(1.0/3, f[1]),
			ui.NewCol(1.0/3, f[2]),
		}
	} else if totalLength%2 == 0 {
		f := getDD(2, val...)
		Col = []ui.GridItem{
			ui.NewCol(1.0/2, f[0]),
			ui.NewCol(1.0/2, f[1]),
		}
	}
	return Col
}

func getDD(size int, val ...interface{}) [][]ui.GridItem {
	ans := make([][]ui.GridItem, 0, size)
	totalRow := len(val) / 3
	for i := range ans {
		ans[i] = make([]ui.GridItem, 0, totalRow)
	}
	for i, v := range val {
		ans[i%3] = append(ans[i%3], ui.NewRow(float64(totalRow)/float64(1), v))
	}
	return ans
}

func InitServer() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	grid := GetDefaultGrid()
	ui.Render(grid)
	EventDeal()
}
