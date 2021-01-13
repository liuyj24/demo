package jun

func resolveAddress(addr []string) string {
	switch len(addr) {
	case 0:
		return ":8080" //default
	case 1:
		return addr[0]
	default:
		panic("too much parameters")
	}
}
