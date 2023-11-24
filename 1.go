package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fp, err := os.Create("./test1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	bufferedWriter := bufio.NewWriter(fp)

	starttime := time.Now()

	for i := 0; i < 100000; i++ {
		data := []byte("hello world\n")
		_, err := bufferedWriter.Write(data)
		if err != nil {
			fmt.Println("写入data中失败")
			return
		}
	}

	bufferedWriter.Flush()

	elapsedtime := time.Since(starttime)
	fmt.Println("写入时间", elapsedtime)

	fp2, err := os.Create("./test2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp2.Close()

	starttime2 := time.Now()

	for j := 0; j < 100000; j++ {
		data2 := []byte("hello world\n")
		_, err := fp2.Write(data2)
		if err != nil {
			fmt.Println("写入data2失败")
			return
		}
	}

	elapsedtime2 := time.Since(starttime2)

	fmt.Println("写入时间", elapsedtime2)

}
