//go:build !cgo
// +build !cgo

package rl

import (
	"unsafe"

	"github.com/gen2brain/raylib-go/raylib/internal/convert"
	"github.com/jupiterrider/ffi"
)

var (
	// dll is the pointer to the shared library
	dll ffi.Lib = loadLibrary()

	// Window-related functions

	initWindow               = dll.MustPrep("InitWindow", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypePointer)
	closeWindow              = dll.MustPrep("CloseWindow", &ffi.TypeVoid)
	setTraceLogCallback      = dll.MustPrep("SetTraceLogCallback", &ffi.TypeVoid, &ffi.TypePointer)
	windowShouldClose        = dll.MustPrep("WindowShouldClose", &ffi.TypeUint8)
	isWindowReady            = dll.MustPrep("IsWindowReady", &ffi.TypeUint8)
	isWindowFullscreen       = dll.MustPrep("IsWindowFullscreen", &ffi.TypeUint8)
	isWindowHidden           = dll.MustPrep("IsWindowHidden", &ffi.TypeUint8)
	isWindowMinimized        = dll.MustPrep("IsWindowMinimized", &ffi.TypeUint8)
	isWindowMaximized        = dll.MustPrep("IsWindowMaximized", &ffi.TypeUint8)
	isWindowFocused          = dll.MustPrep("IsWindowFocused", &ffi.TypeUint8)
	isWindowResized          = dll.MustPrep("IsWindowResized", &ffi.TypeUint8)
	isWindowState            = dll.MustPrep("IsWindowState", &ffi.TypeUint8, &ffi.TypeUint32)
	setWindowState           = dll.MustPrep("SetWindowState", &ffi.TypeVoid, &ffi.TypeUint32)
	clearWindowState         = dll.MustPrep("ClearWindowState", &ffi.TypeVoid, &ffi.TypeUint32)
	toggleFullscreen         = dll.MustPrep("ToggleFullscreen", &ffi.TypeVoid)
	toggleBorderlessWindowed = dll.MustPrep("ToggleBorderlessWindowed", &ffi.TypeVoid)
	maximizeWindow           = dll.MustPrep("MaximizeWindow", &ffi.TypeVoid)
	minimizeWindow           = dll.MustPrep("MinimizeWindow", &ffi.TypeVoid)
	restoreWindow            = dll.MustPrep("RestoreWindow", &ffi.TypeVoid)
	setWindowIcon            = dll.MustPrep("SetWindowIcon", &ffi.TypeVoid, &typeImage)
	setWindowIcons           = dll.MustPrep("SetWindowIcons", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypeSint32)
	setWindowTitle           = dll.MustPrep("SetWindowTitle", &ffi.TypeVoid, &ffi.TypePointer)
	setWindowPosition        = dll.MustPrep("SetWindowPosition", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32)
	setWindowMonitor         = dll.MustPrep("SetWindowMonitor", &ffi.TypeVoid, &ffi.TypeSint32)
	setWindowMinSize         = dll.MustPrep("SetWindowMinSize", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32)
	setWindowMaxSize         = dll.MustPrep("SetWindowMaxSize", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32)
	setWindowSize            = dll.MustPrep("SetWindowSize", &ffi.TypeSint32, &ffi.TypeSint32)
	setWindowOpacity         = dll.MustPrep("SetWindowOpacity", &ffi.TypeVoid, &ffi.TypeFloat)
	setWindowFocused         = dll.MustPrep("SetWindowFocused", &ffi.TypeVoid)
	getWindowHandle          = dll.MustPrep("GetWindowHandle", &ffi.TypePointer)
	getScreenWidth           = dll.MustPrep("GetWindowHandle", &ffi.TypeSint32)
	getScreenHeight          = dll.MustPrep("GetScreenHeight", &ffi.TypeSint32)
	getRenderWidth           = dll.MustPrep("GetRenderWidth", &ffi.TypeSint32)
	getRenderHeight          = dll.MustPrep("GetRenderHeight", &ffi.TypeSint32)
	getMonitorCount          = dll.MustPrep("GetMonitorCount", &ffi.TypeSint32)
	getCurrentMonitor        = dll.MustPrep("GetCurrentMonitor", &ffi.TypeSint32)
	getMonitorPosition       = dll.MustPrep("GetMonitorPosition", &typeVector2, &ffi.TypeSint32)
	getMonitorWidth          = dll.MustPrep("GetMonitorWidth", &ffi.TypeSint32, &ffi.TypeSint32)
	getMonitorHeight         = dll.MustPrep("GetMonitorHeight", &ffi.TypeSint32, &ffi.TypeSint32)
	getMonitorPhysicalWidth  = dll.MustPrep("GetMonitorPhysicalWidth", &ffi.TypeSint32, &ffi.TypeSint32)
	getMonitorPhysicalHeight = dll.MustPrep("GetMonitorPhysicalHeight", &ffi.TypeSint32, &ffi.TypeSint32)
	getMonitorRefreshRate    = dll.MustPrep("GetMonitorRefreshRate", &ffi.TypeSint32, &ffi.TypeSint32)
	getWindowPosition        = dll.MustPrep("GetWindowPosition", &typeVector2)
	getWindowScaleDPI        = dll.MustPrep("GetWindowScaleDPI", &typeVector2)
	getMonitorName           = dll.MustPrep("GetMonitorName", &ffi.TypePointer, &ffi.TypeSint32)
	setClipboardText         = dll.MustPrep("SetClipboardText", &ffi.TypeVoid, &ffi.TypePointer)
	getClipboardText         = dll.MustPrep("GetClipboardText", &ffi.TypePointer)
	getClipboardImage        = dll.MustPrep("GetClipboardImage", &typeImage)
	enableEventWaiting       = dll.MustPrep("EnableEventWaiting", &ffi.TypeVoid)
	disableEventWaiting      = dll.MustPrep("DisableEventWaiting", &ffi.TypeVoid)

	// Cursor-related functions

	showCursor       = dll.MustPrep("ShowCursor", &ffi.TypeVoid)
	hideCursor       = dll.MustPrep("HideCursor", &ffi.TypeVoid)
	isCursorHidden   = dll.MustPrep("IsCursorHidden", &ffi.TypeUint8)
	enableCursor     = dll.MustPrep("EnableCursor", &ffi.TypeVoid)
	disableCursor    = dll.MustPrep("DisableCursor", &ffi.TypeVoid)
	isCursorOnScreen = dll.MustPrep("IsCursorOnScreen", &ffi.TypeUint8)
)

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

// GetMonitorWidth - Get specified monitor width (current video mode used by monitor)
func GetMonitorWidth(monitor int) int {
	var ret ffi.Arg
	m := int32(monitor)
	getMonitorWidth.Call(&ret, &m)
	return int(ret)
}

// GetMonitorHeight - Get specified monitor height (current video mode used by monitor)
func GetMonitorHeight(monitor int) int {
	var ret ffi.Arg
	m := int32(monitor)
	getMonitorHeight.Call(&ret, &m)
	return int(ret)
}

// GetMonitorPhysicalWidth - Get specified monitor physical width in millimetres
func GetMonitorPhysicalWidth(monitor int) int {
	var ret ffi.Arg
	m := int32(monitor)
	getMonitorPhysicalWidth.Call(&ret, &m)
	return int(ret)
}

// GetMonitorPhysicalHeight - Get specified monitor physical height in millimetres
func GetMonitorPhysicalHeight(monitor int) int {
	var ret ffi.Arg
	m := int32(monitor)
	getMonitorPhysicalHeight.Call(&ret, &m)
	return int(ret)
}

// GetMonitorRefreshRate - Get specified monitor refresh rate
func GetMonitorRefreshRate(monitor int) int {
	var ret ffi.Arg
	m := int32(monitor)
	getMonitorRefreshRate.Call(&ret, &m)
	return int(ret)
}

// GetWindowPosition - Get window position XY on monitor
func GetWindowPosition() Vector2 {
	var ret Vector2
	getWindowPosition.Call(&ret)
	return ret
}

// GetWindowScaleDPI - Get window scale DPI factor
func GetWindowScaleDPI() Vector2 {
	var ret Vector2
	getWindowScaleDPI.Call(&ret)
	return ret
}

// GetMonitorName - Get the human-readable, UTF-8 encoded name of the specified monitor
func GetMonitorName(monitor int) string {
	var ret *byte
	m := int32(monitor)
	getMonitorName.Call(&ret, &m)
	return convert.ToString(ret)
}

// SetClipboardText - Set clipboard text content
func SetClipboardText(text string) {
	cText := convert.ToBytePtr(text)
	setClipboardText.Call(nil, &cText)
}

// GetClipboardText - Get clipboard text content
func GetClipboardText() string {
	var ret *byte
	getClipboardText.Call(&ret)
	return convert.ToString(ret)
}

// GetClipboardImage - Get clipboard image content
//
// Only works with SDL3 backend or Windows with RGFW/GLFW
func GetClipboardImage() Image {
	var ret Image
	getClipboardImage.Call(&ret)
	return ret
}

// EnableEventWaiting - Enable waiting for events on EndDrawing(), no automatic event polling
func EnableEventWaiting() {
	enableEventWaiting.Call(nil)
}

// DisableEventWaiting - Disable waiting for events on EndDrawing(), automatic events polling
func DisableEventWaiting() {
	disableEventWaiting.Call(nil)
}

// ShowCursor - Shows cursor
func ShowCursor() {
	showCursor.Call(nil)
}

// HideCursor - Hides cursor
func HideCursor() {
	hideCursor.Call(nil)
}

// IsCursorHidden - Check if cursor is not visible
func IsCursorHidden() bool {
	var ret ffi.Arg
	isCursorHidden.Call(&ret)
	return ret.Bool()
}

// EnableCursor - Enables cursor (unlock cursor)
func EnableCursor() {
	enableCursor.Call(nil)
}

// DisableCursor - Disables cursor (lock cursor)
func DisableCursor() {
	disableCursor.Call(nil)
}

// IsCursorOnScreen - Check if cursor is on the screen
func IsCursorOnScreen() bool {
	var ret ffi.Arg
	isCursorOnScreen.Call(&ret)
	return ret.Bool()
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
