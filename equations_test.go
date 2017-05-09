package sprocket

import "testing"
import "fmt"

var EPSILON float64 = 0.000001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

type centerToCenterTest struct {
	lengthInPitches float64
	pitch           float64
	toothCount1     int
	toothCount2     int
	expected        float64
}

type chainLengthTest struct {
	distance    float64
	pitch       float64
	toothCount1 int
	toothCount2 int
	expected    float64
}

type practicalTest struct {
	length   float64
	expected float64
}

var centerToCenterTests = []centerToCenterTest{
	centerToCenterTest{48.0, 0.25, 20, 15, 3.807302},
	centerToCenterTest{48.0, 0.25, 15, 20, 3.807302},
	centerToCenterTest{100.0, 0.25, 40, 15, 9.007576},
	//centerToCenterTest{60.0, 0.3, 16, 6, 7.348},
}

var chainLengthTests = []chainLengthTest{
	chainLengthTest{3.371, 0.25, 20, 15, 44.514964},
	chainLengthTest{3.371, 0.25, 15, 20, 44.514964},
	chainLengthTest{3.371, 0.25, 15, 25, 47.15585446388746},
}

var practicalTests = []practicalTest{
	practicalTest{0.001, 2.0},
	practicalTest{0.9, 2.0},
	practicalTest{1.0, 2.0},
	practicalTest{1.1, 2.0},
	practicalTest{2.0, 2.0},
	practicalTest{20.001, 22.0},
	practicalTest{20.9, 22.0},
	practicalTest{21.0, 22.0},
	practicalTest{21.1, 22.0},
	practicalTest{22.0, 22.0},
}

func TestCenterToCenterDist(t *testing.T) {
	for i, test := range centerToCenterTests {
		dist := CalcCenterToCenterDist(test.pitch, test.lengthInPitches, test.toothCount1, test.toothCount2)

		if !floatEquals(test.expected, dist) {
			t.Errorf("CalcCenterToCenterDist %d incorrect, expected %f, got %f", i, test.expected, dist)
		}
	}
}

func TestChainLength(t *testing.T) {
	for i, test := range chainLengthTests {
		length := CalcChainLengthInPitches(test.distance, test.pitch, test.toothCount1, test.toothCount2)

		if !floatEquals(test.expected, length) {
			t.Errorf("CalcChainLength %d incorrect, expected %f, got %f", i, test.expected, length)
		}
	}
}

func TestPracticalChainLength(t *testing.T) {
	for i, test := range practicalTests {
		length := NearestPracticalChainLengthInPitches(test.length)
		if length != test.expected {
			t.Errorf("Incorrect length for test %d,  got %f, expected %f", i, length, test.expected)
		}
	}
}

func TestSprocketSolver(t *testing.T) {
	result := SolveSprocketSizeForCenterDist(3.5, 0.25, 11, 18, 36)
	fmt.Printf("Solved %d, %f\n", result.SprocketSize[0], result.LengthError)

	result = SolveSprocketSizeForCenterDist(3.5, 0.25, 9, 18, 36)
	fmt.Printf("Solved %d, %f\n", result.SprocketSize[0], result.LengthError)

	result = SolveBothSprocketSizesForCenterDist(3.5, 0.25, 11, 18, 18, 36)
	fmt.Printf("Solved %d, %d, %f\n", result.SprocketSize[0], result.SprocketSize[1], result.LengthError)
}
