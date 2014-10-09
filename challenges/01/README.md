Да се напише функция `mapSum`, чийто аргументи са друга функция и произволен брой цели числа.
Функцията аргумент от своя страна е функция с един аргумент - цяло число и връщаща цяло
число.

Резултатът от `mapSum` трябва да е сбор на резултатите от прилагането на
функцията аргумент върху всеки от другите аргументи на `mapSum`.

Example

```go
    func cube (a int) int {
        return a * a * a
    }

    result := mapSum(cube, 7, 33)
```

В горния пример `result` трябва да стане 36280.