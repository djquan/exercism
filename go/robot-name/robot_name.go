package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var names = make(map[string]bool)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Robot type represents a Robot with a name
type Robot struct {
	name string
}

/*Name generates a name for a robot if it does not have one.
Returns it's name if it does have one.
*/
func (r *Robot) Name() string {
	if r.name != "" {
		return r.name
	}
	r.Reset()

	return r.name
}

//Reset removes a robots old name and generates a new one.
func (r *Robot) Reset() {
	name := fmt.Sprintf("%c%c%d%d%d",
		letters[rand.Intn(len(letters))],
		letters[rand.Intn(len(letters))],
		rand.Intn(9),
		rand.Intn(9),
		rand.Intn(9),
	)

	if names[name] == true {
		r.Reset()
	} else {
		r.name = name
		names[name] = true
	}
}
