package utils

// IsMapSame return true when m1 == m2.
// If both of them is nil, return true.
// If one is nil, return false.
// When all the key and value is equals, return ture, else return false.
func IsMapSame(m1, m2 map[string]string) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 == nil || m2 == nil {
		return false
	}

	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v1 != v2 {
			return false
		}
	}

	return true
}
