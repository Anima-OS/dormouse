/* Copyright (c) 2017, Mark Bauermeister
 *
 * This software may be modified and distributed under the terms
 * of the BSD license.  See the LICENSE file for details.
 */

package dormouse

import "os"

const escape = "\033"
const code = "1155"

var termcookie = os.Getenv("GTERM_COOKIE")

func main() {
	// Bash: echo "${esc}[?${gterm_code};${gterm_cookie}h<b>Hello World</b>${esc}[?${gterm_code}l

	os.Stdout.Write([]byte(escape +
		"[?1155;" + termcookie +
		"h" +
		"<img src='http://clipartist.net/social/clipartist.net/B/bengal_tiger.svg'/><h2 style='background: linear-gradient(to right, red 0%, orange 22%, yellow 36%, green 50%, blue 74%, indigo 88%, violet 95%)'>" +
		"Welcome to the 21st century. Terminal emulators need not be bland and boring anymore.</h2>" +
		"<span style='background-image: -webkit-gradient(linear, left top, right top, color-stop(0, #f22), color-stop(0.15, #f2f), color-stop(0.3, #22f), color-stop(0.45, #2ff), color-stop(0.6, #2f2),color-stop(0.75, #2f2), color-stop(0.9, #ff2), color-stop(1, #f22) ); -webkit-background-clip: text; color: transparent;'> Rainbows are colorful and scalable and lovely in general</span>" +
		escape +
		"[?1155" +
		"l"))

	render("<b>HELLO WORLD!</b>")
	render("<b style='color:red;'>HELLO ANOTHER WORLD!</b>")
}

// Render HTML to output (internal function).
func render(HTML string) {
	os.Stdout.Write([]byte(escape +
		"[?1155;" + termcookie +
		"h" +
		HTML +
		escape +
		"[?1155" +
		"l"))
}

// DisplayImage renders blockimage to terminal.
func DisplayImage(URL string) {

	HTML := "<img src='" +
		URL +
		"'"

	render(HTML)
}
