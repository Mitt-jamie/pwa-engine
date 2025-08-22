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
		Text(b.Label).
		Class(
			"bg-gray-100 text-gray-900 font-semibold px-6 py-2 rounded-sm " +
				"hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-400 " +
				"transition-colors duration-150",
		)
}
