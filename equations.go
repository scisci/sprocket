package sprocket

import "math"
import "fmt"

// center dist 3.5"

// Calculates the distance between two sprocket centers given their properties
// and that of the chain.
func CalcCenterToCenterDist(pitch, lengthInPitches float64, toothCount1, toothCount2 int) float64 {
	largerSprocket := toothCount1
	smallerSprocket := toothCount2

	if largerSprocket < smallerSprocket {
		largerSprocket, smallerSprocket = smallerSprocket, largerSprocket
	}

	sprocketDif := float64(largerSprocket - smallerSprocket)
	sprocketDifSquared := sprocketDif * sprocketDif

	lengthRemainder := 2*lengthInPitches - float64(smallerSprocket+largerSprocket)
	lengthRemainderSquared := lengthRemainder * lengthRemainder

	c := lengthRemainderSquared - (8.0/(math.Pi*math.Pi))*sprocketDifSquared
	return (pitch / 8) * (lengthRemainder + math.Sqrt(c))
}

// Calculates the number of links of chain needed given the distance between
// two sprockets.
func CalcChainLengthInPitches(distance, pitch float64, toothCount1, toothCount2 int) float64 {
	largerSprocket := toothCount1
	smallerSprocket := toothCount2

	if largerSprocket < smallerSprocket {
		largerSprocket, smallerSprocket = smallerSprocket, largerSprocket
	}

	sprocketDif := float64(largerSprocket - smallerSprocket)
	sprocketDifOver2PI := sprocketDif / (2 * math.Pi)

	return (2*distance)/pitch + float64(smallerSprocket+largerSprocket)/2.0 +
		(pitch*sprocketDifOver2PI*sprocketDifOver2PI)/distance
}

func NearestPracticalChainLengthInPitches(exactLengthInPitches float64) float64 {
	// We round up to the nearest even number
	return math.Ceil(exactLengthInPitches*0.5) * 2.0
}

type SprocketSizeResult struct {
	SprocketSize []int
	LengthError  float64
}

// If you made a design error and can't change the center distance, this can
// help to find a sprocket size.
func SolveSprocketSizeForCenterDist(distance, pitch float64, otherToothCount, minToothCount, maxToothCount int) SprocketSizeResult {
	fmt.Printf("\nSolving 1 Sprockets for dist %f, sprocket %d (%d-%d)\n-------------\n",
		distance, otherToothCount, minToothCount, maxToothCount)

	if minToothCount > maxToothCount {
		minToothCount, maxToothCount = maxToothCount, minToothCount
	}

	result := SprocketSizeResult{
		[]int{-1},
		math.MaxFloat64,
	}

	for toothCount := minToothCount; toothCount <= maxToothCount; toothCount++ {
		length := CalcChainLengthInPitches(distance, pitch, toothCount, otherToothCount)
		nearestLength := NearestPracticalChainLengthInPitches(length)
		lengthError := nearestLength - length

		fmt.Printf("%d: %f\n", toothCount, lengthError)
		if lengthError < result.LengthError {
			result.SprocketSize[0] = toothCount
			result.LengthError = lengthError
		}
	}

	return result
}

// If you made a design error and can't change the center distance, this can
// help to find a sprocket size.
func SolveBothSprocketSizesForCenterDist(distance, pitch float64, minToothCount, maxToothCount, minOtherToothCount, maxOtherToothCount int) SprocketSizeResult {
	fmt.Printf("\nSolving 2 Sprockets for dist %f (%d-%d, %d-%d)\n-------------\n",
		distance, minToothCount, maxToothCount, minOtherToothCount, maxOtherToothCount)

	if minToothCount > maxToothCount {
		minToothCount, maxToothCount = maxToothCount, minToothCount
	}

	if minOtherToothCount > maxOtherToothCount {
		minOtherToothCount, maxOtherToothCount = maxOtherToothCount, minOtherToothCount
	}

	result := SprocketSizeResult{
		[]int{-1, -1},
		math.MaxFloat64,
	}

	for toothCount := minToothCount; toothCount <= maxToothCount; toothCount++ {
		otherResult := SolveSprocketSizeForCenterDist(distance, pitch, toothCount, minOtherToothCount, maxOtherToothCount)
		if otherResult.LengthError < result.LengthError {
			result.SprocketSize[0] = toothCount
			result.SprocketSize[1] = otherResult.SprocketSize[0]
			result.LengthError = otherResult.LengthError
		}
	}

	return result
}
