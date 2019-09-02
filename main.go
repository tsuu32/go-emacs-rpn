// go build -buildmode=c-shared -o go-emacs-rpn.so
package main

// int plugin_is_GPL_compatible;
import "C"

import (
	"github.com/sigma/go-emacs"
)

func init() {
	emacs.Register(initModule)
}

func initModule(env emacs.Environment) {
	env.RegisterFunction("go-emacs-rpn-eval", RpnEval, 1, "eval reverse poland notation by go", nil)
	env.ProvideFeature("go-emacs-rpn")
}

func RpnEval(ctx emacs.FunctionCallContext) (emacs.Value, error) {
	env := ctx.Environment()
	s, err := ctx.GoStringArg(0)
	if err != nil {
		return nil, err
	}

	stack := eval([]byte(s))

	emacs_stack := make([]emacs.Value, 0)

	for i := range stack {
		emacs_stack = append(emacs_stack, env.Int(int64(stack[i])))
	}

	return env.StdLib().List(emacs_stack...), nil
}

func main() {}
