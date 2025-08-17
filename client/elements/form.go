package elements

import (
	_ "embed"
	"encoding/json"
	"log"
	"pwa-engine/client/inputs"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

var contactJSON = []byte(`
{
  "title": "Contact Us",
  "heading": "Get in touch",
  "elements": ["form"],
  "fields": [
    { "id": "first_name", "name": "firstName", "type": "text", "label": "First Name" },
    { "id": "email", "name": "email", "type": "email", "label": "Email" },
    { "id": "submit_btn", "name": "submit", "type": "button", "label": "Send" }
  ]
}
`)

type Form struct {
	app.Compo
	Config *PageConfig // We'll parse the JSON into this type
}

// PageConfig represents the structure of your JSON
type PageConfig struct {
	Heading string `json:"heading"`
	Fields  []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Label string `json:"label"`
		Type  string `json:"type"`
	} `json:"fields"`
}

func (f *Form) OnMount(ctx app.Context) {
	var cfg PageConfig
	if err := json.Unmarshal(contactJSON, &cfg); err != nil {
		log.Println("Failed to parse embedded JSON:", err)
		return
	}
	f.Config = &cfg
}

func (f *Form) Render() app.UI {
	if f.Config == nil {
		return app.Div().Text("Loading...")
	}

	var fields []app.UI
	for _, field := range f.Config.Fields {
		switch field.Type {
		case "text":
			fields = append(fields, &inputs.TextInput{
				ID:    field.ID,
				Name:  field.Name,
				Label: field.Label,
			})
		case "email":
			fields = append(fields, &inputs.EmailInput{
				ID:    field.ID,
				Name:  field.Name,
				Label: field.Label,
			})
		case "button":
			fields = append(fields, &inputs.ButtonInput{
				ID:    field.ID,
				Name:  field.Name,
				Label: field.Label,
			})
		}
	}

	return app.Div().Class("container").Body(
		app.H1().Text(f.Config.Heading),
		app.Form().Body(fields...),
	)
}
