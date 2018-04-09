/* Copyright (c) 2017, Mark Bauermeister
 *
 * This software may be modified and distributed under the terms
 * of the BSD license.  See the LICENSE file for details.
 */

package dormouse

import (
	"fmt"
	"os"
)

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

	Render("<b>HELLO WORLD!</b>")
	Render("<b style='color:red;'>HELLO ANOTHER WORLD!</b>")
}

// Render HTML to output (internal function).
func Render(HTML string) {

	input := fmt.Sprintf(escape+
		"[?1155;%s"+
		"h"+
		HTML+
		escape+
		"[?1155"+
		"l", termcookie)

	os.Stdout.Write([]byte(input))
}

// DisplayImage renders blockimages to the terminal.
func DisplayImage(URL string) {

	HTML := "<img class='gterm-blockimg'  src='" +
		URL +
		"'"

	Render(HTML)
}

// DisplayVideo renders video frames to the terminal/embeds them in the terminal's media player.
func DisplayVideo(URL string) {}

// DisplayText renders stylized text to the terminal (Font ought to be an enumarable list. Color/Size should be optional).
func DisplayText(Text string, Font string, FontSize int, Color string) {
	HTML := fmt.Sprintf("<span style='font-size:%dpx;'>"+
		"%s"+
		"</span", FontSize, Text)

	Render(HTML)
}

// DisplayPDF renders the contents of a PDF to the terminal/the terminal's PDF reader.
func DisplayPDF(URL string) {}

// DisplayEmoji renders emojis to the terminal (Emoji ought to be an enumarable list. Size should be optional + int).
func DisplayEmoji(Emoji emoji, Size int) {

	HTML := fmt.Sprintf("<span style='font-size:%dpx;' class='emoji'>"+
		"%s"+
		"</span", Size, Emoji.Name)

	Render(HTML)
}

// DisplayFrame renders iFrames and their embedded contents to the terminal.
func DisplayFrame(URL string) {}

// DisplayDialog renders a dialog the terminal (Type ought to be an eunamarable list of all basic HTML dialog types + the basic Anima overlay as 0)
// Make sure there can only ever be one dialog open at a time.
func DisplayDialog(Type string) {}

// DisplaySyntax renders input as syntax highlighted code.
func DisplaySyntax(Code string, Language string) {}

// DisplayHotlink renders a clickable hotlink to the terminal (DisplayName should be optional).
func DisplayHotlink(URL string, DisplayName string) {

	HTML := "<a href='" +
		URL + "'>" +
		DisplayName +
		"</a>"

	Render(HTML)
}

// DisplayHTML renders arbitrary HTML to the terminal.
func DisplayHTML(HTML string) {

	Render(HTML)
}
