package flatten

//Flatten takes inputs of slices and flattens them
func Flatten(input interface{}) []interface{} {
	in := input.([]interface{})
	result := make([]interface{}, 0)

	for _, v := range in {
		if v != nil {
			switch v.(type) {
			case []interface{}:
				result = append(result, Flatten(v)...)
			default:
				result = append(result, v)
			}
		}
	}

	return result
}
