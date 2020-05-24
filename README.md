# Golang structs for Open API specification v3

There are many opensource golang packages implement Open API v3 tooling for validation, marshal, and unmarshal functions.

Most of those projects address specific problem, or implement functionality around tooling support, therefore the data structures are bloated with specific method implementations.

Aim of this repository to maintain golang data structure for Open API specification version 3.x without specific tooling functions.

Parsing, validation, and implementation of functions are upto the reference project.

However, please note at the moment this is an experimental repository.

## Project stucture

Project structure is based on [golang-standards](https://github.com/golang-standards/project-layout) reference docs.

### [pkg](https://github.com/Mahifernando/openapispec/tree/master/pkg)

Provides data structures for [Open API specification version 3.0.3](http://spec.openapis.org/oas/v3.0.3)