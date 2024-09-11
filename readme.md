## GoHttpLite

# Description
This is a simple and lightweight library designed to create simple HTTP applications.  
It is more of a side project than an actual library.  
My goal in the development of this project was to gain a deeper understanding of how HTTP works in general.

# Features

**Routing**:
    This library offers robust routing capabilities, enabling the creation of multiple routers to logically separate different parts of your application's functionality. This not only helps in maintaining a clean and modular codebase but also simplifies the management of complex applications by grouping related routes together.

**Variables in Path**:
    With built-in support for extracting variables directly from URL paths, the library allows for dynamic routing. These path variables are seamlessly passed along with the request, making it easy to handle user-specific data or tailor responses based on URL parameters.

**Middleware**:
    Middleware is a key feature that lets you insert custom logic into the request/response cycle. Whether you need to log requests, implement authentication, or modify requests before they reach your handlers, middleware provides a flexible way to manage cross-cutting concerns in your application.
