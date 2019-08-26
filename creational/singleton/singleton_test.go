package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("should not be nil!")
	}
	expectedCounter := counter1

	currentCount := counter1.AddOne()

	if currentCount != 1 {
		t.Errorf("the count must be 1 but it is %d\n", currentCount)
	}
	
	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Error("They are not the same!")
	}

	currentCount = counter2.AddOne()
	
	if currentCount != 2 {
		t.Errorf("the count must be 2 but it is %d\n", currentCount)
	}
	
}
