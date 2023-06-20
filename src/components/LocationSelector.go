package components

import "fmt"

func LocationSelector(name string, number string) string {
	return fmt.Sprintf(`
		<div class="mb-4 p-4 border rounded-lg flex flex-row justify-between">
			<div>
				<h2 class="font-serif text-sm">%s</h2>
				<p class="text-xs">%s</p>
			</div>
			<div class="self-center">
				<svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 rotate-90  text-black" fill="none" viewBox="0 0 24 24" stroke="currentColor">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
				</svg>
			</div>
		</div>
	`, name, number)
}
