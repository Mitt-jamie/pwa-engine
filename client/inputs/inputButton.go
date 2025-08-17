package inputs

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type ButtonInput struct {
	app.Compo
	ID    string
	Name  string
	Label string
}

func (b *ButtonInput) Render() app.UI {
	return app.Button().
		ID(b.ID).
		Name(b.Name).
		Text(b.Label)
}
