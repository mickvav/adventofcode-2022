package main

import "testing"

func TestMove(t *testing.T) {
	a := Number{value: 1, is_moved: false}
	a1 := Number{value: -1, is_moved: false}
	b := Number{value: 2, is_moved: false}
	c := Number{value: 3, is_moved: false}
	type args struct {
		r *Numbers
		i int
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "simple",
			args: args{
				r: &Numbers{
					initial_order: &[]*Number{&a, &b, &c},
					actual_order:  &[]*Number{&a, &b, &c},
					N:             3,
				},
				i: 0,
			},
			expected: " 2 1 3",
		},
		{
			name: "negative",
			args: args{
				r: &Numbers{
					initial_order: &[]*Number{&a1, &b, &c},
					actual_order:  &[]*Number{&a1, &b, &c},
					N:             3,
				},
				i: 0,
			},
			expected: " 2 -1 3",
		},
		{
			name: "positive over the edge",
			args: args{
				r: &Numbers{
					initial_order: &[]*Number{&a, &b, &c},
					actual_order:  &[]*Number{&a, &b, &c},
					N:             3,
				},
				i: 1,
			},
			expected: " 1 2 3",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Move(tt.args.r, tt.args.i)
			if String(tt.args.r) != tt.expected {
				t.Errorf("Unexpected: %s != %s", String(tt.args.r), tt.expected)
			}
		})
	}
}
