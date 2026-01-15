package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter folder name: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	folderName := strings.TrimSpace(input)

	err = os.Mkdir(folderName, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	mainFilePath := filepath.Join(folderName, "main.go")
	mainFileContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello from ` + folderName + `")
}
`

	err = os.WriteFile(mainFilePath, []byte(mainFileContent), 0644)
	if err != nil {
		fmt.Println("Error creating main.go:", err)
		return
	}

	cmd := exec.Command("go", "mod", "init", folderName)
	cmd.Dir = folderName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running go mod init:", err)
		return
	}

	fmt.Println("Project created successfully ✔")
}
