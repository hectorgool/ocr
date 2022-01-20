/*
	go test -v
	t.Logf("%d", i)
	go mod tidy
	https://dev.to/quii/learn-go-by-writing-tests-arrays-and-slices-ahm
*/
package utils

import (
	"ocr/schema"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDisplay(t *testing.T) {

	t.Parallel()

	t.Run("test unknown character ?", func(t *testing.T) {

		// define inputs
		input := schema.Displayx

		// define expected result
		expected := "?"

		// perform test
		actual := ValidateDisplay(input)

		// assert that the actual result is equal to expected
		assert.Equal(t, expected, actual)
	})

	t.Run("test display 0", func(t *testing.T) {

		// define inputs
		input := schema.Display0

		// define expected result
		expected := "0"

		// perform test
		actual := ValidateDisplay(input)

		// assert that the actual result is equal to expected
		assert.Equal(t, expected, actual)
	})

	t.Run("test display 1", func(t *testing.T) {

		// define inputs
		input := schema.Display1

		// define expected result
		expected := "1"

		// perform test
		actual := ValidateDisplay(input)

		// assert that the actual result is equal to expected
		assert.Equal(t, expected, actual)
	})

	t.Run("test display 2", func(t *testing.T) {

		// define inputs
		input := schema.Display2

		// define expected result
		expected := "2"

		// perform test
		actual := ValidateDisplay(input)

		// assert that the actual result is equal to expected
		assert.Equal(t, expected, actual)
	})

	t.Run("test display 3", func(t *testing.T) {

		// define inputs
		input := schema.Display3

		// define expected result
		expected := "3"

		// perform test
		actual := ValidateDisplay(input)

		// assert that the actual result is equal to expected
		assert.Equal(t, expected, actual)
	})

	t.Run("test display 4", func(t *testing.T) {

		// define inputs
		input := schema.Display4

		// define expected result
		expected := "4"

		// perform test
		actual := ValidateDisplay(input)

		// assert that the actual result is equal to expected
		assert.Equal(t, expected, actual)
	})

	t.Run("test display 5", func(t *testing.T) {

		// define inputs
		input := schema.Display5

		// define expected result
		expected := "5"

		// perform test
		actual := ValidateDisplay(input)

		// assert that the actual result is equal to expected
		assert.Equal(t, expected, actual)
	})

	t.Run("test display 6", func(t *testing.T) {

		// define inputs
		input := schema.Display6

		// define expected result
		expected := "6"

		// perform test
		actual := ValidateDisplay(input)

		// assert that the actual result is equal to expected
		assert.Equal(t, expected, actual)
	})

	t.Run("test display 7", func(t *testing.T) {

		// define inputs
		input := schema.Display7

		// define expected result
		expected := "7"

		// perform test
		actual := ValidateDisplay(input)

		// assert that the actual result is equal to expected
		assert.Equal(t, expected, actual)
	})

	t.Run("test display 8", func(t *testing.T) {

		// define inputs
		input := schema.Display8

		// define expected result
		expected := "8"

		// perform test
		actual := ValidateDisplay(input)

		// assert that the actual result is equal to expected
		assert.Equal(t, expected, actual)
	})

	t.Run("test display 9", func(t *testing.T) {

		// define inputs
		input := schema.Display9

		// define expected result
		expected := "9"

		// perform test
		actual := ValidateDisplay(input)

		// assert that the actual result is equal to expected
		assert.Equal(t, expected, actual)
	})

}
