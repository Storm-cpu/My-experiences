### Concept
Strategy Pattern cho phép định nghĩa một nhóm các thuật toán, đóng gói từng thuật toán thành một đối tượng, và làm cho chúng có thể thay thế cho nhau.

Ví dụ: Trong một ứng dụng thanh toán, mình có thể có nhiều phương thức thanh toán khác nhau như PayPal, Credit Card, Bitcoin, v.v. Mỗi phương thức thanh toán này có thể được triển khai như một strategy riêng biệt và có thể được chọn tùy theo ngữ cảnh sử dụng

### Structure
Structure của Strategy pattern bao gồm các thành phần:

![strategy_structure](../../access/strategy_structure.png)

- Context: Là lớp chứa một tham chiếu đến một đối tượng Strategy. Nó có thể thay đổi đối tượng Strategy tại runtime và cho phép client chọn thuật toán cần thiết thông qua phương thức setStrategy.
- Strategy Interface: Là interface định nghĩa một nhóm các thuật toán. Mỗi thuật toán được đóng gói và làm cho chúng có thể thay thế lẫn nhau. Interface này có phương thức execute mà các ConcreteStrategies sẽ triển khai.
- ConcreteStrategies: Đây là các lớp triển khai interface Strategy và cung cấp cài đặt cụ thể cho phương thức execute. Mỗi ConcreteStrategy sẽ có một thuật toán riêng biệt.
- Client: Đây là lớp sử dụng lớp Context và quyết định sử dụng ConcreteStrategy nào. Nó tạo ra đối tượng Strategy mới và thiết lập nó trong Context bằng phương thức setStrategy. Sau đó, nó gọi phương thức doSomething của Context để thực hiện thuật toán.

### Example
```
// Strategy Interface
type PaymentStrategy interface {
    Pay(amount float32)
}

// Concrete Strategies
type CreditCardPayment struct{}
func (c *CreditCardPayment) Pay(amount float32) {
    fmt.Printf("Paid %f using Credit Card\n", amount)
}

type PaypalPayment struct{}
func (p *PaypalPayment) Pay(amount float32) {
    fmt.Printf("Paid %f using Paypal\n", amount)
}

// Context
type PaymentContext struct {
    strategy PaymentStrategy
}
func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
    p.strategy = strategy
}
func (p *PaymentContext) Pay(amount float32) {
    p.strategy.Pay(amount)
}

// Client code
func main() {
    payment := PaymentContext{}
    payment.SetStrategy(&CreditCardPayment{})
    payment.Pay(22.30)

    payment.SetStrategy(&PaypalPayment{})
    payment.Pay(17.50)
}
```

### Applicability
Strategy pattern thường được áp dụng khi:

- Có nhiều class chỉ khác nhau về hành vi. Strategy cho phép cấu hình một class với nhiều hành vi khác nhau.
- Cần tránh sự phụ thuộc vào các phép toán có điều kiện. Thay vì nhiều điều kiện, mỗi hành vi được đóng gói trong một Strategy riêng biệt.
- Khi muốn thay đổi thuật toán trong runtime.

### Pros and Cons
Ưu điểm của Strategy pattern:

- Tăng cường sự linh hoạt và tái sử dụng mã nguồn.
- Có thể thay đổi hành vi của một class mà không cần sửa đổi mã nguồn.
- Tách biệt logic nghiệp vụ và chi tiết cài đặt thuật toán.

Nhược điểm của Strategy pattern:

- Phức tạp hóa mã nguồn do sự gia tăng số lượng class và đối tượng.
- Cần phải quản lý đối tượng Strategy, có thể dẫn đến tăng chi phí về bộ nhớ nếu không được quản lý cẩn thận.