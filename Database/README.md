# Table of Content
1. [ SELECT ](#select)
2. [ WHERE ](#where)
3. [ ORDER BY ](#order_by)
4. [ PROCEDURE, FUNCTION, VIEW AND TRIGGER ](#procedure_function_view_trigger)
5. [ INDEX ](#index)
5. [ DELETE AND TRUNCATE ](#delete_truncate)


<a name="select"></a>
# SELECT

## Syntax
```
SELECT column1, column2, ... FROM table_name;
```

## Example
```
SELECT user_id FROM USER

SELECT user_id, age FROM USER

SELECT * FROM USER
```

<a name="where"></a>
# WHERE

## Syntax
```
SELECT column1, column2, ... FROM table_name
WHERE condition;
```

## Example
```
SELECT * FROM USER
WHERE user_id = 1

SELECT * FROM USER
WHERE username = "wgioan2"
```

<a name="order_by"></a>
# ORDER BY
- Mặc định sẽ sắp xếp ở thứ thự tăng dần

## Syntax
```
SELECT column1, column2, ...
FROM table_name
ORDER BY column1, column2, ... ASC|DESC;
```

## Example
```
SELECT * FROM USER
ORDER BY user_id;

SELECT * FROM USER
ORDER BY user_id DESC;

SELECT * FROM USER
ORDER BY user_id, username;
```

<a name="aon"></a>
# AND, OR, NOT

## Syntax
```
SELECT column1, column2, ... FROM table_name
WHERE condition1 AND/OR/NOT condition2 AND/OR/NOT condition3 ...;
```

## Example
```
SELECT * FROM USER
WHERE user_id > 10
OR age = 20
NOT country = "china";
```

<a name="procedure_function_trigger"></a>
# Procedure, Function and Trigger

## Procedure
### Syntax
```
create [or replace] procedure procedure_name(parameter_list)
language plpgsql
as $$
declare
-- variable declaration
begin
-- stored procedure body
end; $$
```

### Example
```
create or replace procedure checkUserExisted(input_id int)
language plpgsql
as $$
begin
    if not exists (select 1 from USER where user_id = input_id) then
        raise exception 'User with id % does not exist', input_id;
    end if;
end;
$$;
```

## Function
### Syntax
```
create [or replace] function function_name(parameter_list)
returns return_type
language plpgsql
as $$
declare
-- variable declaration
begin
-- function body
return return_value;
end; $$
```

### Example
```
create or replace function getUserAge(input_id int)
returns int
language plpgsql
as $$
declare
    user_age int;
begin
    select age into user_age from USER where user_id = input_id;
    if user_age is null then
        raise exception 'User with id % does not exist', input_id;
    end if;
    return user_age;
end;
$$;
```

## Trigger
### Syntax
```
--Create function first
create [or replace] function trigger_function_name()
returns trigger
language plpgsql
as $$
begin
-- trigger body
return new;
end; $$
--Then create trigger
create trigger trigger_name
{before|after|instead of} {event}
on table_name
for each row
execute function trigger_function_name();
```

### Example
```
--Create function first
create or replace function log_user_update()
returns trigger
language plpgsql
as $$
begin
    raise notice 'User with id % has been updated. Old username: %, New username: %', OLD.user_id, OLD.username, NEW.username;
    return new;
end;
$$;
--Then create trigger
create trigger user_update_logger
after update on USER
for each row
execute function log_user_update();
```

<a name="index"></a>
# INDEX

## Syntax
```
create index [if not exists] index_name
on table_name(column1, column2, ...);
```

## Example
```
create index if not exists gender_age_idx
on card(gender,age);
```

<a name="view_materialized_view"></a>
# View and Materialized View

## View
- Là table ảo dựa theo câu lệnh select
- Giúp thao tác nhanh với dữ liệu
- Nên dùng cho các truy vấn đơn giản và có khả năng tái sử dụng cao

### Syntax
```
create [or replace] view view_name as
select_statement;
```

### Example
```
create or replace view user_overview as
select user_id, username, country from USER;
```

## Materialized 
- Nó giống như view nhưng nó sẽ lưu lại data khi truy vấn
- Phải cập nhật lại data manualy bằng cách dùng refresh. `refresh materialized view view_name;`
- Cần thêm không gian lưu trữ bộ nhớ
- Giúp việc truy vấn data nhanh hơn
- Nên dùng cho những loại dữ liệu ít được update nhưng được truy cập thường xuyên

### Syntax
```
create materialized view view_name as
select_statement
```

### Example
```
create materialized view user_stats as
select user_id, count(*) as total_posts
from posts
group by user_id;
```

<a name="delete_truncate"></a>
# Detele and Truncate

## Delete
- Xóa các hàng dựa theo điều kiện
- Có thể rollback nếu trong transaction
- Chậm hơn khi xóa lượng lớn dữ liệu

### Syntax
```
DELETE FROM table_name [WHERE condition];
```

### Example
```
DELETE FROM USER WHERE country = 'China';
```

## Truncate
- Xóa tất cả dữ liệu của một bảng
- Không thể rollback trừ khi trong transaction được hỗ trợ
- Xóa nhanh một lượng lớn dữ liệu

### Syntax
```
TRUNCATE TABLE table_name;
```

### Example
```
TRUNCATE TABLE USER;
```