package main

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	fastAttack := true
	if knightIsAwake == true {
		fastAttack = false
	} else {
		fastAttack = true
	}
	return fastAttack
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	spy := false
	if knightIsAwake == true || archerIsAwake == true || prisonerIsAwake == true {
		spy = true
	} else {
		spy = false
	}
	return spy
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	signalPrisoner := true
	if archerIsAwake == false && prisonerIsAwake == true {
		signalPrisoner = true
	} else {
		signalPrisoner = false
	}
	return signalPrisoner
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	switch {
	case knightIsAwake == false && archerIsAwake == false && prisonerIsAwake == true:
		return true
	case knightIsAwake == false && archerIsAwake == false && petDogIsPresent == true:
		return true
	case archerIsAwake == false && prisonerIsAwake == true && petDogIsPresent == true:
		return true
	case archerIsAwake == false && prisonerIsAwake == false && petDogIsPresent == true:
		return true
	default:
		return false
	}
}
