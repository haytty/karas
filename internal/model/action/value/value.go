package value

type Value string

func (v *Value) String() string {
	return string(*v)
}
