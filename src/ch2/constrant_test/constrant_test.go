package constrant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstrantWeek(t *testing.T) {

	t.Log("Monday ", Monday)

	t.Log("Tuesday", Tuesday)

	t.Log("Wednesday", Wednesday)

	t.Log("Thursday", Thursday)

	t.Log("Friday", Friday)

	t.Log("Saturday", Saturday)

	t.Log("Sunday", Sunday)
}

func TestConstrantMov(t *testing.T) {

	var testValue int = 7
	//0111
	t.Log("Readable value is :", Readable)
	t.Log("Readable ", testValue&Readable == Readable)

	t.Log("Writable value is :", Writable)
	t.Log("Writable ", testValue&Writable == Writable)

	t.Log("Executable value is :", Executable)
	t.Log("Executable ", testValue&Executable == Executable)
}
