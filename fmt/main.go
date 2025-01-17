package main

import "fmt"

func main() {
	func8()
}

func func1(a ...int) {
	fmt.Printf("%#v\n", a)
}

func func2() {
	s := "../../usage-guide/ones-project/guan-li-gong-zuo-xiang/guan-lian-guan-xi-lei-xing"
	fmt.Printf("%s\n", s)

	i := 5
	j, err := fmt.Sscanf(" ../../usage-guide/ones-project/guan-li-gong-zuo-xiang/guan-lian-guan-xi-lei-xing ", "%5s%d", &s, &i)
	if err != nil {
		panic(err)
	}
	fmt.Println(j)
	fmt.Println(s)
}

func func3() {
	s := "{\"condition_groups\":[[{\"field_uuid\":\"field024\",\"operate\":{\"label\":\"filter.addQueryContent.include\",\"operate_id\":\"include\",\"predicate\":\"in\"},\"value\":[\"%s\",\"%s\"]}]]}"
	placeholder := []interface{}{"11111", "222222"}
	fmt.Println(fmt.Sprintf(s, placeholder...))
	// s := "{\"condition_groups\":[[{\"field_uuid\":\"field024\",\"operate\":{\"label\":\"filter.addQueryContent.include\",\"operate_id\":\"include\",\"predicate\":\"in\"},\"value\":[\"%s\",\"%s\"]}]]}"
	// placeholder := []interface{}{"11111", "222222"}
	// fmt.Println(fmt.Sprintf(s, placeholder...))
}

type integer int

func (i integer) String() string {
	return "hello"
}

func func4() {
	fmt.Println(integer(1))
}

// 0 开头，表明是八进制，但八进制最大的数字是 7，因此编译不通过
//func func5() {
//	fmt.Println(09)
//}

func func6() {
	const X = 7.0
	var x interface{} = X

	if y, ok := x.(int); ok {
		fmt.Println(y)
	} else {
		fmt.Println(58)
		fmt.Println(int(y))
	}
}

type myStringer struct {
	s string
}

func (ms *myStringer) String() string {
	return ms.s
}

func func7() {
	//var s fmt.Stringer
	//s = "s"
	//fmt.Println(s)

	ms := new(myStringer)
	ms.s = "s"
	fmt.Println(ms)
}

func func8() {
	v := 101.35345345
	fmt.Printf("value: %-3.1f", v)
}
