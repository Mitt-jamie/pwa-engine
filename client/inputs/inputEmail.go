package inputs

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type EmailInput struct {
	app.Compo
	ID    string
	Name  string
	Label string
}

func (e *EmailInput) Render() app.UI {
	return app.Div().Body(
		app.Label().For(e.ID).Text(e.Label),
		app.Input().ID(e.ID).Name(e.Name).Type("email"),
	)
}
