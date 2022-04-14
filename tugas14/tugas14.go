package main

import "os/exec"

func main() {
	hasil, _ := exec.Command("ipconfig").Output()
	print(string(hasil))
	hasil, _ = exec.Command("help").Output()
	print(string(hasil))
	hasil, _ = exec.Command("getmac").Output()
	print(string(hasil))
}
