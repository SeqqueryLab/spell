package spell

import (
	"fmt"

	words "github.com/utubun/spell/internal"
)

func Spell() string {
	return fmt.Sprintf("%s %s of %s", words.Adjective(), words.Noun(), words.Noun())
}
