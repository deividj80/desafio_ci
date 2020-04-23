package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(5, 5)

	if result == 10 {
		t.Logf("sum(5,5) sucesso")
	} else {
		t.Errorf("sum(5,5) erro")
	}

}
