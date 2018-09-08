package listops

// IntList provides a type for dealing with an []int
type IntList []int
type predFunc func(int) bool
type unaryFunc func(int) int
type binFunc func(int, int) int

// Foldr applies a function to an IntList starting from the right
func (l IntList) Foldr(fn binFunc, initial int) int {
	for i := l.Length() - 1; i >= 0; i-- {
		initial = fn(l[i], initial)
	}
	return initial
}

// Foldl applies a function to an IntList starting from the left
func (l IntList) Foldl(fn binFunc, initial int) int {
	for _, v := range l {
		initial = fn(initial, v)
	}
	return initial
}

// Filter returns an IntList which only contains values that the match the predicate func
func (l IntList) Filter(fn predFunc) IntList {
	filtered := IntList{}
	for _, v := range l {
		if fn(v) {
			filtered = append(filtered, v)
		}

	}
	return filtered
}

// Length returns the length of an IntList
func (l IntList) Length() int {
	return len(l)
}

// Map returns an IntList with the input function applied to every value
func (l IntList) Map(fn unaryFunc) IntList {
	mapped := make(IntList, l.Length())
	for i, v := range l {
		mapped[i] = fn(v)
	}
	return mapped
}

// Reverse returns a reversed Intlist
func (l IntList) Reverse() IntList {
	length := l.Length()
	reversed := make(IntList, length)
	for i, v := range l {
		reversed[length-i-1] = v
	}

	return reversed
}

// Append returns a new IntList with the original and input values
func (l IntList) Append(list IntList) IntList {
	appended := make(IntList, l.Length()+list.Length())
	copy(appended, l)
	copy(appended[l.Length():], list)
	return appended
}

// Concat combines the IntList with a list of Intlists
func (l IntList) Concat(lists []IntList) IntList {
	c := make(IntList, l.Length())
	copy(c, l)

	for _, v := range lists {
		c = c.Append(v)
	}
	return c
}
