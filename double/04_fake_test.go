package double

import "testing"

type FakeSearcher struct{}

func (fs FakeSearcher) Search(people []*Person, firstName string, lastName string) *Person {
	if len(people) == 0 {
		return nil
	}

	return people[0]
}

func TestFindCallsSearchAndReturnsEmptyStringForNoPerson(t *testing.T) {
	phonebook := &Phonebook{}
	fake := &FakeSearcher{}

	// phonebook is empty, when call Find, which call fake.Search inside, will return people[0]
	phone, _ := phonebook.Find(fake, "Jane", "Doe")

	if phone != "" {
		t.Errorf("Wanted '', got '%s'", phone)
	}
}
