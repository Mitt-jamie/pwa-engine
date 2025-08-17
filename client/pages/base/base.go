package base

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Base struct {
	app.Compo
	BodyContent app.UI
}

func (b *Base) Render() app.UI {
	return app.Div().Body(
		&Nav{},
		app.Div().Class("container").Body(
			b.BodyContent, // <- dynamically rendered content
		),
		&Footer{},
	)
}
