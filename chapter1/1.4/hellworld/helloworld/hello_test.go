package helloworld

import "testing"

func TestPrint(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "golang1"},
		{name: "golang2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print(tt.name)
		})
	}
}
