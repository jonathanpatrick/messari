package aggregate

import (
	"fmt"
	"math"
	"testing"
)

func TestAggregate_AdjustValueForYesterday(t *testing.T) {
	var testName string
	testName = "valid case, simple example"
	t.Run(testName, func(t *testing.T) {
		expected := 100.0
		currentMarketCap := 110.0
		percentChange := 10.0
		actual := adjustValueForYesterday(currentMarketCap, percentChange)

		if math.Round(expected) != math.Round(actual) {
			fmt.Println(fmt.Sprintf("actual=%v did not match expected output=%v", actual, expected))
			t.FailNow()
		}
	})

	testName = "valid case, negative percent change"
	t.Run(testName, func(t *testing.T) {
		expected := 400.0
		currentMarketCap := 300.0
		percentChange := -25.0
		actual := adjustValueForYesterday(currentMarketCap, percentChange)

		if math.Round(expected) != math.Round(actual) {
			fmt.Println(fmt.Sprintf("actual=%v did not match expected output=%v", actual, expected))
			t.FailNow()
		}
	})
}
