package markdown

import (
	"fmt"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	var header, list int
	var strong, em bool
	var html string

	for pos := 0; pos < len(markdown); pos++ {
		switch char := markdown[pos]; char {
		case '#':
			for char == '#' {
				header++
				pos++
				char = markdown[pos]
			}
			html += fmt.Sprintf("<h%d>", header)
		case '*':
			if list == 0 {
				html += "<ul>"
			}
			html += "<li>"
			list++
			pos++
		case '\n':
			if list > 0 {
				html += "</li>"
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
		case '_':
			if pos+1 < len(markdown) && markdown[pos+1] == '_' {
				if strong = !strong; strong {
					html += "<strong>"
				} else {
					html += "</strong>"
				}
				pos++
			} else {
				if em = !em; em {
					html += "<em>"
				} else {
					html += "</em>"
				}
			}
		default:
			html += string(char)
		}
	}

	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	}
	if list > 0 {
		return html + "</li></ul>"
	}
	return "<p>" + html + "</p>"
}
