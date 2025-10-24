OUTPUT_FORMAT("binary");

BASE = 0x00;

SECTIONS
{
    . = BASE;
    .text : {
        . = BASE;
        *(.entry)
        *(.text)
        *(.func)
        *(.rdata)
    }
}