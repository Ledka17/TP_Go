package main

import "testing"

type testflags struct {
	ignoreCase *bool
	uniqueSort *bool
	reverseSort *bool
	numberSort *bool
	columnSort *int
	fileOutput *string
}

type testPairData struct {
	got []string
	expected []string
}
var TestFlags testflags

var testDataSort = []testPairData {
	{[]string {"Mike", "John", "Irma", "Vicky"} , []string {"Irma", "John", "Mike", "Vicky"} },
}

func TestMySort(t *testing.T) {
	*TestFlags.ignoreCase = false
	*TestFlags.columnSort = 1
	*TestFlags.numberSort = false
	*TestFlags.reverseSort = false
	*TestFlags.fileOutput = "stdout"

	for _, pair := range testDataSort {
		v := MySort(pair.got)
		for i, _ := range v {
			if v[i] != pair.expected[i] {
				t.Error(
					"For", pair.got,
					"expected", pair.expected,
					"got", v,
				)
				continue
			}
		}
	}
}