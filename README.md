# EPIC - Extensible Position Independent Code

## Interesting Observations

I tried two approaches to linking functions:

1. With `__attribute__((section(".func")))` on the every function in the code and the `*(.func)` in the linker script.

2. With no attribute in code but with `-ffunction-sections` parameter during compilation and `--gc-sections` during linking.

The first option creates smaller paylaod but doesn't provide dead function elimination.

The second option creates a little bit larger payload but provides dead function elimination.

Why does the second option create a little bit larger payload? In this approach each function is placed in a separate section. This is the outcome of using `-ffunction-sections` parameter. Each section typically requires alignment (often 16 bytes or more). If you have many small functions, you accumulate significant padding after each one.

TODO: Check if this is really the difference between alignment

## Disassembly payload.bin

```bash
objdump -D -b binary -m i386:x86-64 -M intel payload.bin
```

## Linker map

It's possible to generate map of linked sections. Great tool for deep inspection of linker's work. Use `-Map=linker.map` parameter. This file shows which sections (section == function when used with `-ffunction-sections`) are discarded and which are linked into the final payload. It shows the layout of linked sections and their size. Great tool for debugging.

## Dead code elimination doesn't work

Dead code elimination doesn't work for:

1. MinGW PE/COFF target
2. Linker output "binary"

To hack this I use `gcc` with custom linker script (`ld`) to ELF and then extracting PIC `.text` section using `objdump` to final `payload.bin` output. It works like a charm. This way the final payload is smaller then ever. This is the reason why I don't use MinGW for payload building.

## Stack alignment for Win API

Mingw-w64 automatically handles stack alignment when function is defined with `WINAPI` attribute.

## Troubleshooting

1. Clean output/ directory.
