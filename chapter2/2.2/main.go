package main

import (
	"fmt"
	"strconv"
)

func main() {
	//boolDemo()
	//intDemo()
	floatDemo()
}

func boolDemo() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
}

func intDemo() {
	a := 1
	var b int = 0b1100
	var c int = 0o14
	var d int = 0xC

	fmt.Printf("10进制:%d，%T\n", a, a)
	fmt.Printf("2进制数 %b 表示的是: %d \n", b, b)
	fmt.Printf("8进制数 %o 表示的是: %d \n", c, c)
	fmt.Printf("16进制数 %X 表示的是: %d \n", d, d)
}

func floatDemo() {
	// 默认类型
	// var a float32 = math.Pi
	// b := math.Pi
	// fmt.Printf("type of a:%T, value of a:%.4f, b:%T, b:%.5f\n", a, a, b, b)
	// fmt.Printf("%v\n", float64(a))

	// 精度截取问题
	c := 100.00
	ff := 0.310615789
	d := FloatRound(ff, 4)
	e := c * d
	fmt.Println("e:", e)

	f := c * ff
	g := FloatRound(f, 4)
	fmt.Println("g:", g)
}

// 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}
