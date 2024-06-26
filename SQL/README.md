# Table of Content
1. [ SELECT ](#select)
2. [ WHERE ](#where)
3. [ ORDER BY ](#order_by)


<a name="select"></a>
# SELECT

## Syntax
SELECT column1, column2, ... FROM table_name;

## Example
```
SELECT user_id FROM USER

SELECT user_id, age FROM USER

SELECT * FROM USER
```

<a name="where"></a>
# WHERE

## Syntax
SELECT column1, column2, ... FROM table_name
WHERE condition;

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
SELECT column1, column2, ...
FROM table_name
ORDER BY column1, column2, ... ASC|DESC;

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
SELECT column1, column2, ... FROM table_name
WHERE condition1 AND/OR/NOT condition2 AND/OR/NOT condition3 ...;

## Example
```
SELECT * FROM USER
WHERE user_id > 10
OR age = 20
NOT country = "china";
```