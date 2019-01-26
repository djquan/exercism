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
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	name := generateName()
	for names[name] {
		name = generateName()
	}

	r.name = name
	names[name] = true

	return r.name, nil
}

//Reset removes a robots old name
func (r *Robot) Reset() {
	r.name = ""
}

func generateName() string {
	return fmt.Sprintf("%c%c%03d",
		rand.Intn(26)+'A',
		rand.Intn(26)+'A',
		rand.Intn(1000),
	)
}
