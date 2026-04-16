//go:build !cgo
// +build !cgo

package rl

import (
	"unsafe"

	"github.com/gen2brain/raylib-go/raylib/internal/convert"
	"github.com/jupiterrider/ffi"
)

var (
	// raylibDll is the pointer to the shared library
	raylibDll ffi.Lib

	initWindow               ffi.Fun
	closeWindow              ffi.Fun
	setTraceLogCallback      ffi.Fun
	windowShouldClose        ffi.Fun
	isWindowReady            ffi.Fun
	isWindowFullscreen       ffi.Fun
	isWindowHidden           ffi.Fun
	isWindowMinimized        ffi.Fun
	isWindowMaximized        ffi.Fun
	isWindowFocused          ffi.Fun
	isWindowResized          ffi.Fun
	isWindowState            ffi.Fun
	setWindowState           ffi.Fun
	clearWindowState         ffi.Fun
	toggleFullscreen         ffi.Fun
	toggleBorderlessWindowed ffi.Fun
	maximizeWindow           ffi.Fun
	minimizeWindow           ffi.Fun
	restoreWindow            ffi.Fun
	setWindowIcon            ffi.Fun
	setWindowIcons           ffi.Fun
	setWindowTitle           ffi.Fun
	setWindowPosition        ffi.Fun
	setWindowMonitor         ffi.Fun
	setWindowMinSize         ffi.Fun
	setWindowMaxSize         ffi.Fun
	setWindowSize            ffi.Fun
	setWindowOpacity         ffi.Fun
	setWindowFocused         ffi.Fun
	getWindowHandle          ffi.Fun
	getScreenWidth           ffi.Fun
	getScreenHeight          ffi.Fun
	getRenderWidth           ffi.Fun
	getRenderHeight          ffi.Fun
	getMonitorCount          ffi.Fun
	getCurrentMonitor        ffi.Fun
	getMonitorPosition       ffi.Fun
)

func init() {
	raylibDll = loadLibrary()

	initWindow = must(raylibDll.Prep("InitWindow", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypePointer))
	closeWindow = must(raylibDll.Prep("CloseWindow", &ffi.TypeVoid))
	setTraceLogCallback = must(raylibDll.Prep("SetTraceLogCallback", &ffi.TypeVoid, &ffi.TypePointer))
	windowShouldClose = must(raylibDll.Prep("WindowShouldClose", &ffi.TypeUint8))
	isWindowReady = must(raylibDll.Prep("IsWindowReady", &ffi.TypeUint8))
	isWindowFullscreen = must(raylibDll.Prep("IsWindowFullscreen", &ffi.TypeUint8))
	isWindowHidden = must(raylibDll.Prep("IsWindowHidden", &ffi.TypeUint8))
	isWindowMinimized = must(raylibDll.Prep("IsWindowMinimized", &ffi.TypeUint8))
	isWindowMaximized = must(raylibDll.Prep("IsWindowMaximized", &ffi.TypeUint8))
	isWindowFocused = must(raylibDll.Prep("IsWindowFocused", &ffi.TypeUint8))
	isWindowResized = must(raylibDll.Prep("IsWindowResized", &ffi.TypeUint8))
	isWindowState = must(raylibDll.Prep("IsWindowState", &ffi.TypeUint8, &ffi.TypeUint32))
	setWindowState = must(raylibDll.Prep("SetWindowState", &ffi.TypeVoid, &ffi.TypeUint32))
	clearWindowState = must(raylibDll.Prep("ClearWindowState", &ffi.TypeVoid, &ffi.TypeUint32))
	toggleFullscreen = must(raylibDll.Prep("ToggleFullscreen", &ffi.TypeVoid))
	toggleBorderlessWindowed = must(raylibDll.Prep("ToggleBorderlessWindowed", &ffi.TypeVoid))
	maximizeWindow = must(raylibDll.Prep("MaximizeWindow", &ffi.TypeVoid))
	minimizeWindow = must(raylibDll.Prep("MinimizeWindow", &ffi.TypeVoid))
	restoreWindow = must(raylibDll.Prep("RestoreWindow", &ffi.TypeVoid))
	setWindowIcon = must(raylibDll.Prep("SetWindowIcon", &ffi.TypeVoid, &typeImage))
	setWindowIcons = must(raylibDll.Prep("SetWindowIcons", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypeSint32))
	setWindowTitle = must(raylibDll.Prep("SetWindowTitle", &ffi.TypeVoid, &ffi.TypePointer))
	setWindowPosition = must(raylibDll.Prep("SetWindowPosition", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32))
	setWindowMonitor = must(raylibDll.Prep("SetWindowMonitor", &ffi.TypeVoid, &ffi.TypeSint32))
	setWindowMinSize = must(raylibDll.Prep("SetWindowMinSize", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32))
	setWindowMaxSize = must(raylibDll.Prep("SetWindowMaxSize", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32))
	setWindowSize = must(raylibDll.Prep("SetWindowSize", &ffi.TypeSint32, &ffi.TypeSint32))
	setWindowOpacity = must(raylibDll.Prep("SetWindowOpacity", &ffi.TypeVoid, &ffi.TypeFloat))
	setWindowFocused = must(raylibDll.Prep("SetWindowFocused", &ffi.TypeVoid))
	getWindowHandle = must(raylibDll.Prep("GetWindowHandle", &ffi.TypePointer))
	getScreenWidth = must(raylibDll.Prep("GetWindowHandle", &ffi.TypeSint32))
	getScreenHeight = must(raylibDll.Prep("GetScreenHeight", &ffi.TypeSint32))
	getRenderWidth = must(raylibDll.Prep("GetRenderWidth", &ffi.TypeSint32))
	getRenderHeight = must(raylibDll.Prep("GetRenderHeight", &ffi.TypeSint32))
	getMonitorCount = must(raylibDll.Prep("GetMonitorCount", &ffi.TypeSint32))
	getCurrentMonitor = must(raylibDll.Prep("GetCurrentMonitor", &ffi.TypeSint32))
	getMonitorPosition = must(raylibDll.Prep("GetMonitorPosition", &typeVector2, &ffi.TypeSint32))
}

// InitWindow - Initialize window and OpenGL context
func InitWindow(width int32, height int32, title string) {
	cTitle := convert.ToBytePtr(title)
	initWindow.Call(nil, &width, &height, &cTitle)
}

// CloseWindow - Close window and unload OpenGL context
func CloseWindow() {
	closeWindow.Call(nil)
}

// WindowShouldClose - Check if application should close (KEY_ESCAPE pressed or windows close icon clicked)
func WindowShouldClose() bool {
	var ret ffi.Arg
	windowShouldClose.Call(&ret)
	return ret.Bool()
}

// IsWindowReady - Check if window has been initialized successfully
func IsWindowReady() bool {
	var ret ffi.Arg
	isWindowReady.Call(&ret)
	return ret.Bool()
}

// IsWindowFullscreen - Check if window is currently fullscreen
func IsWindowFullscreen() bool {
	var ret ffi.Arg
	isWindowFullscreen.Call(&ret)
	return ret.Bool()
}

// IsWindowHidden - Check if window is currently hidden (only PLATFORM_DESKTOP)
func IsWindowHidden() bool {
	var ret ffi.Arg
	isWindowHidden.Call(&ret)
	return ret.Bool()
}

// IsWindowMinimized - Check if window is currently minimized (only PLATFORM_DESKTOP)
func IsWindowMinimized() bool {
	var ret ffi.Arg
	isWindowMinimized.Call(&ret)
	return ret.Bool()
}

// IsWindowMaximized - Check if window is currently maximized (only PLATFORM_DESKTOP)
func IsWindowMaximized() bool {
	var ret ffi.Arg
	isWindowMaximized.Call(&ret)
	return ret.Bool()
}

// IsWindowFocused - Check if window is currently focused (only PLATFORM_DESKTOP)
func IsWindowFocused() bool {
	var ret ffi.Arg
	isWindowFocused.Call(&ret)
	return ret.Bool()
}

// IsWindowResized - Check if window has been resized last frame
func IsWindowResized() bool {
	var ret ffi.Arg
	isWindowResized.Call(&ret)
	return ret.Bool()
}

// IsWindowState - Check if one specific window flag is enabled
func IsWindowState(flag uint32) bool {
	var ret ffi.Arg
	isWindowState.Call(&ret, &flag)
	return ret.Bool()
}

// SetWindowState - Set window configuration state using flags (only PLATFORM_DESKTOP)
func SetWindowState(flags uint32) {
	setWindowState.Call(nil, &flags)
}

// ClearWindowState - Clear window configuration state flags
func ClearWindowState(flags uint32) {
	clearWindowState.Call(nil, &flags)
}

// ToggleFullscreen - Toggle window state: fullscreen/windowed (only PLATFORM_DESKTOP)
func ToggleFullscreen() {
	toggleFullscreen.Call(nil)
}

// ToggleBorderlessWindowed - Toggle window state: borderless windowed (only PLATFORM_DESKTOP)
func ToggleBorderlessWindowed() {
	toggleBorderlessWindowed.Call(nil)
}

// MaximizeWindow - Set window state: maximized, if resizable (only PLATFORM_DESKTOP)
func MaximizeWindow() {
	maximizeWindow.Call(nil)
}

// MinimizeWindow - Set window state: minimized, if resizable (only PLATFORM_DESKTOP)
func MinimizeWindow() {
	minimizeWindow.Call(nil)
}

// RestoreWindow - Set window state: not minimized/maximized (only PLATFORM_DESKTOP)
func RestoreWindow() {
	restoreWindow.Call(nil)
}

// SetWindowIcon - Set icon for window (single image, RGBA 32bit, only PLATFORM_DESKTOP)
func SetWindowIcon(image Image) {
	setWindowIcon.Call(nil, &image)
}

// SetWindowIcons - Set icon for window (multiple images, RGBA 32bit, only PLATFORM_DESKTOP)
func SetWindowIcons(images []Image, count int32) {
	imagesPtr := &images[0]
	setWindowIcons.Call(nil, &imagesPtr, &count)
}

// SetWindowTitle - Set title for window (only PLATFORM_DESKTOP and PLATFORM_WEB)
func SetWindowTitle(title string) {
	cTitle := convert.ToBytePtr(title)
	setWindowTitle.Call(nil, &cTitle)
}

// SetWindowPosition - Set window position on screen (only PLATFORM_DESKTOP)
func SetWindowPosition(x int, y int) {
	posX, posY := int32(x), int32(y)
	setWindowPosition.Call(nil, &posX, &posY)
}

// SetWindowMonitor - Set monitor for the current window
func SetWindowMonitor(monitor int) {
	m := int32(monitor)
	setWindowMonitor.Call(nil, &m)
}

// SetWindowMinSize - Set window minimum dimensions (for FLAG_WINDOW_RESIZABLE)
func SetWindowMinSize(width int, height int) {
	w, h := int32(width), int32(height)
	setWindowMinSize.Call(nil, &w, &h)
}

// SetWindowMaxSize - Set window maximum dimensions (for FLAG_WINDOW_RESIZABLE)
func SetWindowMaxSize(width int, height int) {
	w, h := int32(width), int32(height)
	setWindowMaxSize.Call(nil, &w, &h)
}

// SetWindowSize - Set window dimensions
func SetWindowSize(width int, height int) {
	w, h := int32(width), int32(height)
	setWindowSize.Call(nil, &w, &h)
}

// SetWindowOpacity - Set window opacity [0.0f..1.0f] (only PLATFORM_DESKTOP)
func SetWindowOpacity(opacity float32) {
	setWindowOpacity.Call(nil, &opacity)
}

// SetWindowFocused - Set window focused (only PLATFORM_DESKTOP)
func SetWindowFocused() {
	setWindowFocused.Call(nil)
}

// GetWindowHandle - Get native window handle
func GetWindowHandle() unsafe.Pointer {
	var ret unsafe.Pointer
	getWindowHandle.Call(&ret)
	return ret
}

// GetScreenWidth - Get current screen width
func GetScreenWidth() int {
	var ret ffi.Arg
	getScreenWidth.Call(&ret)
	return int(ret)
}

// GetScreenHeight - Get current screen height
func GetScreenHeight() int {
	var ret ffi.Arg
	getScreenHeight.Call(&ret)
	return int(ret)
}

// GetRenderWidth - Get current render width (it considers HiDPI)
func GetRenderWidth() int {
	var ret ffi.Arg
	getRenderWidth.Call(&ret)
	return int(ret)
}

// GetRenderHeight - Get current render height (it considers HiDPI)
func GetRenderHeight() int {
	var ret ffi.Arg
	getRenderHeight.Call(&ret)
	return int(ret)
}

// GetMonitorCount - Get number of connected monitors
func GetMonitorCount() int {
	var ret ffi.Arg
	getMonitorCount.Call(&ret)
	return int(ret)
}

// GetCurrentMonitor - Get current monitor where window is placed
func GetCurrentMonitor() int {
	var ret ffi.Arg
	getCurrentMonitor.Call(&ret)
	return int(ret)
}

// GetMonitorPosition - Get specified monitor position
func GetMonitorPosition(monitor int) Vector2 {
	var ret Vector2
	m := int32(monitor)
	getMonitorPosition.Call(&ret, &m)
	return ret
}

// SetTraceLogCallback - Set custom trace log
func SetTraceLogCallback(fn TraceLogCallbackFun) {
	cb := traceLogCallbackWrapper(fn)
	setTraceLogCallback.Call(nil, &cb)
}

// GetMouseDelta - Get mouse delta between frames
func GetMouseDelta() Vector2 {
	return Vector2{}
}

// IsKeyDown - Check if a key is being pressed
func IsKeyDown(key int32) bool {
	return false
}

// IsGamepadAvailable - Check if a gamepad is available
func IsGamepadAvailable(gamepad int32) bool {
	return false
}

// GetGamepadAxisMovement - Get axis movement value for a gamepad axis
func GetGamepadAxisMovement(gamepad int32, axis int32) float32 {
	return 0
}

// IsMouseButtonDown - Check if a mouse button is being pressed
func IsMouseButtonDown(button MouseButton) bool {
	return false
}

// IsKeyPressed - Check if a key has been pressed once
func IsKeyPressed(key int32) bool {
	return false
}

// GetMouseWheelMove - Get mouse wheel movement for X or Y, whichever is larger
func GetMouseWheelMove() float32 {
	return 0
}

// GetFrameTime - Get time in seconds for last frame drawn (delta time)
func GetFrameTime() float32 {
	return 0
}
