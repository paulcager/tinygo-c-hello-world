package main

/*
	#cgo CFLAGS: -I/home/paul/pico/pico-sdk/src/rp2_common/hardware_claim/include -I/home/paul/pico/pico-sdk/src/rp2_common/hardware_irq/include -I/home/paul/pico/pico-sdk/src/host/hardware_sync/include -I/home/paul/pico/pico-sdk/src/rp2040/hardware_structs/include -I/home/paul/pico/pico-sdk/src/rp2_common/hardware_timer/include -I/home/paul/pico/pico-sdk/src/rp2_common/pico_multicore/include -I/home/paul/pico/pico-sdk/src/common/pico_sync/include -I/home/paul/pico/pico-sdk/src/common/pico_time/include -I/home/paul/pico/pico-sdk/src/rp2_common/boot_stage2/asminclude -I/home/paul/pico/pico-sdk/src/rp2040/hardware_regs/include -I/home/paul/pico/pico-sdk/src/rp2_common/hardware_base/include -I/home/paul/pico/pico-sdk/src/common/pico_base/include -I/home/paul/pico/hid-proxy2/build/generated/pico_base -I/home/paul/pico/pico-sdk/src/boards/include -I/home/paul/pico/pico-sdk/src/rp2_common/pico_platform/include -I/home/paul/pico/pico-sdk/src/rp2_common/boot_stage2/include
	#include "pico/multicore.h"
	#include "pico/time.h"
	int my_core_number();
	void launch_core1();
	extern volatile int count;
*/
import "C"
import (
	"fmt"
	"github.com/paulcager/tinygo-c-hello-world/core1"
	"time"
)

//	#cgo LDFLAGS: -L. -l:memmap_core1.ld

func main() {
	time.Sleep(5 * time.Second)
	fmt.Println("OK, I am running on core", C.my_core_number())
	time.Sleep(1 * time.Second)
	C.launch_core1()
	fmt.Println("Core1 running!")

	core1.DoIt()

	for {
		fmt.Printf("\rCount %d\r\n", C.count)
		time.Sleep(1 * time.Second)
	}
}
