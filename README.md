# IdeaBox
## About
Accumulation field of your idea.  
アイデア置き場かつ管理用アプリケーション

## Requirements
```
Golang==1.12.5
echo
```

## How to Use
### Settings
```
$ go get github.com/labstack/echo
$ go get github.com/dgrijalva/jwt-go
```

### Run
```
$ go run main.go
```

## TODO
ユーザ登録の設置とDB作成

## DBMS
### MySQL
If you already have MySQL, this process should be skipped.
```
$ brew install mysql
$ mysql.server start
$ mysql -uroot
mysql> CREATE DATABASE IF NOT EXISTS ideabox;
mysql> CREATE TABLE ideabox.ideas (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, name VARCHAR(100) NOT NULL);
```

When you exit
```
$ mysql.server stop
```


## Tips
### Golang
- blank import  
importを行うライブラリ指定の前に`_`を指定すると、ライブラリ内の`init`関数のみを実行する(`main.go`)

```Go
import (
    _ "github.com/go-sql-driver/mysql"
)
```

### MySQL
- When you reinstall MySQL
```
$ rm -rf /usr/local/var/mysql
$ brew uninstall mysql
$ brew install mysql
```


## Author
[Yusei](https://github.com/index30)