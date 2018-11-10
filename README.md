##### Демо сервис на golang с использованием фреймворка gin

1. Данный сервис может работать сразу на двух портах: 8081 и 8082
http://localhost:8081/
http://localhost:8082/
Есть возможность использовать BasicAuth

2. Добавлен элементарный тест

3. Предусмотрена сборка в один бинарник

assert - для того, чтобы собрать всё в один бинарник (даже html-шаблоны)

go get github.com/stretchr/testify/assert
go-assets-builder html -o assets.go

Собрать всё в один бинарник - вместе с шаблонами
go build -o assets-in-binary

4. Попробовал использовать go модули (https://roberto.selbach.ca/intro-to-go-modules/)
