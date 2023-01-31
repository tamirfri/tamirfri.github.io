package lodash_test

import (
	"testing"

	lodash "github.com/tamirfri/py-run-go-run-js"
)

func TestFind(t *testing.T) {
	testCases := []struct {
		desc           string
		objects        string
		partialObject  string
		expectedResult string
	}{{
		desc:           "find second result",
		objects:        `[{"a":1, "b":2, "d":true}, {"b":2, "c":3, "d":false}, {"a":1, "d":true}]`,
		partialObject:  `{"b": 2, "c": 3}`,
		expectedResult: `{"b":2,"c":3,"d":false}`,
	}}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := lodash.Find(tC.objects, tC.partialObject)
			if err != nil {
				t.Fatal(err)
			}
			if got != tC.expectedResult {
				t.Fatalf("expected: %v, got: %v", tC.expectedResult, got)
			}
		})
	}
}
