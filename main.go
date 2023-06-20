package main

import (
	"context"
	"htmx-cares/src/components"
	"htmx-cares/src/core"
	"htmx-cares/src/models"
	"htmx-cares/src/pages"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	// -----------------------------
	// CONFIGURATION
	// -----------------------------

	// config
	godotenv.Load()
	r := gin.Default()
	// r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

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
	r.POST("/", func(c *gin.Context) {
		b := core.NewGoBuild()
		userModel := models.NewUserModel(c.PostForm("email"), c.PostForm("password"))
		httpErr := userModel.Find(mongoStore.UserCollection)
		if httpErr != nil {
			if httpErr.Code == 500 {
				pages.ServerErrorPage(&b)
				b.Serve(c)
				return
			}
			if httpErr.Code == 400 {
				pages.LoginPage(&b, httpErr.Message)
				b.Serve(c)
				return
			}
		}
		if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(c.PostForm("password"))); err != nil {
			pages.LoginPage(&b, "invalid credentials")
			b.Serve(c)
			return
		}

		pages.LoginPage(&b, "")
		b.Serve(c)
	})

	// signing up our user
	r.POST("/signup", func(c *gin.Context) {
		b := core.NewGoBuild()
		userModel := models.NewUserModel(c.PostForm("email"), c.PostForm("password"))
		httpErr := userModel.Insert(mongoStore.UserCollection)
		if httpErr != nil {
			if httpErr.Code == 500 {
				pages.ServerErrorPage(&b)
				b.Serve(c)
				return
			}
			if httpErr.Code == 400 {
				pages.SignupPage(&b, httpErr.Message)
				b.Serve(c)
				return
			}
		}
		pages.LoginPage(&b, "")
		b.Serve(c)
	})

	// -----------------------------
	// HTMX ACTIONS
	// -----------------------------

	// opening our navigation
	r.POST("/htmx/open_guest_navigation", func(c *gin.Context) {
		b := core.NewGoBuild()
		b.Consume(components.GuestNavOpened())
		b.Serve(c)
	})

	// close the navigation
	r.POST("/htmx/close_guest_navigation", func(c *gin.Context) {
		b := core.NewGoBuild()
		b.Consume(components.GuestNavClosed())
		b.Serve(c)
	})

	// hiding an element
	r.POST("/htmx/hide", func(c *gin.Context) {
		b := core.NewGoBuild()
		b.Serve(c)
	})


	// ------------------------
	// PAGES
	// ------------------------

	// login page
	r.GET("/", func(c *gin.Context) {
		b := core.NewGoBuild()
		pages.LoginPage(&b, "")
		b.Serve(c)
	})

	// sign up page
	r.GET("/signup", func(c *gin.Context) {
		b := core.NewGoBuild()
		pages.SignupPage(&b, "")
		b.Serve(c)
	})

	// application page
	r.GET("/app", func(c *gin.Context) {
		b := core.NewGoBuild()
		b.Consume(`<p>testing</p>`)
		b.Serve(c)
	})

	// ------------------------
	// RUNNING OUR WEBSERVER
	// ------------------------

	// running
	r.Run()


}