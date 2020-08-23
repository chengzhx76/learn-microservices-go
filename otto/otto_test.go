package otto

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"testing"
)

func TestOtto(t *testing.T) {
	vm := otto.New()
	vm.Run(`
		abc = 2 + 2;
		console.log(abc)
	`)
	value, err := vm.Get("abc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

func TestExeFun(t *testing.T) {
	vm := otto.New()
	script := `
		function getName(pre) {
			return pre + " cheng"
		}
	`
	_, err := vm.Run(script)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("====================================")

	result, err := vm.Call("getName", nil, "haha")

	fmt.Println(result)
}

func TestExeInnerFun(t *testing.T) {
	vm := otto.New()
	script := `
		function innerFun() {
			return {
				getName: function(name) {
					return name + " cc"
				}
			}
		}
	`
	_, err := vm.Run(script)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("====================================")

	jsFun, err := vm.Get("innerFun")

	this, err := jsFun.Call(jsFun)

	getNameFun, err := this.Object().Get("getName")

	result_getName, err := getNameFun.Call(this, "haha")

	fmt.Println(result_getName.String())
}

func TestRunFun(t *testing.T) {
	vm := otto.New()

	script := `
		function RunTest() {
			return {
				"name": "cheng",
				"age": 19,
				"sayHi": function(val) {
					return val + "hello world!"
				}
			}
		}
	`

	_, err := vm.Run(script)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("====================================")

	plugin, err := vm.Get("RunTest")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("is-fun->", plugin.IsFunction())

	this, err := plugin.Call(plugin, nil)

	this_value, err := this.Object().Get("name")
	fmt.Println(this_value)

	fmt.Println("====================================")

	//sayHiFun, err := result.Object().Get("sayHi")
	/*sayHiFun, err := this.Object().Get("sayHi")
	returnObj, err := sayHiFun.Call(sayHiFun, "cheng")

	fmt.Println(returnObj.IsString())
	fmt.Println(returnObj.IsFunction())

	sayHiVal, err := sayHiFun.ToString()

	fmt.Println("sayHiVal->", sayHiVal)*/

}
