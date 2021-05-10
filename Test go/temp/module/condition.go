package module

func TestUp(a int, b int) bool {
	res := false
	if a > b {
		res = true
	}
	return res
}

func TestDown(a int, b int) bool {
	res := false
	if a < b {
		res = true

	}
	return res
}

func TestEgal( a int, b int ) bool {
	res := false
	if a == b {
		res = true

	}
	return res
}
