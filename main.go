package main

import (
	"fmt"
	"go-palyground/datastructure/linearlist/linkedlist"
	"log"
	"math"
	"reflect"
	"runtime"
)

func main() {
	//m := map[string]int{"1": 1,"2": 2}
	//mp2 := make(map[int]*[]int)
	//sub := []int{1, 2, 3, 4}
	//mp2[1] = &sub
	//delete(mp2, 1)
	//fmt.Print(mp2)
	//linkedlist_test()
	var bool bool
	//blank := isBlank(reflect.ValueOf(a))
	fmt.Print(bool)
	//go fmt.Print("1")

}

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func linkedlist_test() {
	linkedlist := linkedlist.EmptyLinkedlist()
	linkedlist.TailAdd("1")
	linkedlist.TailAdd("2")
	linkedlist.TailAdd("3")
	linkedlist.TailAdd("4")
	linkedlist.Reverse()
}

func slice_test() {
	//b:= []string{"g", "o", "l", "a", "t", "f"}
	//ints := b[0:3]
	i := []int{1, 2, 3}
	b := make([]int, 3)
	b[0] = 1
	fmt.Print(i[1:])
	//for _, item := range j {
	//	item *= 2
	//	fmt.Print(item)
	//}
	//b = b[:1]
	//fmt.Print(b)'
}

func doSth(a *int) {
	b := a

	fmt.Printf("b: %s \n", b)
	c := &a
	fmt.Printf("c: %s \n", c)
}

func returnVal() {
	sqrt, ok := mySqrt(-0.0001)
	if ok {
		fmt.Printf("%s\n,%s\n", sqrt, ok)
	}
	fmt.Printf("%s\n", &ok)
}

func switch_test() {
	k := 6
	switch k {
	case 4:
		fmt.Printf("%s \n", "was <= 4")
	case 5:
		fmt.Printf("%s \n", "was <= 5")
	case 6:
		fmt.Printf("%s \n", "was <= 6")
		fallthrough
	case 7:
		fmt.Printf("%s\n", "was <= 7")
	case 8:
		fmt.Printf("%s \n", "was <= 8")
	default:
		fmt.Printf("%s \n", "default case")
	}
}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	} // error case
	return math.Sqrt(f), true
}
func init() {
	fmt.Printf("%s \n", "===============================")
}

func defertest() {
	i := 0
	defer fmt.Printf("===========> %d \n", i)
	i++
	return
}

func debug_test() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
}
