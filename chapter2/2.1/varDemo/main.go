package main

import "fmt"

func main() {
	var t1 int
	var f1 float32
	var s1 string
	fmt.Printf("t1:%d, f1:%f, s1:%s\n", t1, f1, s1)

	var t2 = 1
	var f2 = 0.0
	var s2 = "sxx"
	fmt.Printf("t2:%d, f2:%f, s2:%s\n", t2, f2, s2)

	t3 := t2
	f3 := f2
	s3 := s2
	fmt.Printf("t3:%d, f3:%f, s3:%s\n", t3, f3, s3)

	var t4, t5, t6 int = 1, 2, 3
	fmt.Printf("t4:%d, t5:%d, t6:%d\n", t4, t5, t6)

	var t7, f7, s7 = 4, 0.0, "s7"
	fmt.Printf("t7:%d, f7:%f, s7:%s\n", t7, f7, s7)

	var t8 int = 8
	{
		var t8 int = 9
		fmt.Printf("t8:%d\n", t8)
	}
	fmt.Printf("t8:%d\n", t8)

	var t9 = 1
	var t10 = t9
	t9 = t9 + 1
	fmt.Printf("t9:%d, t10:%d\n", t9, t10)

	var t11 = 1
	var t12 = &t11
	t11 = t11 + 1
	fmt.Printf("t11:%d, t12:%d\n", t11, *t12)
	var su int = 0
	sum(1, 2, &su)
	fmt.Printf("su:%d\n", su)

}

func sum(i, j int, su *int) {
	*su = i + j
}
