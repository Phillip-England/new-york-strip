package pages

import (
	"htmx-cares/src/components"
	"htmx-cares/src/core"
	"htmx-cares/src/templates"
)

func LoginPage(snap *core.GoSnap, formErr string) {
	snap.HtmlConsume(components.GuestNavClosed())
	snap.HtmlConsume(components.LoginForm(formErr))
	snap.HtmlInject(templates.BaseTemplate("Log In"))
}