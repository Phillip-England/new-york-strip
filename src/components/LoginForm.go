package components

import "fmt"

func LoginForm(formErr string) (string) {

	errorElement := `
		<div id="login-form-err-container">	
			<p class="text-xs mb-2 text-red-500">` + formErr + `</p>
		</div>
		`

	html := fmt.Sprintf(`
		<form class="p-4 flex flex-col border" method="POST" action="/">
			<h2 class="mb-2 font-serif">Log In</h2>
			%s
			<label class="text-xs">Email</label>
			<input type="text" name="email" class="border mb-2 p-1 text-xs" value="phillip@gmail.com" hx-post="/htmx/hide" hx-target="#login-form-err-container" hx-swap="outerHTML" hx-trigger="input" />
			<label class="text-xs">Password</label>
			<input type="text" name="password" class="border mb-6 p-1 text-xs" value="password" hx-post="/htmx/hide" hx-target="#login-form-err-container" hx-swap="outerHTML" hx-trigger="input" />
			<input type="submit" class="border text-sm bg-red-500 p-2 rounded-lg text-white" />
		</form>
	`, errorElement)
	return html
}