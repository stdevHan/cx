package base

// constant codes
const (
	// https://www.khronos.org/registry/OpenGL/api/GL/glcorearb.h
	/* gl_1_0 */
	CONST_GL_DEPTH_BUFFER_BIT = iota
	CONST_GL_STENCIL_BUFFER_BIT
	CONST_GL_COLOR_BUFFER_BIT
	CONST_GL_FALSE
	CONST_GL_TRUE
	CONST_GL_POINTS
	CONST_GL_LINES
	CONST_GL_LINE_LOOP
	CONST_GL_LINE_STRIP
	CONST_GL_TRIANGLES
	CONST_GL_TRIANGLE_STRIP
	CONST_GL_TRIANGLE_FAN
	CONST_GL_QUADS
	CONST_GL_LESS
	CONST_GL_LEQUAL
	CONST_GL_SRC_ALPHA
	CONST_GL_ONE_MINUS_SRC_ALPHA
	CONST_GL_FRONT
	CONST_GL_BACK
	CONST_GL_FRONT_AND_BACK
	CONST_GL_NO_ERROR
	CONST_GL_INVALID_ENUM
	CONST_GL_INVALID_VALUE
	CONST_GL_INVALID_OPERATION
	CONST_GL_STACK_OVERFLOW
	CONST_GL_STACK_UNDERFLOW
	CONST_GL_OUT_OF_MEMORY
	CONST_GL_LINE_SMOOTH
	CONST_GL_POLYGON_SMOOTH
	CONST_GL_CULL_FACE
	CONST_GL_DEPTH_TEST
	CONST_GL_DITHER
	CONST_GL_BLEND
	CONST_GL_POLYGON_SMOOTH_HINT
	CONST_GL_TEXTURE_2D
	CONST_GL_DONT_CARE
	CONST_GL_UNSIGNED_BYTE
	CONST_GL_FLOAT
	CONST_GL_TEXTURE
	CONST_GL_COLOR
	CONST_GL_RGBA
	CONST_GL_REPLACE
	CONST_GL_NEAREST
	CONST_GL_LINEAR
	CONST_GL_NEAREST_MIPMAP_NEAREST
	CONST_GL_LINEAR_MIPMAP_NEAREST
	CONST_GL_NEAREST_MIPMAP_LINEAR
	CONST_GL_LINEAR_MIPMAP_LINEAR
	CONST_GL_TEXTURE_MAG_FILTER
	CONST_GL_TEXTURE_MIN_FILTER
	CONST_GL_TEXTURE_WRAP_S
	CONST_GL_TEXTURE_WRAP_T
	CONST_GL_REPEAT

	/* gl_1_1 */
	CONST_GL_RGBA8
	CONST_GL_VERTEX_ARRAY

	/* gl_1_2 */
	CONST_GL_TEXTURE_WRAP_R
	CONST_GL_CLAMP_TO_EDGE

	/* gl_1_3 */
	CONST_GL_TEXTURE0
	CONST_GL_MULTISAMPLE_ARB // remove _ARB
	CONST_GL_CLAMP_TO_BORDER

	/* gl_1_4 */
	CONST_GL_MIRRORED_REPEAT

	/* gl_1_5 */
	CONST_GL_ARRAY_BUFFER
	CONST_GL_STREAM_DRAW
	CONST_GL_STREAM_READ
	CONST_GL_STREAM_COPY
	CONST_GL_STATIC_DRAW
	CONST_GL_STATIC_READ
	CONST_GL_STATIC_COPY
	CONST_GL_DYNAMIC_DRAW
	CONST_GL_DYNAMIC_READ
	CONST_GL_DYNAMIC_COPY

	/* gl_2_0 */
	CONST_GL_FRAGMENT_SHADER
	CONST_GL_VERTEX_SHADER

	/* gl_3_0 */
	CONST_GL_FRAMEBUFFER_UNDEFINED
	CONST_GL_FRAMEBUFFER_COMPLETE
	CONST_GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT
	CONST_GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT
	CONST_GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER
	CONST_GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER
	CONST_GL_FRAMEBUFFER_UNSUPPORTED
	CONST_GL_COLOR_ATTACHMENT0
	CONST_GL_FRAMEBUFFER
	CONST_GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE

	/* gl_3_2 */
	CONST_GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS

	/* gl_4_4 */
	CONST_GL_MIRROR_CLAMP_TO_EDGE

	/* Fixed pipeline. Deprecated ? */
	CONST_GL_POLYGON
	CONST_GL_MODELVIEW
	CONST_GL_PROJECTION
	CONST_GL_MODELVIEW_MATRIX
	CONST_GL_LIGHTING
	CONST_GL_LIGHT0
	CONST_GL_AMBIENT
	CONST_GL_DIFFUSE
	CONST_GL_POSITION
	CONST_GL_TEXTURE_ENV
	CONST_GL_TEXTURE_ENV_MODE
	CONST_GL_MODULATE
	CONST_GL_DECAL
	CONST_GL_POINT_SMOOTH

	// glfw
	CONST_GLFW_FALSE
	CONST_GLFW_TRUE
	CONST_GLFW_PRESS
	CONST_GLFW_RELEASE
	CONST_GLFW_REPEAT
	CONST_GLFW_KEY_UNKNOWN
	CONST_GLFW_CURSOR
	CONST_GLFW_STICKY_KEYS
	CONST_GLFW_STICKY_MOUSE_BUTTONS
	CONST_GLFW_CURSOR_NORMAL
	CONST_GLFW_CURSOR_HIDDEN
	CONST_GLFW_CURSOR_DISABLED
	CONST_GLFW_RESIZABLE
	CONST_GLFW_CONTEXT_VERSION_MAJOR
	CONST_GLFW_CONTEXT_VERSION_MINOR
	CONST_GLFW_OPENGL_PROFILE
	CONST_GLFW_OPENGL_COREPROFILE
	CONST_GLFW_OPENGL_FORWARD_COMPATIBLE
	CONST_GLFW_MOUSE_BUTTON_LAST
	CONST_GLFW_MOUSE_BUTTON_LEFT
	CONST_GLFW_MOUSE_BUTTON_RIGHT
	CONST_GLFW_MOUSE_BUTTON_MIDDLE

	// gltext
	CONST_GLTEXT_LEFT_TO_RIGHT
	CONST_GLTEXT_RIGHT_TO_LEFT
	CONST_GLTEXT_TOP_TO_BOTTOM
)

// For the parser. These shouldn't be used in the runtime for performance reasons
var ConstNames map[int]string = map[int]string{
	/* gl_1_0 */
	CONST_GL_DEPTH_BUFFER_BIT:       "gl.DEPTH_BUFFER_BIT",
	CONST_GL_STENCIL_BUFFER_BIT:     "gl.STENCIL_BUFFER_BIT",
	CONST_GL_COLOR_BUFFER_BIT:       "gl.COLOR_BUFFER_BIT",
	CONST_GL_FALSE:                  "gl.FALSE",
	CONST_GL_TRUE:                   "gl.TRUE",
	CONST_GL_POINTS:                 "gl.POINTS",
	CONST_GL_LINES:                  "gl.LINES",
	CONST_GL_LINE_LOOP:              "gl.LINE_LOOP",
	CONST_GL_LINE_STRIP:             "gl.LINE_STRIP",
	CONST_GL_TRIANGLES:              "gl.TRIANGLES",
	CONST_GL_TRIANGLE_STRIP:         "gl.TRIANGLE_STRIP",
	CONST_GL_TRIANGLE_FAN:           "gl.TRIANGLE_FAN",
	CONST_GL_QUADS:                  "gl.QUADS",
	CONST_GL_LESS:                   "gl.LESS",
	CONST_GL_LEQUAL:                 "gl.LEQUAL",
	CONST_GL_SRC_ALPHA:              "gl.SRC_ALPHA",
	CONST_GL_ONE_MINUS_SRC_ALPHA:    "gl.ONE_MINUS_SRC_ALPHA",
	CONST_GL_FRONT:                  "gl.FRONT",
	CONST_GL_BACK:                   "gl.BACK",
	CONST_GL_FRONT_AND_BACK:         "gl.FRONT_AND_BACK",
	CONST_GL_NO_ERROR:               "gl.NO_ERROR",
	CONST_GL_INVALID_ENUM:           "gl.INVALID_ENUM",
	CONST_GL_INVALID_VALUE:          "gl.INVALID_VALUE",
	CONST_GL_INVALID_OPERATION:      "gl.INVALID_OPERATION",
	CONST_GL_STACK_OVERFLOW:         "gl.STACK_OVERFLOW",
	CONST_GL_STACK_UNDERFLOW:        "gl.STACK_UNDERFLOW",
	CONST_GL_OUT_OF_MEMORY:          "gl.OUT_OF_MEMORY",
	CONST_GL_LINE_SMOOTH:            "gl.LINE_SMOOTH",
	CONST_GL_POLYGON_SMOOTH:         "gl.POLYGON_SMOOTH",
	CONST_GL_CULL_FACE:              "gl.CULL_FACE",
	CONST_GL_DEPTH_TEST:             "gl.DEPTH_TEST",
	CONST_GL_DITHER:                 "gl.DITHER",
	CONST_GL_BLEND:                  "gl.BLEND",
	CONST_GL_POLYGON_SMOOTH_HINT:    "gl.POLYGON_SMOOTH_HINT",
	CONST_GL_TEXTURE_2D:             "gl.TEXTURE_2D",
	CONST_GL_DONT_CARE:              "gl.DONT_CARE",
	CONST_GL_UNSIGNED_BYTE:          "gl.UNSIGNED_BYTE",
	CONST_GL_FLOAT:                  "gl.FLOAT",
	CONST_GL_TEXTURE:                "gl.TEXTURE",
	CONST_GL_COLOR:                  "gl.COLOR",
	CONST_GL_RGBA:                   "gl.RGBA",
	CONST_GL_REPLACE:                "gl.REPLACE",
	CONST_GL_NEAREST:                "gl.NEAREST",
	CONST_GL_LINEAR:                 "gl.LINEAR",
	CONST_GL_NEAREST_MIPMAP_NEAREST: "gl.NEAREST_MIPMAP_NEAREST",
	CONST_GL_LINEAR_MIPMAP_NEAREST:  "gl.LINEAR_MIPMAP_NEAREST",
	CONST_GL_NEAREST_MIPMAP_LINEAR:  "gl.NEAREST_MIPMAP_LINEAR",
	CONST_GL_LINEAR_MIPMAP_LINEAR:   "gl.LINEAR_MIPMAP_LINEAR",
	CONST_GL_TEXTURE_MAG_FILTER:     "gl.TEXTURE_MAP_FILTER",
	CONST_GL_TEXTURE_MIN_FILTER:     "gl.TEXTURE_MIN_FILTER",
	CONST_GL_TEXTURE_WRAP_S:         "gl.TEXTURE_WRAP_S",
	CONST_GL_TEXTURE_WRAP_T:         "gl.TEXTURE_WRAP_T",
	CONST_GL_REPEAT:                 "gl.REPEAT",

	/* gl_1_1 */
	CONST_GL_RGBA8:                  "gl.RGBA8",
	CONST_GL_VERTEX_ARRAY:           "gl.VERTEX_ARRAY",

	/* gl_1_2 */
	CONST_GL_TEXTURE_WRAP_R:         "gl.TEXTURE_WRAP_R",
	CONST_GL_CLAMP_TO_EDGE:          "gl.CLAMP_TO_EDGE",

	/* gl_1_3 */
	CONST_GL_TEXTURE0:               "gl.TEXTURE0",
	CONST_GL_MULTISAMPLE_ARB:        "gl.MULTISAMPLE_ARB", // remove _ARB
	CONST_GL_CLAMP_TO_BORDER:        "gl.CLAMP_TO_BORDER",

	/* gl_1_4 */
	CONST_GL_MIRRORED_REPEAT:        "gl.MIRRORED_REPEAT",

	/* gl_1_5 */
	CONST_GL_ARRAY_BUFFER:        "gl.ARRAY_BUFFER",
	CONST_GL_STREAM_DRAW:         "gl.STREAM_DRAW",
	CONST_GL_STREAM_READ:         "gl.STREAM_READ",
	CONST_GL_STREAM_COPY:         "gl.STREAM_COPY",
	CONST_GL_STATIC_DRAW:         "gl.STATIC_DRAW",
	CONST_GL_STATIC_READ:         "gl.STATIC_READ",
	CONST_GL_STATIC_COPY:         "gl.STATIC_COPY",
	CONST_GL_DYNAMIC_DRAW:        "gl.DYNAMIC_DRAW",
	CONST_GL_DYNAMIC_READ:        "gl.DYNAMIC_READ",
	CONST_GL_DYNAMIC_COPY:        "gl.DYNAMIC_COPY",

	/* gl_2_0 */
	CONST_GL_FRAGMENT_SHADER:     "gl.FRAGMENT_SHADER",
	CONST_GL_VERTEX_SHADER:       "gl.VERTEX_SHADER",

	/* gl_3_0 */
	CONST_GL_FRAMEBUFFER_UNDEFINED:                     "gl.FRAMEBUFFER_UNDEFINED",
	CONST_GL_FRAMEBUFFER_COMPLETE:                      "gl.FRAMEBUFFER_COMPLETE",
	CONST_GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT:         "gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT",
	CONST_GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT: "gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT",
	CONST_GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER:        "gl.FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER",
	CONST_GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER:        "gl.FRAMEBUFFER_INCOMPLETE_READ_BUFFER",
	CONST_GL_FRAMEBUFFER_UNSUPPORTED:                   "gl.FRAMEBUFFER_UNSUPPORTED",
	CONST_GL_COLOR_ATTACHMENT0:                         "gl.COLOR_ATTACHMENT0",
	CONST_GL_FRAMEBUFFER:                               "gl.FRAMEBUFFER",
	CONST_GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE:        "gl.FRAMEBUFFER_INCOMPLETE_MULTISAMPLE",

	/* gl_3_2 */
	CONST_GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS:      "gl.FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS",

	/* gl_4_4 */
	CONST_GL_MIRROR_CLAMP_TO_EDGE: "gl.MIRROR_CLAMP_TO_EDGE",

	/* Fixed pipeline. Deprecated ? */
	CONST_GL_POLYGON:             "gl.POLYGON",
	CONST_GL_MODELVIEW:           "gl.MODELVIEW",
	CONST_GL_PROJECTION:          "gl.PROJECTION",
	CONST_GL_MODELVIEW_MATRIX:    "gl.MODELVIEW_MATRIX",
	CONST_GL_LIGHTING:            "gl.LIGHTING",
	CONST_GL_LIGHT0:              "gl.LIGHT0",
	CONST_GL_AMBIENT:             "gl.AMBIENT",
	CONST_GL_DIFFUSE:             "gl.DIFFUSE",
	CONST_GL_POSITION:            "gl.POSITION",
	CONST_GL_TEXTURE_ENV:         "gl.TEXTURE_ENV",
	CONST_GL_TEXTURE_ENV_MODE:    "gl.TEXTURE_ENV_MODE",
	CONST_GL_MODULATE:            "gl.MODULATE",
	CONST_GL_DECAL:               "gl.DECAL",
	CONST_GL_POINT_SMOOTH:        "gl.POINT_SMOOTH",

	// glfw
	CONST_GLFW_FALSE:                     "glfw.False",
	CONST_GLFW_TRUE:                      "glfw.True",
	CONST_GLFW_PRESS:                     "glfw.Press",
	CONST_GLFW_RELEASE:                   "glfw.Release",
	CONST_GLFW_REPEAT:                    "glfw.Repeat",
	CONST_GLFW_KEY_UNKNOWN:               "glfw.KeyUnknown",
	CONST_GLFW_CURSOR:                    "glfw.Cursor",
	CONST_GLFW_STICKY_KEYS:               "glfw.StickyKeys",
	CONST_GLFW_STICKY_MOUSE_BUTTONS:      "glfw.StickyMouseButtons",
	CONST_GLFW_CURSOR_NORMAL:             "glfw.CursorNormal",
	CONST_GLFW_CURSOR_HIDDEN:             "glfw.CursorHidden",
	CONST_GLFW_CURSOR_DISABLED:           "glfw.CursorDisabled",
	CONST_GLFW_RESIZABLE:                 "glfw.Resizable",
	CONST_GLFW_CONTEXT_VERSION_MAJOR:     "glfw.ContextVersionMajor",
	CONST_GLFW_CONTEXT_VERSION_MINOR:     "glfw.ContextVersionMinor",
	CONST_GLFW_OPENGL_PROFILE:            "glfw.Opengl.Profile",
	CONST_GLFW_OPENGL_COREPROFILE:        "glfw.Opengl.Coreprofile",
	CONST_GLFW_OPENGL_FORWARD_COMPATIBLE: "glfw.Opengl.ForwardCompatible",
	CONST_GLFW_MOUSE_BUTTON_LAST:         "glfw.MouseButtonLast",
	CONST_GLFW_MOUSE_BUTTON_LEFT:         "glfw.MouseButtonLeft",
	CONST_GLFW_MOUSE_BUTTON_RIGHT:        "glfw.MouseButtonRight",
	CONST_GLFW_MOUSE_BUTTON_MIDDLE:       "glfw.MouseButtonMiddle",

	// gltext
	CONST_GLTEXT_LEFT_TO_RIGHT:           "gltext.LeftToRight",
	CONST_GLTEXT_RIGHT_TO_LEFT:           "gltext.RightToLeft",
	CONST_GLTEXT_TOP_TO_BOTTOM:           "gltext.TopToBottom",
}

// For the parser. These shouldn't be used in the runtime for performance reasons
var ConstCodes map[string]int = map[string]int{
	/* gl_1_0 */
	"gl.DEPTH_BUFFER_BIT":       CONST_GL_DEPTH_BUFFER_BIT,
	"gl.STENCIL_BUFFER_BIT":     CONST_GL_STENCIL_BUFFER_BIT,
	"gl.COLOR_BUFFER_BIT":       CONST_GL_COLOR_BUFFER_BIT,
	"gl.FALSE":                  CONST_GL_FALSE,
	"gl.TRUE":                   CONST_GL_TRUE,
	"gl.POINTS":                 CONST_GL_POINTS,
	"gl.LINES":                  CONST_GL_LINES,
	"gl.LINE_LOOP":              CONST_GL_LINE_LOOP,
	"gl.LINE_STRIP":             CONST_GL_LINE_STRIP,
	"gl.TRIANGLES":              CONST_GL_TRIANGLES,
	"gl.TRIANGLE_STRIP":         CONST_GL_TRIANGLE_STRIP,
	"gl.TRIANGLE_FAN":           CONST_GL_TRIANGLE_FAN,
	"gl.QUADS":                  CONST_GL_QUADS,
	"gl.LESS":                   CONST_GL_LESS,
	"gl.LEQUAL":                 CONST_GL_LEQUAL,
	"gl.SRC_ALPHA":              CONST_GL_SRC_ALPHA,
	"gl.ONE_MINUS_SRC_ALPHA":    CONST_GL_ONE_MINUS_SRC_ALPHA,
	"gl.FRONT":                  CONST_GL_FRONT,
	"gl.BACK":                   CONST_GL_BACK,
	"gl.FRONT_AND_BACK":         CONST_GL_FRONT_AND_BACK,
	"gl.NO_ERROR":               CONST_GL_NO_ERROR,
	"gl.INVALID_ENUM":           CONST_GL_INVALID_ENUM,
	"gl.INVALID_VALUE":          CONST_GL_INVALID_VALUE,
	"gl.INVALID_OPERATION":      CONST_GL_INVALID_OPERATION,
	"gl.STACK_OVERFLOW":         CONST_GL_STACK_OVERFLOW,
	"gl.STACK_UNDERFLOW":        CONST_GL_STACK_UNDERFLOW,
	"gl.OUT_OF_MEMORY":          CONST_GL_OUT_OF_MEMORY,
	"gl.LINE_SMOOTH":            CONST_GL_LINE_SMOOTH,
	"gl.POLYGON_SMOOTH":         CONST_GL_POLYGON_SMOOTH,
	"gl.CULL_FACE":              CONST_GL_CULL_FACE,
	"gl.DEPTH_TEST":             CONST_GL_DEPTH_TEST,
	"gl.DITHER":                 CONST_GL_DITHER,
	"gl.BLEND":                  CONST_GL_BLEND,
	"gl.POLYGON_SMOOTH_HINT":    CONST_GL_POLYGON_SMOOTH_HINT,
	"gl.TEXTURE_2D":             CONST_GL_TEXTURE_2D,
	"gl.DONT_CARE":              CONST_GL_DONT_CARE,
	"gl.UNSIGNED_BYTE":          CONST_GL_UNSIGNED_BYTE,
	"gl.FLOAT":                  CONST_GL_FLOAT,
	"gl.TEXTURE":                CONST_GL_TEXTURE,
	"gl.COLOR":                  CONST_GL_COLOR,
	"gl.RGBA":                   CONST_GL_RGBA,
	"gl.REPLACE":                CONST_GL_REPLACE,
	"gl.NEREAST":                CONST_GL_NEAREST,
	"gl.LINEAR":                 CONST_GL_LINEAR,
	"gl.NEAREST_MIPMAP_NEAREST": CONST_GL_NEAREST_MIPMAP_NEAREST,
	"gl.LINEAR_MIPMAP_NEAREST":  CONST_GL_LINEAR_MIPMAP_NEAREST,
	"gl.NEAREST_MIPMAP_LINEAR":  CONST_GL_NEAREST_MIPMAP_LINEAR,
	"gl.LINEAR_MIPMAP_LINEAR":   CONST_GL_LINEAR_MIPMAP_LINEAR,
	"gl.TEXTURE_MAG_FILTER":     CONST_GL_TEXTURE_MAG_FILTER,
	"gl.TEXTURE_MIN_FILTER":     CONST_GL_TEXTURE_MIN_FILTER,
	"gl.TEXTURE_WRAP_S":         CONST_GL_TEXTURE_WRAP_S,
	"gl.TEXTURE_WRAP_T":         CONST_GL_TEXTURE_WRAP_T,
	"gl.REPEAT":                 CONST_GL_REPEAT,

	/* gl_1_1 */
	"gl.RGBA8":                  CONST_GL_RGBA8,
	"gl.VERTEX_ARRAY":           CONST_GL_VERTEX_ARRAY,

	/* gl_1_2 */
	"gl.TEXTURE_WRAP_R":         CONST_GL_TEXTURE_WRAP_R,
	"gl.CLAMP_TO_EDGE":          CONST_GL_CLAMP_TO_EDGE,

	/* gl_1_3 */
	"gl.TEXTURE0":               CONST_GL_TEXTURE0,
	"gl.MULTISAMPLE_ARB":        CONST_GL_MULTISAMPLE_ARB, // remove _ARB
	"gl.CLAMP_TO_BORDER":        CONST_GL_CLAMP_TO_BORDER,

	/* gl_1_4 */
	"gl.MIRRORED_REPEAT":        CONST_GL_MIRRORED_REPEAT,

	/* gl_1_5 */
	"gl.ARRAY_BUFFER":        CONST_GL_ARRAY_BUFFER,
	"gl.STREAM_DRAW":         CONST_GL_STREAM_DRAW,
	"gl.STREAM_READ":         CONST_GL_STREAM_READ,
	"gl.STREAM_COPY":         CONST_GL_STREAM_COPY,
	"gl.STATIC_DRAW":         CONST_GL_STATIC_DRAW,
	"gl.STATIC_READ":         CONST_GL_STATIC_READ,
	"gl.STATIC_COPY":         CONST_GL_STATIC_COPY,
	"gl.DYNAMIC_DRAW":        CONST_GL_DYNAMIC_DRAW,
	"gl.DYNAMIC_READ":        CONST_GL_DYNAMIC_READ,
	"gl.DYNAMIC_COPY":        CONST_GL_DYNAMIC_COPY,

	/* gl_2_0 */
	"gl.FRAGMENT_SHADER":     CONST_GL_FRAGMENT_SHADER,
	"gl.VERTEX_SHADER":       CONST_GL_VERTEX_SHADER,

	/* gl_3_0 */
	"gl.FRAMEBUFFER_UNDEFINED":                     CONST_GL_FRAMEBUFFER_UNDEFINED,
	"gl.FRAMEBUFFER_COMPLETE":                      CONST_GL_FRAMEBUFFER_COMPLETE,
	"gl.FRAMEBUFFER_INCOMPLETE_ATTACHMENT":         CONST_GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT,
	"gl.FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT": CONST_GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT,
	"gl.FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER":        CONST_GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER,
	"gl.FRAMEBUFFER_INCOMPLETE_READ_BUFFER":        CONST_GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER,
	"gl.FRAMEBUFFER_UNSUPPORTED":                   CONST_GL_FRAMEBUFFER_UNSUPPORTED,
	"gl.COLOR_ATTACHMENT0":                         CONST_GL_COLOR_ATTACHMENT0,
	"gl.FRAMEBUFFER":                               CONST_GL_FRAMEBUFFER,
	"gl.FRAMEBUFFER_INCOMPLETE_MULTISAMPLE":        CONST_GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE,

	/* gl_3_2 */
	"gl.FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS":      CONST_GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS,

	/* gl_4_4 */
	"gl.MIRROR_CLAMP_TO_EDGE": CONST_GL_MIRROR_CLAMP_TO_EDGE,

	/* Deprecated ? */
	"gl.POLYGON":             CONST_GL_POLYGON,
	"gl.MODELVIEW":           CONST_GL_MODELVIEW,
	"gl.PROJECTION":          CONST_GL_PROJECTION,
	"gl.MODELVIEW_MATRIX":    CONST_GL_MODELVIEW_MATRIX,
	"gl.LIGHTING":            CONST_GL_LIGHTING,
	"gl.LIGHT0":              CONST_GL_LIGHT0,
	"gl.AMBIENT":             CONST_GL_AMBIENT,
	"gl.DIFFUSE":             CONST_GL_DIFFUSE,
	"gl.POSITION":            CONST_GL_POSITION,
	"gl.TEXTURE_ENV":         CONST_GL_TEXTURE_ENV,
	"gl.TEXTURE_ENV_MODE":    CONST_GL_TEXTURE_ENV_MODE,
	"gl.MODULATE":            CONST_GL_MODULATE,
	"gl.DECAL":               CONST_GL_DECAL,
	"gl.POINT_SMOOTH":        CONST_GL_POINT_SMOOTH,

	// glfw
	"glfw.False":                    CONST_GLFW_FALSE,
	"glfw.True":                     CONST_GLFW_TRUE,
	"glfw.Press":                    CONST_GLFW_PRESS,
	"glfw.Release":                  CONST_GLFW_RELEASE,
	"glfw.Repeat":                   CONST_GLFW_REPEAT,
	"glfw.KeyUnknown":               CONST_GLFW_KEY_UNKNOWN,
	"glfw.Cursor":                   CONST_GLFW_CURSOR,
	"glfw.StickyKeys":               CONST_GLFW_STICKY_KEYS,
	"glfw.StickyMouseButtons":       CONST_GLFW_STICKY_MOUSE_BUTTONS,
	"glfw.CursorNormal":             CONST_GLFW_CURSOR_NORMAL,
	"glfw.CursorHidden":             CONST_GLFW_CURSOR_HIDDEN,
	"glfw.CursorDisabled":           CONST_GLFW_CURSOR_DISABLED,
	"glfw.Resizable":                CONST_GLFW_RESIZABLE,
	"glfw.ContextVersionMajor":      CONST_GLFW_CONTEXT_VERSION_MAJOR,
	"glfw.ContextVersionMinor":      CONST_GLFW_CONTEXT_VERSION_MINOR,
	"glfw.Opengl.Profile":           CONST_GLFW_OPENGL_PROFILE,
	"glfw.Opengl.Coreprofile":       CONST_GLFW_OPENGL_COREPROFILE,
	"glfw.Opengl.ForwardCompatible": CONST_GLFW_OPENGL_FORWARD_COMPATIBLE,
	"glfw.MouseButtonLast":          CONST_GLFW_MOUSE_BUTTON_LAST,
	"glfw.MouseButtonLeft":          CONST_GLFW_MOUSE_BUTTON_LEFT,
	"glfw.MouseButtonRight":         CONST_GLFW_MOUSE_BUTTON_RIGHT,
	"glfw.MouseButtonMiddle":        CONST_GLFW_MOUSE_BUTTON_MIDDLE,

	// gltext
	"gltext.LeftToRight":           CONST_GLTEXT_LEFT_TO_RIGHT,
	"gltext.RightToLeft":           CONST_GLTEXT_RIGHT_TO_LEFT,
	"gltext.TopToBottom":           CONST_GLTEXT_TOP_TO_BOTTOM,
}

var Constants map[int]CXConstant = map[int]CXConstant{
	/* gl_1_0 */
	CONST_GL_DEPTH_BUFFER_BIT:       CXConstant{Type: TYPE_I32, Value: FromI32(0x00000100)},
	CONST_GL_STENCIL_BUFFER_BIT:     CXConstant{Type: TYPE_I32, Value: FromI32(0x00000400)},
	CONST_GL_COLOR_BUFFER_BIT:       CXConstant{Type: TYPE_I32, Value: FromI32(0x00004000)},
	CONST_GL_FALSE:                  CXConstant{Type: TYPE_I32, Value: FromI32(0)},
	CONST_GL_TRUE:                   CXConstant{Type: TYPE_I32, Value: FromI32(1)},
	CONST_GL_POINTS:                 CXConstant{Type: TYPE_I32, Value: FromI32(0x0000)},
	CONST_GL_LINES:                  CXConstant{Type: TYPE_I32, Value: FromI32(0x0001)},
	CONST_GL_LINE_LOOP:              CXConstant{Type: TYPE_I32, Value: FromI32(0x0002)},
	CONST_GL_LINE_STRIP:             CXConstant{Type: TYPE_I32, Value: FromI32(0x0003)},
	CONST_GL_TRIANGLES:              CXConstant{Type: TYPE_I32, Value: FromI32(0x0004)},
	CONST_GL_TRIANGLE_STRIP:         CXConstant{Type: TYPE_I32, Value: FromI32(0x0005)},
	CONST_GL_TRIANGLE_FAN:           CXConstant{Type: TYPE_I32, Value: FromI32(0x0006)},
	CONST_GL_QUADS:                  CXConstant{Type: TYPE_I32, Value: FromI32(0x0007)},
	CONST_GL_LESS:                   CXConstant{Type: TYPE_I32, Value: FromI32(0x201)},
	CONST_GL_LEQUAL:                 CXConstant{Type: TYPE_I32, Value: FromI32(0x202)},
	CONST_GL_SRC_ALPHA:              CXConstant{Type: TYPE_I32, Value: FromI32(0x302)},
	CONST_GL_ONE_MINUS_SRC_ALPHA:    CXConstant{Type: TYPE_I32, Value: FromI32(0x303)},
	CONST_GL_FRONT:                  CXConstant{Type: TYPE_I32, Value: FromI32(0x404)},
	CONST_GL_BACK:                   CXConstant{Type: TYPE_I32, Value: FromI32(0x405)},
	CONST_GL_FRONT_AND_BACK:         CXConstant{Type: TYPE_I32, Value: FromI32(0x408)},
	CONST_GL_NO_ERROR:               CXConstant{Type: TYPE_I32, Value: FromI32(0)},
	CONST_GL_INVALID_ENUM:           CXConstant{Type: TYPE_I32, Value: FromI32(0x500)},
	CONST_GL_INVALID_VALUE:          CXConstant{Type: TYPE_I32, Value: FromI32(0x501)},
	CONST_GL_INVALID_OPERATION:      CXConstant{Type: TYPE_I32, Value: FromI32(0x502)},
	CONST_GL_STACK_OVERFLOW:         CXConstant{Type: TYPE_I32, Value: FromI32(0x503)},
	CONST_GL_STACK_UNDERFLOW:        CXConstant{Type: TYPE_I32, Value: FromI32(0x504)},
	CONST_GL_OUT_OF_MEMORY:          CXConstant{Type: TYPE_I32, Value: FromI32(0x505)},
	CONST_GL_LINE_SMOOTH:            CXConstant{Type: TYPE_I32, Value: FromI32(0x0B20)},
	CONST_GL_POLYGON_SMOOTH:         CXConstant{Type: TYPE_I32, Value: FromI32(0x0B41)},
	CONST_GL_CULL_FACE:              CXConstant{Type: TYPE_I32, Value: FromI32(0x0B44)},
	CONST_GL_DEPTH_TEST:             CXConstant{Type: TYPE_I32, Value: FromI32(0x0B71)},
	CONST_GL_DITHER:                 CXConstant{Type: TYPE_I32, Value: FromI32(0x0BD0)},
	CONST_GL_BLEND:                  CXConstant{Type: TYPE_I32, Value: FromI32(0x0BE2)},
	CONST_GL_POLYGON_SMOOTH_HINT:    CXConstant{Type: TYPE_I32, Value: FromI32(0x0C53)},
	CONST_GL_TEXTURE_2D:             CXConstant{Type: TYPE_I32, Value: FromI32(0x0DE1)},
	CONST_GL_DONT_CARE:              CXConstant{Type: TYPE_I32, Value: FromI32(0x1100)},
	CONST_GL_UNSIGNED_BYTE:          CXConstant{Type: TYPE_I32, Value: FromI32(0x1401)},
	CONST_GL_FLOAT:                  CXConstant{Type: TYPE_I32, Value: FromI32(0x1406)},
	CONST_GL_TEXTURE:                CXConstant{Type: TYPE_I32, Value: FromI32(0x1702)},
	CONST_GL_COLOR:                  CXConstant{Type: TYPE_I32, Value: FromI32(0x1800)},
	CONST_GL_RGBA:                   CXConstant{Type: TYPE_I32, Value: FromI32(0x1908)},
	CONST_GL_REPLACE:                CXConstant{Type: TYPE_I32, Value: FromI32(0x1E01)},
	CONST_GL_NEAREST:                CXConstant{Type: TYPE_I32, Value: FromI32(0x2600)},
	CONST_GL_LINEAR:                 CXConstant{Type: TYPE_I32, Value: FromI32(0x2601)},
	CONST_GL_NEAREST_MIPMAP_NEAREST: CXConstant{Type: TYPE_I32, Value: FromI32(0x2700)},
	CONST_GL_LINEAR_MIPMAP_NEAREST:  CXConstant{Type: TYPE_I32, Value: FromI32(0x2701)},
	CONST_GL_NEAREST_MIPMAP_LINEAR:  CXConstant{Type: TYPE_I32, Value: FromI32(0x2702)},
	CONST_GL_LINEAR_MIPMAP_LINEAR:   CXConstant{Type: TYPE_I32, Value: FromI32(0x2703)},
	CONST_GL_TEXTURE_MAG_FILTER:     CXConstant{Type: TYPE_I32, Value: FromI32(0x2800)},
	CONST_GL_TEXTURE_MIN_FILTER:     CXConstant{Type: TYPE_I32, Value: FromI32(0x2801)},
	CONST_GL_TEXTURE_WRAP_S:         CXConstant{Type: TYPE_I32, Value: FromI32(0x2802)},
	CONST_GL_TEXTURE_WRAP_T:         CXConstant{Type: TYPE_I32, Value: FromI32(0x2803)},
	CONST_GL_REPEAT:                 CXConstant{Type: TYPE_I32, Value: FromI32(0x2901)},

	/* gl_1_1 */
	CONST_GL_RGBA8:                  CXConstant{Type: TYPE_I32, Value: FromI32(0x8058)},
	CONST_GL_VERTEX_ARRAY:           CXConstant{Type: TYPE_I32, Value: FromI32(0x8074)},

	/* gl_1_2 */
	CONST_GL_TEXTURE_WRAP_R:         CXConstant{Type: TYPE_I32, Value: FromI32(0x8072)},
	CONST_GL_CLAMP_TO_EDGE:          CXConstant{Type: TYPE_I32, Value: FromI32(0x812F)},

	/* gl_1_3 */
	CONST_GL_TEXTURE0:            CXConstant{Type: TYPE_I32, Value: FromI32(0x84C0)},
	CONST_GL_MULTISAMPLE_ARB:     CXConstant{Type: TYPE_I32, Value: FromI32(0x809D)}, // remove _ARB
	CONST_GL_CLAMP_TO_BORDER:     CXConstant{Type: TYPE_I32, Value: FromI32(0x812D)},

	/* gl_1_4 */
	CONST_GL_MIRRORED_REPEAT:     CXConstant{Type: TYPE_I32, Value: FromI32(0x8370)},

	/* gl_1_5 */
	CONST_GL_ARRAY_BUFFER:        CXConstant{Type: TYPE_I32, Value: FromI32(0x8892)},
	CONST_GL_STREAM_DRAW:         CXConstant{Type: TYPE_I32, Value: FromI32(0x88E0)},
	CONST_GL_STREAM_READ:         CXConstant{Type: TYPE_I32, Value: FromI32(0x88E1)},
	CONST_GL_STREAM_COPY:         CXConstant{Type: TYPE_I32, Value: FromI32(0x88E2)},
	CONST_GL_STATIC_DRAW:         CXConstant{Type: TYPE_I32, Value: FromI32(0x88E4)},
	CONST_GL_STATIC_READ:         CXConstant{Type: TYPE_I32, Value: FromI32(0x88E5)},
	CONST_GL_STATIC_COPY:         CXConstant{Type: TYPE_I32, Value: FromI32(0x88E6)},
	CONST_GL_DYNAMIC_DRAW:        CXConstant{Type: TYPE_I32, Value: FromI32(0x88E8)},
	CONST_GL_DYNAMIC_READ:        CXConstant{Type: TYPE_I32, Value: FromI32(0x88E9)},
	CONST_GL_DYNAMIC_COPY:        CXConstant{Type: TYPE_I32, Value: FromI32(0x88EA)},

	/* gl_2_0 */
	CONST_GL_FRAGMENT_SHADER:     CXConstant{Type: TYPE_I32, Value: FromI32(0x8B30)},
	CONST_GL_VERTEX_SHADER:       CXConstant{Type: TYPE_I32, Value: FromI32(0x8B31)},

	/* gl_3_0 */
	CONST_GL_FRAMEBUFFER_UNDEFINED:                     CXConstant{Type: TYPE_I32, Value: FromI32(0x8219)},
	CONST_GL_FRAMEBUFFER_COMPLETE:                      CXConstant{Type: TYPE_I32, Value: FromI32(0x8CD5)},
	CONST_GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT:         CXConstant{Type: TYPE_I32, Value: FromI32(0x8CD6)},
	CONST_GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT: CXConstant{Type: TYPE_I32, Value: FromI32(0x8CD7)},
	CONST_GL_FRAMEBUFFER_INCOMPLETE_DRAW_BUFFER:        CXConstant{Type: TYPE_I32, Value: FromI32(0x8CDB)},
	CONST_GL_FRAMEBUFFER_INCOMPLETE_READ_BUFFER:        CXConstant{Type: TYPE_I32, Value: FromI32(0x8CDC)},
	CONST_GL_FRAMEBUFFER_UNSUPPORTED:                   CXConstant{Type: TYPE_I32, Value: FromI32(0x8CDD)},
	CONST_GL_COLOR_ATTACHMENT0:                         CXConstant{Type: TYPE_I32, Value: FromI32(0x8CE0)},
	CONST_GL_FRAMEBUFFER:                               CXConstant{Type: TYPE_I32, Value: FromI32(0x8D40)},
	CONST_GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE:        CXConstant{Type: TYPE_I32, Value: FromI32(0x8D56)},

	/* gl_3_2 */
	CONST_GL_FRAMEBUFFER_INCOMPLETE_LAYER_TARGETS:      CXConstant{Type: TYPE_I32, Value: FromI32(0x8DA8)},

	/* gl_4_4 */
	CONST_GL_MIRROR_CLAMP_TO_EDGE: CXConstant{Type: TYPE_I32, Value: FromI32(0x8743)},

	/* Fixed pipeline. Deprecated ? */
	CONST_GL_POLYGON:             CXConstant{Type: TYPE_I32, Value: FromI32(9)},
	CONST_GL_MODELVIEW:           CXConstant{Type: TYPE_I32, Value: FromI32(5888)},
	CONST_GL_PROJECTION:          CXConstant{Type: TYPE_I32, Value: FromI32(5889)},
	CONST_GL_MODELVIEW_MATRIX:    CXConstant{Type: TYPE_I32, Value: FromI32(2982)},
	CONST_GL_LIGHTING:            CXConstant{Type: TYPE_I32, Value: FromI32(2896)},
	CONST_GL_LIGHT0:              CXConstant{Type: TYPE_I32, Value: FromI32(16384)},
	CONST_GL_AMBIENT:             CXConstant{Type: TYPE_I32, Value: FromI32(4608)},
	CONST_GL_DIFFUSE:             CXConstant{Type: TYPE_I32, Value: FromI32(4609)},
	CONST_GL_POSITION:            CXConstant{Type: TYPE_I32, Value: FromI32(4611)},
	CONST_GL_TEXTURE_ENV:         CXConstant{Type: TYPE_I32, Value: FromI32(8960)},
	CONST_GL_TEXTURE_ENV_MODE:    CXConstant{Type: TYPE_I32, Value: FromI32(8704)},
	CONST_GL_MODULATE:            CXConstant{Type: TYPE_I32, Value: FromI32(8448)},
	CONST_GL_DECAL:               CXConstant{Type: TYPE_I32, Value: FromI32(8449)},
	CONST_GL_POINT_SMOOTH:        CXConstant{Type: TYPE_I32, Value: FromI32(2832)},

	// glfw
	CONST_GLFW_FALSE:                     CXConstant{Type: TYPE_I32, Value: FromI32(0)},
	CONST_GLFW_TRUE:                      CXConstant{Type: TYPE_I32, Value: FromI32(1)},
	CONST_GLFW_PRESS:                     CXConstant{Type: TYPE_I32, Value: FromI32(1)},
	CONST_GLFW_RELEASE:                   CXConstant{Type: TYPE_I32, Value: FromI32(0)},
	CONST_GLFW_REPEAT:                    CXConstant{Type: TYPE_I32, Value: FromI32(2)},
	CONST_GLFW_KEY_UNKNOWN:               CXConstant{Type: TYPE_I32, Value: FromI32(-1)},
	CONST_GLFW_CURSOR:                    CXConstant{Type: TYPE_I32, Value: FromI32(208897)},
	CONST_GLFW_STICKY_KEYS:               CXConstant{Type: TYPE_I32, Value: FromI32(208898)},
	CONST_GLFW_STICKY_MOUSE_BUTTONS:      CXConstant{Type: TYPE_I32, Value: FromI32(208899)},
	CONST_GLFW_CURSOR_NORMAL:             CXConstant{Type: TYPE_I32, Value: FromI32(212993)},
	CONST_GLFW_CURSOR_HIDDEN:             CXConstant{Type: TYPE_I32, Value: FromI32(212994)},
	CONST_GLFW_CURSOR_DISABLED:           CXConstant{Type: TYPE_I32, Value: FromI32(212995)},
	CONST_GLFW_RESIZABLE:                 CXConstant{Type: TYPE_I32, Value: FromI32(131075)},
	CONST_GLFW_CONTEXT_VERSION_MAJOR:     CXConstant{Type: TYPE_I32, Value: FromI32(139266)},
	CONST_GLFW_CONTEXT_VERSION_MINOR:     CXConstant{Type: TYPE_I32, Value: FromI32(139267)},
	CONST_GLFW_OPENGL_PROFILE:            CXConstant{Type: TYPE_I32, Value: FromI32(139272)},
	CONST_GLFW_OPENGL_COREPROFILE:        CXConstant{Type: TYPE_I32, Value: FromI32(204801)},
	CONST_GLFW_OPENGL_FORWARD_COMPATIBLE: CXConstant{Type: TYPE_I32, Value: FromI32(139270)},
	CONST_GLFW_MOUSE_BUTTON_LAST:         CXConstant{Type: TYPE_I32, Value: FromI32(7)},
	CONST_GLFW_MOUSE_BUTTON_LEFT:         CXConstant{Type: TYPE_I32, Value: FromI32(0)},
	CONST_GLFW_MOUSE_BUTTON_RIGHT:        CXConstant{Type: TYPE_I32, Value: FromI32(1)},
	CONST_GLFW_MOUSE_BUTTON_MIDDLE:       CXConstant{Type: TYPE_I32, Value: FromI32(2)},

	// gltext
	CONST_GLTEXT_LEFT_TO_RIGHT:           CXConstant{Type: TYPE_I32, Value: FromI32(0)},
	CONST_GLTEXT_RIGHT_TO_LEFT:           CXConstant{Type: TYPE_I32, Value: FromI32(1)},
	CONST_GLTEXT_TOP_TO_BOTTOM:           CXConstant{Type: TYPE_I32, Value: FromI32(2)},
}
