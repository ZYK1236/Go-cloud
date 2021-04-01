/*
 * @Description:
 * @Author: 张宇恺
 * @Date: 2021-04-01 19:02:49
 * @LastEditors: 张宇恺
 * @LastEditTime: 2021-04-01 19:29:51
 */
package main

import (
	"fmt"
	"net/http"

	index "github.com/go_baiduCloud/handler"
)

func main() {
	fmt.Println("starting...")
	http.HandleFunc("/file/upload", index.UploadHandler)
	http.HandleFunc("/file/upload/success", index.UploadHandlerSuccess)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("start server error:", err.Error())
	}
}
