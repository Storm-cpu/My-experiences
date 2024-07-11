# Mục lục
1. [ Repository Pattern ](#repository)
2. [ Unit of Work (UoW) ](#uow)
3. [ Object-relational Mapping (ORM) ](#orm)
4. [ DTO, DAO and Entity ](#dde)
5. [ ACID ](#acid)


<a name="repository"></a>
# Repository Pattern
- Repository là lớp abstractin trung gian nằm giữa business logic và lớp data
- Nó chứa các phương thức để thao tác giữa lớp business logic với lớp data
- Mục đích của nó là tách biệt việc thao tác data ra khỏi lớp bussiness logic để cho những thay đổi không làm ảnh hưởng đến bussiness data

<a name="uow"></a>
# Unit of Work (UoW)
- UoW nó được xem như là một phiên làm việc. Nó sẽ theo dõi và quản lý các thay đổi của phiên làm việc đó được thực hiện thành công hay không hay không thì nó sẽ save, nếu không thì nó sẽ rollback lại
- UoW thường sẽ chứa các repo liên quan đến phiên làm việc đó
- Nó được dùng khi mà mình thao tác trên nhiều loại đối tượng khác nhau trên một phiên làm việc nhằm giữ cho dữ liệu nhất quán

<a name="orm"></a>
# Object-relational Mapping (ORM)
- ORM giúp ánh xạ các đối tượng trong code sang các đối tượng trong csdl
- ORM giúp thực hiện các truy vấn bằng các hàm có sẵng mà không cần viết code SQl
- Nó giúp việc viết code nhanh hơn, dễ hiểu dễ đọc hơn. Nó cũng giúp xử lý các ký tự đặt biệt tránh việc bị tấn công SQL Injection

<a name="dde"></a>
# DTO, DAO and Entity
## Data Transfer Object (DTO)
- DTO dùng để đóng gói dữ liệu và gữi qua các tầng của ứng dụng
- Chỉ chứa một số trường dữ liệu và tránh chứa các thông tin nhạy cảm và không cần thiết khi gữi đi
- Giảm bớt dữ liệu khi truyền tải vì nó sẽ chứa ít thông tin hơn
## Data Access Object (DAO)
- DAO nằm giữa business logic và data như repository
- Nó tập trung truy cập vào một đối tượng cụ thể và thực hiện các truy vấn đơn giản hơn là repository
- DAO tách logic nghiệp vụ ra khỏi các thành phần khác của ứng dụng
## Entity
- Entity là các class tương ứng với các table trong db và có thể map vào db được

<a name="acid"></a>
# ACID
1. Atomicity: Toàn bộ các công việc trong transaction sẽ được thực hiện toàn bộ hoặc không thực hiện cái nào
2. Consistency: Khi một transaction được thực hiện thành công thì csdl vẫn trong trạng thái hợp lệ, vẫn tuân theo các tính chất sẵn có.
3. Isolation: Các transaction sẽ không ảnh hưởng lẫn nhau khi thực hiện đồng thời
4. Durability: Khi transaction được commit thì những thay đổi sẽ được lưu và back up lại được khi gặp sự cố

