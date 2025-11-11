# Project powered by EPIC

This project was built using **EPIC (Extensible Position Independent Code)** — a shellcode development and building toolkit designed for developer experience, predictability, and modularity. _Write code, EPIC will take care of the rest!_

- Repository & Docs: [github.com/Print3M/epic](https://github.com/Print3M/epic)

EPIC is a standalone binary toolkit:

- Latest EPIC Release: [github.com/Print3M/epic/releases/](https://github.com/Print3M/epic/releases/latest)

## Quick Start

Compile your project:

```bash
epic pic-compile project/ -o output/
```

Link and build the payload:

```bash
epic pic-link output/ -o output/ -m hello
```

That's it! You can now take the generated `payload.bin` and inject it into your custom shellcode loader, or simply:

```bash
# Inject PIC payload and compile a simple loader template
epic loader payload.bin -o output/
```

The compiled `loader.exe` is ready to execute.

---

By [Print3M](https://x.com/Print3M_)
