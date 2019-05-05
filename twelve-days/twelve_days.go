package twelve

import (
	"bytes"
	"fmt"
)

var numerals = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var gifts = []string{"a Partridge in a Pear Tree.", "two Turtle Doves, ", "three French Hens, ", "four Calling Birds, ", "five Gold Rings, ", "six Geese-a-Laying, ", "seven Swans-a-Swimming, ", "eight Maids-a-Milking, ", "nine Ladies Dancing, ", "ten Lords-a-Leaping, ", "eleven Pipers Piping, ", "twelve Drummers Drumming, "}

// Song returns the song
func Song() string {
	var buffer bytes.Buffer
	for i := 1; i <= 12; i++ {
		buffer.WriteString(Verse(i))
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// Verse returns one verse.
func Verse(n int) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", numerals[n-1]))
	for i := n - 1; i >= 0; i-- {
		if i == 0 && n > 1 {
			buffer.WriteString("and ")
		}
		buffer.WriteString(gifts[i])
	}
	return buffer.String()
}
