package fabonacci_test

import (
	"testing"
)

func TestFabonacci(t *testing.T) {

	var (
		a int = 1
		b int = 1
	)

	t.Log("  ", a)
	for index := 0; index < 5; index++ {

		t.Log("  ", b)
		temp := a
		a = b
		b = temp + a
		//t.Log("  ", a)

	}
}

func TestExchange(t *testing.T) {

	var (
		a int = 1
		b int = 2
	)

	t.Log("a = ", a)
	t.Log("b = ", b)

	a, b = b, a

	t.Log("a = ", a)
	t.Log("b = ", b)
}
