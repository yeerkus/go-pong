package main

import (
	"fmt"
	"image/color"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WinWidth  = 1200
	WinHeight = 800
)

// Rectangle structure to be used to make the two pads and the ball
type Rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
	Colour color.RGBA
}

func main() {
	rl.InitWindow(WinWidth, WinHeight, "Pong by Bertrand \"Berry\" KARAKE")
	rl.SetTargetFPS(60)

	// Initialize the speed of the pads, the two-dimensional speed of the ball and the score of both players
	var speed float32 = 20
	var ballSpeedX float32 = 12
	var ballSpeedY float32 = -1
	var scoreOne int = 0
	var scoreTwo int = 0
	// var intro int = 0

	// Initialize the ball and the pads. All depend heavily on the Width of the Screen
	Ball := Rectangle{X: WinWidth / 2, Y: WinHeight / 2, Width: 20, Height: 20, Colour: rl.White}
	PlayerOne := Rectangle{X: WinWidth / 10, Y: 200, Width: 20, Height: 200, Colour: rl.White}
	PlayerTwo := Rectangle{X: WinWidth - PlayerOne.X - PlayerOne.Width, Y: 200, Width: 20, Height: 200, Colour: rl.White}

	// Start of the game loop
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawText("Pong by Berry", WinWidth/3, WinHeight/7, 50, rl.White)

		// Check if the ball is off-screen horizontally
		if Ball.X > WinWidth-Ball.Width || Ball.X < 0 {

			// Increase the score of the scoring player
			if Ball.X < 0 {
				scoreTwo++
			} else {
				scoreOne++
			}
			// Initialize the ball in the center with random Y speed
			Ball.X = WinWidth / 2
			Ball.Y = float32(rand.Intn(WinHeight)) - 1
			ballSpeedX = speed / 3
			ballSpeedY = speed / 3
			// Check for collision with the pads then increase the ball's speed in the opposite direction
		} else if (Ball.X < PlayerOne.X && Ball.X > PlayerOne.X-PlayerOne.Width) && (Ball.Y > PlayerOne.Y && Ball.Y < PlayerOne.Y+PlayerOne.Height) {
			// make sure the ball is not too fast
			if ballSpeedX > -15 {
				ballSpeedX *= -1.1
			} else {
				ballSpeedX *= -1
			}

		} else if (Ball.X > PlayerTwo.X && Ball.X < PlayerTwo.X+PlayerTwo.Width) && (Ball.Y > PlayerTwo.Y && Ball.Y < PlayerTwo.Y+PlayerTwo.Height) {
			//^^^^^^^
			if ballSpeedY < 15 {
				ballSpeedX *= -1.1
			} else {
				ballSpeedX *= -1
			}
		}

		// Check if the ball is off-screen vertically, then reverse its Y speed
		if Ball.Y > WinHeight-Ball.Height || Ball.Y < 0 {
			ballSpeedY *= -1
		}

		// Move the ball on screen according to the value of the speed components
		Ball.X += ballSpeedX
		Ball.Y += ballSpeedY

		// Pad movement logic -- make sure they don't go offscreen
		if rl.IsKeyDown(87) && PlayerOne.Y > 0 {
			PlayerOne.Y -= speed
		} else if rl.IsKeyDown(83) && PlayerOne.Y+PlayerOne.Height < WinHeight {
			PlayerOne.Y += speed
		}
		if rl.IsKeyDown(265) && PlayerTwo.Y > 0 {
			PlayerTwo.Y -= speed
		} else if rl.IsKeyDown(264) && PlayerTwo.Y+PlayerTwo.Height < WinHeight {
			PlayerTwo.Y += speed
		}

		// Draw everthing and the changes on the screen
		rl.DrawText(fmt.Sprint(scoreOne), WinWidth/6, WinHeight/5, 100, rl.White)
		rl.DrawText(fmt.Sprint(scoreTwo), 5*WinWidth/6, WinHeight/5, 100, rl.White)

		rl.DrawRectangle(int32(PlayerOne.X), int32(PlayerOne.Y), int32(PlayerOne.Width), int32(PlayerOne.Height), PlayerOne.Colour)
		rl.DrawRectangle(int32(PlayerTwo.X), int32(PlayerTwo.Y), int32(PlayerTwo.Width), int32(PlayerTwo.Height), PlayerTwo.Colour)
		rl.DrawRectangle(int32(Ball.X), int32(Ball.Y), int32(Ball.Width), int32(Ball.Height), rl.White)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
