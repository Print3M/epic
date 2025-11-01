# EPIC — Extensible Position Independent Code Framework

EPIC (Extensible Position Independent Code) – implant development and shellcode-building framework designed for developer experience, predictability, and modularity.

EPIC transforms complex shellcode engineering into a seamless process — building fully position-independent payloads with zero hidden magic, zero heap reliance, and maximal clarity. Whether you’re crafting stealth implants or high-level modular payloads, EPIC ensures your binaries remain elegant, efficient, and exact.

- Built-in modularity – you choose what you want to include.
- Built-in global context support – no memory permission changes!
- Built-in dead-code elimination – the smallest payload on the market.
- Predictable PIC generation — no implicit syscalls, no unexpected code.
- Built-in minimal `libc` and `win32` written for PIC compatibility.
- Built-in mixing C and C++ support.
- More...

// - A stable and convenient tool for building your project. 
TODO: wygodny i przewidywalny CI/CD

## Interesting Observations

I tried two approaches to linking functions:

1. With `__attribute__((section(".func")))` on the every function in the code and the `*(.func)` in the linker script.

2. With no attribute in code but with `-ffunction-sections` parameter during compilation and `--gc-sections` during linking.

The first option creates smaller paylaod but doesn't provide dead function elimination.

The second option creates a little bit larger payload but provides dead function elimination.

Why does the second option create a little bit larger payload? I'm not sure but I have a feeling. In the second approach each function is placed in a separate section. This is the outcome of using `-ffunction-sections` parameter. Each section typically requires some alignment. If you have many small functions, you accumulate significant padding after each one.

## Disassembly payload.bin

```bash
objdump -D -b binary -m i386:x86-64 -M intel payload.bin
```

## Linker map

It's possible to generate map of linked sections. Great tool for deep inspection of linker's work. Use `-Map=linker.map` parameter. This file shows which sections (section == function when used with `-ffunction-sections`) are discarded and which are linked into the final payload. It shows the layout of linked sections and their size. Great tool for debugging.

## Dead code elimination doesn't work

Dead code elimination doesn't work for linker output "binary".

To hack this I use MinGW-w64 toolchain `gcc` with custom linker script (`ld`) to PE and then extracting PIC `.text` section using `objdump` to final `payload.bin` output. It works like a charm. This way the final payload is smaller then ever.

## Stack alignment for Win API

Mingw-w64 automatically handles stack alignment when function is defined with `WINAPI` attribute.

## `main()` and `__main()` functions

If you implement `main()` function no matter how hard I tried it's always treated special by GCC compiler. No matter how many compiler flags, function attributes and linker scripting I used there's always generated unnecessary call to `__main()` at the beginning of `main()`. It means you need to implement this stupid `__main()` somehow, otherwise there's an linker error. The reason for this is behaviour is unknonw and I found no way to disable it.

Solution to this problem is using `main()` as entry point and implement dummy empty `__main() {}` function. It works, but honestly I wanted my code to be as clean as possible with no dummy functions!

Second solution is not to use `main()` at all. Create a not-main function (e.g. `__main_pic`) and use it as a entry point. It works this way.

## Troubleshooting

1. Clean output/ directory.
2. Test monolith version (more reliable).
3. Check if you follow EPIC Guidebook.
4. Run `--debug`.
