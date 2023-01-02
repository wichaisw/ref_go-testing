package variadicsum

import "testing"

func TestSum(t *testing.T) {
	t.Run("should receive multiple parameters", func(t *testing.T) {
		// Arrange
		want := 0
		//Act
		got := sum([]int{}...)
		//Assert
		if got != want {
			t.Errorf("[]int{}... = %d; want %d", got, want)
		}
	})

	t.Run("including sign integer", func(t *testing.T) {
		want := 7
		xs := []int{2, 3, 3, -1}

		got := sum(xs...)

		if got != want {
			t.Errorf("sum([]int{2, 3, 3, -1}...) = %d; want %d", got, want)
		}
	})
}
