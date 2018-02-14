package variables

import "fmt"

func MainInitializers() {
	var c, python, java = true, false, "no!"
	fmt.Println(c, python, java)
}
