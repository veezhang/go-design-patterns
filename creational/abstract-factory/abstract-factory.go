package abstractfactory

import "fmt"

// simple-factory 和 factory-method 中的例子都是一种类型的产品
//
// Abstract Factory 抽象工厂模式 是用于创建多个产品线的情况
// 我们这里以电子设备公司为例子
// 服装厂可以生产CPU也可以生产磁盘
//

// 为了子一个文件中直观的看，这里不放到不同的文件

////////////////////////////////////////////////////////////////////////////////
// electronics-factory.go

// ElectronicsFactory 是一个电子设备公司
type ElectronicsFactory interface {
	CreateCPU() CPU
	CreateDisk() Disk
}

////////////////////////////////////////////////////////////////////////////////
// cpu.go

// CPU
type CPU interface {
	Calc() // 计算
}

////////////////////////////////////////////////////////////////////////////////
// disk.go

// Disk
type Disk interface {
	Save() // 存储
}

////////////////////////////////////////////////////////////////////////////////
// intel.go

// Intel 电子设备公司
type Intel struct {
}

func (factory *Intel) CreateCPU() CPU {
	fmt.Println("Intel:CreateCPU")
	return &IntelCPU{}
}

func (factory *Intel) CreateDisk() Disk {
	fmt.Println("Intel:CreateDisk")
	return &IntelDisk{}
}

////////////////////////////////////////////////////////////////////////////////
// intel-cpu.go

type IntelCPU struct {
}

func (cpu *IntelCPU) Calc() {
	fmt.Println("IntelCPU:Calc")
}

////////////////////////////////////////////////////////////////////////////////
// intel-disk.go

type IntelDisk struct {
}

func (disk *IntelDisk) Save() {
	fmt.Println("IntelDisk:Save")
}

////////////////////////////////////////////////////////////////////////////////
// ibm.go

// IBM 电子设备公司
type IBM struct {
}

func (factory *IBM) CreateCPU() CPU {
	fmt.Println("IBM:CreateCPU")
	return &IBMCPU{}
}

func (factory *IBM) CreateDisk() Disk {
	fmt.Println("IBM:CreateDisk")
	return &IBMDisk{}
}

////////////////////////////////////////////////////////////////////////////////
// ibm-cpu.go

type IBMCPU struct {
}

func (cpu *IBMCPU) Calc() {
	fmt.Println("IBMCPU:Calc")
}

////////////////////////////////////////////////////////////////////////////////
// ibm-disk.go

type IBMDisk struct {
}

func (disk *IBMDisk) Save() {
	fmt.Println("IBMDisk:Save")
}
