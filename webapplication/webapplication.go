package webapplication

import (
	// "bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"html/template"
	// "io"
	"encoding/base64"
	// "gocv.io/x/gocv"
	// "image"
	// "image/color"
	"net/http"
	"os"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("webapplication/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		if err != nil {
			// handle error http.Error() for example
			log.Fatal("ParseForm: ", err)
		}

		// 请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form.Get("username"))
		fmt.Println("password:", r.Form.Get("password"))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20) // limit your max input length!

	err := r.ParseForm() // 解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
	if err != nil {
		// handle error http.Error() for example
		log.Fatal("ParseForm: ", err)
	}
	// in your case file would be fileupload
	image := r.Form.Get("image")
	dist, _ := base64.StdEncoding.DecodeString(image)
	fmt.Println(image)
	f, _ := os.OpenFile("xx.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
	return
}

func Test() {
	http.HandleFunc("/index", sayhelloName) // 设置访问的路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/receiver", ReceiveFile)
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	// xmlfile := "gocv/data/haarcascade_frontalface_default.xml"
	// webcam, _ := gocv.OpenVideoCapture(0)
	// window := gocv.NewWindow("Test")
	// img := gocv.NewMat()
	// classifier := gocv.NewCascadeClassifier()
	// classifier.Load(xmlfile)
	// blue := color.RGBA{0, 0, 255, 0}

	// for {
	// 	webcam.Read(&img)
	// 	rects := classifier.DetectMultiScale(img)
	// 	for _, r := range rects {
	// 		gocv.Rectangle(&img, r, blue, 3)
	// 		size := gocv.GetTextSize("Face", gocv.FontHersheyPlain, 1.2, 2)
	// 		pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
	// 		gocv.PutText(&img, "Face", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
	// 		window.IMShow(img)
	// 		if window.WaitKey(1) >= 0 {
	// 			break
	// 		}
	// 	}
	// }
}
