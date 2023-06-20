package pages

import (
	"htmx-cares/src/components"
	"htmx-cares/src/core"
	"htmx-cares/src/templates"
)

func SignupPage(b *core.GoBuild, formErr string) {
	b.Consume(components.GuestNavClosed())
	b.Consume(components.SignupForm(formErr))
	b.Inject(templates.BaseTemplate("Sign Up"))
}