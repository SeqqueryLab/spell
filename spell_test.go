package spell

import "testing"

func TestSpell(t *testing.T) {
	t.Run("Spell must return random name composed of adjective and noun", func(t *testing.T) {
		want := "empty nothing"
		got := Spell()

		if got != want {
			t.Errorf("Want %s, got %s", want, got)
		}
	})
}
