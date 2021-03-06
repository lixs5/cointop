package cointop

import (
	"github.com/jroimartin/gocui"
)

func (ct *Cointop) getCurrentPage() int {
	return ct.page + 1
}

func (ct *Cointop) getTotalPages() int {
	return (ct.getListCount() / ct.perpage) + 1
}

func (ct *Cointop) getListCount() int {
	return len(ct.allcoins)
}

func (ct *Cointop) cursorDown(g *gocui.Gui, v *gocui.View) error {
	if ct.tableview == nil {
		return nil
	}
	_, y := ct.tableview.Origin()
	cx, cy := ct.tableview.Cursor()
	numRows := len(ct.coins) - 1
	if (cy + y + 1) > numRows {
		return nil
	}
	if err := ct.tableview.SetCursor(cx, cy+1); err != nil {
		ox, oy := ct.tableview.Origin()
		// set origin scrolls
		if err := ct.tableview.SetOrigin(ox, oy+1); err != nil {
			return err
		}
	}
	ct.rowChanged()
	return nil
}

func (ct *Cointop) cursorUp(g *gocui.Gui, v *gocui.View) error {
	if ct.tableview == nil {
		return nil
	}
	ox, oy := ct.tableview.Origin()
	cx, cy := ct.tableview.Cursor()
	if err := ct.tableview.SetCursor(cx, cy-1); err != nil && oy > 0 {
		// set origin scrolls
		if err := ct.tableview.SetOrigin(ox, oy-1); err != nil {
			return err
		}
	}
	ct.rowChanged()
	return nil
}

func (ct *Cointop) pageDown(g *gocui.Gui, v *gocui.View) error {
	if ct.tableview == nil {
		return nil
	}
	ox, oy := ct.tableview.Origin() // this is prev origin position
	cx, _ := ct.tableview.Cursor()  // relative cursor position
	_, sy := ct.tableview.Size()    // rows in visible view
	k := oy + sy
	l := len(ct.coins)
	// end of table
	if (oy + sy + sy) > l {
		k = l - sy
	}
	if err := ct.tableview.SetOrigin(ox, k); err != nil {
		return err
	}
	// move cursor to last line if can't scroll further
	if k == oy {
		if err := ct.tableview.SetCursor(cx, sy-1); err != nil {
			return err
		}
	}
	ct.rowChanged()
	return nil
}

func (ct *Cointop) pageUp(g *gocui.Gui, v *gocui.View) error {
	if ct.tableview == nil {
		return nil
	}
	ox, oy := ct.tableview.Origin()
	cx, _ := ct.tableview.Cursor() // relative cursor position
	_, sy := ct.tableview.Size()   // rows in visible view
	k := oy - sy
	if k < 0 {
		k = 0
	}
	if err := ct.tableview.SetOrigin(ox, k); err != nil {
		return err
	}
	// move cursor to first line if can't scroll further
	if k == oy {
		if err := ct.tableview.SetCursor(cx, 0); err != nil {
			return err
		}
	}
	ct.rowChanged()
	return nil
}

func (ct *Cointop) navigateFirstLine(g *gocui.Gui, v *gocui.View) error {
	if ct.tableview == nil {
		return nil
	}
	ox, _ := ct.tableview.Origin()
	cx, _ := ct.tableview.Cursor()
	if err := ct.tableview.SetOrigin(ox, 0); err != nil {
		return err
	}
	if err := ct.tableview.SetCursor(cx, 0); err != nil {
		return err
	}
	ct.rowChanged()
	return nil
}

func (ct *Cointop) navigateLastLine(g *gocui.Gui, v *gocui.View) error {
	if ct.tableview == nil {
		return nil
	}
	ox, _ := ct.tableview.Origin()
	cx, _ := ct.tableview.Cursor()
	_, sy := ct.tableview.Size()
	l := len(ct.coins)
	k := l - sy
	if err := ct.tableview.SetOrigin(ox, k); err != nil {
		return err
	}
	if err := ct.tableview.SetCursor(cx, sy-1); err != nil {
		return err
	}
	ct.rowChanged()
	return nil
}

func (ct *Cointop) navigatePageFirstLine(g *gocui.Gui, v *gocui.View) error {
	if ct.tableview == nil {
		return nil
	}
	cx, _ := ct.tableview.Cursor()
	if err := ct.tableview.SetCursor(cx, 0); err != nil {
		return err
	}
	ct.rowChanged()
	return nil
}

func (ct *Cointop) navigatePageMiddleLine(g *gocui.Gui, v *gocui.View) error {
	if ct.tableview == nil {
		return nil
	}
	cx, _ := ct.tableview.Cursor()
	_, sy := ct.tableview.Size()
	if err := ct.tableview.SetCursor(cx, (sy/2)-1); err != nil {
		return err
	}
	ct.rowChanged()
	return nil
}

func (ct *Cointop) navigatePageLastLine(g *gocui.Gui, v *gocui.View) error {
	if ct.tableview == nil {
		return nil
	}
	cx, _ := ct.tableview.Cursor()
	_, sy := ct.tableview.Size()
	if err := ct.tableview.SetCursor(cx, sy-1); err != nil {
		return err
	}
	ct.rowChanged()
	return nil
}

func (ct *Cointop) prevPage(g *gocui.Gui, v *gocui.View) error {
	if (ct.page - 1) >= 0 {
		ct.page = ct.page - 1
	}
	ct.updateTable()
	ct.rowChanged()
	return nil
}

func (ct *Cointop) nextPage(g *gocui.Gui, v *gocui.View) error {
	if ((ct.page + 1) * ct.perpage) <= ct.getListCount() {
		ct.page = ct.page + 1
	}
	ct.updateTable()
	ct.rowChanged()
	return nil
}

func (ct *Cointop) firstPage(g *gocui.Gui, v *gocui.View) error {
	ct.page = 0
	ct.updateTable()
	ct.rowChanged()
	return nil
}

func (ct *Cointop) lastPage(g *gocui.Gui, v *gocui.View) error {
	ct.page = ct.getListCount() / ct.perpage
	ct.updateTable()
	ct.rowChanged()
	return nil
}
