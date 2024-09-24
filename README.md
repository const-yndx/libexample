# Lib Example

Библиотека, демонстрирующая версионирование.

Как добавить версию библиотеки:

1. В go.mod называем модуль в формате `github.com/<аккаунт>/<репозиторий>`, например, `github.com/const-yndx/libexample`
2. Пушим изменения на GitHub
3. Создаем релиз, в котором создаем тэг (tag) с названием версии в формате SemVer, например, `v1.0.0`
4. Используем библиотеку в своих проектах: `go get github.com/<аккаунт>/<репозиторий>@v1.0.0`

Под каждую мажорную версию больше 1 необходимо создавать соответствующую папку.
Например, для выпуска версии 2 необходимо создать папку `v2`, и в нее положить весь код новой версии библиотеки.
