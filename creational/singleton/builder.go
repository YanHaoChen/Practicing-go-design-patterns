package singleton

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {}


func (f *ManufacturingDirector) Construct() {

}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {

}


type VehiceProduct struct {
	Wheels int
	Seats int
	Structiure string
}

type CarBuilder struct {}

