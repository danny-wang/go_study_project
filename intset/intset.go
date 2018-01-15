package intset

// Simple bit slice, every bit of each element represent a value of the set.
// When the x bit is set to 1, it means the set includes element x.
import (
	"bytes"
	"fmt"
)

// int 或 uint 类型的长度(32 或 64)
const IntSize = intSize
const intSize = 32 << uint(^uint(0)>>63)

// An IntSet is a set of small non-negative integers.
// It's zero value represents the zero set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/intSize, uint(x%intSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add add the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/intSize, uint(x%intSize)
	//fmt.Println(word,bit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit)
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) *IntSet {
	newSet := s.Copy()
	for i, value := range t.words {
		if i < len(newSet.words) {
			newSet.words[i] |= value
		} else {
			newSet.words = append(newSet.words, value)
		}
	}
	return newSet
}

// IntersectWith returns a IntSet contains elements that s and t both have
func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	newSet := &IntSet{}
	length := len(s.words)
	if len(s.words) >= len(t.words) {
		length = len(t.words)
	}
	for i := 0; i < length; i++ {
		newSet.words = append(newSet.words, s.words[i]&t.words[i])
	}
	return newSet
}

// DifferenceWith returns a IntSet contains elements that is inside s and is not inside t
func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	newSet := s.Copy()
	for i, value := range newSet.words {
		if i < len(t.words) {
			newSet.words[i] &= ^(value & t.words[i])
		} else {
			break
		}
	}
	return newSet
}

// SymmetricDifference returns a IntSet contains elements that are inside s and not inside t
// or that are inside t and not inside s
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	newSet1 := s.DifferenceWith(t)
	newSet2 := t.DifferenceWith(s)
	newSet3 := newSet1.UnionWith(newSet2)
	return newSet3
}

// Elems returns all elements in the set
func (s *IntSet) Elems() []int {
	var res []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, intSize*i+j)
			}
		}
	}
	return res
}

// String returns the set as a string of the form "{1,2,3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				//fmt.Println(buf.Len())
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", intSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// return the number of elements
func (s *IntSet) Len() int {
	length := 0
	for _, v := range s.words {
		if v == 0 {
			continue
		}
		count:=0
		for v !=0 {
			count++
			// a&(a-1)会把最右边的一个1变成0
			v = v&(v-1)
		}
		length += count
		//var sig uint = 1
		//for i := 0; i < intSize; i++ {
		//	if v&sig != 0 {
		//		length++
		//	}
		//	sig = sig << 1
		//}
	}
	return length
}

// remote x form the set
func (s *IntSet) Remove(x int) {
	word, bit := x/intSize, uint(x%intSize)
	s.words[word] &= ^(1 << bit)
}

// remove all element from the set
func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

// return a copy of set
func (s *IntSet) Copy() *IntSet {
	newIntSet := &IntSet{}
	for _, v := range s.words {
		newIntSet.words = append(newIntSet.words, v)
	}
	return newIntSet
}

// return the sum of IntSet and a lot of number
func (s *IntSet) AddAll(nums ...int) int {
	sum := 0
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				sum += intSize*i + j
			}
		}
	}
	for _, v := range nums {
		sum += v
	}
	return sum
}
