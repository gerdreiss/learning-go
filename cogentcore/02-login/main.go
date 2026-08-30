package main

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
)

type user struct {
	Username string
	Password string
}

func main() {
	u := &user{}

	b := core.NewBody("Login")
	pg := core.NewPages(b)

	pg.AddPage("sign-in", func(pg *core.Pages) {
		core.NewForm(pg).SetStruct(u)
		core.NewButton(pg).SetText("Sign in").OnClick(func(e events.Event) {
			pg.Open("home")
		})
	})

	pg.AddPage("home", func(pg *core.Pages) {
		core.NewText(pg).SetText("Welcome, " + u.Username + "!").SetType(core.TextHeadlineSmall)
		core.NewButton(pg).SetText("Sign out").OnClick(func(e events.Event) {
			*u = user{}
			pg.Open("sign-in")
		})
	})

	b.RunMainWindow()
}
