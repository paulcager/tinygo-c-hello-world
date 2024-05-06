#include "pico/platform.h"
#include "pico/time.h"
#include "pico/multicore.h"

#include <stdio.h>

volatile int count;

int my_core_number() {
    return get_core_num();
}

void core1_main() {
    printf("\n");
    printf("In core1_main\r\n");

    while (true) {
        count++;
        //sleep_ms(500);
        int fake = 0;       // To prevent loop being optimised away.
        for (unsigned long i = 0; i < 25000000; i++) {
            fake = (fake < 1) | i;
        }
        if (fake == 22) {
            printf("Slept\r\n");
        }
    }
}

static uint32_t core1_stack[0x0800];

void launch_core1() {
    printf("Exec...\r\n");
	multicore_reset_core1();
	multicore_launch_core1_with_stack(core1_main, core1_stack, 0x0800);
}

void __assert_func(const char *file, int line, const char *func, const char *failedexpr) {
//    printf("assertion \"%s\" failed: file \"%s\", line %d%s%s\n",
//           failedexpr, file, line, func ? ", function: " : "",
//           func ? func : "");

    for(;;) {}
}