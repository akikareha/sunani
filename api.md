# Sunani API

* System API
    - exports: init resize frame
    - imports: halt title cursor
* Console API
    - exports: init get
    - imports: params put
* Canvas API
    - exports: init
    - imports: begin clear color line rect fill_rect path vertex polygon fill_polygon
* Framebuffer API
    - exports: fb_init
    - imports: params paint
* Keyboard API
    - exports: init event
* Mouse API
    - exports: init motion button wheel

## System API exports

init resize frame

### init

If this functions is exported, System API is enabled.
This function is called from host when System API is ready.

```
//export sunani_system_init
func systemInit()
```

### resize

This function is called when host window or tab is resized.
This function is called at least once when app is launched.

```
//export sunani_system_resize
func systemResize(w, h uint32)
```

* w, h: Canvas width and height in pixels.

### frame

This function is called every event loop frame.

```
//export sunani_system_frame
func systemFrame()
```

## System API imports

halt title cursor

```
import "github.com/akikareha/sunani/api/system"
```

### halt

Halts execution of app and close window if possible.

```
system.Halt()
```

### title

Sets title string of window or tab.

```
system.Title(ptr uint32, length uint32)
```

* ptr: Memory address of title string.
* length: Length of title string in bytes.

### cursor

Sets visibility of mouse cursor.

```
system.Cursor(enabled uint32)
```

* enabled:
    - 0: Mouse cursor is hidden.
    - Other than 0: Mouse cursor is shown.

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
import "github.com/akikareha/sunani/api/console"
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

begin clear color line rect fill_rect path vertex polygon fill_polygon

```
import "github.com/akikareha/sunani/api/canvas"
```

### begin

Prepares for starting drawing on canvas.

```
canvas.Begin()
```

### clear

Clears canvas by filling with specified color.

```
canvas.Clear(r, g, b, a uint32)
```

* r, g, b, a: Color.

### color

Sets drawing color.

```
canvas.Color(r, g, b, a uint32)
```

* r, g, b, a: Color.

### line

Draws line.

```
canvas.Line(x1, y1 uint32, x2, y2 uint32)
```

* x1, y1: Start position.
* x2, y2: End position.

### rect

Draws rect.

```
canvas.Rect(x, y uint32, w, h uint32)
```

* x, y: Upper left corner position.
* w, h: Width and height in pixels.

### fill_rect

Draws filled rect.

```
canvas.FillRect(x, y uint32, w, h uint32)
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
canvas.Vertex(x, y uint32)
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

## Framebuffer API imports

params paint

```
import "github.com/akikareha/sunani/api/fb"
```

### params

Sets framebuffer.

```
fb.Params(ptr uint32, width, height uint32)
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
import "github.com/akikareha/sunani/lib"
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
import "github.com/akikareha/sunani/lib"
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
func mouseMotion(x, y uint32)
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
func mouseWheel(xoff uint32, yoff uint32)
```

* xoff, yoff: Scroll amounts.
