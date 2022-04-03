package util

import "testing"

func TestConcatena(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{strs: []string{"Hola ", "mundo"}}, want: "Hola mundo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concatena(tt.args.strs...); got != tt.want {
				t.Errorf("Concatena() = %v, want %v", got, tt.want)
			}
		})
	}
}
