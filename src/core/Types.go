package core

type GoSnap struct {
	Html string
	HtmlBytes []byte
}

func NewGoSnap() (GoSnap) {
	return GoSnap{
		Html: "",
	}
}

func (snap *GoSnap) HtmlGobble(html string) {
	snap.Html = snap.Html + html
}

func (snap *GoSnap) GetHtmlBytes() ([]byte) {
	return []byte(snap.Html)
}