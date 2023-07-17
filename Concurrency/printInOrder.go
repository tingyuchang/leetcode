package Concurrency

import (
	"fmt"
	"sync"
)

type Foo struct {
	wg  sync.WaitGroup
	wg1 sync.WaitGroup
	wg2 sync.WaitGroup
}

func (f *Foo) first() {
	defer f.wg.Done()
	defer f.wg1.Done()
	fmt.Println("first")
}

func (f *Foo) second() {
	defer f.wg.Done()
	defer f.wg2.Done()
	f.wg1.Wait()
	fmt.Println("second")
}

func (f *Foo) third() {
	defer f.wg.Done()
	f.wg2.Wait()
	fmt.Println("third")
}

func (f *Foo) printFoo() {
	defer f.wg.Done()
	fmt.Println("Foo")
	f.wg2.Done()
}

func (f *Foo) printBar() {
	defer f.wg.Done()
	f.wg2.Wait()
	fmt.Println("Bar")
}

func (f *Foo) PrintInOrder(order []int) {

	f.wg.Add(3)
	f.wg1.Add(1)
	f.wg2.Add(1)
	if len(order) != 3 {
		return
	}
	for _, v := range order {
		switch v {

		case 1:
			go f.first()
		case 2:
			go f.second()
		case 3:
			go f.third()
		default:
			break
		}
	}

	f.wg.Wait()
}

func (f *Foo) PrintFooBarAlternately(n int) {
	for i := 0; i < n; i++ {
		f.wg.Add(2)
		f.wg2.Add(1)
		go f.printFoo()
		go f.printBar()
		f.wg.Wait()
	}

}

func TestFoo() {
	foo := Foo{}
	foo.PrintInOrder([]int{3, 2, 1})
}
