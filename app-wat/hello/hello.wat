(module
  ;; --- Sunani Console API ---
  (import "sunani" "console.put"
    (func $console_put (param i32 i32))
  )

  ;; --- Define memory ---
  (memory (export "memory") 1)

  ;; --- Hello World string ---
  (data (i32.const 0) "Hello, World!\n")

  ;; --- sunani_console_init ---
  (func (export "sunani_console_init")
    i32.const 0        ;; ptr
    i32.const 14       ;; length
    call $console_put
  )
)
