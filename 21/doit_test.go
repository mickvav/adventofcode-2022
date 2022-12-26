package main

import "testing"

func TestMonkey_Parse(t *testing.T) {
	type fields struct {
		value    int
		name     string
		known    bool
		operator string
		op1      string
		op2      string
		group    *Monkeys
	}
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			fields: fields{
				value: 10,
				name:  "aaaa",
			},
			wantErr: false,
			args: args{
				line: "aaaa: 10",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Monkey{
				value:    tt.fields.value,
				name:     tt.fields.name,
				known:    tt.fields.known,
				operator: tt.fields.operator,
				op1:      tt.fields.op1,
				op2:      tt.fields.op2,
				group:    tt.fields.group,
			}
			var m1 *Monkey
			var err error
			if m1, err = Parse(tt.args.line); (err != nil) != tt.wantErr {
				t.Errorf("Monkey.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if m1.value != m.value {
				t.Errorf("Value misread: expected %d returned %d", m.value, m1.value)
			}
		})
	}
}
