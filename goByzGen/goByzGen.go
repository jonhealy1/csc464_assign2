package main

import "fmt"
//import "time"
import "sync"

var mutex = &sync.Mutex{}

func BGP(genOrder int, id int, m int, n int, g7 map[int][]int, wg *sync.WaitGroup) (map[int][]int) {

	if m == 0 {
		defer wg.Done()
		mutex.Lock()
		//g7[id] = append(g7[id], genOrder)
		//fmt.Println("id:", id, g7[id])
		mutex.Unlock()
		return g7
	}

	mutex.Lock()
		
	for j := 0; j < n; j++ {
		if j == id {
		} else { 
			g7[j] = append(g7[j], g7[id][0])
		}
	}
	
	mutex.Unlock()
	
	return BGP(genOrder, id, m-1, n, g7, wg)
}

func main() {

	fmt.Println("How many Generals? (One will be the Commander):")
	var numGen int 
	fmt.Scanln(&numGen) 

	fmt.Println("How many Traitors? Or is the Commander a Traitor (enter 0):")
	var numTrait int  
	fmt.Scanln(&numTrait) 

	var genOrder int 
		
	if numTrait != 0 {
		fmt.Println("Does the Commander order ATTACK(enter 1) or RETREAT(0)?:")
		fmt.Scanln(&genOrder) 
	} else {
		genOrder = 1
	}
	var wg sync.WaitGroup
	
	var traitOrder int 

	if genOrder == 1 {
		traitOrder = 0 
	} else if genOrder == 0 {
		traitOrder = 1
	}

	wg.Add(numGen-1)

	g7 := make(map[int][]int)

	if numTrait != 0 {
		for j:=1; j<numGen-numTrait; j++ {
			g7[j] = append(g7[j], genOrder)
		}
	} else {
		for j:=1; j<numGen-numTrait; j++ {
			if j%2 == 0 {
				g7[j] = append(g7[j], 0)
			} else if j%2 == 1 {
				g7[j] = append(g7[j], 1)
			}
		}
	}

	for k:=1; k<numTrait+1; k++ {
		g7[numGen-k] = append(g7[numGen-k], traitOrder)
	}

	var recursion int

	if numTrait == 0 {
		recursion = 1
	} else {
		recursion = numTrait
	}

	for i:=1; i<numGen; i++ {
		go BGP(genOrder, i, recursion, numGen, g7, &wg)
	}

	wg.Wait()

	for i:=1; i<numGen; i++ {
		fmt.Println("Lieutenant", i, "received:", g7[i])
		//fmt.Print("Lieutenant ", i, " ")
		zeroCount := 0
		oneCount := 0

		for _, element := range g7[i] {
			if element == 0 {
				zeroCount++
			} else {
				oneCount++
			}
		}

		if numTrait > 1 {
			if g7[i][0] == 1 {
				oneCount = oneCount - 1
			} else if g7[i][0] == 0 {
				zeroCount = zeroCount - 1
			}
		}

		if numTrait != 0 && g7[i][0] != traitOrder {
			if zeroCount > oneCount {
				fmt.Println("RETREAT")
			} else if oneCount > zeroCount {
				fmt.Println("ATTACK")
			} else {
				fmt.Println("UNDECIDED")
			}
		} else if numTrait != 0 && g7[i][0] == traitOrder {
			fmt.Println("TRAITOR")
		} else {
			if zeroCount > oneCount {
				fmt.Println("RETREAT")
			} else if oneCount > zeroCount {
				fmt.Println("ATTACK")
			} else {
				fmt.Println("UNDECIDED")
			}
		}
	}
}