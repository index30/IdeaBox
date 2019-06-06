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
$ go get github.com/jinzhu/gorm
```

### Run
```
$ go run main.go
```

## TODO
ユーザ登録の設置とDB制作

## DBMS
### MySQL
If you already have MySQL, this process should be skipped.
```
$ brew install mysql
$ mysql.server start
$ mysql -uroot
```

## Tips
### MySQL
- When you reinstall MySQL
```
$ rm -rf /usr/local/var/mysql
$ brew uninstall mysql
$ brew install mysql
```


## Author
[Yusei](https://github.com/index30)