package components

import (
	"bytes"
	"paste/alerts"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func View(title string, elements ...Node) Node {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			Link(Rel("icon"), Type("image/x-icon"), Href("/static/favicon.png")),
			Link(Rel("stylesheet"), Href("/static/css/global.css")),
			Link(Rel("stylesheet"), Href("https://sophed.github.io/iosevka-webfont/3.4.1/iosevka.css")),
		},
		Body: []Node{
			Div(Class("container"),
				Div(Class("content"),
					Group(elements),
				),
				PageFooter(),
			),
		},
	})
}

func ToString(view Node) string {
	buf := new(bytes.Buffer)
	err := view.Render(buf)
	if err != nil {
		alerts.Error(err)
		return ""
	}
	return buf.String()
}
