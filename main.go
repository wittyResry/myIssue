package main

import (
	"fmt"
	"my-go-test/functions"
	"my-go-test/importmath"
	"my-go-test/multipleresults"
	"my-go-test/namedresults"
	"my-go-test/numeric"
	"my-go-test/packages"
	"my-go-test/typeinference"
	"my-go-test/variables"
	"my-go-test/forcontinued"
    "my-go-test/exercise"
    "my-go-test/switchs"
)

func main() {
	fmt.Println("Welcome to the playgroud!")
	importmath.Main()
	packages.Main()
	functions.Main()
	multipleresults.Main()
	namedresults.Main()
	variables.Main()
	variables.MainInitializers()
	variables.MainVariableDeclarations()
	typeinference.MainTypeInference()
	numeric.MainNumericConstants()
    forcontinued.MainForContinued()
    exercise.MainLoopAndFunction()
    switchs.MainSwitchOS()
    switchs.MainSwitchWithNoCondition()
}
