package main

import (
	"htmx-cares/src/components"
	"htmx-cares/src/core"
	"htmx-cares/src/pages"

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
	// router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// mongodb
	// mongoStore, err := database.NewMongoStore()
	// if err != nil {
	// 	log.Fatal("failed to connect to mongo db")
	// }
	// defer mongoStore.Client.Disconnect(context.Background())

	// -----------------------------
	// CONTROLLERS
	// -----------------------------

	// logging in our user
	router.POST("/login", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		if email == "phillip@gmail.com" && password == "password" {
			c.Redirect(303, "/app")
			return
		}
		c.Redirect(303, "/")
	})

	// -----------------------------
	// HTMX ACTIONS
	// -----------------------------

	// opening our navigation
	router.POST("/htmx/open_guest_navigation", func(c *gin.Context) {
		snap := core.NewGoSnap()
		snap.HtmlConsume(components.GuestNavOpened())
		snap.HtmlServe(c)
	})

	// close the navigation
	router.POST("/htmx/close_guest_navigation", func(c *gin.Context) {
		snap := core.NewGoSnap()
		snap.HtmlConsume(components.GuestNavClosed())
		snap.HtmlServe(c)
	})


	// ------------------------
	// PAGES
	// ------------------------

	// login page
	router.GET("/", func(c *gin.Context) {
		snap := core.NewGoSnap()
		snap.HtmlConsume(components.GuestNavClosed())
		snap.HtmlConsume(components.LoginForm())
		snap.HtmlInject(pages.BasePage("Log In"))
		snap.HtmlServe(c)
	})

	// sign up page
	router.GET("/signup", func(c *gin.Context) {
		snap := core.NewGoSnap()
		snap.HtmlConsume(components.GuestNavClosed())
		snap.HtmlConsume(components.SignupForm())
		snap.HtmlInject(pages.BasePage("Sign Up"))
		snap.HtmlServe(c)
	})

	// application page
	router.GET("/app", func(c *gin.Context) {
		snap := core.NewGoSnap()
		snap.HtmlConsume(`<p>testing</p>`)
		snap.HtmlServe(c)
	})

	// ------------------------
	// RUNNING OUR WEBSERVER
	// ------------------------

	// running
	router.Run()


}