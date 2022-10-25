

// main_test.go

package main

import "testing"

func TestSum(t *testing.T) {
    if add(2, 5) != 7 {
        t.Fail()
    }
    if add(2, 100) != 102 {
        t.Fail()
    }
    if add(222, 100) != 322 {
        t.Fail()
    }
}

func TestProduct(t *testing.T) {
    if multiply(2, 5) != 10 {
        t.Fail()
    }
    if multiply(2, 100) != 200 {
        t.Fail()
    }
    if multiply(222, 3) != 666 {
        t.Fail()
    }
}

func TestMagic(t *testing.T) {
	if magic(1, 6) != 5 {
		t.Fail()
	}
	if magic(2, 5) != 3 {
		t.Fail()
	}
	if magic(7, 4) != 32 {
		t.Fail()
	}
	if magic(3, 2) != 12 {
		t.Fail()
	}
}
