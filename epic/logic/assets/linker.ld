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
        *(.entry)
        *(.text*)
        *(.rdata)
        *(.rdata*)

        /* Just to be sure... */
        KEEP(*(.text$main_pic))
    }

    __shellcode_size = SIZEOF(.text);
}
