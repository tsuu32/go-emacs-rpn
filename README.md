# go-emacs-rpn

Eval Reverse Poland Notation in emacs.

This package uses modified [rpn.go](https://gist.github.com/uchan-nos/cd594b3a4c88af136bd4).

## Build
```sh
make
```

## Usage
```sh
emacs -L .
```

and

```emacs-lisp
(require 'go-emacs-rpn)

(go-emacs-rpn-eval "1 1 +")
;; => (2)

(go-emacs-rpn-eval "2 21 * 30 -")
;; => (12)

(go-emacs-rpn-eval "40 2 / 1 2 3 4 + + +")
;; => (20 10)
```

## note
When stack underflow, emacs will crash.
