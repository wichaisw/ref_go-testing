package double

import "testing"

type MockSearcher struct {
	phone         string
	methodsToCall map[string]bool // like spy, but expect dynamic values
}

func (ms *MockSearcher) Search(people []*Person, firstName, lastName string) *Person {
	ms.methodsToCall["Search"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearcher) Create(people []*Person, firstName, lastName string) *Person {
	ms.methodsToCall["Create"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearcher) ExpectToCall(methodName string) {
	if ms.methodsToCall == nil {
		ms.methodsToCall = make(map[string]bool)
	}

	ms.methodsToCall[methodName] = false
}

func (ms *MockSearcher) Verify(t *testing.T) {
	for methodName, called := range ms.methodsToCall {
		if !called {
			t.Errorf("Expected to call '%s', but it wasn't.", methodName)
		}
	}
}

func TestFindCallsSearchAndReturnsPersonUsingMock(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	mock := &MockSearcher{phone: fakePhone}

	// init ms.methodsToCall["Search"] value as false
	mock.ExpectToCall("Search")

	// .Find will call .Search and set ms.methodsToCall["Search"] = true
	// if it doesn't get called,
	phone, _ := phonebook.Find(mock, "John", "Doe")

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}

	// equivalent to toHaveBeenCalled in Jest
	// Go doesn't have that, we have to create Mock Object
	// some language handle this by dependency injection
	mock.Verify(t)
}
