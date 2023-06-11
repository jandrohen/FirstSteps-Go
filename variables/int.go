package variables

import "fmt"

func ShowInt() {
	var intCommon int
	intOff32 := int32(10)
	intOff64 := int64(10)

	fmt.Println("intCommon:", intCommon)
	fmt.Println("intOff32:", intOff32)
	fmt.Println("intOff64:", intOff64)
}
