package vehicle

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()


	if car.Wheels != 4 {
		t.Errorf("Does your car has %d wheel(s)??\n", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("It is a(n) %s!\n", car.Structure)
	}
	if car.Seats != 4 {
		t.Errorf("Does your car has %d seat(s)??\n", car.Seats)
	}
}
