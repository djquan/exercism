package sublist

//Relation can be equal, unequal, sublist, or superlist
type Relation string

//Sublist takes two lists and returns their relationship
func Sublist(listOne, listTwo []int) Relation {
	if isEqual(listOne, listTwo) {
		return "equal"
	}

	if isSublist(listOne, listTwo) {
		return "sublist"
	}

	if isSublist(listTwo, listOne) {
		return "superlist"
	}

	return "unequal"
}

func isSublist(listOne, listTwo []int) bool {
	lenListOne, lenListTwo := len(listOne), len(listTwo)

	for i := 0; i <= lenListTwo-lenListOne; i++ {
		if isEqual(listOne, listTwo[i:i+lenListOne]) {
			return true
		}

	}
	return false
}

func isEqual(listOne, listTwo []int) bool {
	if len(listOne) != len(listTwo) {
		return false
	}

	for i, v := range listOne {
		if v != listTwo[i] {
			return false
		}
	}

	return true
}
