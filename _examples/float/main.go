package main

import (
	"fmt"
	"log"
	"os"

	"github.com/brotholo/gotk4-layer-shell/gtklayershell"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Keyboard struct {
	gridh []*gtk.Grid
	//  gridw  *gtk.Grid
	app    *gtk.Application
	appwin *gtk.ApplicationWindow
	x      int
	y      int
}

func main() {
	NewKeyboard := Keyboard{}
	NewKeyboard.x = 1920
	NewKeyboard.y = 400
	NewKeyboard.gridh = []*gtk.Grid{}
	app := gtk.NewApplication("com.github.diamondburned.gotk4-layer-shell.examples.float", 0)
	app.Connect("activate", NewKeyboard.activate)
	NewKeyboard.app = app

	if code := NewKeyboard.app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

//  var attrs = []*pango.Attribute{
//  pango.NewAttrScale(10),
//  pango.NewAttrWeight(pango.WeightBold),
//  }

func (K *Keyboard) GenKeyboard() {
	row0 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	row1 := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
	row2 := []string{"a", "s", "d", "f", "g", "h", "j", "k", "l"}
	row3 := []string{"z", "x", "c", "v", "b", "n", "m"}
	kdata := []interface{}{row0, row1, row2, row3}

	for _, r := range kdata {
		col := K.x / len(r.([]string))
		grid := gtk.NewGrid()
		grid.SetBaselineRow(len(r.([]string)))
		//  y := K.y / len(r.([]string))
		for nl, l := range r.([]string) {
			k := gtk.NewButton()
			k.SetSizeRequest(col, 100)
			k.SetLabel(l)
			k.SetName(l)
			fmt.Println(0, nl)
			fmt.Println(col, 100)
			//  fmt.Println(int(nr), int(nl), x, y)
			//  grid.Attach(k, nl, 1, x, y)
			grid.Attach(k, nl, 0, 1, 1)
			//  grid.Attach
		}
		//  fmt.Println("GRID LEN", grid.SetBaselineRow()
		K.gridh = append(K.gridh, grid)
	}
	//  return grid
}
func (K *Keyboard) GenRowKeyboard(h int) {
}
func (K *Keyboard) activate(app *gtk.Application) {
	if !gtklayershell.IsSupported() {
		log.Fatalln("gtk-layer-shell not supported")
	}

	//  labelAttrs := pango.NewAttrList()
	//  for _, attr := range attrs {
	//  labelAttrs.Insert(attr)
	//  }

	//  label := gtk.NewLabel("This window floats in the middle.")
	//  label.SetAttributes(labelAttrs)
	//  label.Show()

	K.appwin = gtk.NewApplicationWindow(app)
	K.appwin.SetDefaultSize(1920, 400)
	K.GenKeyboard()

	main_grid := gtk.NewGrid()
	fmt.Println("LEN GRID", len(K.gridh))
	for i := range K.gridh {
		grid := K.gridh[i]
		grid.Show()
		fmt.Println("LEN", grid.BaselineRow())
		main_grid.Attach(grid, 0, i, 1, 1)
	}

	K.appwin.SetChild(main_grid)
	main_grid.Show()
	//  K.appwin.SetChild(K.gridh[0])
	//  K.gridh[0].Show()
	window := &K.appwin.Window

	gtklayershell.InitForWindow(window)
	gtklayershell.SetLayer(window, gtklayershell.LayerShellLayerTop)

	gtklayershell.AutoExclusiveZoneEnable(window)
	edge := gtklayershell.Edge(3)
	gtklayershell.SetAnchor(window, edge, true)
	K.appwin.Show()
}
