# Table of Content
1. [ Repository Pattern ](#repository)
2. [ Unit of Work (UoW) ](#uow)

<a name="repository"></a>
# Repository Pattern
- Repository là lớp abstractin trung gian nằm giữa business logic và lớp data.
- Nó chứa các phương thức để thao tác giữa lớp business logic với lớp data.
- Mục đích của nó là tách biệt việc thao tác data ra khỏi lớp bussiness logic để cho những thay đổi không làm ảnh hưởng đến bussiness data.

<a name="uow"></a>
# Unit of Work (UoW)
- Unit of Work nó được xem như là một phiên làm việc. Nó sẽ theo dõi và quản lý các thay đổi của phiên làm việc đó được thực hiện thành công hay không hay không thì nó sẽ save, nếu không thì nó sẽ rollback lại.
- Unit of Work thường sẽ chứa các repo liên quan đến phiên làm việc đó.
- Nó được dùng khi mà mình thao tác trên nhiều loại đối tượng khác nhau trên một phiên làm việc nhằm giữ cho dữ liệu nhất quán.

