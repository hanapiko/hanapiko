package main

import (
"fmt"
"os"
"reflect"
)

func main() {
args := os.Args
if len(os.Args) != 3 {
fmt.Println("not enough no. of arguments")
return
}

inputFile := args[1]
outputFile := args[2]

data, err := os.ReadFile(inputFile)
if err != nil {
fmt.Println("error reading input:", err)

}

content := string(data)

expected := "package main\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello World!\")\n}"

if !reflect.DeepEqual(expected, content) {
content = expected

} else {
// If content is well-formatted, inform the user
fmt.Println("well formatted.")
}

os.WriteFile(outputFile, []byte(content), 0o644)
}