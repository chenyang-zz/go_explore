package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// 详细看一下http协议
func HttpObservationServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request method: %s\n", r.Method)
	fmt.Printf("request host: %s\n", r.Host) //服务端host
	fmt.Printf("request url: %s\n", r.URL)
	fmt.Printf("request proto: %s\n", r.Proto)
	fmt.Println("request header:")
	for key, values := range r.Header {
		fmt.Printf("%s: %v\n", key, values)
	}
	fmt.Println()

	fmt.Printf("request body: ")
	io.Copy(os.Stdout, r.Body) // 内容拷贝到os.Stdout流里
	fmt.Println()

	w.WriteHeader(http.StatusBadRequest) //设置StatusCode，不设置默认是200
	// w.WriteHeader(http.StatusOK)               //这行是多余的，不起作用，因为之前已经设置过响应码了
	// 必须先设置响应码和响应头，再设置响应体，否则无效。因为当响应体很大时，响应码和响应头会率先通过一次TCP传送发给client，响应体则会通过多次TCP传送发给client
	w.Header().Add("tRAce-id", "4723956498105") //在WriteHeader之前设置Header。header里的key是大小写不敏感的，会自动把每个单词（各单词用-连接）的首字母转为大写，其他字母转为小写
	w.Write([]byte("Hello Boy\n"))
	fmt.Fprint(w, "Hello Girl\n")
	w.Header().Add("uuid", "0987654321") //无效，所有响应头必须在响应体之前设置完毕。
	fmt.Println(strings.Repeat("*", 60))
}

func GetServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request url: %s\n", r.URL)
	params := ParseUrlParams(r.URL.RawQuery)
	fmt.Fprintf(w, "your name is %s, age is %s\n", params["name"], params["age"])
	fmt.Println(strings.Repeat("*", 60))
}

// 流式传输海量数据
func HugeBodyServer(w http.ResponseWriter, r *http.Request) {
	line := []byte("Heavy is the head who wears the crown.\n")
	const R = 10                                                // line重复发送几次
	w.Header().Add("Content-Length", strconv.Itoa(R*len(line))) //先设置Content-Length
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	for i := range R {
		// 触发flush时，content-length不会自动添加
		if _, err := w.Write(line); err != nil { //即使不显式Flush(), Write()的内容足够多(大几K)时也会触发Flush()
			fmt.Printf("%d send error: %s\n", i, err)
			break
		}
		flusher.Flush() //强制write to tcp
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println(strings.Repeat("*", 60))
}

func StartServer() {
	// 路由
	http.HandleFunc("/obs", HttpObservationServer)
	http.HandleFunc("/get", GetServer)
	http.HandleFunc("/stream", HugeBodyServer)

	// 启动Http Server
	if err := http.ListenAndServe("127.0.0.1:5678", nil); err != nil {
		panic(err)
	}
}
