# Golang REST API MySQL service
## Запуск приложения
1. Установить пакеты
```
sudo apt update
sudo apt install golang mysql-server-5.7 mysql-client-5.7 git
```
2. Создать базу и пользователя в MySQL
3. Скачать приложение:
```
go get -v github.com/rubenv/sql-migrate/...
```
и создать конфиг ```dbconfig.yml``` с содержанием:
```
development:
  dialect: mysql
  datasource: dbuser:dbpass@/dbname?parseTime=true
  table: migrations
  dir: migrations
```
4. Скопировать файл миграции в папку migrations
5. Выполнить миграцию:
```
sql-migrate up -env="development"
```
6. Перейти в папку с приложением и выполнить команду (для загрузки зависимостей):
```
go get ./...
```
7. Изменить данные для подключения к БД
8. Запустить приложение командой:
```
go run app.go
```
