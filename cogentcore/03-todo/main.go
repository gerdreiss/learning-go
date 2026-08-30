package main

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
)

type item struct {
	Done bool `display:"checkbox"`
	Task string
}

func main() {
	items := []item{{Task: "Code"}, {Task: "Eat"}}

	b := core.NewBody()

	table := core.NewTable(b).SetSlice(&items)

	addLine := func(e events.Event) {
		table.NewAt(0)
	}
	core.NewButton(b).SetText("Add").SetIcon(icons.Add).OnClick(addLine)

	b.RunMainWindow()
}
