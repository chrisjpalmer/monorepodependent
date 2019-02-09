package main

import "fmt"
import "github.com/chrisjpalmer/monorepotest/pkg2"
import "github.com/chrisjpalmer/monorepotest/pkg1"

func main() {
	fmt.Println("This is app3")
	pkg2.Pkg2Function()
	pkg1.Pkg1Function()
}
