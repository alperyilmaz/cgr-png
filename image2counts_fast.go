package main

import (
	"bytes"
	"fmt"
        //"log"
	"os"
	"strconv"
	"sync"
        "math"

	"github.com/disintegration/imaging"
)

var bytePool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024)
	},
}

//var logger = log.New(log.Writer(), "coordToSeq: ", log.LstdFlags)

func main() {
        //logger.Println("Starting main")
        if len(os.Args) < 2 {
                fmt.Println("No filename provided. Please provide image filename.")
                return
        }

        file := os.Args[1]

        img, err := imaging.Open(file)
        if err != nil {
                fmt.Println(err)
                return
        }
        width := img.Bounds().Dx()
        height := img.Bounds().Dy()

        base := int(log2(float64(width)))

	var buf bytes.Buffer
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			rgb := uint32(r8)<<16 + uint32(g8)<<8 + uint32(b8)

			seq := coordToSeq(x, y, base)
			fmt.Fprintf(&buf, "%s\t%d\n", seq, rgb)
		}
	}
	os.Stdout.Write(buf.Bytes())
}

func log2(x float64) int {
        return int(math.Log2(x))
}

func coordToSeq(x, y, base int) string {
	xstr := fmt.Sprintf("%0"+strconv.Itoa(base)+"b", x)
	ystr := fmt.Sprintf("%0"+strconv.Itoa(base)+"b", y)
        //logger.Printf("x=%d, xstr=%s, y=%d, ystr=%s", x, xstr, y, ystr)
	seq := make([]byte, base)
	for i := 0; i < base; i++ {
		if xstr[i] == '0' {
			if ystr[i] == '0' {
				seq[i] = 'A'
			} else {
				seq[i] = 'C'
			}
		} else {
			if ystr[i] == '0' {
				seq[i] = 'G'
			} else {
				seq[i] = 'T'
			}
		}
	}

	return string(seq)
}
