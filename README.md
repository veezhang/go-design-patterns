# Go Design Patterns

用 golang 实现设计模式

## 设计模式的六大原则

### 开闭原则（Open Closed Principle，OCP）

软件实体应当对扩展开放，对修改关闭（Software entities should be open for extension，but closed for modification）。当应用的需求改变时，在不修改软件实体的源代码或者二进制代码的前提下，可以扩展模块的功能，使其满足新的需求。

### 里氏替换原则（Liskov Substitution Principle，LSP）

继承必须确保超类所拥有的性质在子类中仍然成立（Inheritance should ensure that any property proved about supertype objects also holds for subtype objects）。子类可以扩展父类的功能，但不能改变父类原有的功能。

### 依赖倒置原则（Dependence Inversion Principle，DIP）

高层模块不应该依赖低层模块，两者都应该依赖其抽象；抽象不应该依赖细节，细节应该依赖抽象（High level modules shouldnot depend upon low level modules.Both should depend upon abstractions.Abstractions should not depend upon details. Details should depend upon abstractions）。其核心思想是：要面向接口编程，不要面向实现编程。

### 单一职责原则（Single Responsibility Principle，SRP）

单一职责原则规定一个类应该有且仅有一个引起它变化的原因，否则类应该被拆分（There should never be more than one reason for a class to change）。不要存在多于一个导致类变更的原因，也就是说每个类应该实现单一的职责，如若不然，就应该把类拆分。

### 接口隔离原则（Interface Segregation Principle，ISP）

客户端不应该被迫依赖于它不使用的方法（Clients should not be forced to depend on methods they do not use）。一个类对另一个类的依赖应该建立在最小的接口上（The dependency of one class to another one should depend on the smallest possible interface）。要为各个类建立它们需要的专用接口，而不要试图去建立一个很庞大的接口供所有依赖它的类去调用。

### 迪米特法则（Law of Demeter，LoD）

又叫作最少知识原则（Least Knowledge Principle，LKP)。迪米特法则的定义是：只与你的直接朋友交谈，不跟“陌生人”说话（Talk only to your immediate friends and not to strangers）。其含义是：如果两个软件实体无须直接通信，那么就不应当发生直接的相互调用，可以通过第三方转发该调用。其目的是降低类之间的耦合度，提高模块的相对独立性。

### 合成复用原则（Composite Reuse Principle，CRP）

要尽量先使用组合或者聚合等关联关系来实现，其次才考虑使用继承关系来实现。

## 设计模式

### [创建型模式 (Creational Patterns)](/creational/)

创建型模式的主要关注点是“怎样创建对象？”，它的主要特点是“将对象的创建与使用分离”。这样可以降低系统的耦合度，使用者不需要关注对象的创建细节，对象的创建由相关的工厂来完成。

| 模式 | 状态 | 描述 |
|:----|:---: |:------|
| [Singleton 单例模式](/creational/singleton/) | ✔ | 某个类只能生成一个实例，该类提供了一个全局访问点供外部获取该实例。 |
| [Prototype 原型模式](/creational/prototype/) | ✔ | 将一个对象作为原型，通过对其进行复制而克隆出多个和原型类似的新实例。 |
| [Simple Factory 简单工厂模式](/creational/simple-factory/) | ✔ | 定义一个用于创建产品的接口，根据参数决定生产什么产品。|
| [Factory Method 工厂方法模式](/creational/factory-method/) | ✔ | 定义一个用于创建产品的接口，由子类决定生产什么产品。 |
| [Abstract Factory 抽象工厂模式](/creational/abstract-factory/) | ✔ | 提供一个创建产品族的接口，其每个子类可以生产一系列相关的产品。 |
| [Builder 建造者模式](/creational/builder/) | ✔ | 将一个复杂对象分解成多个相对简单的部分，然后根据不同需要分别创建它们，最后构建成该复杂对象。 |

### [结构型模式 (Structural Patterns)](/structural/)

结构型模式描述如何将类或对象按某种布局组成更大的结构。它分为类结构型模式和对象结构型模式，前者采用继承机制来组织接口和类，后者釆用组合或聚合来组合对象。

| 模式 | 状态 | 描述 |
|:----|:---: |:------|
| [Proxy 代理模式](/structural/proxy/) | ✘ | 为某对象提供一种代理以控制对该对象的访问。即客户端通过代理间接地访问该对象，从而限制、增强或修改该对象的一些特性。 |
| [Adapter 适配器模式](/structural/adapter/) | ✘ | 将一个类的接口转换成客户希望的另外一个接口，使得原本由于接口不兼容而不能一起工作的那些类能一起工作。 |
| [Bridge 桥接模式](/structural/bridge/) | ✘ | 将抽象与实现分离，使它们可以独立变化。它是用组合关系代替继承关系来实现的，从而降低了抽象和实现这两个可变维度的耦合度。 |
| [Decorator 装饰模式](/structural/decorator/) | ✘ | 动态地给对象增加一些职责，即增加其额外的功能。 |
| [Facade 外观模式](/structural/facade/) | ✘ | 为多个复杂的子系统提供一个一致的接口，使这些子系统更加容易被访问。 |
| [Flyweight 享元模式](/structural/flyweight/) | ✘ | 运用共享技术来有效地支持大量细粒度对象的复用。 |
| [Composite 组合模式](/structural/composite/) | ✘ | 将对象组合成树状层次结构，使用户对单个对象和组合对象具有一致的访问性。 |

### [行为性模式 (Behavioral Patterns)](/behavioral/)

行为型模式用于描述程序在运行时复杂的流程控制，即描述多个类或对象之间怎样相互协作共同完成单个对象都无法单独完成的任务，它涉及算法与对象间职责的分配。

| 模式 | 状态 | 描述 |
|:----|:---: |:------|
| [Template Methed 模板方法](/behavioral/template-methed/) | ✘ | 定义一个操作中的算法骨架，将算法的一些步骤延迟到子类中，使得子类在可以不改变该算法结构的情况下重定义该算法的某些特定步骤。 |
| [Strategy 策略模式](/behavioral/strategy/) | ✘ | 定义了一系列算法，并将每个算法封装起来，使它们可以相互替换，且算法的改变不会影响使用算法的客户。 |
| [Command 命令模式](/behavioral/command/) | ✘ | 将一个请求封装为一个对象，使发出请求的责任和执行请求的责任分割开。 |
| [Chain of Responsibility 职责链模式](/behavioral/chain-of-responsibility/) | ✘ | 把请求从链中的一个对象传到下一个对象，直到请求被响应为止。通过这种方式去除对象之间的耦合。 |
| [State 状态模式](/behavioral/state/) | ✘ | 允许一个对象在其内部状态发生改变时改变其行为能力。 |
| [Observer 观察者模式](/behavioral/observer/) | ✘ | 多个对象间存在一对多关系，当一个对象发生改变时，把这种改变通知给其他多个对象，从而影响其他对象的行为。 |
| [Mediator 中介者模式](/behavioral/mediator/) | ✘ | 定义一个中介对象来简化原有对象之间的交互关系，降低系统中对象间的耦合度，使原有对象之间不必相互了解。 |
| [Iterator 迭代器模式](/behavioral/iterator/) | ✘ | 提供一种方法来顺序访问聚合对象中的一系列数据，而不暴露聚合对象的内部表示。 |
| [Visitor 访问者模式](/behavioral/visitor/) | ✘ | 在不改变集合元素的前提下，为一个集合中的每个元素提供多种访问方式，即每个元素有多个访问者对象访问。 |
| [Memento 备忘录模式](/behavioral/memento/) | ✘ | 在不破坏封装性的前提下，获取并保存一个对象的内部状态，以便以后恢复它。 |
| [Interpreter 解释器模式](/behavioral/interpreter/) | ✘ | 提供如何定义语言的文法，以及对语言句子的解释方法，即解释器。 |
