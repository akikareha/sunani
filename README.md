# Sunani

Sunani is a minimal virtual computing environment built around WebAssembly.

It provides a small and portable runtime where almost all application logic
lives inside WebAssembly, while the host only offers a minimal set of services
such as drawing and input handling.

Sunani is designed as an experimental and exploratory project rather than a
full-featured framework.

---

## Concept

The core idea of Sunani is simple:

> **Keep the host small. Push complexity into WebAssembly.**

Sunani deliberately exposes only a minimal host API.
Instead of adding rich functionality to the host,
higher-level abstractions are expected to be implemented
inside the WebAssembly module itself.

This makes Sunani:

- Easy to reason about
- Consistent across platforms
- Suitable for experimentation and learning

---

## Hosts

The same WebAssembly module can run on different Sunani hosts.

Currently supported hosts:

- **Go (native)**  
  A desktop host using Go for windowing, input, and rendering.

- **Web (JavaScript + Canvas)**  
  A browser-based host using HTML5 Canvas and JavaScript.

Both hosts aim to provide the same minimal API surface.

---

## Host API

The Sunani host provides minimal functionality for:

- Basic shape drawing
- Framebuffer-based pixel output
- Keyboard input
- Mouse input
- Basic system lifecycle events

The API is intentionally small so that it can be implemented
consistently on multiple platforms.

See:

* [Sunani API](api.md)

---

## Demos

Sunani includes demos that demonstrate programs running almost entirely
inside WebAssembly.

In these demos:

- Rendering logic lives in WASM
- Input handling is driven from WASM
- The host only forwards events and drawing primitives

See:

- [Sunani Demo](https://akikareha.github.io/sunani/)

---

## Design Goals

Sunani focuses on the following goals:

- Minimal and explicit APIs
- WebAssembly as the primary execution environment
- Portability across hosts
- Encouraging experimentation and creative coding

Sunani does **not** aim to be:

- A general-purpose GUI framework
- A replacement for existing game engines
- A production-ready runtime (at least for now)

---

## Use Cases

Possible use cases include:

- Experimental virtual machines
- Minimal graphical applications
- Creative coding experiments
- Learning projects around WebAssembly
- Prototyping small "OS-like" environments inside WASM

---

## Status

Sunani is an evolving personal project.

APIs may change, features may be added or removed,
and documentation is expected to grow gradually.

---

## License

MIT.

See the [LICENSE](LICENSE) file for details.

---

## Source Code

https://github.com/akikareha/sunani
