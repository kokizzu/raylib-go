package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(800)
	screenHeight := int32(450)

	raylib.InitWindow(screenWidth, screenHeight, "raylib [text] example - rBMF fonts")

	fonts := make([]raylib.SpriteFont, 8)
	fonts[0] = raylib.LoadSpriteFont("fonts/alagard.rbmf")
	fonts[1] = raylib.LoadSpriteFont("fonts/pixelplay.rbmf")
	fonts[2] = raylib.LoadSpriteFont("fonts/mecha.rbmf")
	fonts[3] = raylib.LoadSpriteFont("fonts/setback.rbmf")
	fonts[4] = raylib.LoadSpriteFont("fonts/romulus.rbmf")
	fonts[5] = raylib.LoadSpriteFont("fonts/pixantiqua.rbmf")
	fonts[6] = raylib.LoadSpriteFont("fonts/alpha_beta.rbmf")
	fonts[7] = raylib.LoadSpriteFont("fonts/jupiter_crash.rbmf")

	messages := []string{
		"ALAGARD FONT designed by Hewett Tsoi",
		"PIXELPLAY FONT designed by Aleksander Shevchuk",
		"MECHA FONT designed by Captain Falcon",
		"SETBACK FONT designed by Brian Kent (AEnigma)",
		"ROMULUS FONT designed by Hewett Tsoi",
		"PIXANTIQUA FONT designed by Gerhard Grossmann",
		"ALPHA_BETA FONT designed by Brian Kent (AEnigma)",
		"JUPITER_CRASH FONT designed by Brian Kent (AEnigma)",
	}

	spacings := []int32{2, 4, 8, 4, 3, 4, 4, 1}
	positions := make([]raylib.Vector2, 8)

	var i int32
	for i = 0; i < 8; i++ {
		x := screenWidth/2 - int32(raylib.MeasureTextEx(fonts[i], messages[i], float32(fonts[i].Size*2), spacings[i]).X/2)
		y := 60 + fonts[i].Size + 45*i
		positions[i] = raylib.NewVector2(float32(x), float32(y))
	}

	colors := []raylib.Color{raylib.Maroon, raylib.Orange, raylib.DarkGreen, raylib.DarkBlue, raylib.DarkPurple, raylib.Lime, raylib.Gold, raylib.DarkBrown}

	raylib.SetTargetFPS(60)

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RayWhite)
		raylib.DrawText("free fonts included with raylib", 250, 20, 20, raylib.DarkGray)
		raylib.DrawLine(220, 50, 590, 50, raylib.DarkGray)

		for i = 0; i < 8; i++ {
			raylib.DrawTextEx(fonts[i], messages[i], positions[i], float32(fonts[i].Size*2), spacings[i], colors[i])
		}

		raylib.EndDrawing()
	}

	for i = 0; i < 8; i++ {
		raylib.UnloadSpriteFont(fonts[i])
	}

	raylib.CloseWindow()
}
