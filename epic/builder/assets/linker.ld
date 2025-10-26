/* binary */
OUTPUT_FORMAT(pei-x86-64)
ENTRY(main)

BASE = 0x00;

SECTIONS
{
    . = BASE;
    .text : SUBALIGN(1) {
        . = BASE;
        *(.entry)
        *(.text*)
        *(.rdata)
        *(.rdata*)
    }
}
