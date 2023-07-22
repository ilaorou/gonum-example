package comm

import "os/exec"

func OpenImage(filename string) error {
	cmd := exec.Command("cmd", "/c", "start", filename)
	// 执行命令
	return cmd.Start()
}
