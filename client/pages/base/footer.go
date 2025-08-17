package base

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Footer struct {
	app.Compo
}

func (f *Footer) Render() app.UI {
	return app.Footer().Text("© 2025 PWA Engine")
}
