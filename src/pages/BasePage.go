package pages

func BasePage(title string) (string) {
	page := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<link rel="stylesheet" href="/static/output.css" />
			<title>`+ title +`</title>
		</head>
		<body>
			<div id="go-root"></div>
		</body>
		<script src="https://unpkg.com/htmx.org@1.9.2"></script>
		<script src="https://kit.fontawesome.com/ef0709a418.js" crossorigin="anonymous"></script>
		</html>
	`
	return page
}