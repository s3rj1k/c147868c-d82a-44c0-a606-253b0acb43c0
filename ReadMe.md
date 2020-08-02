## Result

Assignment is implemented as a simple REST backend application that is exposed on port 8080 and written in GoLang.
Golang was chosen because it is my language of a choice for writing code.

Implementation is split into 2 major layers, HTTP layer and algoritm layer.

For HTTP layer `gin-gonic` was chosen for the sake of simplicity.
For algoritm layer simple structure based approach (with methods) was chosen.

Algoritm layer is covered in test as a verification of expected result.

GolangCI is used a as linter to check code quality.

Simple makefile is used to build, update, lint this project.
