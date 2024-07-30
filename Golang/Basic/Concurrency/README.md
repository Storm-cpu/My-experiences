# TABLE OF CONTENT
1. [ Goroutines ](#goroutines)
2. [ Channels ](#channels)
2. [ sync.WaitGroups ](#waitgroups)

<a name="goroutines"></a>
## Goroutines
### Concept 
- Goroutine là một luồng nhẹ trong go. Nó giống như một thread trong những ngôn ngữ khác nhưng nhẹ hơn nhiều.
- Goroutine được quản lý thông qua các công cụ đồng bộ hóa như channels và sync.WaitGroups.
### Syntax
``` go FunctionName() ```
### Example
```
package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(100 * time.Millisecond)
    }
}

func printLetters() {
    for i := 'A'; i <= 'E'; i++ {
        fmt.Printf("%c\n", i)
        time.Sleep(150 * time.Millisecond)
    }
}

func main() {
    go printNumbers() // Khởi chạy goroutine đầu tiên
    go printLetters() // Khởi chạy goroutine thứ hai

    // Đợi để các goroutine hoàn thành
    time.Sleep(1 * time.Second)
    fmt.Println("Main goroutine ends")
}
```
### Use Cases
- Khi một tác vụ có thể được chia thành nhiều luồng để thực hiện tốt hơn
- Khi thực hiện nhiều request đến các API khác nhau
- Chạy các background operations trong một chương trình
- Xử lý nhiều request đồng thời trong các luồng riêng biệt khi các request này không phụ thuộc lẫn nhau

<a name="channels"></a>
## Channels
### Concept
- Channel là một cấu trúc dữ liệu dùng để giao tiếp và đồng bộ hóa giữa các goroutine. 
- Channel cho phép một goroutine nhận và gữi giá trị từ một goroutine khác.
### Syntax
```
// Tạo channel
ch := make(chan int) // Tạo một unbuffered channel và truyền vào kiểu int
ch := make(chan int, 2) // Tạo một buffered channel có buffer với capacity là 2

// Gữi và nhận giá trị
ch <- 10 // Gửi giá trị 10 vào channel ch
value := <-ch // Nhận giá trị từ channel ch và gán cho biến value

// Đóng channel
close(ch)
```
### Example
```
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	fmt.Println("Waiting for value...")
	value := <-ch
	fmt.Printf("Received value: %d\n", value)
}
```
### Use Cases
- Dùng channel khi muốn giao tiếp giữa các goroutine
- Để đồng bộ hóa (Synchronization)

<a name="waitgroups"></a>
## sync.WaitGroups
### Concept
- sync.WaitGroups được dùng để chờ đợi một tập hợp các goroutine hoàn thành công việc giúp đồng bộ hóa các goroutine
### Syntax
```
// Khai báo một "WaitGroup"
var wg sync.WaitGroup

// Thêm n goroutine vào danh sách chờ
wg.Add(n)

// Trong mỗi goroutine, gọi phương thức Done khi công việc hoàn thành để giảm số lượng goroutine đang chờ đợi đi 1.
wg.Done()

// Chờ tất cả goroutine hoàn thành
wg.Wait()
```
### Example
```
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Thông báo khi goroutine hoàn thành
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second) // Giả lập công việc bằng cách ngủ 1 giây
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1) // Tăng số lượng goroutine cần chờ đợi
        go worker(i, &wg) // Khởi chạy goroutine
    }

    wg.Wait() // Chờ tất cả các goroutine hoàn thành
    fmt.Println("All workers done")
}
```
### Use Cases
- Chờ tất cả goroutine hoàn thành trước khi tiếp tục
- Xử lý các tác vụ chạy backgroud
