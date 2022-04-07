package json

import "testing"

type jsonTest struct {
	A int
}

func TestReadFile(t *testing.T) {
	root = "testData/" // temporarily change root path to testData
	tests := []struct {
		name         string
		pathSegments []string
		want         int
	}{
		{"root", []string{"a.json"}, 1},
		{"nested1", []string{"dir", "b.json"}, 2},
		{"nested2", []string{"dir", "c.json"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := jsonTest{}
			res, err := ReadFile[jsonTest](tt.pathSegments...)
			if err != nil {
				t.Errorf("ReadFile() error = %v", err)
			}
			if res.A != tt.want {
				t.Errorf("ReadFile() want:1, got:%v", s.A)
			}
		})
	}
}

func TestReadDir(t *testing.T) {
	root = "testData/" // temporarily change root path to testData
	tests := []struct {
		name         string
		pathSegments []string
		wantSum      int
	}{
		{"root", nil, 8},
		{"dir", []string{"dir"}, 5},
		{"nested", []string{"nested"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := ReadDir[jsonTest](tt.pathSegments...)
			if err != nil {
				t.Errorf("ReadFile() error = %v", err)
			}
			sum := 0
			for _, x := range res {
				sum += x.A
			}
			if sum != tt.wantSum {
				t.Errorf("ReadFile() want:%v, got:%v, all:%v", tt.wantSum, sum, res)
			}
		})
	}
}
