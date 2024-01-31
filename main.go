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
	"math/rand"
	"sync"
	"time"
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

func generator(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		ch <- rand.Int()
	}
	close(ch)
	fmt.Print("generator is done")

}

func consumer(id int, sleep time.Duration, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range ch {
		time.Sleep(sleep)
		fmt.Printf("Consumer %d finished task %d\n", id, task)
	}
	fmt.Printf("Consumer %d is done\n", id)
}

func main() {

	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(3)
	go generator(ch, &wg)
	go consumer(1, 400, ch, &wg)
	go consumer(2, 100, ch, &wg)

	wg.Wait()
}
