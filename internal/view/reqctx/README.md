# reqctx

`reqctx` is a utility package designed to manage request-specific context values
in this project. It helps to encapsulate authentication status and user
information within the Echo request context in a type-safe manner.

## Purpose

When developing web applications, it is common to pass values such as
authentication status and user information through the request lifecycle. Using
Echo's built-in context (`echo.Context`) can lead to potential issues such as
typographical errors, lack of type safety, and reduced code readability.

`reqctx` addresses these issues by providing a structured way to manage these
context values.
