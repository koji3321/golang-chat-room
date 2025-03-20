package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func input() {
	var inp bufio.Reader = *bufio.NewReader(os.Stdin)
	var ip string
	for {
		cmd, _ := inp.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		if cmd == "kapat" {
			fmt.Println("kapatılıyor")
			os.Exit(0)
		} else if cmd == "kick" {
			fmt.Print("kimi atıcaksın>")
			fmt.Scanln(&ip)
			for _, i := range ipler {
				if i == ip {
					a, _ := strconv.Atoi(i)
					liste[a].Close()
					liste = append(liste[:a], liste[a+1:]...)
				}
			}
		} else if cmd == "ban" {
			fmt.Print("kimi banlicaksın>")
			fmt.Scanln(&ip)
			for _, i := range ipler {
				if i == ip {
					a, _ := strconv.Atoi(i)
					liste[a].Write([]byte("öldün çık"))
					liste = append(liste[:a], liste[a+1:]...)
				}
			}
		} else if cmd == "help" {
			fmt.Println("kick=serverdan atar")
			fmt.Println("ban=pcyi kapatır serverdan atar")
			fmt.Println("kapat=serverı kapatır")
		}
	}
}

func senderer(message string) {
	for _, i := range liste {
		i.Write([]byte(message))
	}
}

func receiver(client net.Conn, reader *bufio.Reader) {
	for true {

		message, err := reader.ReadString('\n')

		if err != nil {
			client.Close()
			return
		}
		message = strings.TrimSpace(message)
		senderer(message)
	}
}

var liste []net.Conn
var ipler []string

func main() {
	server, _ := net.Listen("tcp", ":3311")
	defer server.Close()

	fmt.Println("server başladı!" + server.Addr().String())
	go input()

	for true {
		client, _ := server.Accept()
		reader := bufio.NewReader(client)
		nickname, _ := reader.ReadString('\n')
		fmt.Println(client.RemoteAddr().String() + " : " + strings.TrimSuffix(nickname, "\n") + " adıyla bağlandı")
		liste = append(liste, client)
		ipler = append(ipler, client.RemoteAddr().String())
		go receiver(client, reader)
	}
}
