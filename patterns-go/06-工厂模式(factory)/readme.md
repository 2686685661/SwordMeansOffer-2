# 工厂模式

简单工厂和抽象工厂有些区别，除代码结构上区别，主要区别在于使用场景不同

## 简单工厂模式(simple)
用于创建单一的产品，将所有子类创建过程集中于一个工厂中，如要修改，只需要修改一个工厂即可。  
简单工厂经常和单例模式一起使用，例如用简单工厂创建缓存对象（文件缓存），某天需要改用redis缓存，修改工厂即可。  


## 抽象工厂模式(abstract)
常用于创建一整个产品族，而不是单一的产品  
通过选择不同的工厂来达到目的，其优势在于可以通过替换工厂而快速替换整个产品族  
例如例子中美国工厂生产美国girl，中国工厂生产中国girl。   