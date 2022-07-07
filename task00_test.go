package homework

import "testing"

func Test_AverageValueOfArray(t *testing.T) {
	var expected float32 = 8
	input := [15]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	res := average(input)
	if expected != res {
		t.Errorf("result is not correct expected value %f", expected)
	}
}

func Test_SliceReverse(t *testing.T) {
	expected := []int64{15, 5, 2, 1}
	input := []int64{1, 2, 5, 15}
	res := reverse(input)
	for i := 0; i < len(input); i++ {
		if res[i] != expected[i] {
			t.Errorf("result is not correct expected value %d", expected)
		}
	}
}

func Test_sortMapValues_SingleCharacter(t *testing.T) {
	expected := []string{"b", "c", "a"}
	input := map[int]string{2: "a", 0: "b", 1: "c"}
	res := sortMapValues(input)
	for i := 0; i < len(input); i++ {
		if res[i] != expected[i] {
			t.Errorf("result is not correct expected value %s", expected)
		}
	}
}

func Test_sortMapValues_DoubleCharacter(t *testing.T) {
	expected := []string{"bb", "aa", "cc"}
	input := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	res := sortMapValues(input)
	for i := 0; i < len(input); i++ {
		if res[i] != expected[i] {
			t.Errorf("result is not correct expected value %s", expected)
		}
	}
}
