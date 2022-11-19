package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBabyGo(t *testing.T) {
	if BabyGo("Rizki") != "Hello rizki" {

		t.Fatal("Expected Hello Rizki") // to indicate test failed but will not continue and print Log failed reason

		//t.FailNow() // to indicate that the test failed but will stop the test immediately
		//t.Fail() // to indicate test failed but will continue the code
		//t.Error("Expected Hello Rizki") // to indicate test failed and will continue the code and print Log
		//panic("Expected Hello Rizki") // Don't use panic in Unit Test
	}

	fmt.Println("Test BabyGo Success")
}

func TestBabyGoGo(t *testing.T) {
	if BabyGo("Rizki") != "Hello rizki" {
	}

	fmt.Println("Test BabyGo Success")
}

func TestBabyGoGoAssert(t *testing.T) {
	result := BabyGo("Rizki")
	assert.Equal(t, "Hello Rizki", result, "Result should be Hello Rizki")
}

func TestTableBabyGoGo(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Test 1",
			request:  "Rizki",
			expected: "Hello Rizki",
		},
		{
			name:     "Test 2",
			request:  "Fadiah",
			expected: "Hello Fadiah",
		},
		{
			name:     "Test 3",
			request:  "Fatur",
			expected: "Hello Fatur",
		},
		{
			name:     "Test 4",
			request:  "Agus",
			expected: "Hello Agus",
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := BabyGo(tc.request)
			assert.Equal(t, tc.expected, result)
		})
	}
}
