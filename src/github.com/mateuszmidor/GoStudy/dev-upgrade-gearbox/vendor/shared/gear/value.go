package gear

// Value represents gear value
type Value struct {
	value, min, max int
}

// New is constructor
func New(min, max int) Value {
	if min > max {
		panic("Gear range min > max")
	}
	return Value{
		value: min,
		min:   min,
		max:   max,
	}
}

// Set is setter
func (v Value) Set(value int) Value {
	if value < v.min {
		panic("value < min gear")
	}

	if value > v.max {
		panic("value > max gear")
	}

	return Value{
		value: value,
		min:   v.min,
		max:   v.max,
	}
}

// Up increases the gear
func (v Value) Up() Value {
	newValue := v.value + 1
	if newValue > v.max {
		newValue = v.max
	}

	return Value{
		value: newValue,
		min:   v.min,
		max:   v.max,
	}
}

// Down decreases the gear
func (v Value) Down() Value {
	newValue := v.value - 1
	if newValue < v.min {
		newValue = v.min
	}
	return Value{
		value: newValue,
		min:   v.min,
		max:   v.max,
	}
}

// ApplyChange changes the gear according to change
func (v Value) ApplyChange(c Change) Value {
	newValue := v.value + int(c)
	if newValue > v.max {
		newValue = v.max
	}
	if newValue < v.min {
		newValue = v.min
	}
	return Value{
		value: newValue,
		min:   v.min,
		max:   v.max,
	}
}
