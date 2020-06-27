package main

import "fmt"

func main() {
	/*fmt.Print("这是我的go程序")
	test01(11,3)
	fmt.Print(test01(7,9))
	fmt.Print(Person{name: "fan",age: 9})*/
	var ll = []int{2,6,9,22,54,4,15}
	maopao(ll)
}

func test01(i,j int) (int,int) {
	var a1,a2 int = 6,8
	s:="iiiiii"
	fmt.Print(s)
	return a1,a2
}
func maopao(l []int)[]int {
	fmt.Print(l)
	for i:=0;i<len(l)-1;i++{
		for j:=0;j<len(l)-i-1;j++{
			if l[j]>l[j+1]{
				l[j],l[j+1]=l[j+1],l[j]
			}
		}

	}
	fmt.Print(l)
	return l
}
func shujujiegou(){
	a := [3] int{12, 78, 50}
	a2 := [4] int{2,45,5,7}
	fmt.Print(len(a))
	fmt.Print(a2)

	var m = map[string]int{}
	m["aa"]=45

}

type Person struct {
	name string;
	age int;
}