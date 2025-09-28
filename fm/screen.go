package main

import (
	"fmt"
	// "strings"

	"github.com/gbin/goncurses"
)

func render_screen(field *map_t, player *player_t, screen *screen_t) {
	// refresh_screen(map, player, screen);
	// clear_screen(screen)
	goncurses.StdScr().Clear()
	draw_minimap(field, player, screen)
}

// func clear_screen(screen *screen_t) {
// 	for i := 0; i < SCRN_H; i++ {
// 		for j := 0; j < SCRN_W; j++ {
// 			screen.screen[i][j] = SPACE_CHAR
// 		}
// 	}
// }

func draw_minimap(field *map_t, player *player_t, screen *screen_t) {
	for i := 0; i < MAP_HEIGHT; i++ {
		for j := 0; j < MAP_WIDTH; j++ {
			// var cur_char int
			cur_char := int(field.playground[i][j])
			// place_char := SPACE_CHAR

			// if cur_char == SPACE_CHAR || cur_char == OUTER_AREA_CHAR {
			// 	place_char = SPACE_CHAR
			// } else if cur_char == WALL_CHAR {
			// 	place_char = WALL_CHAR
			// } else if cur_char == CORRIDOR_CHAR {
			// 	place_char = CORRIDOR_CHAR
			// }

			screen.screen[i][j] = cur_char
		}
	}

	player_x := player.pos.x
	player_y := player.pos.y

	screen.screen[player_y][player_x] = MINIMAP_PLAYER_CHAR
}

func createStatString(dungeon *dungeon_t, player *player_t, statStr *string) {
	*statStr = fmt.Sprintf("Level: %d Hp: %d/%d Strength %d Agility %d Gold: %d", dungeon.level, player.health, player.maxHealth, player.strength, player.agility, player.gold)
}

func screen_print(screen *screen_t, statStr *string) {
	for i := 0; i < SCRN_H; i++ {
		for j := 0; j < SCRN_W; j++ {
			if i < MAP_HEIGHT && j < MAP_WIDTH {
				goncurses.StdScr().MoveAddChar(i, j, goncurses.Char(screen.screen[i][j]))
			} else if i == MAP_HEIGHT {
				goncurses.StdScr().MovePrint(i, 0, *statStr)

			} else {
				goncurses.StdScr().MoveAddChar(i, j, goncurses.Char(SPACE_CHAR))
				// goncurses.StdScr().MovePrint(i, j, '©')
			}
		}
	}
}
