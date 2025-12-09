package main

import (
	"fmt"
	"sync"
	"task2/mutext"
)

func main() {

	// ---------- 指针
	//题目一
	// a := 10
	// fmt.Println("the value before added=", a)
	// pointer.Add10(&a)
	// fmt.Println("the value after added=", a)
	//----------------------------------------

	//题目二
	// arr := []int{1, 2, 3, 4}
	// fmt.Println("the ele before added=", arr)
	// pointer.ArrMulti2(&arr)
	// fmt.Println("the ele after added=", arr)
	//---------------------------------------

	// --------- goroutine
	//题目一
	// go goroutine.NumberPrinter(false)
	// go goroutine.NumberPrinter(true)
	// time.Sleep(2 * time.Second)
	//-----------------------------

	//题目二
	// coordinator := goroutine.TaskCoordinator{}
	// fc := func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Printf("执行中...; ")
	// 	}
	// }
	// fc2 := func() {
	// 	for i := 0; i < 50; i++ {
	// 		fmt.Printf("执行中...; ")
	// 	}
	// }
	// fc3 := func() {
	// 	for i := 0; i < 100; i++ {
	// 		fmt.Printf("执行中...; ")
	// 	}
	// }
	// frr := []func(){fc, fc2, fc3}
	// coordinator.AddTask(frr...)
	// coordinator.StartAllTask()
	//--------------------------------------------

	//--------------------oop
	//题目一
	// circle := oop.Circle{Radius: 5}
	// rectangle := oop.Rectangle{Length: 3, Width: 4}
	// fmt.Printf("the area of circle=%f,the perimeter of circle=%f \n", circle.Area(), circle.Perimeter())
	// fmt.Printf("the area of rectangle=%f,the perimeter of rectangle=%f \n", rectangle.Area(), rectangle.Perimeter())
	//题目二
	// employee := oop.Employee{Person: oop.Person{Name: "张三", Age: 78}, EmployeeId: 1}
	// employee.PrintInfo()

	//---------------------channel
	//题目一
	// channel.Communication()
	//题目二
	// channel.BufferChan()

	//--------------------锁
	//题目一
	// counter := mutext.Conter{}
	// var wg sync.WaitGroup
	// wg.Add(10)
	// fc := func() {
	// 	defer wg.Done()
	// 	counter.Increment()
	// }
	// for i := 10; i > 0; i-- {
	// 	go fc()
	// }
	// wg.Wait()
	// fmt.Println("the utimate count is", counter.Count)

	//题目二
	counter := mutext.Counter2{}
	var wg sync.WaitGroup
	wg.Add(10)
	fc := func() {
		defer wg.Done()
		counter.Increment()
	}
	for i := 10; i > 0; i-- {
		go fc()
	}
	wg.Wait()
	fmt.Println("the utimate count is", counter.Count.Load())

}
