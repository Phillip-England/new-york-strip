package components

func GuestNavClosed() (string) {
	html := `
	<div id="nav-container">

		<!-- navbar -->
		<div id="nav-guest" class="bg-white fixed h-16 w-screen top-0 left-0 p-4 border-b flex flex-row justify-between z-20">
			<h1 class="text-xl font-serif">Chick-fil-A Tools</h1>
			<div class="flex items-center">
				<i class="fa-bars fa-solid fa-lg" hx-post="/components/open_navigation" hx-trigger="click" hx-target="#nav-container", hx-swap="outerHTML"></i>
			</div>
		</div>
		
		<!-- spacer -->
		<div class="h-16"></div>

	</div>
	`
	return html
}


func GuestNavOpened() (string) {
	html := `
		<div id="nav-container">

			<!-- navbar -->
			<div id="nav-guest" class="bg-white fixed h-16 w-screen top-0 left-0 p-4 flex flex-row justify-between z-20">
				<h1 class="text-xl font-serif">Chick-fil-A Tools</h1>
				<div class="flex items-center">
					<i class="fa-x fa-solid fa-lg" hx-post="/components/close_navigation" hx-trigger="click" hx-target="#nav-container", hx-swap="outerHTML"></i>
				</div>
			</div>
			
			<!-- overlay -->
			<div class="fixed top-0 left-0 w-screen h-screen bg-black opacity-50 z-10" hx-post="/components/close_navigation" hx-trigger="click" hx-target="#nav-container", hx-swap="outerHTML"></div>

			<!-- spacer -->
			<div class="h-16"></div>

			<!-- navigation menu -->
			<nav class="fixed top-16 left-0 w-3/5 h-screen bg-white z-20">
				<ul class="flex flex-col">
					<a class="border p-4" href="/signup">
						<li class="font-serif">Signup</li>
					</a>
					<a class="border p-4" href="/">
						<li class="font-serif">Login</li>
					</a>
				</ul>
			</nav>

		</div>
	`
	return html
}