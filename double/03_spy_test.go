package double

import "testing"

type SpySearcher struct {
	phone           string
	searchWasCalled bool
	whatIsFirstName string
}

// implement Queryer
func (ss *SpySearcher) Search(people []*Person, firstName, lastName string) *Person {
	// if got called, ss.searchWasCalled will be updated to true
	ss.searchWasCalled = true
	ss.whatIsFirstName = firstName
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindCallsSearchAndReturnPerson(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	spy := &SpySearcher{phone: fakePhone}

	phone, _ := phonebook.Find(spy, "Jane", "Doe")

	if !spy.searchWasCalled {
		t.Errorf("Expected to call 'Search' in 'Find', but it wasn't called")
	}

	// we will know that the firstName won't be modified before passed
	if spy.whatIsFirstName != "Jane" {
		t.Errorf("Expected to call 'Search' with '%s' as firstName , but got '%s'", "Jane", spy.whatIsFirstName)
	}

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}
}
