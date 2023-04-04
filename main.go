package main

import (
	"fmt"
	"net/http"
	"os"
	"webchat/glm"
)

func main() {
	http.HandleFunc("/send-message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			message := r.FormValue("message")
			//fmt.Fprintf(w, "问："+message+"\n")

			response, err := glm.Completions("wo", message)
			fmt.Fprintf(w, "金科小兴："+response)

			// 打开文件
			file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer file.Close()

			// 写入文件
			_, err = file.WriteString("message:" + message + "\n")
			_, err = file.WriteString("response:" + response)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 80...")
	http.ListenAndServe(":80", nil)
}
