//go:build wasm

package glfw

var hints = make(map[Hint]int)

type Hint int

const (
	AlphaBits Hint = iota
	DepthBits
	StencilBits
	Samples

	Focused
	Iconified
	Maximized
	Visible
	Hovered
	Resizable
	Decorated
	Floating
	AutoIconify
	CenterCursor
	TransparentFramebuffer
	FocusOnShow
	ScaleToMonitor

	ClientAPI
	ContextVersionMajor
	ContextVersionMinor
	ContextRobustness
	ContextReleaseBehavior
	OpenGLForwardCompatible
	OpenGLDebugContext
	OpenGLProfile

	// goxjs/glfw-specific hints for WebGL.
	PremultipliedAlpha
	PreserveDrawingBuffer
	PreferLowPowerToHighPerformance
	FailIfMajorPerformanceCaveat
)

// NoAPI is the ClientAPI value requesting a context-less window: CreateWindow
// then skips WebGL context creation so the bare canvas can host a WebGPU context
// instead (a canvas holds exactly one context type). Mirrors glfw.NoAPI on
// desktop. Presence in the hints map (not this zero value) is what selects it,
// so an unset ClientAPI still defaults to the WebGL path. (trendvidia: WebGPU.)
const NoAPI int = 0

func WindowHint(target Hint, hint int) {
	hints[target] = hint
}
