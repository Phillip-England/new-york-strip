package pages

import (
	"htmx-cares/src/components"
	"htmx-cares/src/core"
	"htmx-cares/src/templates"
)

func LocationPage(b *core.GoBuild, name string, number string) {
	b.Consume(components.UserNavClosed())
	b.Inject(templates.BaseTemplate(name))
}