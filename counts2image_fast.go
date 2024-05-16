package main

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"bytes"
	"sync"
)

func kmerToCoordinates(kmer string) (int, int) {
	mapping := [256]int{
		'A': 0, 'G': 1, 'C': 2, 'T': 3,
	}
	xBin := 0
	yBin := 0
	for _, c := range kmer {
		xBin = (xBin << 1) | (mapping[c] & 1)
		yBin = (yBin << 1) | ((mapping[c] >> 1) & 1)
	}
	return xBin, yBin
}

func countToRGB(count int) color.Color {
	r := uint8(count >> 16 & 0xff)
	g := uint8(count >> 8 & 0xff)
	b := uint8(count & 0xff)
	return color.RGBA{r, g, b, 255}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide DNA length.")
	}
	lenSeq, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	size := 1 << uint(lenSeq)

	img := image.NewRGBA(image.Rect(0, 0, size, size))
	var wg sync.WaitGroup
	for x := 0; x < size; x++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			for y := 0; y < size; y++ {
				img.Set(x, y, color.RGBA{0, 0, 0, 255}) // set background to black
			}
		}(x)
	}
	wg.Wait()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")
		if len(fields) != 2 {
			log.Fatal("Invalid input line:", line)
		}
		seq := fields[0]
		count, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		x, y := kmerToCoordinates(seq)
		rgb := countToRGB(count)
		img.Set(x, y, rgb)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	buf := &bytes.Buffer{}
	err = png.Encode(buf, img)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, buf)
	if err != nil {
		log.Fatal(err)
	}
}
