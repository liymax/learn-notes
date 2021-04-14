```js
package main

import (
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

//go:embed html/index.html
var indexHtml []byte
//go:embed html/d3.v6.min.js
var d3Js []byte

func main()  {
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}
	addr, err := net.LookupHost(name)
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return
	}
	fmt.Println("http://" + addr[0] + ":9090")
	fmt.Println("http://" + addr[1] + ":9090")
	fmt.Println("http://127.0.0.1:9090")
	r := gin.Default()
	r.GET("/fileList", func(c *gin.Context) {
		pwd, _ := os.Getwd()
		//获取文件或目录相关信息
		fileList, err := ioutil.ReadDir(pwd)
		if err != nil {
			log.Fatal(err)
		}
		fileNames := make([]string, 0)
		for _, file := range fileList {
			fileNames =append(fileNames,file.Name())
		}
		c.JSON(200,gin.H{
			"result": fileNames,
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", indexHtml)
	})
	r.GET("/d3.js", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/javascript", d3Js)
	})
	r.MaxMultipartMemory = 400 << 20  // 400 MiB
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.SaveUploadedFile(file, "./" + file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.Static("/file", "./")
	_ = r.Run(":9090")
}

/**
 * 构建命令
 * go build -o dispatcher.exe .\dispatcher.go
 */
```
