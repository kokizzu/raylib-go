// This example demonstrates the low-level rlgl vertex buffer bindings:
//   - LoadVertexBuffer / LoadVertexBufferElement: upload vertex + index data to GPU
//   - SetVertexAttribute: configure vertex attribute pointers
//   - SetVertexAttributeDefault: provide default values for unused attributes
//   - DrawVertexArrayElements: indexed draw (quad)
//   - DrawVertexArray: non-indexed draw (triangle)
//   - UpdateVertexBuffer: animate vertex colors each frame
package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "rlgl vertex buffer example")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// ---------------------------------------------------------------
	// Quad (indexed draw) - centered on left half of screen
	// Vertices ordered CCW in NDC (since ortho flips Y, screen-CW = NDC-CW,
	// so we use CCW winding via index order)
	// ---------------------------------------------------------------
	quadPositions := []float32{
		// x, y,  z
		150, 150, 0, // 0: top-left
		450, 150, 0, // 1: top-right
		450, 450, 0, // 2: bottom-right
		150, 450, 0, // 3: bottom-left
	}

	quadColors := []float32{
		// r, g, b, a  (one color per corner)
		1, 0, 0, 1, // red
		0, 1, 0, 1, // green
		0, 0, 1, 1, // blue
		1, 1, 0, 1, // yellow
	}

	// Texcoords (all zeros - we only need the default white texture)
	quadTexcoords := []float32{
		0, 0,
		0, 0,
		0, 0,
		0, 0,
	}

	// CCW winding in NDC (with Y-flip ortho: reverse the original CW order)
	quadIndices := []uint16{
		0, 2, 1, // first triangle (CCW in NDC)
		0, 3, 2, // second triangle (CCW in NDC)
	}

	// ---------------------------------------------------------------
	// Quad VAO, VBO, and EBO Creations
	// DONT FORGOT TO CLEANUP YOUR GPU RESOURCES
	// ---------------------------------------------------------------

	// Create quad VAO
	quadVAO := rl.LoadVertexArray()
	rl.EnableVertexArray(quadVAO)

	// Position VBO (attribute 0, vec3)
	quadPosVBO := rl.LoadVertexBuffer(quadPositions, false)
	defer rl.UnloadVertexBuffer(quadPosVBO)
	rl.SetVertexAttribute(0, 3, rl.Float, false, 0, 0)
	rl.EnableVertexAttribute(0)

	// Texcoord VBO (attribute 1, vec2)
	quadTexVBO := rl.LoadVertexBuffer(quadTexcoords, false)
	defer rl.UnloadVertexBuffer(quadTexVBO)
	rl.SetVertexAttribute(1, 2, rl.Float, false, 0, 0)
	rl.EnableVertexAttribute(1)

	// Color VBO (attribute 3, vec4) - dynamic for animation
	quadColVBO := rl.LoadVertexBuffer(quadColors, true)
	defer rl.UnloadVertexBuffer(quadColVBO)
	rl.SetVertexAttribute(3, 4, rl.Float, false, 0, 0)
	rl.EnableVertexAttribute(3)

	// Index buffer (EBO)
	quadIBO := rl.LoadVertexBufferElements(quadIndices, false)
	defer rl.UnloadVertexBuffer(quadIBO)

	rl.DisableVertexArray()

	// ---------------------------------------------------------------
	// Triangle (non-indexed draw) - on the right side
	// Vertices in CCW order in NDC (with Y-flip ortho)
	// ---------------------------------------------------------------
	triPositions := []float32{
		// x, y, z  (CCW in NDC with Y-flip)
		575, 200, 0, // top
		575, 450, 0, // bottom-left
		700, 450, 0, // bottom-right
	}

	triColors := []float32{
		// r, g, b, a
		1.0, 0.0, 1.0, 1.0, // magenta
		0.0, 1.0, 1.0, 1.0, // cyan
		1.0, 1.0, 1.0, 1.0, // white
	}

	triTexcoords := []float32{
		0, 0,
		0, 0,
		0, 0,
	}

	// ---------------------------------------------------------------
  // Triangle VAO, and VBO creation
	// DONT FORGOT TO CLEANUP YOUR GPU RESOURCES
	// ---------------------------------------------------------------

	// Create triangle VAO
	triVAO := rl.LoadVertexArray()
	rl.EnableVertexArray(triVAO)

	// Position VBO (attribute 0, vec3)
	triPosVBO := rl.LoadVertexBuffer(triPositions, false)
	defer rl.UnloadVertexBuffer(triPosVBO)
	rl.SetVertexAttribute(0, 3, rl.Float, false, 0, 0)
	rl.EnableVertexAttribute(0)

	// Texcoord VBO (attribute 1, vec2)
	triTexVBO := rl.LoadVertexBuffer(triTexcoords, false)
	defer rl.UnloadVertexBuffer(triTexVBO)
	rl.SetVertexAttribute(1, 2, rl.Float, false, 0, 0)
	rl.EnableVertexAttribute(1)

	// Color VBO (attribute 3, vec4)
	triColVBO := rl.LoadVertexBuffer(triColors, false)
	defer rl.UnloadVertexBuffer(triColVBO)
	rl.SetVertexAttribute(3, 4, rl.Float, false, 0, 0)
	rl.EnableVertexAttribute(3)

	rl.DisableVertexArray()

	// ---------------------------------------------------------------
	// Shader + MVP setup
	// ---------------------------------------------------------------
	defaultShaderID := rl.GetShaderIdDefault()
	mvpLoc := rl.GetLocationUniform(defaultShaderID, "mvp")
	colDiffuseLoc := rl.GetLocationUniform(defaultShaderID, "colDiffuse")
	defaultTexID := rl.GetTextureIdDefault()

	fmt.Printf("DEBUG: shaderID=%d mvpLoc=%d colDiffuseLoc=%d texID=%d\n",
		defaultShaderID, mvpLoc, colDiffuseLoc, defaultTexID)
	fmt.Printf("DEBUG: quadVAO=%d quadPosVBO=%d quadColVBO=%d quadIBO=%d\n",
		quadVAO, quadPosVBO, quadColVBO, quadIBO)
	fmt.Printf("DEBUG: triVAO=%d triPosVBO=%d triColVBO=%d\n",
		triVAO, triPosVBO, triColVBO)

	// Orthographic projection: screen-space coordinates, origin top-left
  // NOTE: rl.GetCameraMatrix2D/rl.GetCameraMatrix could be used to work with rl camera
  // Need to test...
	mvpMatrix := rl.MatrixOrtho(0, screenWidth, screenHeight, 0, -1, 1)

	// White diffuse color so vertex colors pass through unmodified
	whiteDiffuse := []float32{1, 1, 1, 1}

	// Working buffer for animated quad colors
	animColors := make([]float32, len(quadColors))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// ---------------------------------------------------------------
		// Animated quad colors
		// Thanks Inigo: https://iquilezles.org/articles/palettes/
		// ---------------------------------------------------------------
		t := float32(rl.GetTime())
		for i := range 4 {
			phase := float32(i) * math.Pi / 2.0
			r := 0.5 + 0.5*float32(math.Sin(float64(t*2.0+phase)))
			g := 0.5 + 0.5*float32(math.Sin(float64(t*2.0+phase+math.Pi*2.0/3.0)))
			b := 0.5 + 0.5*float32(math.Sin(float64(t*2.0+phase+math.Pi*4.0/3.0)))
			animColors[i*4+0] = r
			animColors[i*4+1] = g
			animColors[i*4+2] = b
			animColors[i*4+3] = 1.0
		}
		rl.UpdateVertexBuffer(quadColVBO, animColors, 0)

		// Flush raylib's internal batch before custom GL draws
		rl.DrawRenderBatchActive()

		// Set up default shader for our custom VAO draws
		rl.EnableShader(defaultShaderID)
		rl.SetUniformMatrix(mvpLoc, mvpMatrix)
		rl.SetUniform(colDiffuseLoc, whiteDiffuse, int32(rl.ShaderUniformVec4), 1)
		rl.ActiveTextureSlot(0)
		rl.EnableTexture(defaultTexID)

		// Draw quad (indexed)
		rl.EnableVertexArray(quadVAO)
		rl.DrawVertexArrayElements(0, 6, nil)
		rl.DisableVertexArray()

		// Draw triangle (non-indexed)
		rl.EnableVertexArray(triVAO)
		rl.DrawVertexArray(0, 3)
		rl.DisableVertexArray()

		rl.CheckErrors()

		// Restore state for raylib's internal renderer
		rl.DisableShader()
		rl.DisableTexture()
		rl.DrawRenderBatchActive()

		// Text overlays
		rl.DrawText("rlgl Vertex Buffer Bindings Test", 10, 10, 20, rl.DarkGray)
		rl.DrawFPS(int32(rl.GetScreenWidth())-100, 10)

		rl.EndDrawing()
	}
}
