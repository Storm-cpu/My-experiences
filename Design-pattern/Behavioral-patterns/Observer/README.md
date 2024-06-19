### Concept
Observer Pattern tạo điều kiện cho một cách để một đối tượng (gọi là “subject”) thông báo và tự động cập nhật tất cả các đối tượng phụ thuộc (gọi là “observers”) về bất kỳ thay đổi nào trong trạng thái của nó.

### Structure
Structure của Observer pattern bao gồm các thành phần:

![observer_structure](../../access/observer_structure.png)

- Publisher (Subject): Đây là lớp chịu trách nhiệm phát đi thông báo đến các Subscriber khi có sự kiện hoặc thay đổi xảy ra. Trong hình ảnh, “Publisher” có phương thức notifySubscribers(), cho thấy nó sẽ thông báo cho tất cả các Subscriber đã đăng ký về sự thay đổi.
- Subscriber (Observer): Các lớp này đăng ký nhận thông báo từ Publisher và phản ứng lại với thông tin được cung cấp. Mỗi Subscriber có phương thức update(), cho thấy chúng sẽ cập nhật thông tin dựa trên thông báo từ Publisher.

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