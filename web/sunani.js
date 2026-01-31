// ==== Sunani codes ====

const Action = { Unknown: 0, Press: 1, Release: 2 };

const Key = {
  Unknown: 0,
  A: 1, B: 2, C: 3, D: 4, E: 5, F: 6, G: 7, H: 8,
  I: 9, J: 10, K: 11, L: 12, M: 13, N: 14, O: 15, P: 16,
  Q: 17, R: 18, S: 19, T: 20, U: 21, V: 22, W: 23, X: 24,
  Y: 25, Z: 26,

  _0: 27, _1: 28, _2: 29, _3: 30, _4: 31,
  _5: 32, _6: 33, _7: 34, _8: 35, _9: 36,

  Escape: 37,
  Enter: 38,
  Space: 39,
  Tab: 40,
  Backspace: 41,

  Up: 42,
  Down: 43,
  Left: 44,
  Right: 45,
};

const Mouse = { Unknown: 0, Left: 1, Right: 2, Middle: 3 };

// KeyboardEvent.code -> Sunani Key
const keyFromCode = (code) => {
  // Letters: "KeyA".."KeyZ"
  if (code.startsWith("Key") && code.length === 4) {
    const ch = code[3];
    const idx = ch.charCodeAt(0) - 65; // 'A'
    if (0 <= idx && idx < 26) return Key[ch];
  }

  // Digits: "Digit0".."Digit9"
  if (code.startsWith("Digit") && code.length === 6) {
    const d = code[5];
    if ("0" <= d && d <= "9") return Key["_" + d];
  }

  switch (code) {
    case "Escape": return Key.Escape;
    case "Enter": return Key.Enter;
    case "Space": return Key.Space;
    case "Tab": return Key.Tab;
    case "Backspace": return Key.Backspace;
    case "ArrowUp": return Key.Up;
    case "ArrowDown": return Key.Down;
    case "ArrowLeft": return Key.Left;
    case "ArrowRight": return Key.Right;
    default: return Key.Unknown;
  }
};

const mouseFromButton = (button) => {
  // JS: 0=left,1=middle,2=right
  switch (button) {
    case 0: return Mouse.Left;
    case 1: return Mouse.Middle;
    case 2: return Mouse.Right;
    default: return Mouse.Unknown;
  }
};

// ==== WebAssembly ====

let wasm = null;
let running = true;

// memory view (refresh after growth)
let memU8 = null;
const refreshMem = () => {
  const mem = wasm.exports.memory;
  memU8 = new Uint8Array(mem.buffer);
};

// ==== Console ====

const con = document.getElementById("con")
const input = document.getElementById('input');

// console buffer
let consolePtr = 0;
let consoleLength = 0;

// ==== Canvas / Framebuffer ====

const canvas = document.getElementById("c");
const ctx2d = canvas.getContext("2d", { alpha: true });

const fbOffscreen = document.createElement("canvas");
const fbOffCtx = fbOffscreen.getContext("2d", { alpha: true });

let points = [];

// HiDPI resize
const resize = () => {
  const dpr = window.devicePixelRatio || 1;
  const w = Math.max(1, Math.floor(canvas.clientWidth * dpr));
  const h = Math.max(1, Math.floor(canvas.clientHeight * dpr));
  if (canvas.width !== w || canvas.height !== h) {
    canvas.width = w;
    canvas.height = h;
  }
};
window.addEventListener("resize", resize);
resize();

// immediate-mode draw state
let currentColor = "rgba(255,255,255,1)";

// framebuffer state
let fbPtr = 0;
let fbW = 0;
let fbH = 0;
let fbImageData = null;
let rx = 0;
let ry = 0;
let rw = 0;
let rh = 0;

function updateRect() {
  let cw = canvas.width;
  let ch = canvas.height;
  if (fbW == 0 || fbH == 0) {
    rw = 0;
    rh = 0;
  } else {
    const scale = Math.floor(Math.min(cw / fbW, ch / fbH));
    rw = fbW * scale;
    rh = fbH * scale;
  }
  rx = Math.floor((cw - rw) * 0.5);
  ry = Math.floor((ch - rh) * 0.5);
}

// ==== API ====

const importObject = {
  sunani: {
    // ---- std ----
    "std.now": () => {
      return BigInt(Date.now());
    },

    // ---- console ----
    "console.params": (ptr, length) => {
      consolePtr = ptr >>> 0;
      consoleLength = length >>> 0;
    },
    "console.put": (ptr, length) => {
      if (length < 1) {
        return;
      }
      const mem = wasm.exports.memory;
      const bytes = new Uint8Array(mem.buffer, ptr, length);
      let s = new TextDecoder("utf-8").decode(bytes);
      con.textContent += s;
      con.scrollTop = con.scrollHeight;
    },
    "console.wait": () => {
    },
    "console.leave": () => {
    },

    // ---- screen ----
    "screen.halt": () => {
      running = false;
    },
    "screen.title": (ptr, length) => {
      if (length < 1) {
        document.title = "";
        return;
      }
      const mem = wasm.exports.memory;
      const bytes = new Uint8Array(mem.buffer, ptr, length);
      let title = new TextDecoder("utf-8").decode(bytes);
      document.title = title;
    },
    "screen.cursor": (visible) => {
      if (visible === 0) {
        canvas.style.cursor = "none";
      } else {
        canvas.style.cursor = "";
      }
    },
    "screen.clear": (r, g, b, a) => {
      ctx2d.save();
      ctx2d.setTransform(1, 0, 0, 1, 0, 0);
      ctx2d.fillStyle = `rgba(${r},${g},${b},${a/255})`;
      ctx2d.fillRect(0, 0, canvas.width, canvas.height);
      ctx2d.restore();
    },

    // ---- canvas ----
    "canvas.color": (r, g, b, a) => {
      currentColor = `rgba(${r},${g},${b},${a/255})`;
      ctx2d.strokeStyle = currentColor;
      ctx2d.fillStyle = currentColor;
    },
    "canvas.line": (x1, y1, x2, y2) => {
      ctx2d.beginPath();
      ctx2d.moveTo(x1 + 0.5, y1 + 0.5);
      ctx2d.lineTo(x2 + 0.5, y2 + 0.5);
      ctx2d.stroke();
    },
    "canvas.rect": (x, y, w, h) => {
      ctx2d.strokeRect(x + 0.5, y + 0.5, w, h);
    },
    "canvas.fill_rect": (x, y, w, h) => {
      ctx2d.fillRect(x, y, w, h);
    },
    "canvas.path": () => {
      points = [];
    },
    "canvas.vertex": (x, y) => {
      points.push(x);
      points.push(y);
    },
    "canvas.polygon": () => {
      if (points.length < 2) {
        return;
      }
      ctx2d.beginPath();
      ctx2d.moveTo(points[0] + 0.5, points[1] + 0.5);
      for (let i = 2; i < points.length; i += 2) {
        ctx2d.lineTo(points[i] + 0.5, points[i + 1] + 0.5);
      }
      ctx2d.closePath();
      ctx2d.stroke();
    },
    "canvas.fill_polygon": () => {
      if (points.length < 2) {
        return;
      }
      ctx2d.beginPath();
      ctx2d.moveTo(points[0], points[1]);
      for (let i = 2; i < points.length; i += 2) {
        ctx2d.lineTo(points[i], points[i + 1]);
      }
      ctx2d.closePath();
      ctx2d.fill();
    },

    // ---- fb ----
    "fb.params": (ptr, width, height) => {
      fbPtr = ptr >>> 0;
      fbW = width >>> 0;
      fbH = height >>> 0;

      // Canvas ImageData buffer (clamped RGBA)
      fbImageData = new ImageData(fbW, fbH);

      updateRect();
    },
    "fb.paint": () => {
      if (!fbPtr || !fbW || !fbH || !fbImageData) return;

      refreshMem();

      let len = fbW * fbH * 4;
      if (len < 0) {
        len = -len;
      }
      if (len < 1) {
        return;
      }
      const src = memU8.subarray(fbPtr, fbPtr + len);
      fbImageData.data.set(src);

      // ImageData -> offscreen canvas -> magnified draw
      fbOffscreen.width = fbW;
      fbOffscreen.height = fbH;
      fbOffCtx.putImageData(fbImageData, 0, 0);

      ctx2d.imageSmoothingEnabled = false; // crisp pixel
      ctx2d.drawImage(
        fbOffscreen,
        0, 0, fbW, fbH,
        rx, ry, rw, rh
      );

      // for compatibility
      currentColor = "rgba(255,255,255,1)";
      ctx2d.strokeStyle = currentColor;
      ctx2d.fillStyle = currentColor;
    },
  }
};

// ==== load wasm ====

function getCurrentScript() {
  const url = import.meta.url;
  return [...document.scripts].find(s => s.src === url);
}

const script = getCurrentScript();
const wasmUrl = script?.dataset.wasm ?? "./default.wasm";

const bytes = await (await fetch(wasmUrl)).arrayBuffer();
const { instance } = await WebAssembly.instantiate(bytes, importObject);
wasm = instance;
refreshMem();

const call = (name, ...args) => {
  const fn = wasm.exports[name];
  if (typeof fn === "function") return fn(...args);
};

// init exports (call if exists)
call("sunani_init");
call("sunani_std_init");
call("sunani_console_init");
call("sunani_screen_init");
call("sunani_canvas_init");
call("sunani_fb_init");
call("sunani_key_init");
call("sunani_mouse_init");

function notifyResize() {
  const rect = canvas.getBoundingClientRect();
  const dpr = window.devicePixelRatio || 1;

  const width  = Math.floor(rect.width  * dpr);
  const height = Math.floor(rect.height * dpr);

  canvas.width  = width;
  canvas.height = height;

  call("sunani_screen_resize", width, height);
}

function notifyRect() {
  call("sunani_fb_rect", rx, ry, rw, rh);
}

const ro = new ResizeObserver(() => {
  notifyResize();
  updateRect();
  notifyRect();
});
ro.observe(canvas);

input.addEventListener("keydown", (e) => {
  if (e.key === "Enter") {
    e.preventDefault();
    const line = input.value;
    input.value = "";

    const encoder = new TextEncoder();
    const bytes = encoder.encode(line);

    refreshMem();
    memU8.set(bytes, consolePtr);

    con.textContent += line + "\n";
    con.scrollTop = con.scrollHeight;

    call("sunani_console_get", consolePtr, bytes.length);
  }
});

// ==== input wiring ====

window.addEventListener("keydown", (e) => {
  if (e.repeat) return;
  const k = keyFromCode(e.code);
  call("sunani_key_event", k, Action.Press);
});
window.addEventListener("keyup", (e) => {
  const k = keyFromCode(e.code);
  call("sunani_key_event", k, Action.Release);
});

// add pointer lock or capture if needed
canvas.addEventListener("mousemove", (e) => {
  const rect = canvas.getBoundingClientRect();
  const dpr = window.devicePixelRatio || 1;
  const x = (e.clientX - rect.left) * dpr;
  const y = (e.clientY - rect.top) * dpr;
  call("sunani_mouse_motion", x, y);
});

canvas.addEventListener("mousedown", (e) => {
  canvas.focus?.();
  const b = mouseFromButton(e.button);
  call("sunani_mouse_button", b, Action.Press);
});
canvas.addEventListener("mouseup", (e) => {
  const b = mouseFromButton(e.button);
  call("sunani_mouse_button", b, Action.Release);
});
canvas.addEventListener("contextmenu", (e) => {
  e.preventDefault();
});

canvas.addEventListener("wheel", (e) => {
  e.preventDefault();
  // GLFW-like
  let xoff = Math.trunc(e.deltaX / 100);
  let yoff = -Math.trunc(e.deltaY / 100);
  call("sunani_mouse_wheel", xoff, yoff);
}, { passive: false });

// init
notifyResize();
notifyRect();

call("sunani_start");

// ==== main loop ====

function frame() {
  if (!running) return;
  resize();
  call("sunani_screen_frame");
  requestAnimationFrame(frame);
}
requestAnimationFrame(frame);

