package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/qiuyesuifeng/task"
)

func Example1() {
	tk := task.NewTask("taska", "0/5 * * * * *", func(a interface{}) error { fmt.Printf("%s hello world\n", time.Now().String()); return nil })
	task.AddTask("taska", tk)
	task.StartTask()
	time.Sleep(60 * time.Second)
	task.StopTask()
}

func Example2() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	tk1 := task.NewTask("tk1", "0 17 * * * *", func(a interface{}) error { fmt.Println(time.Now().String() + " tk1"); return nil })
	tk2 := task.NewTask("tk2", "0,10,20 * * * * *", func(a interface{}) error { fmt.Println(time.Now().String() + " tk2"); wg.Done(); return nil })
	tk3 := task.NewTask("tk3", "0 2 * * * *", func(a interface{}) error { fmt.Println(time.Now().String() + " tk3"); wg.Done(); return nil })
	tk4 := task.NewTask("tk2", "0/2 * * * * *", func(a interface{}) error { fmt.Println(time.Now().String() + " tk4"); return nil })

	task.AddTask("tk1", tk1)
	task.AddTask("tk2", tk2)
	task.AddTask("tk3", tk3)
	task.AddTask("tk4", tk4)

	task.StartTask()
	defer task.StopTask()

	select {
	case <-time.After(200 * time.Second):
		fmt.Println("200 seconds happen")
	case <-wait(wg):
	}
}

func wait(wg *sync.WaitGroup) chan bool {
	ch := make(chan bool)
	go func() {
		wg.Wait()
		ch <- true
	}()
	return ch
}

func main() {
	//Example1()

	Example2()
}
