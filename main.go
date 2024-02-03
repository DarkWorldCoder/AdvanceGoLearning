// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	fmt.Println("Hello World")

// 	// argsWithProg := os.Args
// 	// argsWithOutProg := os.Args[1:]
// 	// arg := os.Args[3]

// 	// fmt.Println(argsWithProg)
// 	// fmt.Println(argsWithOutProg)
// 	// fmt.Println(arg)

// 	// Convert string to int
// 	num1, err := strconv.ParseInt(os.Args[1], 10, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	num2, err := strconv.ParseInt(os.Args[2], 10, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(num1 + num2)
// }

// package main

// import "fmt"

// func doit(operator func(int, int) int, a int, b int) int {
// 	return operator(a, b)
// }

// func sum(a int, b int) int {
// 	return a + b
// }

// func multiply(a int, b int) int {
// 	return a * b
// }

// func main() {
// 	c := doit(sum, 5, 6)
// 	fmt.Println(c)
// 	c = doit(multiply, 5, 6)
// 	fmt.Println(c)

// }

// closure

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// func accumulator(increment int) func() int {
// 	i := 0
// 	return func() int {
// 		i += increment
// 		return i
// 	}
// }

// func main() {
// 	a := accumulator(1)
// 	b := accumulator(2)

// 	fmt.Println("a", "b")
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(a(), b())
// 	}
// }

// func a(i int) {
// 	i = 0
// }

// func b(i *int) {
// 	*i = 0
// }

// func main() {
// 	x := 5
// 	a(x)
// 	fmt.Println(x)
// 	b(&x)
// 	fmt.Println(x)
// }

// func main() {
// 	var a int
// 	fmt.Println(a)

// 	var b *int
// 	fmt.Println(b)

// 	var c bool
// 	fmt.Println(c)

// 	var d func()
// 	fmt.Println(d)

// 	var e string
// 	fmt.Println(e)
// }

// var Musketeers = []string{
// 	"Athos", "Porthos", "Aramis", "D'Artagnan",
// }

// func GetMusketeer(id int) (string, error) {
// 	if id < 0 || id >= len(Musketeers) {
// 		return "", errors.New("Invalid id")
// 	}
// 	return Musketeers[id], nil
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	id := rand.Int() % 6

// 	mosq, err := GetMusketeer(id)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(mosq)
// }

// type Rectangle struct {
// 	Height int
// 	Width  int
// }

// func (r Rectangle) Enlarge(factor int) {
// 	r.Height *= factor
// 	r.Width *= factor
// }

// func (r *Rectangle) Scale(factor int) {
// 	r.Height *= factor
// 	r.Width *= factor
// }

// func main() {
// 	r := Rectangle{10, 5}
// 	r.Enlarge(10)
// 	fmt.Println(r)
// 	r.Scale(10)
// 	fmt.Println(r)
// }

// type Rectangle struct {
// 	Height int
// 	Width  int
// }

// func (r Rectangle) Surface() int {
// 	return r.Height * r.Width
// }

// type Box struct {
// 	Rectangle
// 	depth int
// }

// func (b Box) Volume() int {
// 	return b.Surface() * b.depth
// }

// func main() {
// 	b := Box{Rectangle{10, 5}, 3}
// 	fmt.Printf("%+v\n", b)
// 	fmt.Println(b.Volume())

// }

// type T struct {
// 	B int
// 	C string
// }

// type S struct {
// 	C string
// 	D T
// 	E map[string]int
// }

// func printerReflect(offset int, typeOfX reflect.Type) {
// 	indent := strings.Repeat(" ", offset)
// 	fmt.Printf("%s %:s %s\n", indent, typeOfX.Kind(), typeOfX.Name())

// 	if typeOfX.Kind() == reflect.Struct {
// 		for i := 0; i < typeOfX.NumField(); i++ {
// 			innerIndent := strings.Repeat(" ", offset+4)
// 			f := typeOfX.Field(i)
// 			if f.Type.Kind() == reflect.Struct {
// 				fmt.Printf("%s %s:\n", innerIndent, f.Name)
// 				printerReflect(offset+4, f.Type)
// 			} else {
// 				fmt.Printf("%s %s %s\n", innerIndent, f.Name, f.Type.Name())
// 			}

// 		}
// 	}

// 	fmt.Printf("%s } %s\n", indent)
// }

// func main() {
// 	X := S{
// 		"root",
// 		T{42, "forty two"},
// 		make(map[string]int),
// 	}

// 	printerReflect(0, reflect.TypeOf(X))
// }

// func ShowIt() {
// 	for {
// 		time.Sleep(time.Millisecond * 100)
// 		fmt.Println("Here it is!!")
// 	}
// }

// func main() {
// 	go ShowIt()

// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Millisecond * 250)
// 		fmt.Println("Where is it?")
// 	}
// }

// func ShowIt(t time.Duration, num int) {
// 	for {
// 		time.Sleep(t)
// 		fmt.Println("Here it is!!")
// 	}
// }

// func main() {
// 	go ShowIt(time.Second, 100)
// 	go ShowIt(time.Millisecond*100, 50)
// 	go ShowIt(time.Millisecond*10, 10)

// 	time.Sleep(time.Second * 5)
// }

// func generator(ch chan int) {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		time.Sleep(time.Millisecond * 500)
// 		sum += i
// 	}
// 	ch <- sum
// }

// func main() {
// 	ch := make(chan int)
// 	go generator(ch)
// 	fmt.Println("main waits for result…")
// 	result := <-ch

// 	fmt.Println(result)
// }

// func generator(ch chan int) {
// 	fmt.Println("generator waits for n")
// 	n := <-ch

// 	fmt.Println("n is", n)
// 	sum := 0
// 	for i := 0; i < n; i++ {
// 		sum += i
// 	}
// 	ch <- sum

// }

// func main() {
// 	ch := make(chan int)
// 	go generator(ch)
// 	fmt.Println("main waits for result…")
// 	ch <- 10
// 	result := <-ch
// 	fmt.Println(result)
// }

// func MrA(ch chan string) {
// 	time.Sleep(time.Millisecond * 200)
// 	ch <- "Hello"
// }

// func MrB(ch chan string) {
// 	time.Sleep(time.Millisecond * 10)
// 	ch <- "World"
// }

// func main() {
// 	ch := make(chan string)

// 	go MrA(ch)
// 	go MrB(ch)

// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)

// }

// func sender(out chan int) {
// 	for i := 0; i < 10; i++ {
// 		time.Sleep(time.Millisecond * 500)
// 		out <- i
// 	}
// 	close(out)
// 	fmt.Println("sender is done")
// }

// func main() {
// 	ch := make(chan int)

// 	go sender(ch)

// 	for {
// 		num, found := <-ch
// 		if found {
// 			fmt.Println("Received", num)
// 		} else {
// 			fmt.Println("Channel closed!")
// 			break
// 		}
// 	}
// }

// func receiver(input <-chan int) {
// 	fmt.Print(input)
// 	for i := range input {
// 		fmt.Println(i)
// 	}
// }

// func sender(output chan<- int, n int) {
// 	for i := 0; i < n; i++ {
// 		time.Sleep(time.Millisecond * 500)
// 		output <- i
// 	}
// 	close(output)
// }

// func main() {
// 	ch := make(chan int)

// 	go sender(ch, 10)
// 	go receiver(ch)

// 	time.Sleep(time.Second * 5)
// 	fmt.Println("main is done")
// }

// func sendNumber(out chan int) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Millisecond * 300)
// 		out <- i
// 	}
// 	fmt.Println("sender is done")
// }

// func sendMsgs(out chan string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Millisecond * 100)
// 		out <- strconv.Itoa(i)
// 	}
// 	fmt.Println("no more msgs")
// }

// func main() {

// 	numbers := make(chan int)
// 	msgs := make(chan string)

// 	go sendNumber(numbers)
// 	go sendMsgs(msgs)

// 	for i := 0; i < 10; i++ {
// 		select {
// 		case num := <-numbers:
// 			fmt.Println("Received number", num)
// 		case msg := <-msgs:
// 			fmt.Println("Received msg", msg)
// 		}
// 	}
// }

// func sendNumbers(out chan int) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Millisecond * 300)
// 		out <- i
// 	}

// 	fmt.Println("No more numbers")
// 	close(out)
// }

// func sendMsgs(out chan string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Millisecond * 100)
// 		out <- strconv.Itoa(i)
// 	}

// 	fmt.Println("No more msgs")
// 	close(out)
// }

// func main() {
// 	numbers := make(chan int)
// 	msgs := make(chan string)

// 	go sendNumbers(numbers)

// 	go sendMsgs(msgs)

// 	closedNums, closedMsgs := false, false

// 	for !closedNums || !closedMsgs {
// 		select {
// 		case num, found := <-numbers:
// 			if found {
// 				fmt.Println("Received number", num)
// 			} else {
// 				closedNums = true
// 			}
// 		case msg, found := <-msgs:
// 			if found {
// 				fmt.Println("Received msg", msg)
// 			} else {
// 				closedMsgs = true
// 			}
// 		}
// 	}
// }

// func main() {
// 	ch := make(chan int)

// 	select {
// 	case i := <-ch:
// 		fmt.Println("Received", i)
// 	default:
// 		fmt.Println("No value received")
// 	}

// 	select {
// 	case ch <- 42:
// 		fmt.Println("Sent value 42")
// 	default:
// 		fmt.Println("No value sent")
// 	}

// 	select {
// 	case i := <-ch:
// 		fmt.Println("Received", i)
// 	default:
// 		fmt.Println("No value received")
// 	}
// }

// func generator(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Millisecond * 500)
// 		ch <- rand.Int()
// 	}
// 	close(ch)
// 	fmt.Print("generator is done")

// }

// func consumer(id int, sleep time.Duration, ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for task := range ch {
// 		time.Sleep(sleep)
// 		fmt.Printf("Consumer %d finished task %d\n", id, task)
// 	}
// 	fmt.Printf("Consumer %d is done\n", id)
// }

// func main() {

// 	ch := make(chan int, 10)
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go generator(ch, &wg)
// 	go consumer(1, 400, ch, &wg)
// 	go consumer(2, 100, ch, &wg)

// 	wg.Wait()
// }

// func worker(x *int) {
// 	for {
// 		time.Sleep(time.Millisecond * 100)
// 		*x += 1
// 	}
// }

// func main() {
// 	timer := time.NewTimer(time.Second * 2)
// 	ticker := time.NewTicker(time.Second)

// 	x := 0

// 	go worker(&x)

// 	for {
// 		select {
// 		case <-timer.C:
// 			fmt.Println("time -> %d\n", x)
// 		case <-ticker.C:
// 			fmt.Printf("ticket -> %d\n", x)
// 		}
// 		if x >= 10 {
// 			break
// 		}
// 	}

// }

// func reaction(t *time.Ticker) {
// 	for {
// 		select {
// 		case x := <-t.C:
// 			fmt.Println("Tick at", x)
// 		}
// 	}
// }

// func slowReaction(t *time.Timer) {
// 	select {
// 	case x := <-t.C:
// 		fmt.Println("Tick at", x)

// 	}
// }

// func main() {
// 	quick := time.NewTicker(time.Second)
// 	slow := time.NewTimer(time.Second * 5)
// 	stopper := time.NewTimer(time.Second * 4)

// 	go reaction(quick)
// 	go slowReaction(slow)
// 	<-stopper.C
// 	quick.Stop()

// 	stopped := slow.Stop()

// 	fmt.Println("slow timer stopped?", stopped)

// }

// func setter(id int, c *int32, ctx context.Context) {
// 	fmt.Printf("Setter %d started\n", id)
// 	t := time.NewTicker(time.Millisecond * 10)
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Printf("Setter %d is done\n", id)
// 			return
// 		case <-t.C:
// 			atomic.AddInt32(c, 1)
// 		}
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	var c int32 = 0
// 	for i := 0; i < 10; i++ {
// 		go setter(i, &c, ctx)
// 		// fmt.Println("Setter", i, "started")
// 	}

// 	time.Sleep(time.Second * 1)
// 	fmt.Println("Counter:", c)
// 	cancel()
// 	time.Sleep(time.Second)
// }

// func work(i int, info chan int) {
// 	t := time.Duration(i*100) * time.Millisecond
// 	time.Sleep(t)
// 	info <- i
// }

// func main() {
// 	d := time.Millisecond * 1000

// 	ch := make(chan int)

// 	i := 0

// 	for {
// 		ctx, cancel := context.WithTimeout(context.Background(), d)
// 		go work(i, ch)
// 		select {
// 		case x := <-ch:
// 			fmt.Println("work", x, "done")
// 		case <-ctx.Done():
// 			fmt.Println("work timed out")
// 		}

// 		if ctx.Err() != nil {
// 			cancel()
// 			return
// 		}
// 		cancel()
// 		i++

// 	}
// }

// func accum(c *uint32, ctx context.Context) {
// 	t := time.NewTicker(time.Millisecond * 250)

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Accumulator is done")
// 			return
// 		case <-t.C:
// 			atomic.AddUint32(c, 1)
// 		}
// 	}
// }

// func main() {
// 	d := time.Now().Add(time.Second * 3)
// 	ctx, cancel := context.WithDeadline(context.Background(), d)
// 	defer cancel()

// 	var counter uint32 = 0
// 	for i := 0; i < 10; i++ {
// 		go accum(&counter, ctx)
// 	}

// 	<-ctx.Done()
// 	fmt.Println("Counter:", counter)

// }

// var first int

// func setter(i int, ch chan bool, once *sync.Once) {
// 	t := rand.Uint32() % 300
// 	time.Sleep(time.Duration(t) * time.Millisecond)
// 	once.Do(func() {
// 		first = i
// 	})

// 	ch <- true
// 	fmt.Println("Setter", i, "done")

// }

// func main() {

// 	var once sync.Once

// 	ch := make(chan bool)

// 	for i := 0; i < 10; i++ {
// 		go setter(i, ch, &once)
// 	}

// 	for i := 0; i < 10; i++ {
// 		<-ch
// 	}
// 	fmt.Println("First:", first)

// }

// func writer(x map[int]int, factor int, m *sync.Mutex) {
// 	i := 1

// 	for {
// 		time.Sleep(time.Second)
// 		m.Lock()
// 		x[i] = x[i-1] * factor
// 		m.Unlock()
// 		i++
// 	}
// }

// func reader(x map[int]int, m *sync.Mutex) {
// 	for {
// 		time.Sleep(time.Second)
// 		// m.Lock()
// 		fmt.Println(x)
// 		// m.Unlock()
// 	}
// }

// func main() {
// 	x := make(map[int]int)
// 	x[0] = 1

// 	var m sync.Mutex

// 	go writer(x, 2, &m)
// 	go writer(x, 3, &m)
// 	go reader(x, &m)

// 	time.Sleep(time.Second * 10)
// 	go writer(x, 3, &m)

// 	time.Sleep(time.Second * 4)
// }

// func increaser(counter *int32) {
// 	for {
// 		atomic.AddInt32(counter, 2)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func decreaser(counter *int32) {
// 	for {
// 		atomic.AddInt32(counter, -1)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func main() {
// 	var counter int32 = 0
// 	go increaser(&counter)
// 	go decreaser(&counter)
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Millisecond * 500)
// 		fmt.Println(atomic.LoadInt32(&counter))
// 	}
// 	fmt.Print(atomic.LoadInt32(&counter))
// }

// type Monitor struct {
// 	ActiveUsers int
// 	Requests    int
// }

// func updater(monitor atomic.Value, m *sync.Mutex) {
// 	for {
// 		time.Sleep(time.Millisecond * 500)
// 		fmt.Println("Updating monitor")
// 		m.Lock()
// 		current := monitor.Load().(*Monitor)
// 		current.ActiveUsers += 100
// 		current.Requests += 1000
// 		monitor.Store(current)
// 		m.Unlock()
// 	}
// }

// func observe(monitor atomic.Value) {
// 	for {
// 		fmt.Println("Observing monitor")
// 		time.Sleep(time.Second)
// 		current := monitor.Load()
// 		fmt.Println(current)
// 	}
// }

// func main() {
// 	var monitor atomic.Value
// 	monitor.Store(&Monitor{ActiveUsers: 100, Requests: 1000})
// 	m := sync.Mutex{}

// 	go updater(monitor, &m)
// 	go observe(monitor)
// 	time.Sleep(time.Second * 10)
// }

// type MyReader struct {
// 	data string
// 	from int
// }

// type MyWriter struct {
// 	data string
// 	size int
// }

// func (r *MyReader) Read(p []byte) (int, error) {
// 	if p == nil {
// 		return -1, errors.New("nil target array")
// 	}
// 	if len(r.data) <= 0 || r.from == len(r.data) {
// 		return 0, io.EOF
// 	}

// 	n := len(r.data) - r.from
// 	if len(p) < n {
// 		n = len(p)
// 	}
// 	for i := 0; i < n; i++ {
// 		b := byte(r.data[r.from])
// 		p[i] = b
// 		r.from++
// 	}
// 	if r.from == len(r.data) {
// 		return n, io.EOF
// 	}
// 	return n, nil
// }

// func (mw *MyWriter) Write(p []byte) (int, error) {
// 	if len(p) == 0 {
// 		return 0, io.EOF
// 	}
// 	n := mw.size
// 	var err error = nil
// 	if len(p) < mw.size {
// 		n = len(p)
// 	} else {
// 		err = errors.New(" p is larger than size")
// 	}
// 	mw.data = mw.data + string(p[0:n])
// 	return n, err
// }

// func main() {
// 	target := make([]byte, 5)
// 	empty := MyReader{}
// 	n, err := empty.Read(target)
// 	fmt.Println(n, err, string(target))
// 	mr := MyReader{"This is to fight with the world", 4}
// 	mw := MyWriter{"hellow world", 5}
// 	_, errs := mw.Write(target)
// 	if errs != nil {
// 		fmt.Println(errs)
// 	}
// 	n, err = mr.Read(target)
// 	for err == nil {
// 		fmt.Println(n, err, string(target))
// 		n, err = mr.Read(target)
// 	}
// 	fmt.Println(n, err, string(target))

// }

// func main() {
// 	msg := "Save the world with go"
// 	filePath := "/tmp/msg"

// 	err := ioutil.WriteFile(filePath, []byte(msg), 0644)
// 	if err != nil {
// 		panic(err)
// 	}

// 	read, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("File content: %s\n", read)
// }


func main(){
	filePath := "/tmp/msg"

	msg := []string{
		"Rule","the","world"
	}

	f,err := os.Create(filePath)
	if err != nil{
		panic(err)
	}

	defer f.Close()

	for _,s := range msg{
		f.WriteString()
	}
}