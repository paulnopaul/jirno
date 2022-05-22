package gocui_cli

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func RunApp() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	_, err = fmt.Fprintln(v, "Hello world!")
	if err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
