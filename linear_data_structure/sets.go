// Package linear_data_structure
// Time    : 2021/5/11 9:22 上午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package linear_data_structure

// Set class
type Set struct {
	integerMap map[int]bool
}

// New to create the map of integer and bool
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// ContainsElement judge the element whether in set
func (set *Set) ContainsElement(element int) bool {
	_, ok := set.integerMap[element]
	return ok
}

// AddElement adds the element to the set
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

// DeleteElement deletes the element from the set
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

// Intersect method returns the set which intersects with anotherSet
func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()
	var value int
	for value, _ = range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

// Union method returns the set which is union of the set with anotherSet
func (set *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()
	var value int
	for value, _ = range set.integerMap {
		unionSet.AddElement(value)
	}
	for value, _ = range anotherSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}

// Diff method returns the set different part to the anotherSet

// IsSubSet method returns the anotherSet whether is subset of the set
