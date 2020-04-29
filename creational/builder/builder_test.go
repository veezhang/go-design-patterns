package builder

import "testing"

// go test -v ./creational/builder/...

func TestElectronicsFactory(t *testing.T) {
	var builder ComputerBuilder
	var director *ComputerDirector
	var computer Computer

	builder = &Intel{}
	director = NewComputerDirector(builder)
	computer = director.Construct()
	computer.CPU.Calc()
	computer.Disk.Save()

	builder = &IBM{}
	director = NewComputerDirector(builder)
	computer = director.Construct()
	computer.CPU.Calc()
	computer.Disk.Save()
}
