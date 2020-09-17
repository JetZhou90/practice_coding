package main

/*
#include <windows.h>
int KeyDown(int key) {
    if (key > 96 && key < 123)  key -= 32;
    return (GetKeyState(key) < 0) ? 1 : 0;
}
*/
// import "C"
import (
	"fmt"
	// "golang.org/x/net/html"
	// "github.com/go-vgo/robotgo"
	// "image"
	"image/color"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path/filepath"
	// "practice/btree"
	// "practice/linklist"
	// "practice/stack_queue"
	"practice/crawl"
	// "practice/webapplication"
	"sync"
	// "time"
	// tg "github.com/galeone/tfgo"
	// "github.com/galeone/tfgo/image"
	// "github.com/galeone/tfgo/image/filter"
	// "github.com/galeone/tfgo/image/padding"
	// tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func main() {
	// webapplication.Test()
	crawl.Test()

	// root := tg.NewRoot()
	// grayImg := image.Read(root, "xx.png", 1)
	// grayImg = grayImg.Scale(0, 255)

	// // Edge detection using sobel filter: convolution
	// Gx := grayImg.Clone().Convolve(filter.SobelX(root), image.Stride{X: 1, Y: 1}, padding.SAME)
	// Gy := grayImg.Clone().Convolve(filter.SobelY(root), image.Stride{X: 1, Y: 1}, padding.SAME)
	// convoluteEdges := image.NewImage(root.SubScope("edge"), Gx.Square().Add(Gy.Square().Value()).Sqrt().Value()).EncodeJPEG()

	// Gx = grayImg.Clone().Correlate(filter.SobelX(root), image.Stride{X: 1, Y: 1}, padding.SAME)
	// Gy = grayImg.Clone().Correlate(filter.SobelY(root), image.Stride{X: 1, Y: 1}, padding.SAME)
	// correlateEdges := image.NewImage(root.SubScope("edge"), Gx.Square().Add(Gy.Square().Value()).Sqrt().Value()).EncodeJPEG()

	// results := tg.Exec(root, []tf.Output{convoluteEdges, correlateEdges}, nil, &tf.SessionOptions{})

	// file, _ := os.Create("convolved.png")
	// file.WriteString(results[0].Value().(string))
	// file.Close()

	// file, _ = os.Create("correlated.png")
	// file.WriteString(results[1].Value().(string))
	// file.Close()
	// a := [][]int{{0}}
	// fmt.Println(a)
	// var b []int = []int{1, 2, 3, 4}
	// var c []int = []int{3, 4}
	// a = append(a, b)
	// fmt.Println(a)
	// a = append(a, c)
	// fmt.Println(a)
	// btree.Test()
	// webapplication.Test()
	// Pipline(10)
	// var cmb = bank{}
	// go func() {
	// 	Deposit(&cmb, 200)              // A1
	// 	fmt.Println("=", Balance(&cmb)) // A2
	// }()
	// go Deposit(&cmb, 200)
	// iter_dir(nil)
	// exit_go()
	// launch(1000000)
	// naturals := make(chan int)
	// squares := make(chan int)
	// go counter(naturals, 10)
	// go squarer(squares, naturals)
	// printer(squares)
	// Pipline(10)
	// var cp ColoredPoint
	// cp.X = 1
	// fmt.Println(cp.Point.X) // "1"
	// cp.Point.Y = 2
	// fmt.Println(cp.Y) // "
	// red := color.RGBA{255, 0, 0, 255}
	// blue := color.RGBA{0, 0, 255, 255}
	// var p = ColoredPoint{Point{1, 1}, red}
	// var q = ColoredPoint{Point{5, 4}, blue}
	// fmt.Println(p.Distance(q.Point)) // "5"
	// r := &Point{1, 2}
	// r.ScaleBy(2)
	// fmt.Println(*r)
	// p := Point{1, 2}
	// (&p).ScaleBy(2)
	// fmt.Println(p) // "{2, 4}"
	// a := squares(4)
	// fmt.Println(a())
	// fmt.Println(a())
	// fmt.Println(sum(1, 2))
	// url := "http://www.baidu.com"
	// Url_get(url)
	// perim := Path{
	// 	{1, 1},
	// 	{5, 1},
	// 	{5, 4},
	// 	{1, 1},
	// }
	// fmt.Println(perim.Distance())
}

/*
函数定义
*/

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/*
匿名函数
*/
func squares(val int) func() int {
	var x int = val
	return func() int {
		x++
		return x * x
	}
}

/*
可变参数
*/
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

/*
Deferred函数
Panic异常
*/
func Url_get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

/*
Recover捕获异常
*/
// func soleTitle(doc *html.Node) (title string, err error) {
// 	type bailout struct{}
// 	defer func() {
// 		switch p := recover(); p {
// 		case nil: // no panic
// 		case bailout{}: // "expected" panic
// 			err = fmt.Errorf("multiple title elements")
// 		default:
// 			panic(p) // unexpected panic; carry on panicking
// 		}
// 	}()
// 	return "", nil
// }

/*
方法定义
*/
type Point struct{ X, Y float64 }

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

/*
基于指针对象的方法
*/
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

/*
通过嵌入结构体来扩展类型
*/
type ColoredPoint struct {
	Point
	Color color.RGBA
}

// func (p ColoredPoint) Distance(q Point) float64 {
// 	return p.Point.Distance(q)
// }

/*
	位运算

	&      位运算 AND
	|      位运算 OR
	^      位运算 XOR
	&^     位清空 (AND NOT)
	<<     左移
	>>     右移

	1<<N = 2^N			---------------	1 左移多少位等于2的多少次方
	1024>>N = 1024/2^N	---------------	右移N位 相当于除以2的N次方

*/

/*
接口
*/

/*
Channels

ch = make(chan int)    // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3

ch <- x  // a send statement
x = <-ch // a receive expression in an assignment statement
<-ch     // a receive statement; result is discarded


串联的Channels（Pipeline）
*/
func Pipline(max int) {
	naturals := make(chan int)
	squares := make(chan int)
	// Counter
	go func() {
		for x := 1; x < max; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

/*
单方向channel
*/
func counter(out chan<- int, max int) {
	for x := 0; x < max; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

/*
带缓存的channel
ch = make(chan string, 3)


并发的循环

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
    type item struct {
        thumbfile string
        err       error
    }

    ch := make(chan item, len(filenames))
    for _, f := range filenames {
        go func(f string) {
            var it item
            it.thumbfile, it.err = thumbnail.ImageFile(f)
            ch <- it
        }(f)
    }

    for range filenames {
        it := <-ch
        if it.err != nil {
            return nil, it.err
        }
        thumbfiles = append(thumbfiles, it.thumbfile)
    }

    return thumbfiles, nil
}

sync.WaitGroup方法

func makeThumbnails6(filenames <-chan string) int64 {
    sizes := make(chan int64)
    var wg sync.WaitGroup // number of working goroutines
    for f := range filenames {
        wg.Add(1)
		// worker

		将f的值作为一个显式的变量传给了函数，而不是在循环的闭包中声明
		防止循环变量快照问题
		确保使用的f是当go语句执行时的“当前”那个f

        go func(f string) {
            defer wg.Done()
            thumb, err := thumbnail.ImageFile(f)
            if err != nil {
                log.Println(err)
                return
            }
            info, _ := os.Stat(thumb) // OK to ignore error
            sizes <- info.Size()
        }(f)
    }

    // closer
    go func() {
        wg.Wait()
        close(sizes)
    }()

    var total int64
    for size := range sizes {
        total += size
    }
    return total
}

select的使用方法
**/

// func key_down(key string, abort *chan int) chan int {
// 	ok := int(C.KeyDown('q'))
// 	if ok == 1 {
// 		*abort <- 1
// 		return *abort
// 	}
// 	return *abort
// }

// func launch(max_num int) {
// 	fmt.Println("Commencing countdown.  Press q to abort.")
// 	tick := time.Tick(1 * time.Second)
// 	abort := make(chan int, 1)
// 	for countdown := max_num; countdown > 0; countdown-- {
// 		select {
// 		case <-tick:
// 			go fmt.Println(countdown)
// 			// Do nothing.
// 		case <-key_down("q", &abort):
// 			fmt.Println("Launch aborted!")
// 			return
// 		}
// 	}
// }

/*
并发目录遍历
*/
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func iter_dir(roots []string) {
	printDiskUsage := func(nfiles, nbytes int64) {
		fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
	}
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// Print the results.
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)

}

/*
并发退出
*/
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

/*
sync.Mutex互斥锁
Go的mutex不能重入
*/
var (
	mu sync.Mutex // guards balance
	// balance int
)

type bank struct {
	balance int
}

func Deposit(bankName *bank, amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(bankName, amount)
	// mu.Unlock()
}

// func Balance(bankName *bank) int {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	b := bankName.balance
// 	return b
// }

/*
当一个goroutine获得了一个互斥锁时，它能断定被互斥锁保护的变量正处于不变状态（译注：即没有其他代码块正在读写共享变量），
在其获取并保持锁期间，可能会去更新共享变量，这样不变性只是短暂地被破坏，然而当其释放锁之后，锁必须保证共享变量重获不变性并且多个goroutine按顺序访问共享变量。
尽管一个可以重入的mutex也可以保证没有其它的goroutine在访问共享变量，但它不具备不变性更深层含义。
一个通用的解决方案是将一个函数分离为多个函数，比如我们把Deposit分离成两个：一个不导出的函数deposit，
这个函数假设锁总是会被保持并去做实际的操作，另一个是导出的函数Deposit，这个函数会调用deposit，但在调用前会先去获取锁。
同理我们可以将Withdraw也表示成这种形式：
*/

func deposit(bankName *bank, amount int) { bankName.balance += amount }
func Withdraw(bankName *bank, amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(bankName, -amount)
	if bankName.balance < 0 {
		deposit(bankName, amount)
		return false // insufficient funds
	}
	return true
}

/*
sync.RWMutex读写锁
*/
var murw sync.RWMutex
var balance int

func Balance(bankName *bank) int {
	murw.RLock() // readers lock
	defer murw.RUnlock()
	return bankName.balance
}

/*
sync.Once惰性初始化
*/
var icons map[string]string
var loadIconsOnce sync.Once

// Concurrency-safe.
func Icon(name string) string {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
func loadIcons() {
	icons = map[string]string{
		"spades.png":   "spades.png",
		"hearts.png":   "hearts.png",
		"diamonds.png": "diamonds.png",
		"clubs.png":    "clubs.png",
	}
}

// func exit_go() {
// 	go func() {
// 		os.Stdin.Read(make([]byte, 1)) // read a single byte
// 		close(done)
// 	}()
// }
