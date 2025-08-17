package middleware

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"pwa-engine/client/elements"
	"pwa-engine/client/pages/base"
)

func RegisterRoutes() {
	app.Route("/", func() app.Composer {
		return &base.Base{
			BodyContent: app.Div().Body(
				app.H1().Text("Welcome to the PWA Engine"),
			),
		}
	})
	app.Route("/contact", func() app.Composer {
		return &base.Base{
			BodyContent: &elements.Form{},
		}
	})
}
