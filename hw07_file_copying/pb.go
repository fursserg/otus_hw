package main

import (
	"fmt"
)

type progressBar struct {
	Total     int64
	BarLength int

	current int64
}

func (p *progressBar) Add(n int) {
	p.current += int64(n)

	percent := float64(p.current) / float64(p.Total) * 100

	// Рассчитываем количество заполненных и пустых блоков
	filledBlocks := int(float64(p.BarLength) * percent / 100)
	emptyBlocks := p.BarLength - filledBlocks

	fmt.Printf("\r[")
	for i := 0; i < filledBlocks; i++ {
		fmt.Print("=")
	}
	for i := 0; i < emptyBlocks; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("] %.2f%%", percent)
}

func (p *progressBar) End() {
	p.current = 0

	fmt.Println()
}
