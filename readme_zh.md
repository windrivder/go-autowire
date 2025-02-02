# Go-AutoWire
> 使用注解自动生成wire依赖注入文件

[English](./readme.md)


项目需要依赖google的依赖注入框架 [wire](https://github.com/google/wire)

但是极大地简化了使用步骤 极大增大了可以使用场景和降低使用门槛

## 安装

安装`wire`:

```sh
go get github.com/google/wire/cmd/wire
```

then 

安装`gutowire`

```sh
go get github.com/windrivder/go-autowire/cmd/gutowire
```

确保`$GOPATH/bin`已经添加到环境变量`$PATH`

## 使用示范


如果我们需要构造一个动物园 我们有以下的结构体定义

```go
package example

type Zoo struct{ 
    Cat         Cat
    Dog         Dog
    FlyAnimal FlyAnimal
}

type Cat struct{
}

type FlyAnimal interface{
    Fly()
}

type Bird struct{
}

func (b Bird)Fly(){
}

type Dog struct{
}
```

如果使用原生`wire`，需要自己手动维护以下一份文档 

```go
package example_zoo

import (
	"github.com/google/wire"
)

var zooSet = wire.NewSet(
	wire.Struct(new(Zoo), "*"),
)

var animalsSet = wire.NewSet(
	wire.Struct(new(Cat), "*"),
	wire.Struct(new(Dog), "*"),

	wire.Struct(new(Bird), "*"),
	wire.Bind(new(FlyAnimal), new(Bird)),
)

var sets = wire.NewSet(zooSet, animalsSet)

func InitZoo() Zoo {
	panic(wire.Build(sets))
}
```

如果你的依赖关系有变动 那么所有的变更都需要在这一份文档进行修改

如果你的动物园越来越大 那么这个文件也会越来越复杂

### 使用`gutowire`

通过写注解的方式 将依赖注入的声明转移到组件定义处

```go
package example

// @autowire.init(set=zoo)
// .init 表示会基于Zoo生成一个实例化入口
// set=zoo代表该组件会被收集到zoo容器
type Zoo struct{ 
    Cat         Cat
    Dog         Dog
    FlyAnimal FlyAnimal
}

// @autowire(set=animals)
type Cat struct{
}


type FlyAnimal interface{
    Fly()
}

// 无参数名的FlyAnimal代表Bird作为FlyAnimal接口实现注入
// @autowire(set=animals,FlyAnimal)
type Bird struct{
}

func (b Bird)Fly(){
}

// @autowire(set=animals)
type Dog struct{
}
```


`@autowire.init` 的`.init`表示作为根对象，将会自动生成实例化入口

```go
// Code generated by go-autowire. DO NOT EDIT.

// +build wireinject
//
package example_zoo

import "github.com/google/wire"

func InitializeZoo() (*Zoo, func(), error) {
	panic(wire.Build(Sets))
}
```

执行命令

```sh
gutowire -s ./example_zoo ./example_zoo
```

`-s` 代表依赖项的搜索范围（文件夹目录）

然后你会发现 使用`wire`所有需要手写的东西都会自动生成了

可以看以下示例：
 
 - [example_zoo](./example_zoo) 
 - [example](./example) 复杂构造的生成
