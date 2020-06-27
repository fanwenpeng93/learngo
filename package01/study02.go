package main

import "fmt"

func main() {
	mapp()

}

//数组
func shuzu()  {
	//数组
	shuzu:=[...]int {9,4,2,44,12}
	var a [3]int
	a[0] = 1

	//声明数组的同时初始化
	b := [3]int{1,2,3}
	c := [2][2] int {{1,2},{3,4}}
	//声明一个不定长数组
	d := [...]int {1,2,3,4,5,6}
	fmt.Println(a,b,c,d)
	fmt.Println(shuzu)

}

//Map
func mapp()  {
	m:=map[string]int{}
	m["age"]=23
	delete(m,"age")
	mlen:=len(m)
	fmt.Println(mlen)
	m02:=make(map[string] int)
	m03:=new(map[string] int)

	fmt.Println(m02)
	fmt.Println(m03)

	m02["age02"]=66
	fmt.Println(m)
	fmt.Println(m02)


}

//切片
func  qie()  {
	q:=[]int{1,4,5,6,2,12,54,16,88}
	fmt.Println(q)

}
