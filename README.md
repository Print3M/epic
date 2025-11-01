# EPIC

EPIC stands for *Extensible Position Independent Code*.

// TODO: IMG ze schematem działania

// TODO: Short description - what problem it solves.
// TODO: Features

## Quick Start

```bash
# 1. Create initial project structure
mkdir project/
epic init project/

# 2. Write your code!

# 3. Compile PIC code
mkdir output/
epic pic-compile project/ -o output/ -m hello

# 4. Link PIC code into standalone `payload.bin`
epic pic-link output/ -o output/ -m hello
```

Great! The job is done. At this point you are ready to take the generated `payload.bin` and inject it into your custom shellcode loader or...

```bash
# 5. [optional] Inject PIC payload and compile a simple loader template
epic loader payload.bin -o output/
```

The compiled `loader.exe` file is ready to be executed! If your payload works in this case, it means it will work everywhere.

## Documentation

### EPIC Commands

#### `init <path>`

Example: `epic init project/`

Create a project structure. It's a way to start a new PIC project. It includes basic example of usage. `<path>` is the target where project structure is created. It creates project in C++ but you can change it to C just changing file extensions and removing C++ features like namespaces. Built-in EPIC headers are compatbile with C and C++. IMPORTANT: It creates directories structure. Set `<path>` to a separate folder to keep things clean.

#### `pic-compile <path>`

Example: `epic pic-compile project/ -o objects/`

Compile all source files from project `<path>` and save object files in the `--output <path>`. IMPORTANT: The output structure of object files directly mimics project structure. Save them rather in a separate folder (let's say `output/`) just to keep things clean. It compiles all modules.

Flags:

* `-o / --output <path>` [required] - path where the compiled objects file will be saved.

#### `pic-link <path>`

Example: `epic pic-link objects/ -o output/`

Link core and selected modules from `<path>` together into a standalone PIC payload. The `<path>` in this command is the output path of the `pic-compile` command. It links only those modules specified in `-m` flag.

IMPORTANT: It creates also folder `assets/` in the output path where linker map, linker script and intermediate executable is stored. Just for a debugging purposes if you want to investigate what exactly is linked into your payload.

Flag:

* `-o / --output <path>` [required] - path where the output payload will be saved.
* `-m / --modules <modules>` - comma-separated list of modules to be linked. Modules are named after their folders in `modules/` directory.

#### `loader <path>`

Example: `epic loader output/payload.bin -o output/`

Inject your `<path>` payload into loader template and compile to Windows executable. It's a great way to test and debug your payload in the wild quickly. If your PIC payload works with loader template it should work with your own custom loader.

IMPORTANT: It creates also folder `assets/` in the output path where the source `loader.c` is saved. This is the file that being compiled. You can check how it looks like and what it does.

Flags:

* `-o / --output <path>` [required] - path where the output loader executable is saved.

#### `monolith <path>`

Example: `epic monolith project/ -o output/`

Compile your project into standard non-PIC executable. In EPIC context it's called `monolith`. It's completely separated process from `pic-compile` and `pic-link`. Because it's a standard non-PIC executable you can use standard libc functions like `printf()` for debugging. Monolith is used for debugging and testing whether your payload does anything at all when you see nothing on the screen.

It compiles all modules at once.

Flags:

* `-o / --output <path>` [required] - path where the output monolith executable is saved.

#### Global flags

In addition to flags specific to a given command, there are several global flags that can be used with any command. All of them are optional.

* `--debug` - enable verbose debug mode.
* `--mingw-w64-gcc <path>` - specify path to MinGW-w64 GCC tool.
* `--mingw-w64-ld <path>` - specify path to MinGW-w64 ld tool.
* `--mingw-w64-objcopy <path>` - specify path to MinGW-w64 objcopy tool.
* `--no-banner` - disable EPIC banner.
* `--no-color` - disable colors output.

### EPIC Coding Guide

#### Where to start?

Use `epic init <path>` to create basic and correct EPIC project structure with all features, header files and entry point.

Created structure:

```text
core/main.cpp    <-- Entry point to your code
include/
    libc/*       <-- All libc headers (PIC-compatbile)
    win32/*      <-- All WinAPI headers (PIC-compatible)
    epic.h       <-- All EPIC macros
modules/
    <your_module_1>/*
    <your_module_2>/*
    ...
```

Entry point to your shellcode is placed in `core/main.cpp`. `main_pic()` is the function where you can start placing your code.

```c
// EPIC: Entry point
SECOND_STAGE void main_pic() {
    // Your code here!
}
```

**IMPORTANT 1**: Nigdy nie implementuj funkcji o nazwie `main()`! Spowoduje to dziwne błędy. Szerzej wyjaśniam to w FAQ na dole.

**IMPORTANT 2**: Nie usuwaj funkcji `__main_pic()` i `WinMain`. Są one niezbędne do prawidłowej kompilacji.

#### Global variables

W kodzie PIC nie możesz używać globalnych read-write variables:

```c
// Global read-write variables are fobidden!
int counter = 5;

void inc_counter() {
    counter += 1;
}
```

Możesz używać jedynie globalnych stałych i stałych literałów:

```c
// This constant is allowed
const char *name = "EPIC";

void calc() {
    // This constant is also allowed
    exec("calc.exe");
}
```

Globalne zmienne są jednak przydatne i potrafią znacznie uprościć kod. Na szczęście EPIC ma rozwiązanie!

```c
typedef struct {
    const char *name;
} Context;

void child_func() {
    Context* ctx = (Context *) GET_GLOBAL(); // Get global variable

    exec(ctx->name);
}

SECOND_STAGE void main_pic() {
    Context ctx;
    SAVE_GLOBAL(ctx); // Make local variable global

    ctx.name = "calc.exe";

    child_func();
}
```

What sorcery is this?! It's a pretty simple compiler trick. EPIC uses one of the CPU registers to keep pointer to your local variable: `SAVE_GLOBAL(var)`. As long as your local variable is present on the stack you can access it using `GET_GLOBAL()` macro.

> **IMPORTANT 1**: Your variable must be on the stack, so the function that initialized your "global" variable cannot return. Therefore, I recommend always using `SAVE_GLOBAL(var)` right in `main_pic()`. This way you make it available to all other functions throughout the entire execution of your shellcode.

> **IMPORTANT 2**: You can use `SAVE_GLOBAL(var)` only once. This mechanism should be used exclusively for maintaining the global context. However, the size and content of the global context size is up to you.

#### Modularity

Podstawowym featurem narzędzia EPIC jest modularność. Wystarczy raz skompilować kod (`pic-compile`) a następnie możesz łączyć ze sobą klocki podczas linkowania (`pic-link`) jak chcesz. Świetnie, ale jak ten sam kod może działać w obu przypadkach?

Otóż po pierwsze wszystkie eksportowane funkcje z modułu oznacz makrem `MODULE`. Wszystkie, bez wyjątku, to pozwoli Ci uniknąć debugowania w przyszłości. Przykładowo plik `modules/exec/exec.h` będzie wyglądał tak:

```c
#include <epic.h>

MODULE int happy_little_function(char *arg);
```

That's it. Teraz w innym miejscu w kodzie możesz sprawdzić czy moduł (a dokładnie dana funkcja) został załadowana.

```c
#include <epic.h>
#include <modules/exec/exec.h>

void test() {
    if (EXISTS(happy_litle_function)) {
        // Module is linked and function is available!
    }
}
```

To potężny, ale bardzo prosty w użyciu mechanizm, dzięki któremu twój PIC projekt może być całkowicie modularny.

> **IMPORTANT**: Funkcja eksportowana z modułu MUSI być oznaczona tagiem  `MODULE`, inaczej pojawią się dziwne błędy, a makro `EXISTS(func)` nie zadziała.

Każdy moduł musi znajdować się w osobnym folderze w `modules/`. Nazwy folderów to jednocześnie nazwy modułów, których używamy następnie z poleceniem `pic-link` w parametrze `-m` (np. `-m exec`). Zauważ, że pisząc modularny kod wystarczy, że raz skompilujemy projekt (`pic-compile`), a następnie możemy  modyfikować nasz PIC payload używając wielokrotnie `pic-link` z innym zestawem modułów!

#### Header files (`libc` and `win32`)

Nie możesz używać biblioteki standardowej `libc`. Nie masz dostępu do żadnych normalnych funkcji tj. `printf()` lub `malloc()`, wszystko musisz sam sobie znaleźć w pamięci. Dlaczego? Szczegółowo wyjaśniłem to [w tym artykule](https://print3m.github.io/blog/x64-winapi-shellcoding).

Nie możesz używać domyślnych header files dla MinGW. Dlaczego? Long story short, jest to jeden wielki bloat, który dodaje fragmenty kodu, które powodują znacznie więcej problemów w przypadku PIC niż korzyści. Więcej o tym piszę w FAQ.

Na szczęście EPIC daje Ci minimalistyczną PIC-compatible implementację `libc` i bardzo podstawowe headery `win32`. Znajdują się one w folderze `include/*` twojego projektu. Są one wygodne i w pełni bezpieczne do użycia w kodzie PIC. Oczywiście możesz je modyfikować w swoim projekcie jak chcesz, a nawet nie musisz w ogóle ich używać.

Wszystkie makra samego EPIC znajdują się w `include/epic.h`.

Przykład importu w dowolnym miejscu projektu:

```c
#include <libc/stdint.h>
#include <win32/windows.h>
#include <epic.h>
```

#### Symbole preprocesora

Kiedy używasz polecenia `monolith` kompilator automatycznie definiuje symbol preprocesora `MONOLITH`. Możesz go użyć do pisania kodu tylko pod kompilację monolith. Przykład:

```c
#ifdef MONOLITH
#include <stdio.h>
#endif

void func() {
    #ifdef MONOLITH
    printf("This is printed only when compiled as monolith!");
    #endif
}
```

Kiedy używasz `pic-compile` kompilator definiuje symbol preprocesora `PIC`. Jest to mniej przydatne, ale możesz go użyć do wykluczenia czegoś z kompilacji `monolith`.

#### Mixing C and C++

It's possible to use C, C++ or both in the same project. Just calling C++ functions from C remember to use `extern "C"` in the header of C++ function to avoid name mangling and linking errors. I generally recommend to stick to one of these languages for the entire project but you are free human being.

### Troubleshooting

1. Clean directory with your object file and run `pic-compile` + `pic-link` again.
2. Test monolith version (more reliable).
3. Check if you follow EPIC Shellcoding Guide.
4. Run `--debug`.
5. Check linker map.

### EPIC Limitations

* Supported platform: x86-64
* C / C++ languages

### FAQ

#### Funkcja importowana z modułu się nie wykonuje

Sprawdź linker map file w `assets/` (po wykonaniu `pic-link`) czy ta funkcja w ogóle jest linkowana do finalnego PIC payloadu. Jeśli jest linkowana, to problem na 95% jest w twoim kodzie. To znaczy, że z jakiegoś powodu doszło do dead code elimination. Sprawdź czy nie wywołujesz funkcji C++ z kodu C bez `extern "C"`.

#### Czy muszę używać inline functions w shellcode?

Nie. Stosuj funkcje jak chesz, tak jak w normalnym kodzie, pamiętaj tylko, żeby funkcje eksportowane z modułu oznaczać makrem `MODULE`.

### Co to jest `monolith`?

Monolith is simply a compilation of your entire project, just like any other non-PIC code. Monolith is a standard `.exe` file that can use the standard library, global variables, and so on. So why use `monolith`? For debugging purposes only :)

#### Why are global variables not usable in PIC payload?

Zmienne globalne potrzebują pamięci typu RW, żeby działać. Istotą shellcode'u jest to, że wystarczy zaalokować pamięć RX i go wykonać, nie trzeba martwić się o specjalne sekcje RW. Istnieją pewne tricki np. wykorzystany w Stardust project, który polega na tym, że shellcode już w trakcie wykonania sam sobie zmienia uprawnienia do sekcji `.bss` i `.data`, że by móc korzystać ze zmiennych globalnych.

To podejście ma jednak wady, których chciałem uniknąć:

1. Podejście wymaga dodatkowego kodu ;p
2. Shellcode musi wykonać funkcję Windows API do zmiany uprawnień stron pamięci. Jest to dodatkowa informacja dla kernela. Chciałem tego uniknąć.
3. Po trzecie shellcode musi być trochę większy, żeby sekcja `.bss` i `.data` zaczynały się od nowej strony pamięci.

Po to stworzyłem CPU-based global variable...

#### Why are `pic-compile` and `pic-link` separate commands?

Możesz raz skompilować, ale wielokrotnie linkować do różnych modułów tworząc za każdym razem unikatowy shellcode.

#### Why is PIC extracted from a PE file?

Dead code elimination doesn't work for linker output "binary".

To hack this I use MinGW-w64 toolchain `gcc` with custom linker script (`ld`) to PE and then extracting PIC `.text` section using `objcopy` to final `payload.bin` output. It works like a charm. This way the final payload is smaller then ever.

#### Do I have to manually align the stack before calling a Windows API?

Nie, Mingw robi to automatycznie. Musisz jednak oznaczyć każdą funkcję Windows API makrem `WINAPI`.

#### Why is the entry point called `__main_pic` and not simply `main()`?

If you implement `main()` function no matter how hard I tried it's always treated special by GCC compiler. No matter how many compiler flags, function attributes and linker scripting I used there's always generated unnecessary call to `__main()` at the beginning of `main()`. It means you need to implement this stupid `__main()` somehow, otherwise there's an linker error. The reason for this is behaviour is unknown and I found no way to disable it.

Solution to this problem is using `main()` as entry point and implement dummy empty `__main() {}` function. It works, but honestly I wanted my code to be as clean as possible with no dummy functions!

Another solution is not to use `main()` at all. I created `__main_pic()` function and EPIC uses it as an entry point. It works flawlessly.

#### Why does EPIC implement its own `libc` and `win32` headers instead of using those from MinGW?

Oczywiście shellcode nie może mieć żadnych zależności, ale teoretycznie EPIC mógłby używać definicji typów i makr z domyślnych plików nagłówkowych MinGW, prawda? Nie, ponieważ domyślne header files są jednym wielkim bloatem, śmietnikiem i dodają kod bez Twojej wiedzy. Powodują błędy w kompilacji, nawet jeśli używasz tylko definicji typów i makr. Poza tym zachęcają do używania funkcji, których nie możesz użyć, np, `printf()`.

Just including default Windows MinGW-w64 header files throws some errors during compilation. For example, it requires SSE to be enabled to compile successfully and I want it to be disabled. This is the reason why I don't use default Windows headers but include only custom ones.

#### Can I check exactly which functions are linked to the PIC payload?

Tak.

It's possible to generate map of linked sections. Great tool for deep inspection of linker's work. Use `-Map=linker.map` parameter. This file shows which sections (section == function when used with `-ffunction-sections`) are discarded and which are linked into the final payload. It shows the layout of linked sections and their size. Great tool for debugging.

#### Can I manually disassemble the PIC payload?

Yes. Use this command:

```bash
objdump -D -b binary -m i386:x86-64 -M intel payload.bin
```

## Credits

* c-to-shellcode.py
* Stardust
