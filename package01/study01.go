package main

import "fmt"

func main() {
	p:=&Pepol{"fff",18,"lll"}
	//test02(p)
	//p.fangfa(8)
	//bianlaing()
	jgtzz(p)

	
}
func bianlaing()  {
	//变量定义的几种
	var i int= 9
	j:=87
	var(
		a1=9
		a2=7
	)
	//常量
	const chang  = "changlaing00000"
	fmt.Print(i,j)
	fmt.Print(a1,a2)
	fmt.Print(chang)
}

//结构体
type Pepol struct {
	name string
	age int
	home string
}

//使用结构体
func test02(pepol Pepol)  {
	pepol.age=15
	pepol.name="fan"
	pepol.home="zjk"
	fmt.Print(pepol)
}

//函数
func hanshu(i int) int  {
	j:=i+6
	return j
}

//方法
//理解：结构体相当于"类"，方法相当于类方法，p相当于实例本身self
func (p Pepol) fangfa(i int)  {
	fmt.Print("fangfa")
	fmt.Print(p.home)
	fmt.Print(i)
}

//结构体指针
func jgtzz(pepol *Pepol)  {
	fmt.Println(pepol)
	fmt.Println(*pepol)
	//这里注意，不用*pepolhome
	fmt.Println(pepol.home)
	//fmt.Print(pepol.home,"/n")

}

