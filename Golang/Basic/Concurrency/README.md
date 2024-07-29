# TABLE OF CONTENT
1. [ Goroutine ](#goroutine)
2. [ Channels ](#channels)
2. [ sync.WaitGroups ](#waitgroups)

<a name="goroutine"></a>
## Goroutine
### Concept 
- Goroutine là một luồng nhẹ trong go. Nó giống như một thread trong những ngôn ngữ khác nhưng nhẹ hơn nhiều (5kb).
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

    // Đợi một chút để các goroutine hoàn thành
    time.Sleep(1 * time.Second)
    fmt.Println("Main goroutine ends")
}
```
- Trong ví dụ trên, hàm printNumbers và printLetters được chạy đồng thời như hai goroutine. Goroutine chính (main) đợi một chút để hai goroutine kia hoàn thành công việc của nó. 
### Use Cases
- Khi một tác vụ có thể được chia thành nhiều luồng để thực hiện tốt hơn
- Khi thực hiện nhiều request đến các API khác nhau
- Chạy các background operations trong một chương trình
- Xử lý nhiều request đồng thời trong các luồng riêng biệt khi các request này không phụ thuộc lẫn nhau

<a name="channels"></a>
## Channel

<a name="waitgroups"></a>
