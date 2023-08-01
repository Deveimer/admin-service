package utils

type Rule interface {
	Check(value interface{}) bool
}

// Validate checks a value against rules. If anyone of the rule fails, it returns false.
func Validate(value interface{}, rules ...Rule) bool {
	for _, rule := range rules {
		if rule.Check(value) == false {
			return false
		}
	}

	return true
}

type Enum struct {
	Values []interface{}
}

func (e Enum) Check(v interface{}) bool {
	for _, i := range e.Values {
		if i == v {
			return true
		}
	}
	return false
}
