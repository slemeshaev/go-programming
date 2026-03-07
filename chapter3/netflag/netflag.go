package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"

	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"

	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"

	fmt.Printf("KiB = %d\n", KiB)
	fmt.Printf("MiB = %d\n", MiB)
	fmt.Printf("GiB = %d\n", GiB)
	fmt.Printf("TiB = %d\n", TiB)
	fmt.Printf("PiB = %d\n", PiB)
	fmt.Printf("EiB = %d\n", EiB)

	fmt.Printf("ZiB = %g\n", float64(ZiB)) // 1 << 70
	fmt.Printf("YiB = %g\n", float64(YiB)) // 1 << 80
}
