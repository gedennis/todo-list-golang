# todo-list-golang

Todo list demo backend with Golang.

## Usage

下载代码：

```bash
$ git clone git@github.com:gedennis/todo-list-golang.git
```

### 启动 MySQL

```bash
$ docker run --name mysql -e MYSQL_ROOT_PASSWORD=root123 -p 3306:3306 -d mysql
```

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

### 安装依赖

```bash
$ cd todo-list-golang
$ go install
```

### 运行

```bash
go run main.go
```

系统在 `localhost:8090`, 然后通过 postman 测试。

## 参考

- [go-gin-gorm-todos](https://github.com/satoshiyamamoto/go-gin-gorm-todos)
