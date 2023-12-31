package pages

import (
	"htmx-cares/src/components"
	"htmx-cares/src/core"
	"htmx-cares/src/templates"
)

func LoginPage(b *core.GoBuild, formErr string) {
	b.Consume(components.GuestNavClosed())
	b.Consume(components.LoginForm(formErr))
	b.Inject(templates.BaseTemplate("Log In"))
}