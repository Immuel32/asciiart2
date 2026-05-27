package main

func MergeBanner(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	m1 := base
	m2 := priority
	m3 := map[rune][]string{}

	for i, v := range m1 {
		m3[i] = v
	}
	for i, v := range m2 {
		m3[i] = v
	}
	return m3
}
