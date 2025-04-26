package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

/*
type OrderedMap struct { ... }

func NewOrderedMap() OrderedMap                      // создать упорядоченный словарь
func (m \*OrderedMap) Insert(key, value int)          // добавить элемент в словарь
func (m \*OrderedMap) Erase(key int)                  // удалить элемент из словари
func (m \*OrderedMap) Contains(key int) bool          // проверить существование элемента в словаре
func (m \*OrderedMap) Size() int                      // получить количество элементов в словаре
func (m \*OrderedMap) ForEach(action func(int, int))  // применить функцию к каждому элементу словаря от меньшего к большему
*/

type OrderedMap struct {
	l    *OrderedMap
	r    *OrderedMap
	node *struct {
		key int
		val int
	}
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		// set:  false,
	}
}

func (m *OrderedMap) Insert(key, value int) {
	if m.node == nil {
		m.node = &struct {
			key int
			val int
		}{
			key: key,
			val: value,
		}
		return
	}
	if key > m.node.key {
		if m.r == nil {
			r := NewOrderedMap()
			m.r = &r
		}
		m.r.Insert(key, value)
	} else if key < m.node.key {
		if m.l == nil {
			l := NewOrderedMap()
			m.l = &l
		}
		m.l.Insert(key, value)
	}
}

func (m *OrderedMap) Erase(key int) {
	if m == nil || m.node == nil {
		return
	}

	if m.node.key == key {
		m.node = nil
	} else if key < m.node.key {
		m.l.Erase(key)
	} else if key > m.node.key {
		m.r.Erase(key)
	}
}

func (m *OrderedMap) Contains(key int) bool {
	if m == nil || m.node == nil {
		return false
	}
	if m.node.key == key {
		return true
	} else if key < m.node.key {
		return m.l.Contains(key)
	} else if key > m.node.key {
		return m.r.Contains(key)
	}
	return false
}

func (m *OrderedMap) Size() int {
	if m == nil {
		return 0
	}
	if m.node == nil {
		return m.l.Size() + m.r.Size()
	}
	return 1 + m.l.Size() + m.r.Size()
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	if m == nil || m.node == nil {
		return
	}
	m.l.ForEach(action)
	action(m.node.key, m.node.val)
	m.r.ForEach(action)
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
