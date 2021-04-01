/*
 * @Description: 上传文件的处理方法
 * @Author: 张宇恺
 * @Date: 2021-04-01 19:03:22
 * @LastEditors: 张宇恺
 * @LastEditTime: 2021-04-01 19:16:21
 */
package index

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

/**
 * @name:
 * @msg:
 * @param {http.ResponseWriter} writer: 向用户返回数据的数据对象
 * @param {*http.Request} req: 用户发的请求（引用）
 * @return {*}
 */
func UploadHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("进来了 UploadHandler")
	if req.Method == "GET" {
		// 返回 html，注意 ioutil.ReadFile 指向的目录是项目根目录
		res, err := ioutil.ReadFile("./static/view/upload.html")
		if err != nil {
			fmt.Println("ioutil.ReadFile err:", err)
			io.WriteString(writer, "server response error")
			return
		}

		io.WriteString(writer, string(res))
	} else {
		// FormFile 函数返回第一个文件，对应 key("file") 值
		file, head, err := req.FormFile("file")
		if err != nil {
			fmt.Println("req.FormFile err:", err.Error())
			return
		}
		defer file.Close()

		// 创建一个文件接受文件流，传入的参数是接受文件的路径
		newFile, err := os.Create("file/" + head.Filename)
		if err != nil {
			fmt.Println("os.Create err:", err.Error())
			return
		}
		defer newFile.Close()

		// 将 FormFile 函数返回的文件写到 newFile 中
		_, error := io.Copy(newFile, file)
		if error != nil {
			fmt.Println("io.Copy err:", error.Error())
			return
		}

		// 上传成功，重定向到成功页面路由
		http.Redirect(writer, req, "/file/upload/success", http.StatusFound)
	}
}

func UploadHandlerSuccess(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("进来了 UploadHandlerSuccess")
	io.WriteString(writer, "Upload success!")
}
