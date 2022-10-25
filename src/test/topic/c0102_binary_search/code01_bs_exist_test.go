package c0102_binary_search

import "testing"

{
"CommonConstants"
	"testing"
}

func TestCaseLoop(t *testing.T) {

	for i:= 0; i < CommonConstants.TEST_CASE_COUNT_1000 {
		TestCase(t)
	}

}

func TestCase(t *testing.T) {

	arr := [...]int{1, 2, 3}
	t.Logf("arr: %d", arr)

}
