# Итоговое задание
Планировщик задач
Выполнены только базовые задания.

Структура проекта:

В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.

Директория `web` содержит файлы фронтенда.

crud.go - файл с фукциями, которые делают SQL запросы к базе данных

db.go - файл, в котором функции создания и проверки базы данных

nextdate_api.go - обрабатывает GET-запросы к api/nextdate и возвращает дату следующего выполнения задачи

nextdate.go - рассчитвает и проверяет дату следующего выполнения задачи

task_delete.go - удаляет задачу

task_done.go - делает задачу выполненной

task_get.go -возвращает задачу по ее id

task_hendle.go - в зависимости от метода вызывает нужные функции

task_post.go - добавляет задачу в базу данных

task_put.go - редактирует задачу

tasks_get.go - вовзращает список ближайших задач

