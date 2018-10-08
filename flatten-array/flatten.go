package flatten

//Flatten ...
func Flatten(array interface{}) []interface{} {

	switch array.(type) {

	case []interface{}:

		res := []interface{}{}

		for _, stuff := range array.([]interface{}) {

			for _, elem := range Flatten(stuff) {
				res = append(res, elem)
			}
		}

		return res

	case nil:
		return []interface{}{}
	default:
		return []interface{}{array}
	}

}
