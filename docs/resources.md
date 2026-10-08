# Полезные ссылки

[К описанию курса](../README.md)

Это навигация, а не требование прочитать всё целиком. Для первой работы сначала нужны установка, первая программа и публикация. Ссылки на источники последующих тем можно пока не открывать.

## Пошаговые инструкции курса

| Материал | Когда открывать |
|---|---|
| [Как пользоваться GitHub и получать файлы](how-to-use.md) | Впервые открыли репозиторий |
| [Установка Go и редактора](environment.md) | Ещё не настроено окружение |
| [Словарь](glossary.md) | Встретился незнакомый термин |
| [Первая практика](../practices/01-introduction/README.md) | Создать программу своими руками |
| [Вторая практика](../practices/02-control-flow-functions/README.md) | Освоить условия, циклы, функции и коллекции |
| [Задания второй практики](../practices/02-control-flow-functions/assignments.md) | Реализовать сводку статусов задач |
| [Примеры второй практики](../practices/02-control-flow-functions/examples/README.md) | Запустить четыре мини-программы |
| [Подробный синтаксис и ввод](../practices/01-introduction/syntax-and-input.md) | Повторить переменные и подготовиться к конвертеру |
| [Задания](../practices/01-introduction/assignments.md) | Написать карточку и конвертер |
| [Публикация проекта](first-project-github.md) | Начать с Git и отправить код на GitHub |
| [Шаблон README](../practices/01-introduction/templates/project-readme.md) | Описать собственный проект |
| [Шаблон .gitignore](../practices/01-introduction/templates/.gitignore) | Не добавлять лишние файлы в Git |

## Начало работы с Go

| Материал | Когда открывать |
|---|---|
| [Загрузить Go](https://go.dev/dl/) | Выбрать стабильный SDK для своего компьютера |
| [Установка Go](https://go.dev/doc/install) | Установить SDK и проверить `go version` |
| [Первая программа](https://go.dev/doc/tutorial/getting-started) | Повторить создание модуля и запуск |
| [Go Playground](https://go.dev/play/) | Проверить короткий пример в браузере |
| [A Tour of Go](https://go.dev/tour/welcome/1) | Закрепить основы интерактивными примерами |

## Редакторы

| Материал | Назначение |
|---|---|
| [VS Code: загрузка](https://code.visualstudio.com/Download) | Установка редактора |
| [VS Code: Windows](https://code.visualstudio.com/docs/setup/windows) | Выбор установщика |
| [VS Code: macOS](https://code.visualstudio.com/docs/setup/mac) | Установка на Mac |
| [VS Code: Linux](https://code.visualstudio.com/docs/setup/linux) | Установка для конкретного дистрибутива |
| [Терминал VS Code](https://code.visualstudio.com/docs/terminal/basics) | Открытие терминала и работа с командами |
| [Go в VS Code](https://code.visualstudio.com/docs/languages/go) | Поддержка языка |
| [Расширение Go](https://marketplace.visualstudio.com/items?itemName=golang.Go) | Официальная карточка `golang.Go` |
| [Документация расширения](https://github.com/golang/vscode-go) | Инструменты, подсказки и диагностика |
| [GoLand: установка](https://www.jetbrains.com/help/go/installation-guide.html) | Альтернативная среда |
| [GoLand: проект и SDK](https://www.jetbrains.com/help/go/installing-and-configuring-goland.html) | Создание проекта и выбор SDK |

## Git и GitHub

| Материал | Назначение |
|---|---|
| [Git: установка](https://git-scm.com/install/) | Выбрать установщик своей системы |
| [Создание аккаунта GitHub](https://docs.github.com/en/account-and-profile/how-tos/account-management/creating-an-account-on-github) | Регистрация и подтверждение почты |
| [Адрес автора коммитов](https://docs.github.com/en/account-and-profile/how-tos/email-preferences/setting-your-commit-email-address) | Настроить почту, в том числе приватный `noreply` |
| [Git в VS Code](https://code.visualstudio.com/docs/sourcecontrol/quickstart) | Локальная история и первая публикация |
| [GitHub в VS Code](https://code.visualstudio.com/docs/sourcecontrol/github) | Вход через браузер и работа с репозиторием |
| [Pro Git на русском](https://git-scm.com/book/ru/v2) | Коммиты, ветки и удалённый репозиторий |
| [Публикация через командную строку](https://docs.github.com/en/migrations/importing-source-code/using-the-command-line-to-import-source-code/adding-locally-hosted-code-to-github) | Альтернативный путь для дальнейшего изучения |

## Язык и стандартная библиотека

| Материал | Темы |
|---|---|
| [Переменные](https://go.dev/tour/basics/8) | Объявление через `var` |
| [Краткое объявление](https://go.dev/tour/basics/10) | Конструкция `:=` |
| [Базовые типы](https://go.dev/tour/basics/11) | `int`, `float64`, `bool`, `string` |
| [Нулевые значения](https://go.dev/tour/basics/12) | Переменные без инициализатора |
| [Преобразования](https://go.dev/tour/basics/13) | Явное преобразование чисел |
| [Константы](https://go.dev/tour/basics/15) | Объявление через `const` |
| [fmt](https://pkg.go.dev/fmt) | Вывод, форматирование и ввод |
| [Команда go](https://pkg.go.dev/cmd/go) | `run`, `build`, `fmt`, `test` и другие команды |
| [Спецификация](https://go.dev/ref/spec) | Уточнение правил; не первое чтение |
| [Effective Go](https://go.dev/doc/effective_go) | Приёмы написания понятного кода после освоения основ |
| [Go Modules](https://go.dev/ref/mod) | Модули и зависимости |
| [Управляющие конструкции](https://go.dev/tour/flowcontrol/1) | `if`, `switch`, `for` |
| [Функции](https://go.dev/tour/basics/4) | Параметры и возвращаемые значения |
| [Срезы](https://go.dev/tour/moretypes/7) | `[]T`, `append`, `range` |
| [Карты](https://go.dev/tour/moretypes/19) | `map`, ключи и значения |
| [errors](https://pkg.go.dev/errors) | Проверка и сопоставление ошибок |
| [encoding/json](https://pkg.go.dev/encoding/json) | Чтение и формирование JSON |
| [testing](https://pkg.go.dev/testing) | Модульные тесты |

## Базы данных

| Материал | Темы |
|---|---|
| [PostgreSQL: загрузка](https://www.postgresql.org/download/) | Подготовка к теме баз данных |
| [PostgreSQL: учебник](https://www.postgresql.org/docs/current/tutorial.html) | Таблицы и SQL |
| [GORM](https://gorm.io/docs/) | Модели и взаимодействие с БД |

## Серверная разработка

| Материал | Темы |
|---|---|
| [net/http](https://pkg.go.dev/net/http) | HTTP-сервер стандартной библиотеки |
| [Gin](https://gin-gonic.com/en/docs/) | Маршрутизация и JSON API |
| [httptest](https://pkg.go.dev/net/http/httptest) | Проверка HTTP-обработчиков |
| [Конкурентность: A Tour of Go](https://go.dev/tour/concurrency/1) | Goroutine и channel |
| [sync](https://pkg.go.dev/sync) | Синхронизация |
| [context](https://pkg.go.dev/context) | Время жизни операций |

Большинство первоисточников на английском. Можно переводить пояснения средствами браузера; команды, имена функций и код переводить не нужно. В шагах установки даются официальные страницы выбора, а не ссылки на устаревающие конкретные версии установщиков.
