# AGENTS.md

## The repository purpose
* this repository is dedicated for small and simple Golang demo programs that demonstrate golang-ecosystem, tools, libraries, techniques
* in short: it's dedicated for studying Go language.

## The repository structure
* as a rule of thumb: every small demo program resides within it's own directory right under the repo's root, and is has it's own dedicated go module files: go.mod and go.sum. Example:
    ```sh
    ./errgroup/
    ├── go.mod
    ├── go.sum
    └── main.go
    ```

* collections of related demo programs should reside under dedicated parent folders. Example of such collections are:
    ```sh
    ./
    ├── GoProgrammingBlueprints/ # (programs from book on golang)
    ├── TheGoProgrammingLanguage/ #(programs from another book on golang)
    └── jose/ #(programs demonstrating JSON Web Encryption, Signature, Token)
  ```
* exceptions - there is a "src/" directory in the repo's root - it's legacy, from the times Go didn't support modules, it should not be used for new go demos