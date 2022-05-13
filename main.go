package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

const testfilePath = "tmp/testfile"

func main() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case sig := <-sigs:
			fmt.Println(sig)
			return
		case <-ticker.C:
			err := os.WriteFile(testfilePath, []byte("Some fake content"), 0755)
			if err != nil {
				fmt.Printf("error writing test file: %v\n", err)
				continue
			}
			fmt.Println("created testfile")

			err = os.Remove(testfilePath)
			if err != nil {
				fmt.Printf("error removing test file: %v\n", err)
				continue
			}
			fmt.Println("removed testfile")

			out, err := exec.Command("df", "-h").Output()
			if err != nil {
				fmt.Printf("error running df: %v\n", err)
				continue
			}
			fmt.Println(out)
		}
	}
}
