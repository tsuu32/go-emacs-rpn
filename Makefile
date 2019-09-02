GO ?= go
EMACS ?= emacs

all: go-emacs-rpn.so

go-emacs-rpn.so: main.go rpn.go
	$(GO) build -buildmode=c-shared -o $@

clean:
	$(RM) go-emacs-rpn.so

test: go-emacs-rpn.so
	$(EMACS) -Q --batch -L . -l go-emacs-rpn --eval \
	'(message (format "%s" (go-emacs-rpn-eval "2 21 * 30 -")))'
