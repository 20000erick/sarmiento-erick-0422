package util

import "testing"

func TestEncryptMd5(t *testing.T) {
	type args struct {
		texto string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{texto: "Hola"}, want: "f688ae26e9cfa3ba6235477831d5122e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptMd5(tt.args.texto); got != tt.want {
				t.Errorf("EncryptMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptMorse(t *testing.T) {
	type args struct {
		texto string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{texto: "Hola"}, want: ".... --- .-.. .-"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptMorse(tt.args.texto); got != tt.want {
				t.Errorf("EncryptMorse() = %v, want %v", got, tt.want)
			}
		})
	}
}
