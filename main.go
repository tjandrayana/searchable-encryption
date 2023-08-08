// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"github.com/tjandrayana/searchable-encryption/transform"
)

type TestCase struct {
	Identifier string
	PlainText  string
	Type       string
}

func main() {

	testcases := []TestCase{
		{
			Identifier: "Test Case 1",
			PlainText:  "karet pedurenan 81",
			Type:       "Address",
		},
		{
			Identifier: "Test Case 2",
			PlainText:  "pedurenan 81",
			Type:       "Address",
		},
		{
			Identifier: "Test Case 3",
			PlainText:  "6281234566789",
			Type:       "Phone Number",
		},
		{
			Identifier: "Test Case 4",
			PlainText:  "6281234566788",
			Type:       "Phone Number",
		},
		{
			Identifier: "Test Case 5",
			PlainText:  "Jl. Jend. Sudirman Kav. 44-46, Jakarta 10210",
			Type:       "Address",
		},
		{
			Identifier: "Test Case 6",
			PlainText:  "Jl. Jend. Sudirman Kav. 1, Jakarta 10220",
			Type:       "Address",
		},
		{
			Identifier: "Test Case 7",
			PlainText:  "Sudirman",
			Type:       "Address",
		},
	}

	for _, tc := range testcases {
		fmt.Printf("%s\nPlainText:%s\nCipher Text:%s\nType:%s\n==========\n\n", tc.Identifier, tc.PlainText, transform.Transform(tc.PlainText), tc.Type)
	}
}
