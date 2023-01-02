package teardown

import "testing"

func setup(t *testing.T) func() {
	t.Log("setup")

	return func() {
		t.Log("teardown")
	}
}

func TestTeardown(t *testing.T) {
	teardown := setup(t)
	// do teardown later after all test end
	defer teardown() // or t.Cleanup(teardown)

	t.Run("subtest", func(t *testing.T) {
		t.Log("subtest")
	})
}
