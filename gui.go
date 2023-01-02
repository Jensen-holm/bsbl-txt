package main

import "github.com/marcusolsson/tui-go"

type Tui struct {
	widgets []*tui.Widget
	boxes   []*tui.Box
	ui      tui.UI
}

func NewTui() *Tui {
	return &Tui{
		widgets: nil,
		boxes:   nil,
		ui:      nil,
	}
}

func (t *Tui) AddWidgets(widgets ...*tui.Widget) {
	for _, s := range widgets {
		t.widgets = append(t.widgets, s)
	}
}

func (t *Tui) Init() {
	for _, s := range t.widgets {
		t.ui, _ = tui.New(*s)
	}
	for _, b := range t.boxes {
		t.ui, _ = tui.New(b)
	}
}

func (t *Tui) Run() {
	if err := t.ui.Run(); err != nil {
		panic(err)
	}
}
