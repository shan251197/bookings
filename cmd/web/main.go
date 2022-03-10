package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/shan251197/bookings/internal/config"
	"github.com/shan251197/bookings/internal/handler"
	"github.com/shan251197/bookings/internal/models"
	"github.com/shan251197/bookings/internal/render"
)

const port = ":8000"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	gob.Register(models.Reservation{})

	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println("starting application on port", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}

// type Person struct {
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	HasDog    bool   `json:"-"`
// }

// func main() {
// 	myJson := `
// 	[{
// 		"first_name":"sam",
// 		"last_name":"jack"

// 	},
// 	{
// 		"first_name":"shan",
// 		"last_name":"s"

// 	}]`

// 	var unmarshalled []Person

// 	json.Unmarshal([]byte(myJson), &unmarshalled)

// 	// if err != nil {
// 	// 	log.Println("Error unmarshalled", err)
// 	// }

// 	log.Printf("unmarshalled: %+v", unmarshalled)

// 	var mySlice []Person

// 	var m1 Person
// 	m1.FirstName = "muthu"
// 	m1.LastName = "kumar"
// 	m1.HasDog = true

// 	mySlice = append(mySlice, m1)

// 	var m2 Person
// 	m2.FirstName = "suresh"
// 	m2.LastName = "a"
// 	m2.HasDog = true
// 	mySlice = append(mySlice, m2)

// 	newJson, err := json.MarshalIndent(mySlice, "", "")

// 	if err != nil {
// 		log.Println("error marshalling", err)
// 	}

// 	log.Println(string(newJson))
// }

// func main() {
// 	//channel
// 	numCh := make(chan int)
// 	oddToMerger := make(chan int)
// 	evenToSquare := make(chan int, 10)
// 	squareToMerger := make(chan int)

// 	mergerToPrint := make(chan int)
// 	done := make(chan struct{})

// 	//goroutine
// 	go counter(numCh)
// 	go oddEvenSplitter(numCh, oddToMerger, evenToSquare)
// 	go square(evenToSquare, squareToMerger)
// 	go merger(oddToMerger, squareToMerger, mergerToPrint)
// 	go printer(mergerToPrint, done)
// 	<-done

// }
// func counter(out chan int) {
// 	for i := 0; i < 10; i++ {
// 		out <- i
// 	}
// 	close(out)
// }
// func square(in chan int, out chan int) {
// 	for a := range in {
// 		// time.Sleep(1 * time.Second)
// 		out <- a * a

// 	}

// 	close(out)

// }
// func oddEvenSplitter(in chan int, odd chan int, even chan int) {
// 	for a := range in {
// 		if a%2 == 0 {
// 			even <- a
// 		} else {
// 			odd <- a
// 		}
// 	}
// 	close(even)
// 	close(odd)
// }
// func merger(oddIn chan int, evenIn chan int, out chan int) {

// 	i := 0

// 	for {
// 		select {
// 		case a, ok := <-oddIn:
// 			if ok {
// 				out <- a
// 			} else {
// 				i++
// 			}
// 		case b, ok := <-evenIn:
// 			if ok {
// 				out <- b
// 			} else {
// 				i++
// 			}
// 		}
// 		if i >= 2 {
// 			break
// 		}

// 	}
// 	close(out)

// }
// func printer(in chan int, done chan struct{}) {
// 	for a := range in {
// 		fmt.Println(a)
// 	}
// 	done <- struct{}{}
// }

// func delayPrint(d chan int) {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("hi")
// 	d <- 50000

// }

// func f(n int) {
// 	for i := 0; i < n; i++ {
// 		fmt.Println("f ", i)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func g(n int) {
// 	for i := 0; i < n; i++ {
// 		fmt.Println("g:", i)
// 		time.Sleep(time.Second * 1)
// 	}
// }
// func feedmonkey(c chan string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Feeding monkey with Banana", i)
// 		c <- fmt.Sprintf("banana %d", i)
// 	}
// }

// func monkeyeat(c chan string)

// func main() {
// 	// go f(5)
// 	// go g(10)
// 	var c chan string
// 	feedmonkey(c)
// 	banana := <-c
// 	fmt.Println(banana)
// 	// 	fmt.Println("Called a goroutine")
// 	// 	fmt.Println("Press return to terminate any time")
// 	// 	input := ""
// 	// 	fmt.Scanln(&input)
// 	// 	fmt.Println(input)
// }

// }

//

// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool
// 	Swap(i, j int)
// }

// type Rectangle struct {
// 	length int
// 	width  int
// }

// func (rect Rectangle) area() int {
// 	return rect.length * rect.width
// }

// func (rect Rectangle) String() string {
// 	return fmt.Sprintf("length: %d, width: %d, area: %d",
// 		rect.length, rect.width, rect.area())
// }

// type Rectangles []Rectangle

// func (rectangles Rectangles) Len() int {
// 	return len(rectangles)
// }

// func (rectangles Rectangles) Less(i, j int) bool {
// 	return rectangles[i].area() < rectangles[j].area()
// }
// func (rectangles Rectangles) Swap(i, j int) {
// 	rectangles[i], rectangles[j] = rectangles[j], rectangles[i]
// }

// func main() {

// 	rect := Rectangles{
// 		Rectangle{10, 6},
// 		Rectangle{10, 4},
// 		Rectangle{10, 8},
// 	}

// 	for _, v := range rect {
// 		log.Println(v)
// 	}

// 	log.Println("------------------------")

// 	sort.Sort(rect)

// 	for _, v := range rect {
// 		log.Println(v)
// 	}

// }

// type Payee interface {
// 	salary() float32
// }

// type Employee struct {
// 	IdNumber   string
// 	Name       string
// 	BaseSalary float32
// }

// type Staff struct {
// 	Employee
// 	Allowance float32
// }

// func main() {

// 	var staff1 = new(Staff)

// 	staff1.IdNumber = "E01"
// 	staff1.Name = "sam"
// 	staff1.BaseSalary = 10000
// 	staff1.Allowance = 1500
// 	// staff1 := Staff{
// 	// 	Employee:  Employee{"E02", "muthu", 11000},
// 	// 	Allowance: 1000,
// 	// }

// 	staff2 := Staff{
// 		Employee: Employee{
// 			IdNumber:   "E02",
// 			Name:       "suresh",
// 			BaseSalary: 15000,
// 		},
// 		Allowance: 5000,
// 	}
// 	staff3 := Staff{
// 		Employee:  Employee{"E02", "muthu", 11000},
// 		Allowance: 1000,
// 	}
// 	staff4 := Staff{
// 		Employee:  Employee{"E02", "balaji", 8000},
// 		Allowance: 1000,
// 	}

// 	var final []float32

// 	final = append(final,
// 		calculate(staff1),
// 		calculate(staff2),

// 		calculate(staff3),
// 		calculate(staff4))

// 	// calculate(staff1)
// 	// calculate(staff2)

// 	// calculate(staff3)
// 	// calculate(staff4)

// 	// log.Println(&staff1)
// 	log.Println(final)

// }

// func (staff Staff) salary() float32 {
// 	return staff.BaseSalary + staff.Allowance
// }
// func (staff Staff) EmployeeName() string {
// 	return staff.Name
// }

// func calculate(p Payee) float32 {
// 	return p.salary()
// }
