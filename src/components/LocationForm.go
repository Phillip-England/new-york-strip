package components

import "fmt"

func LocationForm(formErr string) (string) {

	errorElement := `
		<div id="login-form-err-container">	
			<p class="text-xs mb-2 text-red-500">` + formErr + `</p>
		</div>
		`

	html := fmt.Sprintf(`
		<form class="p-4 flex flex-col border" method="POST" action="/">
			<h2 class="mb-2 font-serif">Create a Location</h2>
			%s
			<label class="text-xs">Name</label>
			<input type="text" name="name" class="border mb-2 p-1 text-xs" value="Chick-fil-A Southroads" hx-post="/htmx/hide" hx-target="#login-form-err-container" hx-swap="outerHTML" hx-trigger="input" />
			<label class="text-xs">Number</label>
			<input type="text" name="number" class="border mb-6 p-1 text-xs" value="03253" hx-post="/htmx/hide" hx-target="#login-form-err-container" hx-swap="outerHTML" hx-trigger="input" />
			<input type="submit" class="border text-sm bg-red-500 p-2 rounded-lg text-white" />
		</form>
	`, errorElement)
	return html
}