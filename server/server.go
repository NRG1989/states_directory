package main

import (
	"fmt"
	"net"
	"strings"
)

var dict = map[string]string{
	"Albania":                "Tirana",
	"Andorra":                "Andorra la Vella",
	"Austria":                "Vienna",
	"Belarus":                "Minsk",
	"Belgium":                "Brussels",
	"Bosnia and Herzegovina": "Sarajevo",
	"Bulgaria":               "Sofia",
	"Croatia":                "Zagreb",
	"Czechia":                "Prague",
	"Denmark":                "Copenhagen",
	"Estonia":                "Tallinn",
	"Finland":                "Helsinki",
	"France":                 "Paris",
	"Germany":                "Berlin",
	"Greece":                 "Athens",
	"Hungary":                "Budapest",
	"Iceland":                "Reykjavik",
	"Ireland":                "Dublin",
	"Italy":                  "Rome",
	"Latvia":                 "Riga",
	"Liechtenstein":          "Vaduz",
	"Lithuania":              "Vilnius",
	"Luxembourg":             "Luxembourg",
	"Malta":                  "Valletta",
	"Moldova":                "Chisinau",
	"Monaco":                 "Monaco",
	"Montenegro":             "Podgorica",
	"Netherlands":            "Amsterdam",
	"North Macedonia":        "Skopje",
	"Norway":                 "Oslo",
	"Poland":                 "Warsaw",
	"Portugal":               "Lisbon",
	"Romania":                "Bucharest",
	"Russia":                 "Moscow",
	"San Marino":             "San Marino",
	"Serbia":                 "Belgrade",
	"Slovakia":               "Bratislava",
	"Slovenia":               "Ljubljana",
	"Spain":                  "Madrid",
	"Sweden":                 "Stockholm",
	"Switzerland":            "Bern",
	"Ukraine":                "Kiev",
	"United Kingdom":         "London",
}

func main() {
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {

		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		source := string(input[0 : n-1])
		source = strings.Title(source)

		target, ok := dict[strings.Title(source)]
		if ok == false {
			target = "undefined"
		}

		conn.Write([]byte(target))
	}
}
