##### Демо сервис на golang с использованием фреймворка gin

##### 1. Данный сервис может работать сразу на двух портах: 8081 и 8082
> http://localhost:8081/

> http://localhost:8082/

Есть возможность использовать BasicAuth

##### 2. Добавлен элементарный тест

##### 3. Перед сборкой выполнить:

assert - для того, чтобы собрать всё в один бинарник (даже html-шаблоны)

```
go get github.com/stretchr/testify/assert
go-assets-builder html -o assets.go
```

Собрать всё в один бинарник (вместе с шаблонами) и положить этот бинарник в каталог build
```
go build -o build/go-gin-demo
```

##### 4. Попробовал использовать go модули (https://roberto.selbach.ca/intro-to-go-modules/)

##### 5. Запустил ab benchmark test на ноутбуке (Dell Inspiron 15 7000 Gaming, intel Core i7, 4Gb, nVidia GeForce GTX 1050 Ti)

```
ab.exe -n 10000 -c 100 "http://localhost:8081/"
rps: 6765.68
```

```
ab.exe -n 10000 -c 100 "http://localhost:8082/"
rps: 6358.08
```
