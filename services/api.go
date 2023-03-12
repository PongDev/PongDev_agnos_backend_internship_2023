package services

type APIService struct{}

func countMissingMixType(input string) int {
	missingUppercase := 1
	missingLowercase := 1
	missingDigit := 1

	for _, char := range input {
		switch {
		case char >= 'A' && char <= 'Z':
			missingUppercase = 0
		case char >= 'a' && char <= 'z':
			missingLowercase = 0
		case char >= '0' && char <= '9':
			missingDigit = 0
		}
	}
	return missingUppercase + missingLowercase + missingDigit
}

func (a *APIService) FindMinimuRecommendationmSteps(input string) int {
	requiredAddOrReplace := countMissingMixType(input)
	addOps := 0
	removeOps := 0
	replaceOps := 0
	outputLen := len(input)
	prevChar := rune(0)
	prevCount := 0

	for _, char := range input {
		if char == prevChar {
			prevCount++
		} else {
			prevChar = char
			prevCount = 1
		}

		if prevCount > 2 {
			switch {
			case outputLen < 6:
				addOps++
				if requiredAddOrReplace > 0 {
					requiredAddOrReplace--
				}
				outputLen++
				prevCount = 1
			case outputLen >= 20:
				removeOps++
				outputLen--
				prevCount -= 1
			default:
				replaceOps++
				if requiredAddOrReplace > 0 {
					requiredAddOrReplace--
				}
				prevCount = 1
				prevChar = 0
			}
		}
	}
	if outputLen < 6 {
		addOps += 6 - outputLen
	}
	if outputLen > 19 {
		removeOps += outputLen - 19
	}
	return addOps + removeOps + replaceOps + requiredAddOrReplace
}
