package matchers

import (
	"fmt"
	"reflect"
)

type HaveKeyMatcher struct {
	key interface{}
}

func HaveKey(key interface{}) HaveKeyMatcher {
	return HaveKeyMatcher{
		key: key,
	}
}

func (m HaveKeyMatcher) Match(actual interface{}) (interface{}, error) {
	t := reflect.TypeOf(actual)
	if t.Kind() != reflect.Map {
		return nil, fmt.Errorf("'%v' (%T) is not a Map", actual, actual)
	}

	if t.Key() != reflect.TypeOf(m.key) {
		return nil, fmt.Errorf("'%v' (%T) has a Key type of %v not %T", actual, actual, t.Key(), m.key)
	}

	v := reflect.ValueOf(actual)
	value := v.MapIndex(reflect.ValueOf(m.key))
	if !value.IsValid() {
		return nil, fmt.Errorf("unable to find key %v in %v", m.key, actual)
	}

	return value.Interface(), nil
}