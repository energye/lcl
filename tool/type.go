// Copyright © yanghy. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.

package tool

import (
	"sort"
)

// Null 空结构体
type Null struct{}

type Array[T comparable] struct {
	values []T
}

type ArrayMap[K, T comparable] struct {
	keys   []K
	values map[K]T
}

type HashSet[T comparable] struct {
	values map[T]Null
}

type ArraySet struct {
	keys []string
	hash map[string]bool
}

type HashMap[K, T comparable] struct {
	values map[K]T
}

func NewArray[T comparable]() *Array[T] {
	return &Array[T]{values: make([]T, 0)}
}

func (m *Array[T]) Add(value T) {
	m.values = append(m.values, value)
}

func (m *Array[T]) Adds(value ...T) {
	for _, val := range value {
		m.values = append(m.values, val)
	}
}

func (m *Array[T]) Get(index int) (T, bool) {
	if index < len(m.values) {
		return m.values[index], true
	}
	var n T
	return n, false
}

func (m *Array[T]) Len() int {
	return len(m.values)
}

func (m *Array[T]) Values() []T {
	return m.values
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{values: make(map[T]Null)}
}

func NewHashSetByValues[T comparable](values ...T) *HashSet[T] {
	result := &HashSet[T]{values: make(map[T]Null)}
	result.Add(values...)
	return result
}

func (m *HashSet[T]) Values() (values []T) {
	for key, _ := range m.values {
		values = append(values, key)
	}
	return
}

func (m *HashSet[T]) Contains(key T) bool {
	if _, ok := m.values[key]; ok {
		return ok
	}
	return false
}

func (m *HashSet[T]) Add(value ...T) {
	for _, v := range value {
		m.values[v] = Null{}
	}
}

func (m *HashSet[T]) Iterate(fn func(key T) bool) {
	for key, _ := range m.values {
		if fn(key) {
			break
		}
	}
}

func NewArraySet() *ArraySet {
	return &ArraySet{keys: make([]string, 0), hash: make(map[string]bool)}
}

func (m *ArraySet) ContainsKey(key string) bool {
	if _, ok := m.hash[key]; ok {
		return true
	}
	return false
}

func (m *ArraySet) Add(key string) {
	if _, ok := m.hash[key]; !ok {
		m.keys = append(m.keys, key)
		m.hash[key] = true
	}
}

func (m *ArraySet) Adds(keys ...string) {
	for _, key := range keys {
		if _, ok := m.hash[key]; !ok {
			m.keys = append(m.keys, key)
			m.hash[key] = true
		}
	}
}

func (m *ArraySet) Values() (result []string) {
	return m.keys
}

func (m *ArraySet) Iterate(fn func(key string) bool) {
	if fn == nil {
		return
	}
	for _, key := range m.keys {
		if fn(key) {
			break
		}
	}
}

func (m *ArraySet) Sort() {
	sort.Strings(m.keys)
}

func (m *ArrayMap[K, T]) Del(key K) {
	if _, ok := m.values[key]; ok {
		delete(m.values, key)
		for idx, tmpKey := range m.keys {
			if key == tmpKey {
				m.keys = append(m.keys[:idx], m.keys[idx+1:]...)
				break
			}
		}
	}
}

func NewArrayMap[K, T comparable]() *ArrayMap[K, T] {
	return &ArrayMap[K, T]{keys: make([]K, 0), values: make(map[K]T)}
}

func (m *ArrayMap[K, T]) ContainsKey(key K) bool {
	if _, ok := m.values[key]; ok {
		return true
	}
	return false
}

func (m *ArrayMap[K, T]) ContainsValue(value T) bool {
	for _, val := range m.values {
		if val == value {
			return true
		}
	}
	return false
}

func (m *ArrayMap[K, T]) Add(key K, value T) {
	if m.values == nil {
		m.values = make(map[K]T)
	}
	if _, ok := m.values[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.values[key] = value
}

func (m *ArrayMap[K, T]) Get(key K) T {
	return m.values[key]
}

func (m *ArrayMap[K, T]) Keys() []K {
	return m.keys
}

func (m *ArrayMap[K, T]) Values() (result []T) {
	for _, key := range m.keys {
		result = append(result, m.values[key])
	}
	return
}
func (m *ArrayMap[K, T]) Iterate(fn func(key K, value T) bool) {
	if fn == nil {
		return
	}
	for _, key := range m.keys {
		if fn(key, m.values[key]) {
			break
		}
	}
}

func NewHashMap[K, T comparable]() *HashMap[K, T] {
	return &HashMap[K, T]{values: make(map[K]T)}
}

func (m *HashMap[K, T]) ContainsKey(key K) bool {
	if _, ok := m.values[key]; ok {
		return true
	}
	return false
}

func (m *HashMap[K, T]) ContainsValue(value T) (key K, ok bool) {
	for k, val := range m.values {
		if val == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func (m *HashMap[K, T]) Add(key K, value T) {
	if m.values == nil {
		m.values = make(map[K]T)
	}
	m.values[key] = value
}

func (m *HashMap[K, T]) Clear() {
	m.values = make(map[K]T)
}

func (m *HashMap[K, T]) Get(key K) T {
	return m.values[key]
}

func (m *HashMap[K, T]) Values() map[K]T {
	return m.values
}

func (m *HashMap[K, T]) Iterate(fn func(key K, value T) bool) {
	if fn == nil {
		return
	}
	for key, val := range m.values {
		if fn(key, val) {
			break
		}
	}
}

func (m *HashMap[K, T]) Remove(name K) {
	delete(m.values, name)
}
