package main

import "fmt"

func main() {
	//panduan(9)
	ii:=[3] int{3,5,66}
	xunhuan(ii)
	
}


//if判断
func panduan(i int)  {
	if i>19{
		fmt.Println("dayu19")
	}else {
		fmt.Println("nono")
	}
}

//for循环
func xunhuan(ii [3]int)  {
	for sy,i:=range ii{
		fmt.Println(sy,i)
	}

	for i2:=0;i2<len(ii);i2++{
		fmt.Println(ii[i2])
	}

}
