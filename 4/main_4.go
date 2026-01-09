/*
На вход подаются два неупорядоченных слайса строк.
Например:
```
slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
slice2 := []string{"banana", "date", "fig"}
```
Напишите функцию, которая возвращает слайс строк, содержащий элементы, которые есть в первом слайсе, но отсутствуют во втором.

* Напишите unit тесты к созданным функциям
*/

package main

func FindCommon(slice1, slice2 []string) []string {

	seen := make(map[string]bool)

	var result []string

	for _, v := range slice2 {
		seen[v] = true
	}

	for _, v := range slice1 {
		if _, ok := seen[v]; ok {
			result = append(result, v)
		}
	}

	return result

}
