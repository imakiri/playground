package service

func IsNilSafe(l ...interface{}) bool {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			return false
		}
	}
	return true
}

func IsNilSafeEx(l ...interface{}) (b []bool) {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			b = append(b, false)
		} else {
			b = append(b, true)
		}
	}
	return
}
