package main

import "github.com/MaxHalford/eaopt"

func CrossGenFloat64Slice(maleGen, femaleGen *eaopt.Float64Slice) {
	if shouldSwap() {
		SwapGenFloat64Slice(maleGen, femaleGen)
	}
}

func CrossGenIntSlice(maleGen, femaleGen *eaopt.IntSlice) {
	if shouldSwap() {
		SwapGenIntSlice(maleGen, femaleGen)
	}
}

func SwapGenFloat64Slice(maleGen, femaleGen *eaopt.Float64Slice) {
	maleValue := *maleGen
	*maleGen = *femaleGen
	*femaleGen = maleValue
}

func SwapGenIntSlice(maleGen, femaleGen *eaopt.IntSlice) {
	maleValue := *maleGen
	*maleGen = *femaleGen
	*femaleGen = maleValue
}

func MutateLittleFloat64(current float64, minMax MinMaxFloat64) float64 {
	if !shouldMutate() {
		return current
	}

	value := MutatePercentFloat(current)
	result := current - value

	if minMax.min > result || result > minMax.max {
		return current
	}

	return result
}

func MutateLittleInt(current int, minMax MinMaxInt) int {
	if !shouldMutate() {
		return current
	}

	value := MutatePercentInt(current)
	result := current - value

	if minMax.min > result || result > minMax.max {
		return current
	}

	return result
}

func MutatePercentFloat(current float64) float64 {
	dir := 1.0
	mutatePercent := GetRandFloat64(0, 100)
	mutateValue := (current * mutatePercent) / 100

	if GetRandInt(0, 1) == 1 {
		dir = -1.0
	}

	return dir * mutateValue
}

func MutatePercentInt(current int) int {
	dir := 1
	mutatePercent := int(GetRandInt(0, 100))
	mutateValue := int((current * mutatePercent) / 100)

	if GetRandInt(0, 1) == 1 {
		dir = -1
	}

	return dir * mutateValue
}

func shouldSwap() bool {
	return GetRandInt(0, 1) == 1
}

func shouldMutate() bool {
	return GetRandInt(0, 5) == 1
}
