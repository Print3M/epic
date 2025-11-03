OUTPUT_FORMAT(pei-x86-64);

/* Entry function */
ENTRY(__main_pic);

BASE = 0x00;

SECTIONS
{
    . = BASE;

    /* Subalign to minimize the gap between functions */
    .text : SUBALIGN(1) { 
        . = BASE;
        *(.start_addr)

        *(.entry)
        *(.text*)
        *(.rdata)
        *(.rdata*)
        
        *(.end_addr)

        /* Just to be sure... */
        KEEP(*(.text$main_pic))
    }
}
