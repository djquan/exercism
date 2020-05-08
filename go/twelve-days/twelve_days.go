package twelve

import (
	"bytes"
	"fmt"
	"strings"
)

type DayGift struct {
	day  string
	gift string
}

var giftList = []DayGift{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Song provides the lyrics for the Twelve Days of Christmas song.
func Song() string {
	var buffer bytes.Buffer

	for i := 1; i <= 12; i++ {
		fmt.Fprintf(&buffer, "%v\n", Verse(i))
	}

	return strings.TrimSpace(buffer.String())
}

// Verse provides the appropriate lyric for the given day of Christmas.
func Verse(n int) string {
	var buffer bytes.Buffer
	result := make([]string, 0, n)

	fmt.Fprintf(&buffer, "On the %v day of Christmas my true love gave to me:", giftList[n-1].day)

	for i := n - 1; i >= 0; i-- {
		if n != 1 && i == 0 {
			result = append(result, fmt.Sprintf("and %v", giftList[i].gift))
			break
		}

		result = append(result, giftList[i].gift)
	}

	fmt.Fprintf(&buffer, " %v.", strings.Join(result, ", "))

	return buffer.String()
}
