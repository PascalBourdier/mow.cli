package cli

import (
	"fmt"
	"reflect"
	"strings"
)

type arg struct {
	name   string
	desc   string
	envVar string

	value reflect.Value
}

type BoolArg struct {
	BoolParam

	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument
	EnvVar string
	// The argument's inital value
	Value bool
}

type StringArg struct {
	StringParam

	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument
	EnvVar string
	// The argument's inital value
	Value string
}
type IntArg struct {
	IntParam

	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument
	EnvVar string
	// The argument's inital value
	Value int
}
type StringsArg struct {
	StringsParam

	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The argument's inital value
	Value []string
}
type IntsArg struct {
	IntsParam

	// The argument name as will be shown in help messages
	Name string
	// The argument description as will be shown in help messages
	Desc string
	// A space separated list of environment variables names to be used to initialize this argument.
	// The env variable should contain a comma separated list of values
	EnvVar string
	// The argument's inital value
	Value []int
}

/*
BoolArg defines a boolean argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to a bool) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) BoolArg(name string, value bool, desc string) *bool {
	return c.mkArg(arg{name: name, desc: desc}, value).(*bool)
}

/*
StringArg defines a string argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to a string) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringArg(name string, value string, desc string) *string {
	return c.mkArg(arg{name: name, desc: desc}, value).(*string)
}

/*
IntArg defines an int argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to an int) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntArg(name string, value int, desc string) *int {
	return c.mkArg(arg{name: name, desc: desc}, value).(*int)
}

/*
StringsArg defines a string slice argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to a string slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) StringsArg(name string, value []string, desc string) *[]string {
	return c.mkArg(arg{name: name, desc: desc}, value).(*[]string)
}

/*
IntsArg defines an int slice argument on the command c named `name`, with an initial value of `value` and a description of `desc` which will be used in help messages.

The result should be stored in a variable (a pointer to an int slice) which will be populated when the app is run and the call arguments get parsed
*/
func (c *Cmd) IntsArg(name string, value []int, desc string) *[]int {
	return c.mkArg(arg{name: name, desc: desc}, value).(*[]int)
}

func (arg *arg) match(args []string, c parseContext) (bool, int) {
	if len(args) == 0 {
		return false, 0
	}
	if strings.HasPrefix(args[0], "-") {
		return false, 0
	}
	c.args[arg] = append(c.args[arg], args[0])
	return true, 1
}

func (a *arg) String() string {
	return fmt.Sprintf("ARG(%s)", a.name)
}

func (a *arg) get() interface{} {
	return a.value.Elem().Interface()
}

func (a *arg) set(s string) error {
	return vset(a.value, s)
}

func (c *Cmd) mkArg(arg arg, defaultvalue interface{}) interface{} {
	value := reflect.ValueOf(defaultvalue)
	res := reflect.New(value.Type())

	vinit(res, arg.envVar, defaultvalue)

	arg.value = res

	c.args = append(c.args, &arg)
	c.argsIdx[arg.name] = &arg

	return res.Interface()
}
