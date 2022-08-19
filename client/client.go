package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Nice to meet you again! Ask me any european countries and I give you the capitals.")
	defer conn.Close()
	time.Sleep(time.Second)
	for {

		fmt.Print("What is the European country? ")
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if input == "" {
			input = " "
		}

		if n, err := conn.Write([]byte(input)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("the capital of %v: ", input)
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		fmt.Print(string(buff[0:n]))
		fmt.Println()

	}
}
