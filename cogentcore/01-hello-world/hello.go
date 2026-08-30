package main

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
)

func main() {
	b := core.NewBody()
	tf := core.NewTextField(b).SetPlaceholder("Name")
	core.NewButton(b).SetText("Greet").OnClick(func(e events.Event) {
		core.MessageSnackbar(b, "Hello, "+tf.Text()+"!")
	})
	b.RunMainWindow()
}
