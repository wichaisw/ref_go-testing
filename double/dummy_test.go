package double

import "testing"

type DummySearcher struct{}

// Dummy has same interface as Queryer just to make it able to be compiled
func (ds DummySearcher) Search(people []*Person, firstName string, lastName string) *Person {
	return &Person{}
}

func TestFind(t *testing.T) {
	t.Run("should return error when FirstName or LastName is empty", func(t *testing.T) {
		phonebook := &Phonebook{}
		want := ErrMissingArgs

		dd := DummySearcher{}
		_, got := phonebook.Find(dd, "", "")

		if got != want {
			t.Errorf("Want '%s', got '%s'", want, got)
		}
	})
}
