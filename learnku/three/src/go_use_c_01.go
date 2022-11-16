package src
package rand

// #includ <stdlib.h>
import "C"

func Random() int {
	return int(C.random())
}


