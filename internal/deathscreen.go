package internal

import (
	"alicevszombies/internal/util"

	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func onPlayerDeath(world *World) {
	world.paused = true
	world.uistate.isDeathScreen = true

	// Delete all velocities to only animate player death effect
	for id := range world.velocity {
		world.velocity[id] = rl.Vector2Zero()
	}
	newDeathEffect(world, "player", world.position[world.player])
	delete(world.texture, world.player)
}

func updateDeathScreen(world *World) {
	updateDeathEffects(world)
	updateVelocity(world)

	if rl.IsKeyPressed(rl.KeyEscape) {
		// Goes back to main menu, as on game start
		world.Reset()
	}
}

func renderDeathScreen(world *World) {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.ColorAlpha(rl.Black, 0.7))
	pos := util.HalfScreenSize()
	util.DrawTextCenteredSpaced("You Died!", 256, pos, 16)
	pos.Y += 128
	util.DrawTextCenteredSpaced("Reached Wave "+fmt.Sprint(world.enemySpawner.wave), 64, pos, 4)
	pos.Y += 256
	util.DrawTextCenteredSpaced("ESC = Main Menu", 64, pos, 4)
}
