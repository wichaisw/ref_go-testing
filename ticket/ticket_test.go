package ticket

import "testing"

func TestPrice(t *testing.T) {
	// test table
	tests := []struct {
		name string
		age  int
		want float64
	}{
		{name: "Free Ticket when age is 0", age: 0, want: 0.0},
		{"Free Ticket when age is 3", 3, 0.0},
		{"Ticket $15 when age is 4", 4, 15.0},
		{"Ticket $15 when age is 15", 15, 15.0},
		{"Ticket $30 when age is 16", 16, 30.0},
		{"Ticket $30 when age is 50", 50, 30.0},
		{"Ticket $5 when age is over 50", 51, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Price(tt.age)

			if got != tt.want {
				t.Errorf("Price(%d) = %f, want %f", tt.age, tt.want, got)
			}
		})
	}
}
