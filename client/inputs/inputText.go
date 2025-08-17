package inputs

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type TextInput struct {
	app.Compo
	ID    string
	Name  string
	Label string
}

func (t *TextInput) Render() app.UI {
	return app.Div().Body(
		app.Label().For(t.ID).Text(t.Label),
		app.Input().ID(t.ID).Name(t.Name).Type("text"),
	)
}
