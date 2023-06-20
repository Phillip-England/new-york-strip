package pages

import (
	"htmx-cares/src/components"
	"htmx-cares/src/core"
	"htmx-cares/src/templates"
)

func LocationSelectionPage(b *core.GoBuild) {
	b.Consume(components.UserNavClosed())
	b.Consume(components.LocationForm(""))
	b.Consume(components.LocationSelectorList())
	b.Inject(templates.BaseTemplate("Log In"))
}