package main

import "sync"
type FooBar struct {
	n int
    fooLock sync.Mutex
	barLock sync.Mutex
}

func NewFooBar(n int) *FooBar {
    fb := &FooBar{n: n}
    fb.barLock.Lock() 
	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
    for i := 0; i < fb.n; i++ {
		fb.fooLock.Lock()
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.barLock.Unlock()
	}
}

func (fb *FooBar) Bar(printBar func()) {
    for i := 0; i < fb.n; i++ {
		fb.barLock.Lock()
        // printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.fooLock.Unlock()
	}
}
