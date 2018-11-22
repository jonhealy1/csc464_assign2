package main

import "fmt"
import "time"
import "sync"

var wg sync.WaitGroup

type person struct {
	a int
	b int
	c int
}

func sendAB(sender *person, recvAB chan int) {
	sender.a++
	recvAB <- sender.a
	recvAB <- sender.b
	recvAB <- sender.c
}

func sendAC(sender *person, recvAC chan int) {
	sender.a++
	recvAC <- sender.a
	recvAC <- sender.b
	recvAC <- sender.c
}

func sendBA(sender *person, recvBA chan int) {
	sender.b++
	recvBA <- sender.a
	recvBA <- sender.b
	recvBA <- sender.c
}

func sendBC(sender *person, recvBC chan int) {
	sender.b++
	recvBC <- sender.a
	recvBC <- sender.b
	recvBC <- sender.c
}

func sendCA(sender *person, recvCA chan int) {
	sender.c++
	recvCA <- sender.a
	recvCA <- sender.b
	recvCA <- sender.c
}

func sendCB(sender *person, recvCB chan int) {
	sender.c++
	recvCB <- sender.a
	recvCB <- sender.b
	recvCB <- sender.c
}

func receiveAB(sender *person, recvAB chan int) {
	
	tempA := <-recvAB
	tempB := <-recvAB
	tempC := <-recvAB
	fmt.Println("msg from A:", tempA, tempB, tempC, "to B:", sender.a, sender.b, sender.c)
	if tempA > sender.a {
		sender.a = tempA
	}
	if tempC > sender.c {
		sender.c = tempC
	}
	sender.b++
	fmt.Println("B updated to:", sender.a, sender.b, sender.c)
}

func receiveAC(sender *person, recvAC chan int) {
	
	tempA := <-recvAC
	tempB := <-recvAC
	tempC := <-recvAC
	fmt.Println("msg from A:", tempA, tempB, tempC, "to C:", sender.a, sender.b, sender.c)
	if tempA > sender.a {
		sender.a = tempA
	}
	if tempB > sender.b {
		sender.b = tempB
	}
	sender.c++
	fmt.Println("C updated to:", sender.a, sender.b, sender.c)
}

func receiveBC(sender *person, recvBC chan int) {
	
	tempA := <-recvBC
	tempB := <-recvBC
	tempC := <-recvBC
	fmt.Println("msg from B:", tempA, tempB, tempC, "to C:", sender.a, sender.b, sender.c)
	if tempA > sender.a {
		sender.a = tempA
	}
	if tempB > sender.b {
		sender.b = tempB
	}
	sender.c++
	fmt.Println("C updated to:", sender.a, sender.b, sender.c)
}

func receiveBA(sender *person, recvBA chan int) {
	
	tempA := <-recvBA
	tempB := <-recvBA
	tempC := <-recvBA
	fmt.Println("msg from B:", tempA, tempB, tempC, "to A:", sender.a, sender.b, sender.c)
	if tempB > sender.b {
		sender.b = tempB
	}
	if tempC > sender.c {
		sender.c = tempC
	}
	sender.a++
	fmt.Println("A updated to:", sender.a, sender.b, sender.c)
}

func receiveCA(sender *person, recvCA chan int) {
	
	tempA := <-recvCA
	tempB := <-recvCA
	tempC := <-recvCA
	fmt.Println("msg from C:", tempA, tempB, tempC, "to A:", sender.a, sender.b, sender.c)
	if tempB > sender.b {
		sender.b = tempB
	}
	if tempC > sender.c {
		sender.c = tempC
	}
	sender.a++
	fmt.Println("A updated to:", sender.a, sender.b, sender.c)
}

func receiveCB(sender *person, recvCB chan int) {
	
	tempA := <-recvCB
	tempB := <-recvCB
	tempC := <-recvCB
	fmt.Println("msg from C:", tempA, tempB, tempC, " to B:", sender.a, sender.b, sender.c)
	if tempA > sender.a {
		sender.a = tempA
	}
	if tempC > sender.c {
		sender.c = tempC
	}
	sender.b++
	fmt.Println("B updated to:", sender.a, sender.b, sender.c)
}


func main() {

	wg.Add(3)
	var mutex = &sync.Mutex{}

	fmt.Println("How many iterations?")
	var counter int   
    fmt.Scanln(&counter) 

	aB := make(chan int)
	bC := make(chan int)
	aC := make(chan int)

	personA := person{a: 0, b: 0, c:0}
	personB := person{a: 0, b: 0, c:0}
	personC := person{a: 0, b: 0, c:0}
	
	As := 0
	Ar := 0 
	Bs := 0
	Br := 0
	Cs := 0
	Cr := 0

	// A function
	go func() {
		defer wg.Done()
		for i:=0; i<counter; i++ {
			receiveBA(&personA, aB)
			Ar++
			time.Sleep(time.Millisecond)
			sendAB(&personA, aB)
			As++
			time.Sleep(time.Millisecond)
			sendAC(&personA, aC)
			As++
			time.Sleep(time.Millisecond)
		}
	}()

	// B function
	go func() {
		defer wg.Done()
		for i:=0; i<counter; i++ {
			receiveCB(&personB, bC)
			Br++
			time.Sleep(time.Millisecond)
			sendBA(&personB, aB)
			Bs++
			time.Sleep(time.Millisecond)
			receiveAB(&personB, aB)
			Br++
			time.Sleep(time.Millisecond)
			sendBC(&personB, bC)
			Bs++
			time.Sleep(time.Millisecond)
		}
	}()

	// C function
	go func() {
		defer wg.Done()
		for i:=0; i<counter; i++ {
			mutex.Lock()
			sendCB(&personC, bC)
			Cs++
			mutex.Unlock()
			time.Sleep(time.Millisecond)
			receiveAC(&personC, aC)
			Cr++
			time.Sleep(time.Millisecond)
			receiveBC(&personC, bC)
			Cr++
			time.Sleep(time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println()
	fmt.Println("final value for A:", personA)
	fmt.Println("final value for B:", personB)
	fmt.Println("final value for C:", personC)
	fmt.Println()
	fmt.Println("A sent", As, "messages and received", Ar, "for a total of", Ar+As)
	fmt.Println("B sent", Bs, "messages and received", Br, "for a total of", Br+Bs)
	fmt.Println("C sent", Cs, "messages and received", Cr, "for a total of", Cr+Cs)
	fmt.Println()
}