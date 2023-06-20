package main

import (
	"context"
	"fmt"
	"htmx-cares/src/components"
	"htmx-cares/src/core"
	"htmx-cares/src/models"
	"htmx-cares/src/pages"
	"htmx-cares/src/templates"
	"log"

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
	mongoStore, err := core.NewMongoStore()
	if err != nil {
		log.Fatal("failed to connect to mongo db")
	}
	defer mongoStore.Client.Disconnect(context.Background())

	// -----------------------------
	// CONTROLLERS
	// -----------------------------

	// logging in our user
	router.POST("/login", func(c *gin.Context) {
		snap := core.NewGoSnap()
		userModel := models.NewUserModel(c.PostForm("email"), c.PostForm("password"))
		httpErr := userModel.Insert(mongoStore.UserCollection)
		if httpErr != nil {
			if httpErr.Code == 500 {
				c.Request.URL.Path = "/"
				router.HandleContext(c)
				return
			}
			if httpErr.Code == 400 {
				pages.LoginPage(&snap, httpErr.Message)
				snap.HtmlServe(c)
				return
			}
		}
		fmt.Println(userModel)
		c.Redirect(303, "/app")
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
		pages.LoginPage(&snap, "")
		snap.HtmlServe(c)
	})

	// sign up page
	router.GET("/signup", func(c *gin.Context) {
		snap := core.NewGoSnap()
		snap.HtmlConsume(components.GuestNavClosed())
		snap.HtmlConsume(components.SignupForm())
		snap.HtmlInject(templates.BaseTemplate("Sign Up"))
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