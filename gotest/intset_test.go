package gotest

import (
	"awesomeProject/intset"
	"fmt"
	"testing"
)

func TestIntset(t *testing.T) {
	fmt.Println(intset.IntSize)
	var x, y intset.IntSet
	x.Add(1)
	x.Add(9)
	x.Add(111)
	//fmt.Println(x.String())
	if x.Len() != 3 {
		t.Error("Len func is wrong")
	}
	y.Add(8)
	y.Add(999)
	if y.Len() != 2 {
		t.Error("Len func is wrong")
	}
	tt:= x.UnionWith(&y)
	fmt.Println(tt.String())
	if !tt.Has(999) {
		t.Error("UnionWith or Has func is wrong")
	}
	tt.Remove(111)
	tt.Remove(999)
	if tt.Len() != 3 {
		t.Error("Remove func is wrong")
	}
	x.Clear()
	if x.Len() != 0 {
		t.Error("Clear func is wrong")
	}
	z := y.Copy()
	fmt.Println(z)
	if z.Len() != 2 || !z.Has(8) || !z.Has(999) {
		t.Error("Copy func is wrong")
	}
}

func TestIntsetAddAll(t *testing.T) {
	var x intset.IntSet
	x.Add(1)
	x.Add(9)
	x.Add(111)
	if x.AddAll(1, 2, 3) != 127 {
		t.Error("AddAll func is wrong")
	}
}

func TestIntersectWithAndDifferenceWith(t *testing.T) {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(9)
	x.Add(111)

	y.Add(8)
	y.Add(999)
	y.Add(9)
	y.Add(999999999)
	fmt.Println("@@@@@")
	z := x.IntersectWith(&y)
	fmt.Println("!!!!!!")
	fmt.Println(z)
	if z.Len() != 1 || !z.Has(9) {
		t.Error("IntersectWith func is wrong")
	}

	z = x.DifferenceWith(&y)
	fmt.Println(z)
	if z.Len() != 2 || !z.Has(1) || !z.Has(111) {
		t.Error("DifferenceWith func is wrong")
	}
}

func TestSymmertricDifference(t *testing.T) {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(9)
	x.Add(111)

	y.Add(8)
	y.Add(999)
	y.Add(9)
	y.Add(999999999)
	fmt.Println("@@@@@")
	z := x.SymmetricDifference(&y)
	fmt.Println("!!!!!!")
	fmt.Println(z)
	if z.Len() != 5 || !z.Has(1) {
		t.Error("SymmertricDifference func is wrong")
	}
	for _,v :=range z.Elems(){
		fmt.Println(v)
	}

}
