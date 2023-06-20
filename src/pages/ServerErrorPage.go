package pages

import (
	"htmx-cares/src/components"
	"htmx-cares/src/core"
	"htmx-cares/src/templates"
)

func ServerErrorPage(b *core.GoBuild) {
	b.Consume(components.GuestNavClosed())
	b.Inject(templates.BaseTemplate("Server Error"))
}