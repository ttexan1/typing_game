package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// 標準入力の受付
	stdinCh := stdInStream(os.Stdin)

	// 引数の処理
	t := 1
	flag.IntVar(&t, "t", t, "use")
	flag.Parse()
	fileName := flag.Args()[0]
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	words := strings.Split(string(data[:]), "\n")
	rand.Seed(time.Now().Unix())
	currentWord := words[rand.Intn(len(words))]
	fmt.Print(currentWord, ">")

	score := 0
	timer := time.After(time.Duration(t) * time.Second)
	finished := false
	for !finished {
		select {
		case given := <-stdinCh:
			if given == currentWord {
				currentWord = words[rand.Intn(len(words))]
				score++
			}
			fmt.Print(currentWord, ">")
		case <-timer:
			fmt.Println("Time Over")
			finished = true
		default:

		}
	}
	fmt.Println("Score:", score)
}

func stdInStream(r io.Reader) <-chan string {
	ch1 := make(chan string)
	go func() {
		sc := bufio.NewScanner(r)
		// sc.Split(bufio.ScanWords)
		for sc.Scan() {
			ch1 <- sc.Text()
		}
	}()
	return ch1
}
