package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	ten                  = 10
	one_billion          = 1000000000
	one_hundred          = 100
	one_thoundsand       = 1000
	one_million          = 1000000
	ten_million          = 10000000
	one_hundred_million  = 100000000
	five_hundred_million = 500000000
)

const maxRW int = 1 << 30
const newLine byte = 10

func main() {
	start := time.Now()

	err := Process(one_billion)
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Process took %s", time.Since(start))
}

func Process(target int) error {
	file, err := os.Create("result.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriterSize(file, 2*maxRW)

	pool := []byte("0123456789")

	lenTarget := len(fmt.Sprintf("%s", target))
	numByte := make([]byte, lenTarget+1)
	//numByte := make([]byte, 11)
	lenNum := 1
	i := 1
	var j int

	for target > 0 {
		numByte[lenNum] = newLine
		j = lenNum-1
		for i <= 9 { // except `9` because `9 + 1 = 10`
			numByte[j] = pool[i]
			if w.Buffered()+lenNum+1 >= maxRW {
				diffLen := maxRW - w.Buffered()
				_, _ = w.Write(numByte[:diffLen])
				_ = w.Flush()
				_, _ = w.Write(numByte[diffLen : lenNum+1])
			} else {
				_, _ = w.Write(numByte[:lenNum+1])
			}
			i++
			target--
		}

		for numByte[j] == pool[9] {
			numByte[j] = pool[0]
			if j == 0 {
				lenNum++
				numByte[lenNum-1] = pool[0]
				numByte[lenNum] = newLine
				break
			}
			j--
		}

		numByte[j]++
		i = 1
		target--
		if w.Buffered()+lenNum+1 >= maxRW {
			diffLen := maxRW - w.Buffered()
			_, _ = w.Write(numByte[:diffLen])
			_ = w.Flush()
			_, _ = w.Write(numByte[diffLen : lenNum+1])
		} else {
			_, _ = w.Write(numByte[:lenNum+1])
		}
	}
	_ = w.Flush()

	return nil
}
