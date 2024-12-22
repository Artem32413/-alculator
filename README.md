# Калькулятор


## Сервис подсчёта арифметических выражений
___
[![Typing SVG](https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=21&pause=10000&color=00D7FF&repeat=false&width=530&lines=%D0%97%D0%B0%D0%BF%D1%83%D1%81%D0%BA+%D0%B8+%D1%83%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B0+%D0%BF%D1%80%D0%BE%D0%B5%D0%BA%D1%82%D0%B0)](https://git.io/typing-svg)

Для запуска программы можно использовать:
>calc.exe
>
Или в командной строке прописать:
>go run main.go
>
___
[![Typing SVG](https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=21&pause=10000&color=00D7FF&repeat=false&width=530&lines=%D0%9F%D1%80%D0%B8%D0%BC%D0%B5%D1%80+%D0%B2%D0%B2%D0%BE%D0%B4%D0%B0++%D1%81+%D0%BF%D0%BE%D0%BC%D0%BE%D1%89%D1%8C%D1%8E+Postman)](https://git.io/typing-svg)

У сервиса 1 endpoint с url-ом /api/v1/calculate. Отправим на этот url POST-запрос с телом:
```
{
    "Expression": "2+2"
}
```
В ответе получаем HTTP-ответ с телом:
```
{
    "Result": "4"
}
```
и кодом 200, если выражение вычислено успешно, либо HTTP-ответ с телом:
```
{
    "Error": "2+2j"
}
```
и кодом 422, если входные данные не соответствуют требованиям приложения — например, кроме цифр и разрешённых операций ты ввёл символ английского алфавита.

Ещё один вариант HTTP-ответа:
```
{
    "Error": "Internal server error"
}
```
и код 500 в случае какой-либо иной ошибки («Что-то пошло не так»).
____
[![Typing SVG](https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=21&pause=10000&color=00D7FF&repeat=false&width=530&lines=%D0%9F%D1%80%D0%B8%D0%BC%D0%B5%D1%80+%D0%B2%D0%B2%D0%BE%D0%B4%D0%B0++%D1%81+%D0%BF%D0%BE%D0%BC%D0%BE%D1%89%D1%8C%D1%8E+curl)](https://git.io/typing-svg)

```
curl --location 'http://localhost:8080/api/v1/calculate' \
--header '7: application/gzip' \
--header 'Content-Type: application/json' \
--data '{
    "expression": "2+2"
}
'
```
___
