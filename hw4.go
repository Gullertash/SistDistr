/* Ex1

package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool 
}


func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	maxId := id
	
	out <- Msg{Value: id, Done: false}

	for {
		msg := <-in

		
		if msg.Done {
			out <- msg
			break
		}

		// Gelen ID bizim bildiğimizden büyükse günceller ve iletiriz
		if msg.Value > maxId {
			maxId = msg.Value
			out <- Msg{Value: maxId, Done: false}
		} else if msg.Value == maxId {
			// Gelen ID kendi ID'mize eşitse, mesaj tüm halkayı dolaşıp bize dönmüştür.
			// Yani en büyük ID biziz!
			*isLeader = true
			fmt.Printf("Nod %d: Am fost ales lider!\n", id)
			out <- Msg{Value: id, Done: true}
			break
		}
	}
}

// returneaza id-ul liderului
func RingFlooding(ids []int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}

	channels := make([]chan Msg, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan Msg, 10)
	}

	isLeader := make([]bool, n)
	var wg sync.WaitGroup
	wg.Add(n)

	
	for i := 0; i < n; i++ {
		go node(ids[i], channels[i], channels[(i+1)%n], &isLeader[i], &wg)
	}

	wg.Wait()

	
	for i := 0; i < n; i++ {
		if isLeader[i] {
			return ids[i]
		}
	}

	return -1
}

func main() {
	ids := []int{5, 3, 8, 1, 4}
	fmt.Printf("Noduri in inel: %v\n", ids)

	leader := RingFlooding(ids)
	fmt.Printf("Liderul selectat este ID: %d\n", leader)
}

*/




// ex 2
/*

package main

import (
	"fmt"
	"sync"
)


type GraphMsg struct {
	Value int
}


func graphNode(id int, inChannels []<-chan GraphMsg, outChannels []chan<- GraphMsg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	maxId := id
	round := 0


	for _, outCh := range outChannels {
		outCh <- GraphMsg{Value: id}
	}

	
	for {
		updated := false
		round++

	
		for _, inCh := range inChannels {
			select {
			case msg := <-inCh:
				if msg.Value > maxId {
					maxId = msg.Value
					updated = true
				}
			}
		}

		
		if updated {
			for _, outCh := range outChannels {
				outCh <- GraphMsg{Value: maxId}
			}
		}


		
		if maxId == id && round > 2 {
			*isLeader = true
			break
		}
	}
}

func GeneralizedFlooding(ids []int, adjacencyList [][]int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}

	
	channels := make([]chan GraphMsg, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan GraphMsg, 100)
	}

	isLeader := make([]bool, n)
	var wg sync.WaitGroup
	wg.Add(n)

	
	for i := 0; i < n; i++ {
		var inChs []<-chan GraphMsg
		var outChs []chan<- GraphMsg

		
		for j := 0; j < n; j++ {
			for _, neighbor := range adjacencyList[j] {
				if neighbor == i {
					inChs = append(inChs, channels[j])
				}
			}
		}

		
		for _, neighbor := range adjacencyList[i] {
			outChs = append(outChs, channels[neighbor])
		}

		go graphNode(ids[i], inChs, outChs, &isLeader[i], &wg)
	}

	wg.Wait()

	
	for i := 0; i < n; i++ {
		if isLeader[i] {
			return ids[i]
		}
	}

	return -1
}

func main() {
	ids := []int{10, 25, 15, 30}
	adjacencyList := [][]int{
		{1, 2}, 
		{2, 3}, 
		{0, 3}, 
		{0, 1}, 
	}

	fmt.Printf("Graf Nodlari: %v\n", ids)
	leader := GeneralizedFlooding(ids, adjacencyList)
	fmt.Printf("Graf Flooding Sonucu Secilen Lider ID: %d\n", leader)
}

*/

//Ex3
/*

package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type ConsensusMsg struct {
	Val float64
}


func consensusNode(id int, initialVal float64, inChs []<-chan ConsensusMsg, outChs []chan<- ConsensusMsg, resultChan chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()

	currentVal := initialVal
	deg := float64(len(outChs)) 
	
	for step := 0; step < 50; step++ {
	
		for _, outCh := range outChs {
			outCh <- ConsensusMsg{Val: currentVal}
		}

	
		sumIncoming := 0.0
		for _, inCh := range inChs {
			msg := <-inCh
			sumIncoming += msg.Val
		}

	
		currentVal = (1.0 / (deg + 1.0)) * (currentVal + sumIncoming)
		time.Sleep(10 * time.Millisecond) // Simülasyon senkronizasyonu
	}

	resultChan <- currentVal
}

func AverageFlooding(initialValues []float64, adjacencyList [][]int) []float64 {
	n := len(initialValues)
	channels := make([]chan ConsensusMsg, n)
	for i := 0; i < n; i++ {
		channels[i] = make(chan ConsensusMsg, 100)
	}

	resultChans := make([]chan float64, n)
	for i := 0; i < n; i++ {
		resultChans[i] = make(chan float64, 1)
	}

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		var inChs []<-chan ConsensusMsg
		var outChs []chan<- ConsensusMsg

		for j := 0; j < n; j++ {
			for _, neighbor := range adjacencyList[j] {
				if neighbor == i {
					inChs = append(inChs, channels[j])
				}
			}
		}

		for _, neighbor := range adjacencyList[i] {
			outChs = append(outChs, channels[neighbor])
		}

		go consensusNode(i, initialValues[i], inChs, outChs, resultChans[i], &wg)
	}

	wg.Wait()

	finalValues := make([]float64, n)
	for i := 0; i < n; i++ {
		finalValues[i] = <-resultChans[i]
	}

	return finalValues
}

func main() {

	initialStates := []float64{10.0, 20.0, 30.0, 40.0}
	adjacencyList := [][]int{
		{1, 2},
		{0, 3},
		{0, 3},
		{1, 2},
	}

	fmt.Printf("Baslangic Durumlari (Oylar): %v\n", initialStates)
	consensusResults := AverageFlooding(initialStates, adjacencyList)

	fmt.Println("Konsensus Sonrasi Dugum Degerleri:")
	for i, val := range consensusResults {
		fmt.Printf("Nod %d -> %.2f (Yuvarlanmis: %.0f)\n", i, val, math.Round(val))
	}
}

*/
