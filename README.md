# EPIC - Extensible Position Independent Code

```bash
make
```

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

It's possible to generate map of linked sections. Great tool for deep inspection of linker's work. Use `-Map=linker.map` parameter.

## Dead code elimination doesn't work

It doesn't work probably because of  `OUTPUT_FORMAT("binary");`. 