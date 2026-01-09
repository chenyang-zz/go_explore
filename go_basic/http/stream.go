package http

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

var contents = []string{"蜀道之难，难于上青天！", "蚕丛及鱼凫，开国何茫然！", "尔来四万八千岁，不与秦塞通人烟。", "西当太白有鸟道，可以横绝峨眉巅。", "地崩山摧壮士死，然后天梯石栈相钩连。", "上有六龙回日之高标，下有冲波逆折之回川。", "黄鹤之飞尚不得过，猿猱欲度愁攀援。"}

func ChunkedTransfer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("Transfer-Encoding", "chunked") // 数据分块传送。一般情况下， client接收数据长度到达Content-Length后才开始渲染，但由于此处我们是分批渲染，所以不需要（也不能）设置Content-Length
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<html><body><ol>")
	for _, chunk := range contents {
		fmt.Fprintf(w, "<li>%s</li>", chunk) // 发送一块数据
		flusher.Flush()                      // 强制数据立刻发给对方
		time.Sleep(time.Second)              // 故意卡顿一下，前端分段渲染效果更明显
	}
	fmt.Fprintf(w, "</ol></body></html>")
}

// Server-Sent Events。基于HTTP的单向数据流技术，服务端可实时向客户端推送数据。客户端在连接断开后会自动尝试重连
func SSE(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/event-stream; charset=utf-8") //标识响应为事件流。charset=utf-8是为了解决中文乱码
	//w.Header().Add("Cache-Control", "no-cache")                        //防止浏览器缓存响应，确保实时性
	//w.Header().Add("Connection", "keep-alive")                         //保持连接开放，支持持续流式传输

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	for _, chunk := range contents {
		fmt.Fprintf(w, "data: %s\n\n", chunk) // 发送一条数据
		flusher.Flush()                       // 强制数据立刻发给对方
		time.Sleep(time.Second)               // 故意卡顿一下，前端分段渲染效果更明显
	}

	fmt.Fprint(w, "data: [DONE]\n\n") // 结束标志
	flusher.Flush()
}

func StartHttpServerStream() {
	// 路由
	http.HandleFunc("/chunk", ChunkedTransfer)
	http.HandleFunc("/server_source_event", SSE)
	http.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./sse.html")
		if err != nil {
			fmt.Println("create template failed:", err)
			return
		}
		tmpl.Execute(w, map[string]string{"url": "https://5678-firebase-goexplore-1767890864420.cluster-qxqlf3vb3nbf2r42l5qfoebdry.cloudworkstations.dev/server_source_event"})
	})

	// 启动Http Server
	if err := http.ListenAndServe("127.0.0.1:5678", nil); err != nil {
		panic(err)
	}
}
