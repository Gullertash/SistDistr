package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


type Point struct {
	X float64
	Y float64
}


func readInput(filename string) []Point {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return nil
	}


	_, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		panic(err)
	}

	var points []Point
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			x, _ := strconv.ParseFloat(parts[0], 64)
			y, _ := strconv.ParseFloat(parts[1], 64)
			points = append(points, Point{X: x, Y: y})
		}
	}
	return points
}

func main() {
	
	points := readInput("input.txt")
	
	commRange := 15.0 
	maxIters := 100   
	n := len(points)

	// LCM döngüsü (Look-Compute-Move)[cite: 1]
	for iter := 0; iter < maxIters; iter++ {
		newPoints := make([]Point, n)
		copy(newPoints, points)

		for i := 0; i < n; i++ {
			var neighbors []Point
			
		
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				dist := math.Sqrt(math.Pow(points[i].X-points[j].X, 2) + math.Pow(points[i].Y-points[j].Y, 2))
				if dist <= commRange {
					neighbors = append(neighbors, points[j])
				}
			}

		
			
			
			sumX := points[i].X
			sumY := points[i].Y
			for _, nb := range neighbors {
				sumX += nb.X
				sumY += nb.Y
			}
			count := float64(1 + len(neighbors))
			
	
			newPoints[i] = Point{X: sumX / count, Y: sumY / count}
		}

		points = newPoints

	
		maxDist := 0.0
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				d := math.Sqrt(math.Pow(points[i].X-points[j].X, 2) + math.Pow(points[i].Y-points[j].Y, 2))
				if d > maxDist {
					maxDist = d
				}
			}
		}

		if maxDist < 1e-4 {
			fmt.Printf("Rendezvous %d. iterasyonda sağlandı!\n", iter+1)
			break
		}
	}

	
	fmt.Println("Son Pozisyonlar:")
	for i, p := range points {
		fmt.Printf("Robot %d: X=%.4f, Y=%.4f\n", i+1, p.X, p.Y)
	}
}
