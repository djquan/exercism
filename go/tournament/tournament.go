// Package tournament provides a method for calculating the results of a tournament.
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	victories, defeats, ties int
}

func (t *team) score() int {
	return (3 * t.victories) + t.ties
}

func (t *team) played() int {
	return t.victories + t.defeats + t.ties
}

type kv struct {
	key   string
	value int
}

var fmtString = "%-31s| %2d | %2d | %2d | %2d | %2d\n"

// Tally calculates the outcome of a tournament
func Tally(reader io.Reader, writer io.Writer) error {
	teams := make(map[string]team)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		m := scanner.Text()
		match := strings.Split(m, ";")

		if match[0] == "" || strings.HasPrefix(match[0], "#") {
			continue
		} else if len(match) != 3 {
			return fmt.Errorf("Unable to parse match: %s", m)
		}

		home, away, outcome := teams[match[0]], teams[match[1]], match[2]
		switch outcome {
		case "win":
			home.victories++
			away.defeats++
		case "loss":
			away.victories++
			home.defeats++
		case "draw":
			home.ties++
			away.ties++
		default:
			return fmt.Errorf("Unknown game outcome for %s", m)
		}

		teams[match[0]], teams[match[1]] = home, away
	}

	s := make([]kv, len(teams))
	for n, t := range teams {
		s = append(s, kv{
			fmt.Sprintf(fmtString, n, t.played(), t.victories, t.ties, t.defeats, t.score()),
			t.score(),
		})
	}
	sort.Slice(s, func(i, j int) bool {
		a, b := s[i].value, s[j].value
		if a != b {
			return a > b
		}
		return s[i].key < s[j].key
	})

	io.WriteString(writer, "Team                           | MP |  W |  D |  L |  P\n")
	for _, t := range s {
		io.WriteString(writer, t.key)
	}
	return nil
}
