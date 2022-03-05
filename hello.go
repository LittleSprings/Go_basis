package main

//	hello.go归属于包main，每个文件都必须归属于一个包
import "fmt"

//	表示导入fmt包

func hello() {
	const (
		a = 1    //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
