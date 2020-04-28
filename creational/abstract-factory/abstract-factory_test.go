package abstractfactory

import "testing"

// go test -v ./creational/abstract-factory/...

func TestElectronicsFactory(t *testing.T) {
	var factory ElectronicsFactory
	var cpu CPU
	var disk Disk

	factory = &Intel{}
	cpu = factory.CreateCPU()
	cpu.Calc()
	disk = factory.CreateDisk()
	disk.Save()

	factory = &IBM{}
	cpu = factory.CreateCPU()
	cpu.Calc()
	disk = factory.CreateDisk()
	disk.Save()
}
