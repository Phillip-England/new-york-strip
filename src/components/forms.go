package components

func LoginForm() (string) {
	html := `
		<form class="p-4 flex flex-col border" method="POST" action="/login">
			<h2 class="mb-2 font-serif">Log In</h2>
			<label class="text-xs">Email</label>
			<input type="text" name="email" class="border mb-2 p-1 text-xs" value="phillip@gmail.com" />
			<label class="text-xs">Password</label>
			<input type="text" name="password" class="border mb-6 p-1 text-xs" value="password" />
			<input type="submit" class="border text-sm bg-red-500 p-2 rounded-lg text-white" />
		</form>
	`
	return html
}

func SignupForm() (string) {
	html := `
		<form class="p-4 flex flex-col border" method="POST" action="/signup">
			<h2 class="mb-2 font-serif">Log In</h2>
			<label class="text-xs">Email</label>
			<input type="text" class="border mb-2 p-1 text-xs" value="phillip@gmail.com" />
			<label class="text-xs">Password</label>
			<input type="text" class="border mb-6 p-1 text-xs" value="password" />
			<input type="submit" class="border text-sm bg-red-500 p-2 rounded-lg text-white" />
		</form>
	`
	return html
}