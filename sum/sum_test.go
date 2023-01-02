package sum

import "testing"

func TestSum(t *testing.T) {
	t.Log("testing sum!")

	t.Run("should return 3 when input 1 and 2", func(t *testing.T) {
		// Arrange (mocking phase)
		want := 3

		// Act
		got := sum(1, 2)

		// Assert
		// if this method is called, the test will fail
		if got != 3 {
			t.Errorf("sum(1, 2) = %d; want %d", got, want)
		}
	})
	t.Run("should return 1 when input 1 and 0", func(t *testing.T) {
		want := 1

		got := sum(1, 0)

		if got != 1 {
			t.Errorf("sum(0,1) = %d; want %d", got, want)
		}
	})

	t.Run("should return minus 2 when input minus 1 and minus 1", func(t *testing.T) {
		want := -2

		got := sum(-1, -1)

		if got != -2 {
			t.Errorf("sum(-1,-1) = %d; want %d", got, want)
		}
	})
}
