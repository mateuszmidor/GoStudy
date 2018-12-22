# This is the Go workspace: GoStudy. 

To let Go know the workspace location, do:  
    export GOPATH="/home/user/SoftwareDevelopment/GoStudy"  

To let shell know go tools, do:
    export PATH="/home/user/SoftwareDevelopment/GoStudy/bin:$PATH" 

To fix breaking builds eg "go can't find package fmt", do:  
    unset GOROOT  

More info at:  
+ https://golang.org/doc/code.html  
+ https://www.youtube.com/watch?v=XCsL89YtqCs  

To build the hello example without outputting any file: 
    go build github.com/mateuszmidor/GoStudy/hello/main 
 
To install the hello example under workspace bin/: 
    go install github.com/mateuszmidor/GoStudy/hello/main 

To run the hello example:
    go run github.com/mateuszmidor/GoStudy/hello/main 

To test the hello string package: 
    go test github.com/mateuszmidor/GoStudy/hello/string 