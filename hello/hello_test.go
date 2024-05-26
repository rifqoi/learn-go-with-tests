package hello

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name     string
		language string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hello 'English'",
			args: args{"Rifqi", "English"},
			want: "Hello, Rifqi",
		},
		{
			name: "hello 'Spanish'",
			args: args{"Rifqi", "Spanish"},
			want: "Hola, Rifqi",
		},
		{
			name: "hello 'Indonesia'",
			args: args{"Rifqi", "Indonesia"},
			want: "Assalamualaikum, Rifqi",
		},
		{
			name: "name is empty",
			args: args{"", "English"},
			want: "Hello, World",
		},
		{
			name: "language is empty",
			args: args{"Rifqi", ""},
			want: "Hello, Rifqi",
		},
		{
			name: "name and language is empty",
			args: args{"", ""},
			want: "Hello, World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name, tt.args.language); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
