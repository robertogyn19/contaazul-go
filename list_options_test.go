package contaazul

import (
	"fmt"
	"net/url"
	"testing"
	"time"
)

type listOptionsTest1 struct {
	Enabled bool      `structs:"enabled"`
	Name    string    `structs:"name"`
	Page    int       `structs:"page"`
	Start   time.Time `structs:"start"`
}

type listOptionsTest2 struct {
	Enabled bool      `structs:"enabled,omitempty"`
	Name    string    `structs:"name,omitempty"`
	Page    int       `structs:"page,omitempty"`
	Start   time.Time `structs:"start,omitempty"`
}

func TestAddOptions(t *testing.T) {
	baseUrlTest := "https://domain.com/v1/test"
	urlFull := fmt.Sprintf("%s?enabled=true&name=a&page=1&start=2020-08-12T09:45:21Z", baseUrlTest)

	tests := []struct {
		lot      interface{}
		expected string
	}{
		{
			lot:      listOptionsTest1{},
			expected: fmt.Sprintf("%s?enabled=false&name=&page=0", baseUrlTest),
		},
		{
			lot: listOptionsTest1{
				Enabled: true,
				Name:    "a",
				Page:    1,
				Start:   time.Date(2020, 8, 12, 9, 45, 21, 0, time.UTC),
			},
			expected: urlFull,
		},
		{
			lot:      listOptionsTest2{},
			expected: fmt.Sprintf("%s", baseUrlTest),
		},
		{
			lot: listOptionsTest2{
				Enabled: true,
				Name:    "a",
				Page:    1,
				Start:   time.Date(2020, 8, 12, 9, 45, 21, 0, time.UTC),
			},
			expected: urlFull,
		},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			actual := structToQueryValues(test.lot, baseUrlTest)
			var err error
			actual, err = url.QueryUnescape(actual)

			if err != nil {
				t.Fatal(err)
			}

			if actual != test.expected {
				t.Errorf("\n(actual)   %s\n(expected) %s", actual, test.expected)
			}
		})
	}
}
