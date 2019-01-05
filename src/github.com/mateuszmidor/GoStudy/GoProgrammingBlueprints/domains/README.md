# domainfinder 
This app takes a word as an input and converts it into a bunch of possible domain names.  
It uses the following tools:
  - synonyms - generates a bunch of synonyms of a given word
  - sprinkle - generates a word mutation by adding prefixes and postfixes
  - coolify - generates a word mutation by adding/removing vowels
  - domanify - modifies word to make a valid domain name that ends with .com or .net
  - available - checks if given domain is available
Each tool can be used seperately or piped ( | ) to make a chain

# Build
To build all the tools at once, you can run ./buildall.sh

# Run
To generate domains, run ./domainfinder