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
    }
}
