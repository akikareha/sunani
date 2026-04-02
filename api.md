# Sunani API

* Initialization API
    - exports: init start
* Standard API
    - imports: now
* Console API
    - exports: init get
    - imports: params put
* Screen API
    - exports: init resize frame
    - imports: halt title cursor clear
* Canvas API
    - exports: init
    - imports: color line rect fill_rect path vertex polygon fill_polygon
* Framebuffer API
    - exports: init rect
    - imports: params paint
* Keyboard API
    - exports: init event
* Mouse API
    - exports: init motion button wheel

## Initialization API exports

init start

### init

Called from host before host is initialized.

```
//export sunani_init
func init()
```

### start

Called from host after host is initialized.

```
//export sunani_start
func start()
```

## Standard API exports

init

### init

If this functions is exported, Standard API is enabled.
This function is called from host when Standard API is ready.

```
//export sunani_std_init
func stdInit()
```

## Standard API imports

now

```
import "tea.kareha.org/loom/sunani/api/std"
```

### now

Returns current UNIX time in ms.

```
std.Now() int64
```

## Console API exports

init get

### init

If this functions is exported, Console API is enabled.
This function is called from host when Console API is ready.

```
//export sunani_console_init
func consoleInit()
```

### get

This function is called when line is entered on STDIN or text field.

```
//export sunani_console_get
func consoleGet(ptr uint32, length uint32)
```

* ptr: Memory address of input string.
* length: Length of input string in bytes.

## Console API imports

params put

```
import "tea.kareha.org/loom/sunani/api/console"
```

### params

Sets buffer for storing input string.

```
console.Params(ptr uint32, length uint32)
```

* ptr: Memory address of input string buffer.
* length: Length of input string buffer in bytes.

### put

Puts string to STDOUT or text output area.

```
console.Put(ptr uint32, length uint32)
```

* ptr: Memory address of string to put.
* length: Length of string to put in bytes.

### wait

Waits until leave is called.

```
console.Wait()
```

### leave

Leaves from wait.

```
console.Leave()
```

## Screen API exports

init resize frame

### init

If this functions is exported, Screen API is enabled.
This function is called from host when Screen API is ready.

```
//export sunani_screen_init
func screenInit()
```

### resize

This function is called when host window or tab is resized.
This function is called at least once when app is launched.

```
//export sunani_screen_resize
func screenResize(w, h int32)
```

* w, h: Canvas width and height in pixels.

### frame

This function is called every event loop frame.

```
//export sunani_screen_frame
func screenFrame()
```

## Screen API imports

halt title cursor clear

```
import "tea.kareha.org/loom/sunani/api/screen"
```

### halt

Halts execution of app and close window if possible.

```
screen.Halt()
```

### title

Sets title string of window or tab.

```
screen.Title(ptr uint32, length uint32)
```

* ptr: Memory address of title string.
* length: Length of title string in bytes.

### cursor

Sets visibility of mouse cursor.

```
screen.Cursor(visible uint32)
```

* visible:
    - 0: Mouse cursor is invisible.
    - Other than 0: Mouse cursor is visible.

### clear

Clears screen by filling with specified color.

```
screen.Clear(r, g, b, a uint32)
```

* r, g, b, a: Color.

## Canvas API exports

init

### init

```
//export sunani_canvas_init
func canvasInit()
```

If this functions is exported, Canvas API is enabled.
This function is called from host when Canvas API is ready.

## Canvas API imports

color line rect fill_rect path vertex polygon fill_polygon

```
import "tea.kareha.org/loom/sunani/api/canvas"
```

### color

Sets drawing color.

```
canvas.Color(r, g, b, a uint32)
```

* r, g, b, a: Color.

### line

Draws line.

```
canvas.Line(x1, y1 int32, x2, y2 int32)
```

* x1, y1: Start position.
* x2, y2: End position.

### rect

Draws rect.

```
canvas.Rect(x, y int32, w, h int32)
```

* x, y: Upper left corner position.
* w, h: Width and height in pixels.

### fill_rect

Draws filled rect.

```
canvas.FillRect(x, y int32, w, h int32)
```

* x, y: Upper left corner position.
* w, h: Width and height in pixels.

### path

Starts drawing polygon.

```
canvas.Path()
```

### vertex

Adds vertex to drawing polygon.

```
canvas.Vertex(x, y int32)
```

* x, y: Vertex.

### polygon

Finishes adding vertices and draws polygon.

```
canvas.Polygon()
```

### fill_polygon

Finishes adding vertices and draws filled polygon.

```
canvas.FillPolygon()
```

## Framebuffer API exports

init

### init

If this functions is exported, Framebuffer API is enabled.
This function is called from host when Framebuffer API is ready.

```
//export sunani_fb_init
func fbInit()
```

### rect

This function is called when host window or tab is resized.
This function is called at least once when app is launched.

```
//export sunani_fb_rect
func fbRect(x, y int32, w, h int32)

## Framebuffer API imports

params paint

```
import "tea.kareha.org/loom/sunani/api/fb"
```

### params

Sets framebuffer.

```
fb.Params(ptr uint32, width, height int32)
```

* ptr: Memory address of framebuffer.
* width, height: Width and height of framebuffer in pixels.

### paint

Paints framebuffer on screen.

```
fb.Paint()
```

## Keyboard API exports

init event

```
import "tea.kareha.org/loom/sunani/lib"
```

### init

If this functions is exported, Keyboard API is enabled.
This function is called from host when Keyboard API is ready.

```
//export sunani_key_init
func keyInit()
```

### event

This function is called when key pressed or released.

```
//export sunani_key_event
func keyEvent(key uint32, action uint32)
```

* key: Key code.
* action: Indicates pressed or released.

## Mouse API exports

init motion button wheel

```
import "tea.kareha.org/loom/sunani/lib"
```

### init

If this functions is exported, Mouse API is enabled.
This function is called from host when Mouse API is ready.

```
//export sunani_mouse_init
func mouseInit()
```

### motion

This function is called when mouse cursor moved.

```
//export sunani_mouse_motion
func mouseMotion(x, y int32)
```

* x, y: Mouse cursor position on canvas in pixels

### button

This function is called when mouse button is pressed or released.

```
//export sunani_mouse_button
func mouseButton(button uint32, action uint32)
```

* button: Mouse button.
* action: Indicates pressed or released.

### wheel

This function is called when mouse wheel is moved.

```
//export sunani_mouse_wheel
func mouseWheel(xoff, yoff int32)
```

* xoff, yoff: Scroll amounts.
