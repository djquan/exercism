package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const MaxNames = 26 * 26 * 1000

var names = map[string]bool{"": true}

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

	for names[r.name] {
		if len(names) > MaxNames {
			return "", fmt.Errorf("No more possible names to generate")
		}

		r.name = fmt.Sprintf("%c%c%03d",
			rand.Intn(26)+'A',
			rand.Intn(26)+'A',
			rand.Intn(1000),
		)
	}

	names[r.name] = true
	return r.name, nil
}

//Reset removes a robots old name
func (r *Robot) Reset() {
	r.name = ""
}
