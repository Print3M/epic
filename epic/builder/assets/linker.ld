/* binary */
OUTPUT_FORMAT("elf64-x86-64")
ENTRY(main)

BASE = 0x00;

SECTIONS
{
    . = BASE;
    .text : {
        . = BASE;
        *(.entry)
        *(.text.*)
        /* *(.func) */
        *(.rodata)
    }
}
