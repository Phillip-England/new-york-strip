package components

import "fmt"

func LocationSelectorList() (string) {
	return fmt.Sprintf(`
		<div class="p-4">
			%s
			%s
		</div>
	`, LocationSelector("Chick-fil-A Utica", "03334"), LocationSelector("Chick-fil-A Oceiana", "33323"))
}