package main

import "fmt"

type VisitorFunc func() error

type VisitorList []Visitor

type Visitor interface {
	Visit(VisitorFunc) error
}

func (l VisitorList) Visit(fn VisitorFunc) error {
	for i := range l {
		if err := l[i].Visit(func() error {
			fmt.Println("In VisitorList before fn")
			fn()
			fmt.Println("In VisitorList after fn")
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

type Visitor1 struct{}

func (v Visitor1) Visit(fn VisitorFunc) error {
	fmt.Println("In Visitor1 before fn")
	fn()
	fmt.Println("In Visitor1 after fn")
	return nil
}

type Visitor2 struct {
	visitor Visitor
}

func (v Visitor2) Visit(fn VisitorFunc) error {
	v.visitor.Visit(func() error {
		fmt.Println("In Visitor2 before fn")
		fn()
		fmt.Println("In Visitor2 after fn")
		return nil
	})
	return nil
}

type Visitor3 struct {
	visitor Visitor
}

func (v Visitor3) Visit(fn VisitorFunc) error {
	v.visitor.Visit(func() error {
		fmt.Println("In Visitor3 before fn")
		fn()
		fmt.Println("In Visitor3 after fn")
		return nil
	})
	return nil
}

/*
Visitor Example 代码示例中定义了Visitor接口，增加了VisitorList对象，该对象相当于多个visitor匿名函数的集合。
另外定义了3个visitor类，分别实现了visitor方法，在每一个visitor执行之前（before）和之后（after）分别print信息。
代码执行结果如下：
In Visitor1 before fn
In Visitor2 before fn
In Visitor3 before fn
In VisitorList before fn
In visitorFunc
In VisitorList after fn
In Visitor3 after fn
In Visitor2 after fn
In Visitor1 after fn


通过Visitor代码示例的输出，能够更好地理解Visitor的多层嵌套关系。在main函数中，首先将Visitor1嵌入VisitorList中，
VisitorList是Visitor的集合，可存放多个Visitor。然后将VisitorList嵌入Visitor2中，接着将Visitor2嵌入Visitor3中。
最终形成VisitorList{Visitor3{Visitor2{Visitor1}}}的嵌套关系。

根据输出结果，最先执行的是Visitor1中fn匿名函数之前的代码，然后是VisitorList、Visitor2和Visitor3中fn匿名函数之前的代码。
紧接着执行VisitFunc（visitor.Visit）。最后执行Visitor3、Visitor2、VisitorList、Visitor1的fn匿名函数之后的代码。
整个多层嵌套关系的执行过程有些类似于递归操作。


多层嵌套关系理解起来有点困难，如果读者看过电影《盗梦空间》的话，该过程可以类比为其中的场景。
每次执行Visitor相当于进入盗梦空间中的另一层梦境，在触发执行了visitFunc return后，就开始从每一层梦境中苏醒过来。
*/
func main() {

	var visitor Visitor
	var visitors []Visitor

	visitor = Visitor1{}
	visitor = Visitor2{visitor}
	visitor = Visitor3{visitor}

	visitors = append(visitors, visitor)
	visitor = VisitorList(visitors)
	visitor.Visit(func() error {
		fmt.Println("In visitorFunc")
		return nil
	})

}
