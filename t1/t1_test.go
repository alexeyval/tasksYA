package t1

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

func TestKolya(t *testing.T) {
	s := bytes.Buffer{}
	s.WriteString(`3 3 12
DISABLE 1 2
DISABLE 2 1
DISABLE 3 3
GETMAX
RESET 1
RESET 2
DISABLE 1 2
DISABLE 1 3
DISABLE 2 2
GETMAX
RESET 3
GETMIN`)
	s2 := bytes.Buffer{}
	s2.WriteString(`2 3 9
DISABLE 1 1
DISABLE 2 2
RESET 2
DISABLE 2 1
DISABLE 2 3
RESET 1
GETMAX
DISABLE 2 1
GETMIN`)
	type args struct {
		n       int
		m       int
		q       int
		scanner *bufio.Scanner
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{
			name:    "Пример 1",
			args:    args{scanner: bufio.NewScanner(&s)},
			wantRes: []int{1, 2, 1},
		},
		{
			name:    "Пример 2",
			args:    args{scanner: bufio.NewScanner(&s2)},
			wantRes: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Kolya(tt.args.scanner); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Kolya() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
