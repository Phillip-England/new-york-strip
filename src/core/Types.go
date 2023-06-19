package core

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type GoSnap struct {
	Html string
	HtmlBytes []byte
}

func NewGoSnap() (GoSnap) {
	return GoSnap{
		Html: "",
	}
}

func (snap *GoSnap) HtmlConsume(html string) {
	snap.Html = snap.Html + html
}

func (snap *GoSnap) GetHtmlBytes() ([]byte) {
	return []byte(snap.Html)
}

func (snap *GoSnap) HtmlInject(page string) {
	const divID = "go-root"
	injectedPage := strings.Replace(page, "<div id=\""+divID+"\"></div>", "<div id=\""+divID+"\">"+snap.Html+"</div>", 1)
	snap.Html = injectedPage
}

func (snap *GoSnap) HtmlServe(c *gin.Context) {
	c.Data(200, "text/html; charset=utf-8", snap.GetHtmlBytes())
}
