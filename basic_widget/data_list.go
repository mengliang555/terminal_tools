package basicwidget

import (
	datalist "terminal_tools/data_list"

	"github.com/gizak/termui/v3/widgets"
)

func GetMacTermianlUsage() *widgets.List {
	list := widgets.NewList()
	list.Title = "mac usage"
	list.Rows = datalist.MacUsage
	return list
}

func init() {
	InitGrid(GetDefaultGrid(), GetMacTermianlUsage())
}
