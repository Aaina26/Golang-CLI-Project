package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	/*
	   *********************
	   Step1: Display Prompt
	   *********************
	   Output:
	   Infosys Cloud CLI
	   ------------------
	   =>
	   Toolkit:
	   Package-fmt|Function-Println
	*/
	fmt.Println("Infosys Cloud CLI")
	fmt.Println("------------------")
	fmt.Print("=>")
	/*
	   *********************************
	   Step 2: Accept Command Line Input
	   *********************************
	   Input:command string
	   Example:  =>exit
	   Toolkit:
	   Package-bufio|Function-NewScanner|Output Variable-scanner
	   Package-bufio|struct-Scanner|Function-Scan
	   Package-bufio|struct-Scanner|Function-Text|Output Variable-userInput
	*/
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput := scanner.Text()
		if strings.Compare("exit", userInput) == 0 {
			os.Exit(0)
		}
		commandString := strings.Split(userInput, " ")
		if len(commandString) < 2 {
			fmt.Println("Argument is required")
		} else if strings.Compare("checkconnect", commandString[0]) == 0 {
			commandOutput := execPINGCommand(commandString[1])
			fmt.Println(commandOutput)
		} else {
			fmt.Println("Command not supported by the CLI")
		}
		fmt.Print("=>")
	}
	/*
	   ******************************
	   Step 3: Implement exit Command
	   ******************************
	   Toolkit:
	   Package-strings|Function-Compare
	   Package-os|Function-Exit
	*/

	/*
	   ******************************************
	   Step 4: Implement checkconnect command
	   ******************************************
	   Toolkit:
	   Package-strings|Function-Split|Output Variable-commandString
	   Package-strings|Function-Compare
	   UserDefined Function-execPINGCommand|Input Variable-commandString|Output Variable-commandOutput
	*/
}

func execPINGCommand(cmdParam string) string {
	/*
	   ***************************************
	   Step 5: Execute Windows OS command ping
	   ***************************************
	   Toolkit:
	   Package-os/exec|Function-Command|Output Variables-c,err
	   Package-os/exec|struct-Cmd|Function-CombinedOutput
	*/
	if c, err := exec.Command("C:\\Windows\\System32\\ping", "-n", "3", cmdParam).CombinedOutput(); err != nil {
		s := "Error! Command Execution failed"
		fmt.Println(err)
		return s
	} else {
		return string(c)
	}
}
