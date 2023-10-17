# Doodocs-tech-task
Этот микросервис реализует функциональность управления списком задач.

# Требования к запуску
## Установка и Запуск приложения

1. Клонируйте репозиторий проекта на свой компьютер:
```shell
git clone https://github.com/zsandibe/tech-task
```
2. Перейдите в директорию проекта:
```shell
cd ./tech-task
```
3. Запустите приложение с помощью команды:
```shell
make run
```



### API Endpoints
#### ** Формат обмена данными JSON.**

## Получение информации по zip

1. Метод: POST
- URL: /api/archive/information
- Тело запроса:

```json
{
    "file_name":"test.zip",
    "archive_size": 35
	"total_size": 38
	"total_files": 1
	"files" : {
        "file_path": "photo.png",
	    "size": 3,
        "mime_type": "image/png",
    }
}
```

## Cоздание zip file

2. Метод: POST
    
- URL: /api/archive/files
- Тело запроса:
```json
{
    "file_name":"test.zip",
    "archive_size": 35
	"total_size": 38
	"total_files": 1
	"files" : {
        "file_path": "photo.png",
	    "size": 3,
        "mime_type": "image/png",
    }
}
```


## Отправка файла на почту

3. Метод: POST

- URL: /api/mail/file
- Тело запроса:

```json
{
   "file": "test.pdf",
   "emails": "test@gmail.com",
}
```


##  Тестирование
Запуск unit тестов
```shell
make test
```
