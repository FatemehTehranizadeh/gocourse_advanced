package main

import (
	"fmt"
	"sync"
	"time"
)

// func main(){
// 	var x int
// 	wg := sync.WaitGroup{}
// 	mutex := sync.Mutex{}
// 	now := time.Now()

// 	for i:=0; i < 20; i++ {
// 		wg.Add(1)
// 		go write(&x, &wg, &mutex)
// 	}

// 	for i:=0; i < 100; i++ {
// 		wg.Add(1)
// 		go read(&x, &wg, &mutex)
// 	}

// 	wg.Wait()

// 	fmt.Println("The final value of x is: ", x)
// 	fmt.Printf("%v is passed.\n", time.Since(now).Milliseconds()) //1051

// }

// func read(x *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
// 	defer wg.Done()
// 	mutex.Lock()
// 	time.Sleep(time.Millisecond * 10)
// 	// fmt.Println("The value of x is: ", *x)
// 	mutex.Unlock()
// }

// func write(x *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
// 	defer wg.Done()
// now := time.Now()
// 	mutex.Lock()
// 	*x++
// 	mutex.Unlock()
// fmt.Printf("%v is passed.\n", time.Since(now).Milliseconds())

// 	// fmt.Println("The value of x is: ", *x)
// }

func main() {
	var x int
	wg := sync.WaitGroup{}
	mutex := sync.RWMutex{}
	now := time.Now()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go write(&x, &wg, &mutex)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read(&x, &wg, &mutex)
	}

	wg.Wait()

	fmt.Println("The final value of x is: ", x)
	fmt.Printf("%v is passed.\n", time.Since(now).Milliseconds()) //20

}

func read(x *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()
	mutex.RLock()
	time.Sleep(time.Millisecond * 10)
	// fmt.Println("The value of x is: ", *x)
	mutex.RUnlock()
}

func write(x *int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	defer wg.Done()
	now := time.Now()

	mutex.Lock()
	*x++
	mutex.Unlock()
	fmt.Printf("%v is passed.\n", time.Since(now).Milliseconds())
	// fmt.Println("The value of x is: ", *x)
}

/* تفاوت mutex و RWMutex
میوتکس به طور کلی اجازه می‌ده که فقط یک گوروتین در آنِ واحد به داده‌ای دسترسی داشته باشه (چه بخواد بخونه، چه بخواد بنویسه).
وقتی یک گوروتین به داده‌ها دسترسی پیدا می‌کنه، بقیه گوروتین‌ها تا زمانی که کارش تموم بشه، نمی‌تونند به اون داده‌ها دسترسی داشته باشن.
یعنی میوتکس همیشه یک نفره هست!
RWMutex:

RWMutex دو حالت متفاوت داره:
خواندن (Read Lock): وقتی چند گوروتین می‌خوان، به همدیگه اجازه می‌ده که همزمان به داده‌ها دسترسی داشته باشن.
نوشتن (Write Lock): فقط یک گوروتین می‌تونه بنویسه و در این حالت، هیچ گوروتین دیگه‌ای نمی‌تونه همزمان بخونه یا بنویسه.
در حالت خواندن، RWMutex اجازه می‌ده که چندین گوروتین به داده‌ها دسترسی داشته باشن، اما در حالت نوشتن، دسترسی همه رو قطع می‌کنه.
خلاصه:
Mutex یک Lock کلی برای دسترسی به داده‌ها است، یعنی فقط یک گوروتین می‌تونه دسترسی داشته باشه.
RWMutex اجازه می‌ده که چند گوروتین همزمان بخونن، ولی وقتی یک گوروتین می‌خواد بنویسه، بقیه رو مسدود می‌کنه.
*/