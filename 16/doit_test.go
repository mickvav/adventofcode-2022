package main

import (
	"testing"
)

func TestState_Open(t *testing.T) {
	type fields struct {
		valvestatus uint64
		actorpos    string
	}
	type args struct {
		v string
		m *Map
	}
	tests := []struct {
		name                string
		fields              fields
		args                args
		expectedvalvestatus uint64
	}{
		{
			name: "simple",
			fields: fields{
				valvestatus: 0,
				actorpos:    "AA",
			},
			args: args{
				v: "AA",
				m: &Map{
					valvenames:   []string{"AA"},
					valvenumbers: map[string]int{"AA": 0},
					valves:       map[string]int{"AA": 1},
				},
			},
			expectedvalvestatus: 1,
		},
		{
			name: "simple2",
			fields: fields{
				valvestatus: 0,
				actorpos:    "BB",
			},
			args: args{
				v: "BB",
				m: &Map{
					valvenames:   []string{"AA", "BB"},
					valvenumbers: map[string]int{"AA": 0, "BB": 1},
					valves:       map[string]int{"AA": 1, "BB": 2},
				},
			},
			expectedvalvestatus: 2,
		},
		{
			name: "simple3",
			fields: fields{
				valvestatus: 1,
				actorpos:    "BB",
			},
			args: args{
				v: "BB",
				m: &Map{
					valvenames:   []string{"AA", "BB"},
					valvenumbers: map[string]int{"AA": 0, "BB": 1},
					valves:       map[string]int{"AA": 1, "BB": 2},
				},
			},
			expectedvalvestatus: 3,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &State{
				valvestatus: tt.fields.valvestatus,
				actorpos:    tt.fields.actorpos,
			}
			s.Open(tt.args.v, tt.args.m)
			if tt.expectedvalvestatus != s.valvestatus {
				t.Errorf("Wrong valve status: %d != %d", tt.expectedvalvestatus, s.valvestatus)
			}
		})
	}
}

func TestState_Income(t *testing.T) {
	type fields struct {
		valvestatus uint64
		actorpos    string
	}
	type args struct {
		m *Map
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "simple",
			fields: fields{
				valvestatus: 0b10,
				actorpos:    "AA",
			},
			args: args{
				m: &Map{
					valves:       map[string]int{"AA": 0, "BB": 5},
					valvenames:   []string{"AA", "BB"},
					valvenumbers: map[string]int{"AA": 0, "BB": 1},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &State{
				valvestatus: tt.fields.valvestatus,
				actorpos:    tt.fields.actorpos,
			}
			if got := s.Income(tt.args.m); got != tt.want {
				t.Errorf("State.Income() = %v, want %v", got, tt.want)
			}
		})
	}
}
