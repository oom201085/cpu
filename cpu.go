
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* Initialize variables */
var (
	cpu1  string
	cpu2  string
	ready []string
	io1   []string
	io2   []string
	io3   []string
	io4   []string
)

func main() {
	initialized() // do initialized() function
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "new":
			for i := range commandx {
				if i == 0 {
					continue
				}
				commandNew(commandx[i])
			}
		case "terminate":
			for i := range commandx {
				if i == 0 {
					continue
				}
				commandTerminate(commandx[i])
			}
		case "expire":
			for i := range commandx {
				if i == 0 {
					continue
				}
				commandExpire(commandx[i])
			}
		case "io1":
			for i := range commandx {
				if i == 0 {
					continue
				}
				commandIo1(commandx[i])
			}
		case "io2":
			for i := range commandx {
				if i == 0 {
					continue
				}
				commandIo2(commandx[i])
			}
		case "io3":
			for i := range commandx {
				if i == 0 {
					continue
				}
				commandIo3(commandx[i])
			}
		case "io4":
			for i := range commandx {
				if i == 0 {
					continue
				}
				commandIo4(commandx[i])
			}
		case "io1x":
			commandIo1x()
		case "io2x":
			commandIo2x()
		case "io3x":
			commandIo3x()
		case "io4x":
			commandIo4x()
		default:
			fmt.Printf("\nCommand error.\n")
		}
	}
}

func initialized() {
	cpu1 = ""
	cpu2 = ""
	ready = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)
}

func showProcess() {
	fmt.Printf("\n-----------\n")
	fmt.Printf("cpu1  -> %s\n", cpu1)
	fmt.Printf("cpu2  -> %s\n", cpu2)
	fmt.Printf("Ready -> ")
	for i := range ready {
		fmt.Printf("%s ", ready[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 1 -> ")
	for i := range io1 {
		fmt.Printf("%s ", io1[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 2 -> ")
	for i := range io2 {
		fmt.Printf("%s ", io2[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 3 -> ")
	for i := range io3 {
		fmt.Printf("%s ", io3[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 4 -> ")
	for i := range io4 {
		fmt.Printf("%s ", io4[i])
	}
	fmt.Printf("\n\nCommand > ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func commandNew(p string) {
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}

}
func commandTerminate(p string) {
	switch p {
	case "cpu1":
		if cpu1 != "" {
			cpu1 = deleteQueue(ready)
		}
	case "cpu2":
		if cpu2 != "" {
			cpu2 = deleteQueue(ready)
		}
	default:
		fmt.Printf("\nArgument error at \"%s\"\n", p)
	}
}

func commandExpire(p string) {
	switch p {
	case "cpu1":
		p := deleteQueue(ready)
		if p == "" {
			return
		}
		insertQueue(ready, cpu1)
		cpu1 = p
	case "cpu2":
		p := deleteQueue(ready)
		if p == "" {
			return
		}
		insertQueue(ready, cpu2)
		cpu2 = p
	default:
		fmt.Printf("\nArgument error at \"%s\"\n", p)
	}
}

func commandIo1(p string) {
	switch p {
	case "cpu1":
		insertQueue(io1, cpu1)
		cpu1 = ""
		commandExpire(p)
	case "cpu2":
		insertQueue(io1, cpu2)
		cpu2 = ""
		commandExpire(p)
	default:
		fmt.Printf("\nArgument error at \"%s\"\n", p)
	}
}

func commandIo2(p string) {
	switch p {
	case "cpu1":
		insertQueue(io2, cpu1)
		cpu1 = ""
		commandExpire(p)
	case "cpu2":
		insertQueue(io2, cpu2)
		cpu2 = ""
		commandExpire(p)
	default:
		fmt.Printf("\nArgument error at \"%s\"\n", p)
	}
}

func commandIo3(p string) {
	switch p {
	case "cpu1":
		insertQueue(io3, cpu1)
		cpu1 = ""
		commandExpire(p)
	case "cpu2":
		insertQueue(io3, cpu2)
		cpu2 = ""
		commandExpire(p)
	default:
		fmt.Printf("\nArgument error at \"%s\"\n", p)
	}
}

func commandIo4(p string) {
	switch p {
	case "cpu1":
		insertQueue(io4, cpu1)
		cpu1 = ""
		commandExpire(p)
	case "cpu2":
		insertQueue(io4, cpu2)
		cpu2 = ""
		commandExpire(p)
	default:
		fmt.Printf("\nArgument error at \"%s\"\n", p)
	}
}

func commandIo1x() {
	p := deleteQueue(io1)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

func commandIo2x() {
	p := deleteQueue(io2)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

func commandIo3x() {
	p := deleteQueue(io3)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

func commandIo4x() {
	p := deleteQueue(io4)
	if p == "" {
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

func insertQueue(q []string, data string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			break
		}
	}
}

func deleteQueue(q []string) string {
	result := q[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
	}
	q[9] = ""
	return result
}
