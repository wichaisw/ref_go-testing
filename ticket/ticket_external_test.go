// black box testing
package ticket_test

import (
	"testing"

	"github.com/wichaisw/ref_go-testing/ticket"
)

func TestTicket(t *testing.T) {
	t.Run("should return 0 when age is 3", func(t *testing.T) {
		want := 0.0

		got := ticket.Price(3)

		if got != want {
			t.Errorf("ticket.Price(3) = %f; want %f", got, want)
		}
	})
}
