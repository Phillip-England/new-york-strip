package core

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type GoBuild struct {
	Html string
}

func NewGoBuild() (GoBuild) {
	return GoBuild{
		Html: "",
	}
}

func (b *GoBuild) Consume(html string) {
	b.Html = b.Html + html
}

func (b *GoBuild) GetHtmlBytes() ([]byte) {
	return []byte(b.Html)
}

func (b *GoBuild) Inject(template string) {
	const divID = "go-root"
	injectedPage := strings.Replace(template, "<div id=\""+divID+"\"></div>", "<div id=\""+divID+"\">"+b.Html+"</div>", 1)
	b.Html = injectedPage
}

func (b *GoBuild) Serve(c *gin.Context) {
	c.Data(200, "text/html; charset=utf-8", b.GetHtmlBytes())
}
