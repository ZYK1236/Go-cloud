package meta

import "fmt"

// 文件元信息结构
type FileMeta struct {
	FileSha  string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMeta map[string]FileMeta

// 初始化
func init() {
	fileMeta = make(map[string]FileMeta)
}

func UpdateFileMeta(newFileMeta FileMeta) {
	fileMeta[newFileMeta.FileSha] = newFileMeta
}

func GetFileMeta(fileSha string) FileMeta {
	return fileMeta[fileSha]
}

func LogFileMeta() {
	fmt.Println("当前元信息:")
	for key, value := range fileMeta {
		fmt.Printf("key: %v, value: %+v\n", key, value)
	}
}
