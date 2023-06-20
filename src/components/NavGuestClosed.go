package components

func GuestNavClosed() (string) {
	html := `
	<div id="nav-container">

		<!-- navbar -->
		<div id="nav-guest" class="bg-white fixed h-16 w-screen top-0 left-0 p-4 border-b flex flex-row justify-between z-20">
			<h1 class="text-xl font-serif">Chick-fil-A Tools</h1>
			<div class="flex items-center">
				<i class="fa-bars fa-solid fa-lg" hx-post="/htmx/open_guest_navigation" hx-trigger="click" hx-target="#nav-container", hx-swap="outerHTML"></i>
			</div>
		</div>
		
		<!-- spacer -->
		<div class="h-16"></div>

	</div>
	`
	return html
}