// Incrementor package copyright 2019 Dmitry <vathsven@pm.me>.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"testing"
	"incrementor"
)

// Работает ли GetNumber на первом пуске
// Должно сработать значение ноль по дефолту
func TestGetNumber(t *testing.T) {
	var i = &incrementor.Number{}
	init := i.GetNumber()
	if init != 0 {
		t.Errorf("GetNumber() = %d; want 0", init)
	}
}

// Проверка инкремента числа.
// Здесь же автоматом проверяется условие
// установки начального максимального-допустимого значения числа
func TestIncrementNumber(t *testing.T) {
	var i = &incrementor.Number{}
	i.IncrementNumber()
	i.IncrementNumber()
	init := i.GetNumber()
	if init != 2 {
		t.Errorf("IncrementNumber() = %d; want 2", init)
	}
}

// Проверяем, если число достигло 
// установленного максимального значения
// установленного SetMaximumValue
func TestSetMaximumValueOnMatch(t *testing.T) {
	var i = &incrementor.Number{}
	i.SetMaximumValue(5)
	i.IncrementNumber()
	i.IncrementNumber()
	i.IncrementNumber()
	i.IncrementNumber()
	i.IncrementNumber()
	init := i.GetNumber()
	if init != 0 {
		t.Errorf("SetMaximumValue(5) = %d; want 0", init)
	}
}

// Проверка срабатывания условия:
// Если при смене максимального значения число резко начинает
// превышать максимальное значение, то число надо обнулить.
func TestSetMaximumValueWhenOutOfRange(t *testing.T) {
	var i = &incrementor.Number{}
	i.IncrementNumber()
	i.IncrementNumber()
	i.IncrementNumber()
	i.IncrementNumber()
	i.IncrementNumber()
	i.IncrementNumber()
	i.IncrementNumber()
	i.SetMaximumValue(5)
	init := i.GetNumber()
	if init != 0 {
		t.Errorf("SetMaximumValue(5) = %d; want 0", init)
	}
}
