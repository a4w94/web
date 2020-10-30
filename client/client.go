package main

///hi test

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = "8000"

	addr = flag.String("addr", "localhost:"+port, "http service address")
)

func main() {

	conn, err := net.Dial("tcp", "localhost:"+port)
	fmt.Println(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//獲取服務端訊息
	go ioCopy(os.Stdout, conn)
	//將使用者輸入的文字訊息傳送到到服務端
	ioCopy(conn, os.Stdin)
}

func ioCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
