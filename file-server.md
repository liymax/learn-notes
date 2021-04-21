```go
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
/*//go:embed html/d3.v6.min.js
var d3Js []byte*/

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
			if !file.IsDir() {
				fileNames =append(fileNames,file.Name())
			}
		}
		c.JSON(200,gin.H{
			"result": fileNames,
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", indexHtml)
	})
	/*r.GET("/d3.js", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/javascript", d3Js)
	})*/
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
 * go build -o file-server.exe .\file-server.go
 */
 
```

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>File sending and receiving</title>
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=0">
    <script src="https://d3js.org/d3.v6.min.js"></script>
    <style>
        .file-list{
            display: flex;
            flex-direction: column;
            padding: 3vw;
            align-items: flex-start;
        }
        .file-list > a{
            margin-bottom: 2.5vw;
            font-size: 4.5vw;
            word-break: break-all;
        }
        .upload-btn {
            position: fixed;
            width: 18vw;
            height: 10vw;
            right: 5vw;
            bottom: 5vw;
            text-align: center;
            line-height: 10vw;
            font-size: 5vw;
            color: #fff;
            background-color: #1890ff;
            border-radius: 1vw;
        }
        .upload-btn > input {
            opacity: 0;
            position: absolute;
            width: 100%;
            height: 100%;
            top:0;
            left: 0;
        }
        #progress{
            display: none;
            padding: 2vw 3vw 0;
            justify-content: space-between;
            color: darkgreen;
        }
        #progress > label{
            max-width: 50vw;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
        }
        #progress > em{
            font-style: normal;
        }
        #notify{
            display: none;
            position: fixed;
            top:50%;
            left: 50%;
            transform: translate(-50%, -50%);
            border-radius: 1vw;
            padding: 3vw 4vw;
            max-width: 80vw;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;
            background-color: #e6f7ff;
            font-size: 5vw;
            color: #333;
        }
    </style>
</head>
<body>
    <div id="progress">
        <label></label><em></em><span></span>
    </div>
    <div id="fileList" class="file-list"></div>
    <div class="upload-btn"><input id="fileInput" type="file"/>上传</div>
    <div id="notify">
        <span></span>
    </div>
    <script>
        let isUploading = false
        function makeList(data) {
            d3.select('#fileList').selectAll('a').data(data)
            .join('a').text(d => d)
            .attr('href', d => '/file/' + d)
            .attr('download', d => d)
        }
        function initUploadBtn() {
            const btn =  d3.select('.upload-btn')
            const dsp = {x:0, y: 0}
            let once = false
            btn.call(
                d3.drag().on('drag', (e) => {
                    const x = e.x - dsp.x
                    const y = e.y - dsp.y
                    btn.style('transform', `translate(${x}px,${y}px)`)
                }).on('start', (e) => {
                    if (once) return
                    Object.assign(dsp, { x: e.x, y: e.y })
                    once = true
                })
            )
        }
        function notify(msg) {
            const notify = d3.select('#notify')
            notify.selectChild('span').text(msg)
            notify.style('display','block')
            setTimeout(() => {
                notify.style('display','none')
            }, 2600)
        }
        function fitch(url, opts={}, onProgress) {
            return new Promise((res, rej) => {
                const xhr = new XMLHttpRequest()
                xhr.open(opts.method || 'get', url)
                for (let k in opts.headers||{})
                    xhr.setRequestHeader(k, opts.headers[k])
                xhr.onload = e => res(e.target.responseText)
                xhr.onerror = rej;
                if (xhr.upload && onProgress)
                    xhr.upload.onprogress = onProgress
                xhr.send(opts.body)
            });
        }
        function makeProgress(name, per, speed) {
            const p = d3.select('#progress')
            p.style('display', 'flex')
            p.selectChild('label').text(name)
            if (speed >= 1024) {
                speed /= 1024
                p.selectChild('em').text(speed.toFixed(2)+'Mb/s')
            } else {
                p.selectChild('em').text(speed.toFixed(2)+'Kb/s')
            }
            p.selectChild('span').text(per+'%')
            if(per === 100) {
                setTimeout(() => {
                    p.style('display', 'none')
                }, 1500)
            }
        }
        function initUpload() {
            d3.select('#fileInput').on('change', function (e){
                if (isUploading) {
                    return notify('当前上传中，请稍后！')
                }
                const file = e.target.files[0]
                const formData = new FormData()
                formData.append('file', file)
                let prevTime = Date.now(), prevLoaded = 0
                isUploading = true
                fitch('/upload', { method: 'post', body: formData}, (evt) => {
                    if (evt.lengthComputable) {
                        const curTime = Date.now()
                        const curLoaded = evt.loaded
                        const per = Math.round(curLoaded / evt.total * 100)
                        const speed = (curLoaded-prevLoaded)/(curTime-prevTime)*(1000/1024)
                        makeProgress(file.name, per, speed)
                        prevTime = curTime
                        prevLoaded = curLoaded
                    }
                }).then(res => {
                    console.log(res)
                    getFileList()
                }).catch(err => {
                    console.log(err)
                    notify('上传失败')
                }).finally(()=> isUploading = false)
            })
        }

        const ignores = ['dispatcher.exe']
        function getFileList() {
            fetch('/fileList').then(res => {
                return res.json()
            }).then(data => {
                const downloads = data.result.filter(e => !ignores.includes(e))
                makeList(downloads)
            })
        }
        getFileList()
        initUpload()
        initUploadBtn()
    </script>
</body>
</html>
```
