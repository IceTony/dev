# Golang REST API MySQL service
## Запуск приложения
1. Установить пакеты
```
sudo apt update
sudo apt install golang mysql-server-5.7 mysql-client-5.7 git
```
2. Создать базу и пользователя в MySQL
3. Выполнить импорт таблицы:
```
mysql database < dump.sql
```
4. Перейти в папку с приложением и выполнить команду (для загрузки зависимостей):
```
go get ./...
```
5. В методе Database() изменить данные для подключения к базе
6. Запустить приложение командой:
```
go run app.go
```
