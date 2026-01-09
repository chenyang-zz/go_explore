package http

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 详细看一下http协议
func HttpObservationClient() {
	fmt.Println(strings.Repeat("*", 30) + "GET" + strings.Repeat("*", 30))
	if resp, err := http.Get("http://127.0.0.1:5678/obs?name=chenyang"); err != nil {
		panic(err)
	} else {
		defer resp.Body.Close() //注意：一定要调用resp.Body.Close()，否则会协程泄漏（同时引发内存泄漏）
		fmt.Printf("response proto: %s\n", resp.Proto)
		if major, minor, ok := http.ParseHTTPVersion(resp.Proto); ok {
			fmt.Printf("http major version %d, http minor version %d\n", major, minor)
		}

		fmt.Printf("response status: %s\n", resp.Status)
		fmt.Printf("response statusCode: %d\n", resp.StatusCode)
		fmt.Println("response header")
		for key, values := range resp.Header {
			fmt.Printf("%s: %v\n", key, values)
			if key == "Date" {
				if tm, err := http.ParseTime(values[0]); err == nil {
					fmt.Printf("server time %s\n", tm.Format("2006-01-02 15:04:05"))
				}
			}
		}

		fmt.Println("response body:")
		io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
		os.Stdout.WriteString("\n\n")
	}
}

func GetClient() {
	fmt.Println(strings.Repeat("*", 30) + "GET" + strings.Repeat("*", 30))
	if resp, err := http.Get("http://127.0.0.1:5678/get?" + EncodeUrlParams(map[string]string{
		"name": "zcy",
		"age":  "18",
	})); err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("response status: %s\n", resp.Status)
		fmt.Println("response body:")
		// io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
		if body, err := io.ReadAll(resp.Body); err == nil {
			fmt.Print(string(body))
		}
		os.Stdout.WriteString("\n\n")
	}
}

func HugeBodyClient() {
	fmt.Println(strings.Repeat("*", 30) + "GET HUGE BODY" + strings.Repeat("*", 30))
	if resp, err := http.Get("http://127.0.0.1:5678/stream"); err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()

		fmt.Println("response header")
		for key, values := range resp.Header {
			fmt.Printf("%s: %v\n", key, values)
			if key == "Date" {
				if tm, err := http.ParseTime(values[0]); err == nil {
					fmt.Printf("server time %s\n", tm.Format("2006-01-02 15:04:05"))
				}
			}
		}

		headerKey := http.CanonicalHeaderKey("content-length") // 正规化之后是Content-Length
		if ls, exists := resp.Header[headerKey]; exists {
			if l, err := strconv.Atoi(ls[0]); err == nil {
				fmt.Printf("Content-Length=%d\n", l)
				total := 0
				reader := bufio.NewReader(resp.Body)
				for total < l {
					if bs, err := reader.ReadBytes('\n'); err == nil {
						total += len(bs)
						fmt.Printf("进度 %.2f%%, 内容 %s", 100*float64(total)/float64(l), string(bs)) // bs末尾包含了\n
					} else {
						if err == io.EOF {
							if len(bs) > 0 { // 即使读到末尾了，本次read也可能读出了内容
								total += len(bs)
								fmt.Printf("进度 %.2f%%, 内容 %s", 100*float64(total)/float64(l), string(bs))
							}
						} else {
							fmt.Printf("read response body error: %s\n", err)
						}
						break
					}
					if total >= l/2 {
						break
					}
				}
			}
		}

	}
}
