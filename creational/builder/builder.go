package builder

import "fmt"

// 还是拿 abstract-factory 中的例子来说
// 现在我们模拟组装电脑的情形

// 为了子一个文件中直观的看，这里不放到不同的文件

////////////////////////////////////////////////////////////////////////////////
// computer-director.go

// NewComputerDirector 创建 ComputerDirector
func NewComputerDirector(builder ComputerBuilder) *ComputerDirector {
	return &ComputerDirector{builder: builder}
}

// ComputerDirector 指挥者/导演
type ComputerDirector struct {
	builder ComputerBuilder
}

// Construct 组装 Computer
func (c *ComputerDirector) Construct() Computer {
	fmt.Println("ComputerDirector:Construct")
	c.builder.BuildCPU()
	c.builder.BuildDisk()
	return c.builder.GetComputer()
}

////////////////////////////////////////////////////////////////////////////////
// computer.go

// Computer 产品
type Computer struct {
	CPU  CPU  // CPU 部件
	Disk Disk // Disk 部件
}

////////////////////////////////////////////////////////////////////////////////
// computer-builder.go

// ComputerBuilder 抽象建造者
type ComputerBuilder interface {
	BuildCPU()             // 构建 CPU 部件
	BuildDisk()            // 构建 Disk 部件
	GetComputer() Computer // 获取最终的 Computer
}

// BaseComputerBuilder 基础的抽象建造者，包含 Computer 字段，并提供 GetComputer 方法
type BaseComputerBuilder struct {
	Computer
}

func (builder *BaseComputerBuilder) GetComputer() Computer {
	fmt.Println("BaseComputerBuilder:GetComputer")
	return builder.Computer
}

////////////////////////////////////////////////////////////////////////////////
// cpu.go

// CPU 部件
type CPU interface {
	Calc() // 计算
}

////////////////////////////////////////////////////////////////////////////////
// disk.go

// Disk 部件
type Disk interface {
	Save() // 存储
}

////////////////////////////////////////////////////////////////////////////////
// intel.go

// Intel 具体建造者
type Intel struct {
	BaseComputerBuilder
}

func (builder *Intel) BuildCPU() {
	fmt.Println("Intel:BuildCPU")
	builder.Computer.CPU = &IntelCPU{}
}

func (builder *Intel) BuildDisk() {
	fmt.Println("Intel:BuildDisk")
	builder.Computer.Disk = &IntelDisk{}
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

// IBM 具体建造者
type IBM struct {
	BaseComputerBuilder
}

func (builder *IBM) BuildCPU() {
	fmt.Println("IBM:BuildCPU")
	builder.Computer.CPU = &IBMCPU{}
}

func (builder *IBM) BuildDisk() {
	fmt.Println("IBM:BuildDisk")
	builder.Computer.Disk = &IBMDisk{}
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
