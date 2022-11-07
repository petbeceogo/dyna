package dyna

func IntValue(val interface{}) (int, error) {
	if intVal, ok := val.(int); ok {
		return intVal, nil
	}
	if floatVal, ok := val.(float64); ok {
		return int(floatVal), nil
	}

	return 0, ErrCast
}

func DoubleValue(val interface{}) (float64, error) {
	if floatVal, ok := val.(float64); ok {
		return floatVal, nil
	}
	if intVal, ok := val.(int); ok {
		return float64(intVal), nil
	}

	return 0, ErrCast
}

func StringValue(val interface{}) (string, error) {
	if stringVal, ok := val.(string); ok {
		return stringVal, nil
	}

	return "", ErrCast
}

func IntMapValue(m map[string]interface{}, key string) (int, error) {
	if val, ok := m[key]; ok {
		return IntValue(val)
	}

	return 0, ErrNoKey
}

func DoubleMapValue(m map[string]interface{}, key string) (float64, error) {
	if val, ok := m[key]; ok {
		return DoubleValue(val)
	}

	return 0, ErrNoKey
}

func StringMapValue(m map[string]interface{}, key string) (string, error) {
	if val, ok := m[key]; ok {
		return StringValue(val)
	}

	return "", ErrNoKey
}

func DoubleArrayValue(arrVal []interface{}) ([]float64, error) {
	arr := make([]float64, len(arrVal))
	for i, v := range arrVal {
		dv, err := DoubleValue(v)
		if err != nil {
			return nil, err
		}
		arr[i] = dv
	}

	return arr, nil
}

func DoubleArrayMapValue(m map[string]interface{}, key string) ([]float64, error) {
	if val, ok := m[key]; ok {
		if arrVal, ok := val.([]interface{}); ok {
			return DoubleArrayValue(arrVal)
		}

		return nil, ErrCast
	}

	return nil, ErrNoKey
}
