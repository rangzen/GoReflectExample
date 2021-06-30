# Introduction

Here is an example of the use of the reflection package.

The program is an Extract Transform Load (ETL) service.
It gets data from a repository.
It calls a lot of functions from the repository and returns data. 

# Source code

Open `service/service_test.go` to see the example.

You can see that I’m using `package service_test` instead of `package service` to be sure that I’m accessing public API.

In the initial tests (see the commented code at the end of the file), I have a lot of duplication (ok, only two but imagine a lot of use cases).
How can I write a helper to decrease this?
How can I include the function to called by service in the helper?
Yes, this idea may seem strange, but it’s just a support.

In the new version (the current not commented code), you can study the helper function.
The explanations are inline.

You can see that now, the test functions are clear and the [Arrange, Act, Assert](https://www.qwant.com/?q=arrange+act+assert) structure is visible again.

Please, don’t hesitate to suggest improvements, especially for the english or for additional explanations.

# Links

## Reflection

* [Documentation of the official package reflect](https://golang.org/pkg/reflect/)
* [The Laws of Reflection from the Go Blog](https://blog.golang.org/laws-of-reflection)

## Packages

* [Testify: test framework](https://github.com/stretchr/testify)
* [mockery: mock generator](https://github.com/vektra/mockery)