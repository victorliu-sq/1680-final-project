package controlPkg

import (
	"bufio"
	"fmt"
	"os"
	"snowcast/pkg/protocol/CLI"
	"snowcast/pkg/protocol/command"
	"strconv"
)

func (cc *ClientControl) ScanClientCLI(clientCLIChan chan CLI.ClientCLI) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if num, err := strconv.Atoi(line); err == nil {
				fmt.Println("Try to set station", line)
				clientCLI := CLI.NewClientCLI(command.TypeSetStation, uint16(num))
				clientCLIChan <- *clientCLI
			} else if len(line) == 1 && line[0] == 'q' {
				clientCLI := CLI.NewClientCLI(CLI.TypeQuitClient, uint16(0))
				clientCLIChan <- *clientCLI
			}
		}
	}
}
