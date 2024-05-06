package core1

/*
	#cgo CFLAGS: -I/home/paul/pico/pico-sdk/src/rp2_common/hardware_claim/include -I/home/paul/pico/pico-sdk/src/rp2_common/hardware_irq/include -I/home/paul/pico/pico-sdk/src/host/hardware_sync/include -I/home/paul/pico/pico-sdk/src/rp2040/hardware_structs/include -I/home/paul/pico/pico-sdk/src/rp2_common/hardware_timer/include -I/home/paul/pico/pico-sdk/src/rp2_common/pico_multicore/include -I/home/paul/pico/pico-sdk/src/common/pico_sync/include -I/home/paul/pico/pico-sdk/src/common/pico_time/include -I/home/paul/pico/pico-sdk/src/rp2_common/boot_stage2/asminclude -I/home/paul/pico/pico-sdk/src/rp2040/hardware_regs/include -I/home/paul/pico/pico-sdk/src/rp2_common/hardware_base/include -I/home/paul/pico/pico-sdk/src/common/pico_base/include -I/home/paul/pico/hid-proxy2/build/generated/pico_base -I/home/paul/pico/pico-sdk/src/boards/include -I/home/paul/pico/pico-sdk/src/rp2_common/pico_platform/include -I/home/paul/pico/pico-sdk/src/rp2_common/boot_stage2/include
*/
import "C"

import (
	"github.com/paulcager/tinygo-c-hello-world/core1/common/pico_time"
	"github.com/paulcager/tinygo-c-hello-world/core1/common/pico_util"
	"github.com/paulcager/tinygo-c-hello-world/core1/rp2_common/hardware_irq"
	"github.com/paulcager/tinygo-c-hello-world/core1/rp2_common/hardware_sync"
	"github.com/paulcager/tinygo-c-hello-world/core1/rp2_common/hardware_timer"
	"github.com/paulcager/tinygo-c-hello-world/core1/rp2_common/pico_multicore"
)

func init() {
	pico_multicore.Init()
	hardware_irq.Init()
	hardware_sync.Init()
	pico_time.Init()
	hardware_timer.Init()
	pico_util.Init()
}

func DoIt() int {
	return 42
}
