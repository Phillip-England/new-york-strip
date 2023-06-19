package main

import (
	"fmt"
	"htmx-cares/src/components"
	"htmx-cares/src/core"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// -----------------------------
	// CONFIGURATION
	// -----------------------------

	// config
	godotenv.Load()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// mongodb
	// mongoStore, err := database.NewMongoStore()
	// if err != nil {
	// 	log.Fatal("failed to connect to mongo db")
	// }
	// defer mongoStore.Client.Disconnect(context.Background())

	// -----------------------------
	// PAGES
	// -----------------------------

	// signup page
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// login page
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(200, "page_signup.html", nil)
	})

	// location page
	router.GET("/locations", func(c *gin.Context) {
		c.HTML(200, "page_location.html", nil)
	})

	// -----------------------------
	// CONTROLLERS
	// -----------------------------

	// logging in our user
	router.POST("/login", func(c *gin.Context) {
		fmt.Println("hitem!")
	})

	// -----------------------------
	// HTMX ACTIONS
	// -----------------------------

	// this is a quick and dirty way to hide an element from the sreen by replacing it with empty html
	router.POST("/hide", func(c *gin.Context) {
		html := ``
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	})

	// opens our users navigation
	// this navigation will read the http-cookies along with the request
	// these cookies will determine which navigation menu is rendered
	router.POST("/components/open_navigation", func(c *gin.Context) {

		// if the user is not authenticated, sent back the guest naviation
		html, err := os.ReadFile("templates/nav_guest_opened.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read HTML file")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", html)

	})

	// close the navigation
	// this navigation will read the http-cookies along with the request
	// these cookies will determine which nvaigtation menu is rendered
	router.POST("/components/close_navigation", func(c *gin.Context) {

		// if the user is not authenticated, send back the closed guest navigation
		html, err := os.ReadFile("templates/nav_guest_closed.html")
		if err != nil {
			// Handle the error if the file couldn't be read
			c.String(http.StatusInternalServerError, "Failed to read HTML file")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", html)

	})


	// ------------------------
	// TESTING GOSNAP
	// ------------------------

	router.GET("/test", func(c *gin.Context) {
		snap := core.NewGoSnap()
		components.LoginForm(&snap)
		c.Data(200, "text/html; charset=utf-8", snap.GetHtmlBytes())
	})

	// ------------------------
	// RUNNING OUR WEBSERVER
	// ------------------------

	// running
	router.Run()


}