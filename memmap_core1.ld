SECTIONS
{

    /* .stack*_dummy section doesn't contains any symbols. It is only
     * used for linker to calculate size of stack sections, and assign
     * values to stack symbols later
     *
     * stack1 section may be empty/missing if platform_launch_core1 is not used */

    /* by default we put core 0 stack at the end of scratch Y, so that if core 1
     * stack is not used then all of SCRATCH_X is free.
     */



    /* stack limit is poorly named, but historically is maximum heap ptr */
    __StackOneTop = 0x20040000 + 4k;
    __StackOneBottom = __StackOneTop - SIZEOF(.stack1_dummy);

    /* todo assert on extra code */
}

