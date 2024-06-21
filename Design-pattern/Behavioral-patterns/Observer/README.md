### Concept
Observer Pattern thiết kế cho các đối tượng (observers) đăng ký để nhận thông báo từ một đối tượng khác (subject), khi có sự thay đổi trạng thái.

Ví dụ: Khi chơi game sẽ có một nhân vật chính (subject) và mình muốn các hệ thống khác nhau trong trò chơi như điểm số, máu, kho đồ (observers) được cập nhật mỗi khi nhân vật chính thay đổi trạng thái, như là mất máu hoặc thu thập vật phẩm.

Lý do dùng observer pattern ở đây là để giữ cho mã nguồn gọn gàng và dễ bảo trì, bởi vì mình không cần phải viết mã cứng để thông báo cho từng hệ thống riêng lẻ mỗi khi có sự kiện xảy ra. Thay vào đó, các hệ thống tự đăng ký để nhận thông báo và tự xử lý chúng khi cần thiết.

### Structure
Structure của Observer pattern bao gồm các thành phần:

![observer_structure](../../access/observer_structure.png)

- Publisher (Subject): Đây là lớp chịu trách nhiệm phát đi thông báo đến các Subscriber khi có sự kiện hoặc thay đổi xảy ra. Trong hình ảnh, “Publisher” có phương thức notifySubscribers(), cho thấy nó sẽ thông báo cho tất cả các Subscriber đã đăng ký về sự thay đổi.
- Subscriber (Observer): Các lớp này đăng ký nhận thông báo từ Publisher và phản ứng lại với thông tin được cung cấp. Mỗi Subscriber có phương thức update(), cho thấy chúng sẽ cập nhật thông tin dựa trên thông báo từ Publisher.

### Example
```
// Subject interface
type Subject interface {
    RegisterObserver(o Observer)
    RemoveObserver(o Observer)
    NotifyObservers()
}

// Observer interface
type Observer interface {
    Update(subject Subject)
}

// GameCharacter là subject
type GameCharacter struct {
    observers []Observer
    health    int
}

func (g *GameCharacter) RegisterObserver(o Observer) {
    g.observers = append(g.observers, o)
}

func (g *GameCharacter) RemoveObserver(o Observer) {
    var indexToRemove int
    for i, observer := range g.observers {
        if observer == o {
            indexToRemove = i
            break
        }
    }
    g.observers = append(g.observers[:indexToRemove], g.observers[indexToRemove+1:]...)
}

func (g *GameCharacter) NotifyObservers() {
    for _, observer := range g.observers {
        observer.Update(g)
    }
}

func (g *GameCharacter) TakeDamage(amount int) {
    g.health -= amount
    g.NotifyObservers()
}

// HealthSystem là một observer
type HealthSystem struct{}

func (h *HealthSystem) Update(subject Subject) {
    // Cập nhật hệ thống sức khỏe dựa trên trạng thái của nhân vật chính
    if character, ok := subject.(*GameCharacter); ok {
        fmt.Printf("HealthSystem: Character has %d health remaining\n", character.health)
    }
}

func main() {
    character := &GameCharacter{health: 100}
    healthSystem := &HealthSystem{}

    character.RegisterObserver(healthSystem)

    // Nhân vật chính nhận sát thương và hệ thống sức khỏe được thông báo
    character.TakeDamage(10)
}
```

### Applicability
Observer pattern thường được áp dụng khi:

- Sử dụng khi một số đối tượng trong ứng dụng phải quan sát những đối tượng khác nhưng chỉ trong một khoảng thời gian giới hạn hoặc trong các trường hợp cụ thể.
- Khi muốn một đối tượng có khả năng thông báo cho một nhóm các đối tượng khác về sự thay đổi của nó mà không cần quan tâm đến việc nhóm đó gồm những đối tượng nào.

### Pros and Cons
Ưu điểm của Observer:

- Dễ dàng thêm hoặc bớt subscribers mà không ảnh hưởng đến publisher.
- Tạo một mối quan hệ một-đến-nhiều giữa các đối tượng, sao cho khi một đối tượng thay đổi trạng thái, tất cả phụ thuộc của nó được thông báo và cập nhật tự động.
- Giảm sự phụ thuộc lẫn nhau giữa các đối tượng, làm cho mã nguồn dễ bảo trì và mở rộng hơn.
- Cho phép các đối tượng có thể tương tác với nhau mà không cần biết chi tiết về nhau, tăng tính mô-đun và tái sử dụng mã nguồn.

Nhược điểm của Observer:

- Cần quản lý đăng ký và hủy đăng ký cẩn thận để tránh rò rỉ bộ nhớ.
- Nếu có nhiều subscribers, việc thông báo có thể chậm và tốn kém về hiệu suất.
- Thứ tự thông báo đến các subscribers có thể khó kiểm soát, dẫn đến khó khăn trong việc debug và bảo trì.