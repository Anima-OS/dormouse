/* Copyright (c) 2017, Mark Bauermeister
 *
 * This software may be modified and distributed under the terms
 * of the BSD license.  See the LICENSE file for details.
 */

package dormouse

// Emoji type
type emoji struct {
	Name   string
	symbol string
}

// Emojis collection
var Emojis = struct {
	Bear     emoji
	Dog      emoji
	Cat      emoji
	Mouse    emoji
	TestTEST emoji
}{
	Bear:     emoji{Name: "🐻", symbol: "A"},
	Dog:      emoji{Name: "🐶", symbol: "T"},
	Cat:      emoji{Name: "🐱"},
	Mouse:    emoji{Name: "🐭"},
	TestTEST: emoji{Name: "🐭🐱🐶"},
}
