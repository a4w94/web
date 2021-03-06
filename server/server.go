package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		//log.Fatal()列印錯誤資訊並呼叫os.Exit(1)，終止程式
		log.Fatal(err)
	}
	// http.HandleFunc("/", myweb)

	// //browser.OpenURL("http://localhost:8000")
	// err1 := http.ListenAndServe(":8000", nil)

	// if err1 != nil {
	// 	fmt.Println("服務器錯誤", err1)
	// }

	//廣播，傳送訊息到所有客戶端
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		//每個客戶端一個goroutine
		go handleConn(conn)
	}
}

//channel的三種類型(只發送、只接受、即傳送也接受)
//這裡client只發送不接受
//只接受 type client <-chan string
//即傳送也接受 type client chan string
type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-message:
			for cli := range clients {
				//這裡的cli就是handleConn裡的ch channel，
				//writeToCLient goroutine一直在監聽ch channel，讀取channel中的內容，並寫入客戶端連線
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	//寫入訊息到客戶端的連線
	go writeToCLient(conn, ch)

	who := conn.RemoteAddr().String()
	//當客戶端連線過來時，給客戶端一條訊息
	//注意，這時的ch會立馬被writeToCLient goroutine讀取，併發送到當前客戶端
	//所以已連線的其他客戶端不會接受到該條訊息
	ch <- "You are " + who

	//這裡的message channel會被broadcaster讀取，廣播給所有已連線的客戶端
	//注意，這時當前客戶端還沒給entering，所以當前客戶端不會接受到該條訊息
	message <- who + " are arrived"
	//將當前客戶端傳送給entering channel，broadcaster會將當前客戶端新增到已連線的客戶端集合中
	entering <- ch

	input := bufio.NewScanner(conn)
	//阻塞監聽客戶端輸入
	for input.Scan() {
		//獲取客戶端輸入，併發送到message channel，然後broadcaster會將它廣播給所有連線的客戶端
		//因為這時，當前客戶端已經新增到clients集合中，所以當前客戶端也會接受到訊息
		message <- who + ": " + input.Text()
	}

	//客戶端斷開連線
	leaving <- ch
	message <- who + " are left"
	conn.Close()
}

func writeToCLient(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
