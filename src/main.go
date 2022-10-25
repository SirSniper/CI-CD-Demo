// main.go

package main

import (
    "strconv"

    "github.com/astaxie/beego"
)


func main() {
    /* This would match routes like the following:
       /sum/3/5
       /product/6/23
       ...
    */
    beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
    beego.Run()
}

type mainController struct {
    beego.Controller
}


func (c *mainController) Get() {

    //Obtain the values of the route parameters defined in the route above
    operation := c.Ctx.Input.Param(":operation")
    num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
    num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

    //Set the values for use in the template
    c.Data["operation"] = operation
    c.Data["num1"] = num1
    c.Data["num2"] = num2
    c.TplName = "results.html"

    // Perform the calculation depending on the 'operation' route parameter
    switch operation {
    case "sum":
        c.Data["result"] = add(num1, num2)
    case "product":
        c.Data["result"] = multiply(num1, num2)
		case "magic":
        c.Data["result"] = magic(num1, num2)
    default:
        c.TplName = "not_found.html"
    }
}

func add(n1, n2 int) int {
    return n1 + n2
}

func multiply(n1, n2 int) int {
    return n1 * n2
}

func magic(n1, n2 int) int {
	if (n1 * 6 <= n2) {
		return n2 - n1
	} else if (n2 > n1) {
		return n1 - n2
	} else if (n1 * 2 - 5 > n2) {
		return n2 + 4 * n1 // should be n2 + 4 * n1
	} else {
		return n1 * n2 * n2
	}
}
