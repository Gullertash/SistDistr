package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


type Info struct {
	ID    int
	Value int
}


type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
}


func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {

		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}
		for _, d := range dist {
			if d == -1 {
				return -1 
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	readNextLine := func() string {
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				return line
			}
		}
		return ""
	}

	
	headerLine := readNextLine()
	if headerLine == "" {
		return
	}
	headerParts := strings.Fields(headerLine)
	n, _ := strconv.Atoi(headerParts[0])
	s, _ := strconv.Atoi(headerParts[1])


	tokenLine := readNextLine()
	tokenParts := strings.Fields(tokenLine)
	tokens := make([]int, n)
	for i := 0; i < n; i++ {
		tokens[i], _ = strconv.Atoi(tokenParts[i])
	}


	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		line := readNextLine()
		fields := strings.Fields(line)
		var neighbors []int
		for _, f := range fields {
			nbr, _ := strconv.Atoi(f)
			neighbors = append(neighbors, nbr)
		}
		adj[i] = neighbors
	}


	diam := diametru(adj)


	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = Node{
			ID:        i,
			Neighbors: adj[i],
			M:         make(Memory),
		}
		nodes[i].M[Info{ID: i, Value: tokens[i]}] = struct{}{}
	}

	maxIter := (s + 1) * diam

	
	for t := 0; t <= maxIter; t++ {
		
		incoming := make([]Memory, n)
		for i := range incoming {
			incoming[i] = make(Memory)
		}

		for i := 0; i < n; i++ {
	
			for _, nbr := range nodes[i].Neighbors {
				for info := range nodes[i].M {
					incoming[nbr][info] = struct{}{}
				}
			}
		}

		
		for i := 0; i < n; i++ {
			for info := range incoming[i] {
				nodes[i].M[info] = struct{}{}
			}
		}
	}

	
	var finalInfos []Info
	for info := range nodes[0].M {
		finalInfos = append(finalInfos, info)
	}

	
	sort.Slice(finalInfos, func(i, j int) bool {
		return finalInfos[i].ID < finalInfos[j].ID
	})

	
	var outputParts []string
	for _, inf := range finalInfos {
		outputParts = append(outputParts, fmt.Sprintf("(%d, %d)", inf.ID, inf.Value))
	}
	fmt.Println(strings.Join(outputParts, " "))
}
