assert - для того, чтобы собрать всё в один бинарник (даже html-шаблоны)

go get github.com/stretchr/testify/assert
go-assets-builder html -o assets.go

Собрать всё в один бинарник - вместе с шаблонами
go build -o assets-in-binary
