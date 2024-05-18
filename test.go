package main

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

func main() {
	// 数值初始化 init1()
	//a, b := init1()
	//fmt.Println(a, b)
	interfaceTest()
}

// 格式输出
func format() {
	fmt.Println("hello,world")
	var code = 123
	var date = "2024-3-20"
	var url = "Code=%d&date=%s"
	var target_url = fmt.Sprintf(url, code, date)
	fmt.Println(target_url)
}

// 数值初始化,返回值个数定义
func init1() (int, int) {
	a := 123
	b := 345
	return a, b
}

func constTest() {
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	fmt.Println(a, b, c)
}

func iotaTest() {
	const (
		conut = iota
		a     = "abc"
		b     = len(a)
		c     = unsafe.Sizeof(a)
		d     = iota
	)
	fmt.Println(a, b, c, d)
}

// 取地址
func addr() {
	a := 123
	fmt.Printf("%p\n", &a)
}

// channel  通道，并开启协程
func channelTest() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("收到", msg1)
		case msg2 := <-c2:
			fmt.Println("收到", msg2)
		}
	}
}

// java while在GO中的写法
func forTest() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}

// 无限循环
func forTest1() {
	for {
		fmt.Println("无线循环")
	}
}

// range的解释：
/*
range 关键字不是一个函数，而是一个语言结构。它的行为取决于它所用于的数据类型。

对于数组和切片：range 会返回两个值，分别是索引和对应索引处的元素值。

对于映射：range 会返回两个值，分别是键和对应键所映射的值。

对于通道：range 会迭代通道中的元素值，不过它只返回一个值，即通道中的元素值。
*/
//foreach循环
func forEachTest() {
	strings := []string{"google", "wy"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
func makeTest() {
	m1 := make(map[int]float32)
	m1[0] = 0.0
	m1[1] = 1.0
	m1[2] = 2.0
	for key, value := range m1 {
		fmt.Println("key is ", key, "value is ", value)
		//fmt.Printf("key is %d  value is %.2f\n", key, value)
	}
}
func array() {
	//arr := [...]int{1, 2, 3, 4, 5}
	//指定下标初始化
	//arr1 := [...]int{1: 3, 3: 4}
	arr2 := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	arr2 = append(arr2, row1)
	arr2 = append(arr2, row2)
	//二维数组初始化
	arr3 := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	fmt.Println(len(arr3))

}

func structTest() {
	type Book struct {
		name     string
		page     int
		language string
	}
	//var book1 Book
	book1 := Book{}

	book1.page = 500
	book1.name = "wy"
	book1.language = "汉语"

	bookPtr := &book1

	fmt.Println(bookPtr.name)

}

func strconvTest() {
	s := "10"
	a := 1
	i, err := strconv.Atoi(s)
	k := strconv.Itoa(a)
	fmt.Println(i, err)
	fmt.Printf("%q", k)
	num := 3.14
	str := strconv.FormatFloat(num, 'g', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num, str)
}

// 接口
type Phone interface {
	call()
}
type Nokia struct {
}
type IPhone struct {
}

func (nokia Nokia) call() {
	fmt.Println("nokia")
}

func (iPhone IPhone) call() {
	fmt.Println("iPhone")
}
func interfaceTest() {
	var phone Phone
	phone = new(Nokia)
	phone.call()
	phone = new(IPhone)
	phone.call()
}
