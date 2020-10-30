package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"

	"github.com/Luxurioust/excelize"
)

var (
	port = "8000"
)

func main() {

	http.HandleFunc("/entrance", entrance)
	http.HandleFunc("/entrance/topic", input)

	fmt.Println("server 開啟  http://localhost:8000/entrance")
	//browser.OpenURL("http://localhost:8000/entrance")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("server fault", err)
	}

}

func entrance(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/test.html")
	if err != nil {
		log.Fatal("網頁載入錯誤", err)

	}

	t.Execute(w, nil)
}

func input(w http.ResponseWriter, r *http.Request) {

	topic := SendTopicFromExcel()
	t, err := template.ParseFiles("./html/test.html")

	if err != nil {
		log.Fatal("網頁載入錯誤", err)
	}
	r.ParseForm()
	if r.Method == "GET" {
		t, err1 := template.ParseFiles("./html/test.html")
		log.Fatal(t.Execute(w, nil))
		if err1 != nil {
			log.Fatal(err1)
		}
	} else {
		result := map[string]interface{}{
			"topictitle": topic[1],
			"choose_A":   topic[2],
			"choose_B":   topic[3],
			"choose_C":   topic[4],
			"choose_D":   topic[5],
		}

		t.Execute(w, result)
		if r.Form["select"][0] == topic[6] {
			fmt.Println("Right")
			// fmt.Fprintln(w, "Right")
		} else {
			fmt.Println("Wrong")
		}

	}
	fmt.Println(r.Form)
}

func GetExceltest() [][]string {
	xlsx, err := excelize.OpenFile("./test.xlsx")

	if err != nil {
		log.Fatal("excel檔案開啟錯誤", err)
	}
	row, err1 := xlsx.GetRows("test")
	if err1 != nil {
		log.Fatal("表格開啟錯誤", err1)

	}

	//fmt.Println(row)

	return row
}

func SendTopicFromExcel() []string {
	topicsource := GetExceltest()
	seed := rand.Intn(len(topicsource))
	if seed == 0 {
		seed = 1
	}
	topic := topicsource[seed]

	//fmt.Println(result)
	return topic

}
