package main

import "fmt"

const LINUX_REBOOT_MAGIC1 uintptr = 0xfee1dead
const LINYX_REBOOT_MAGIC2 uintptr = 672274793
const LINUX_REBOOT_CMD_RESTART uintptr = 0x1234567

func main() {
	//syscall.Syscall(syscall.SYS_REBOOT,
	//	LINUX_REBOOT_MAGIC1,
	//	LINYX_REBOOT_MAGIC2,
	//	LINUX_REBOOT_CMD_RESTART)
	fmt.Println(123)
}
