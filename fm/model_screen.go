package main

import (
	"fmt"
)

func render_screen(field *map_t, player *player_t, screen *screen_t, dungeon *dungeon_t) {

	clear_screen()
	draw_minimap(field, player, screen, dungeon)
}

func inRoom(dungeon *dungeon_t, player *player_t) (bool, position_t, position_t) {
	for i := range dungeon.sequence {
		if dungeon.sequence[i] != nil {
			if player.pos.y < dungeon.sequence[i].bot_right.y &&
				player.pos.x > dungeon.sequence[i].top_left.x &&
				player.pos.y > dungeon.sequence[i].top_left.y &&
				player.pos.x < dungeon.sequence[i].bot_right.x {
				return true, dungeon.sequence[i].top_left, dungeon.sequence[i].bot_right
			}
		}
	}
	var fakePos, fakePos2 position_t
	return false, fakePos, fakePos2
}

func draw_minimap(field *map_t, player *player_t, screen *screen_t, dungeon *dungeon_t) {
	OK, topLeft, botRight := inRoom(dungeon, player)
	viewed := ComputeFOV(field, player.pos, 6)
	for i := 0; i < MAP_HEIGHT; i++ {
		for j := 0; j < MAP_WIDTH; j++ {
			var cur_char int
			var pos position_t
			pos.y, pos.x = i, j
			if viewed[pos] {
				cur_char = int(field.playground[i][j])
			} else if field.visited[pos] {
				cur_char = '#'
			} else {
				cur_char = ' '
			}
			screen.screen[i][j] = cur_char
		}
	}

	player_x := player.pos.x
	player_y := player.pos.y

	screen.screen[player_y][player_x] = MINIMAP_PLAYER_CHAR

	if OK {
		for i := topLeft.y; i <= botRight.y; i++ {
			for j := topLeft.x; j <= botRight.x; j++ {
				if field.playground[i][j] == WALL_CHAR {
					var pos position_t
					pos.y, pos.x = i, j
					field.visited[pos] = true
				}
				screen.screen[i][j] = field.playground[i][j]
			}
		}
	}
}

func createStatString(dungeon *dungeon_t, player *player_t, statStr *string) {
	*statStr = fmt.Sprintf("Level %d Hp %d/%d Strength %d Agility %d Gold %d", dungeon.level, player.health, player.maxHealth, player.strength, player.agility, player.gold)
}
