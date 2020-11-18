# todo-list-golang

Todo list demo backend with Golang.

## 准备

### 创建数据库

```sql
# create DB
CREATE DATABASE todoDB;

# create table
CREATE TABLE todos (
	id INT(11) not null auto_increment,
	title VARCHAR(100) not null,
	description VARCHAR(300),
	status TINYINT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
	PRIMARY KEY(id)
);
```
