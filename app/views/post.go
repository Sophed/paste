package views

import (
	"paste/app/components"

	. "maragu.dev/gomponents"
)

func Post(id, content string) Node {
	return components.View(id,
		Raw(content),
	)
}
