package services

import "testing"

func TestFindMinimuRecommendationmSteps(t *testing.T) {
	service := &APIService{}

	testDataSet := []struct {
		input    string
		expected int
	}{
		{input: "aA1", expected: 3},
		{input: "1445D1cd", expected: 0},
		{input: "123456", expected: 2},
		{input: "111222", expected: 2},
		{input: "1112345", expected: 2},
		{input: "12345678901234567890", expected: 3},
		{input: "1234567890123456789", expected: 2},
		{input: "abc4567890123456777", expected: 1},
		{input: "A1111", expected: 1},
		{input: "A111", expected: 2},
		{input: "1111111111111111111111111", expected: 12},
	}

	for _, data := range testDataSet {
		actual := service.FindMinimuRecommendationmSteps(data.input)
		if actual != data.expected {
			t.Errorf("Input: %s, Expected: %d, Got: %d", data.input, data.expected, actual)
		}
	}
}
