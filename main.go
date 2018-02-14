package main

import (
	"fmt"
	"my-go-test/defers"
	"my-go-test/exercise"
	"my-go-test/forcontinued"
	"my-go-test/functions"
	"my-go-test/importmath"
	"my-go-test/multipleresults"
	"my-go-test/namedresults"
	"my-go-test/numeric"
	"my-go-test/packages"
	"my-go-test/pointers"
	"my-go-test/structs"
	"my-go-test/switchs"
	"my-go-test/typeinference"
	"my-go-test/variables"
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
	defers.MainDefer()
	defers.MainDeferStack()
	pointers.MainPointer()
	structs.MainStruct()
	structs.MainStructPoint()
	fmt.Printf("%T", "ss")
}
