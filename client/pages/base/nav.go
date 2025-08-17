package base

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Nav struct {
	app.Compo
}

func (n *Nav) Render() app.UI {
	return app.Nav().Body(
		app.A().Href("/").Text("Home"),
		app.A().Href("/contact").Text("Contact"),
	)
}
