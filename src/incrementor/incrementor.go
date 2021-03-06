// Incrementor package copyright 2019 Dmitry <vathsven@pm.me>.
// Т.е. класс очень простой. А теперь сложность: оно должно быть сделано очень
// хорошо. Т.е. максимально качественно, как только можно. Код должен быть
// идеальным, все должно быть покрыто unit тестами. Классы и все методы должны быть
// полностью покрыты понятной (т.е. полезной, а не для отписки) javadoc (для Java) или
// аналогом для Swift и Golang документацией. В общем, нужно сделать такой код,
// который каждый разработчик мечтает получить на поддержку -- идеальный (насколько
// кандидат способен).
package incrementor

// Максимально-допустимое число типа int
const MaxInt = int(^uint(0) >> 1)

// Incrementor is an interface
type Incrementor interface {
	GetNumber() int
	IncrementNumber() interface{}
	SetMaximumValue(int) interface{}
}

// Возвращает текущее число. В самом начале это ноль.
func (n *Number) GetNumber() int {
	return n.M
}

// Увеличивает текущее число на один. После каждого вызова этого
// метода getNumber() будет возвращать число на один больше.
func (n *Number) IncrementNumber() interface{} {
	n.M += 1
	// Максимальное число не задано или равно нулю
	// Устанавливаем максимально-допустимое int
	if n.MaximumValue == 0 {
		n.MaximumValue = MaxInt
	}
	// Число достигло максимального заданного значения ?
	if n.M == n.MaximumValue {
		n.M = 0
	}
	return n.M
}

// Устанавливает максимальное значение текущего числа.
// Когда при вызове incrementNumber() текущее число достигает
// этого значения, оно обнуляется, т.е. getNumber() начинает
// снова возвращать ноль, и снова один после следующего
// вызова incrementNumber() и так далее.
// По умолчанию максимум -- максимальное значение int.
// Если при смене максимального значения число резко начинает
// превышать максимальное значение, то число надо обнулить.
// Нельзя позволять установить тут число меньше нуля.
func (n *Number) SetMaximumValue(maximumValue int) interface{} {
	// при смене максимального значения число резко начинает 
	// превышать максимальное значение -> обнулить.
	if n.M > maximumValue {
		n.M = 0
	}
	n.MaximumValue = maximumValue
	return 0
}

// Число и его максимальное значение
type Number struct {
	M int
	MaximumValue int
}
