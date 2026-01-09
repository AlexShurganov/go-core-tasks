/*
* Напишите unit тесты к созданным функциям
 */

package main

type StringIntMap struct {
	data map[string]int
}

func (sim *StringIntMap) Add(key string, value int) {
	sim.data[key] = value
}

func (sim *StringIntMap) Remove(key string) {
	delete(sim.data, key)
}

func (sim *StringIntMap) Exists(key string) bool {
	_, ok := sim.data[key]
	return ok
}

func (sim *StringIntMap) Get(key string) (int, bool) {
	i, ok := sim.data[key]
	return i, ok
}

func (sim *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int, len(sim.data))

	for k, v := range sim.data {
		newMap[k] = v
	}
	return newMap
}
