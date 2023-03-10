package value

type Value string

func NewValue(s string) *Value {
	v := Value(s)
	return &v
}

func (v *Value) String() string {
	return string(*v)
}
