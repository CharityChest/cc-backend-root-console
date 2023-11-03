package array

func FilterStrings(data []string, f func(string) bool) []string {

	fltd := make([]string, 0)

	for _, e := range data {

		if f(e) {
			fltd = append(fltd, e)
		}
	}

	return fltd
}
