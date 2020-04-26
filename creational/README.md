# 创建型模式 (Creational Patterns)

创建型模式的主要关注点是“怎样创建对象？”，它的主要特点是“将对象的创建与使用分离”。这样可以降低系统的耦合度，使用者不需要关注对象的创建细节，对象的创建由相关的工厂来完成。

| 模式 | 状态 | 描述 |
|:----|:---: |:------|
| [Singleton 单例模式](/creational/singleton/) | ✔ | 某个类只能生成一个实例，该类提供了一个全局访问点供外部获取该实例。 |
| [Prototype 原型模式](/creational/prototype/) | ✘ | 将一个对象作为原型，通过对其进行复制而克隆出多个和原型类似的新实例。 |
| [Simple Factory 简单工厂模式](/creational/simple-factory/) | ✘ | 定义一个用于创建产品的接口，根据参数决定生产什么产品。|
| [Factory Method 工厂方法模式](/creational/factory-method/) | ✘ | 定义一个用于创建产品的接口，由子类决定生产什么产品。 |
| [Abstract Factory 抽象工厂模式](/creational/abstract-factory/) | ✘ | 提供一个创建产品族的接口，其每个子类可以生产一系列相关的产品。 |
| [Builder 建造者模式](/creational/builder/) | ✘ | 将一个复杂对象分解成多个相对简单的部分，然后根据不同需要分别创建它们，最后构建成该复杂对象。 |
