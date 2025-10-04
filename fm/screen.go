package main

import (
	"container/list"
	"fmt"

	// "strings"

	"github.com/gbin/goncurses"
)

func render_screen(field *map_t, player *player_t, screen *screen_t, dungeon *dungeon_t) {
	// refresh_screen(map, player, screen);
	// clear_screen(screen)
	goncurses.StdScr().Clear()
	draw_minimap(field, player, screen, dungeon)
}

//	func clear_screen(screen *screen_t) {
//		for i := 0; i < SCRN_H; i++ {
//			for j := 0; j < SCRN_W; j++ {
//				screen.screen[i][j] = SPACE_CHAR
//			}
//		}
//	}
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
			// cur_char := int(field.playground[i][j])
			var pos position_t
			pos.y, pos.x = i, j
			if viewed[pos] == true {
				cur_char = int(field.playground[i][j])
			} else if field.visited[pos] {
				cur_char = '#'
			} else {
				cur_char = ' '
			}
			// if OK {
			// 	if pos.y > topLeft.y && pos.y < botRight.y && pos.x > topLeft.x &&
			// 		pos.x < botRight.x {
			// 		cur_char = int(field.playground[i][j])
			// 	}
			// }
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
	*statStr = fmt.Sprintf("Level: %d Hp: %d/%d Strength %d Agility %d Gold: %d", dungeon.level, player.health, player.maxHealth, player.strength, player.agility, player.gold)
}

func screen_print(screen *screen_t, statStr *string, msgStr *list.List) {
	for i := 0; i < SCRN_H; i++ {
		for j := 0; j < SCRN_W; j++ {
			if i < MAP_HEIGHT && j < MAP_WIDTH {
				goncurses.StdScr().MoveAddChar(i, j, goncurses.Char(screen.screen[i][j]))
			} else if i == MAP_HEIGHT {
				goncurses.StdScr().MovePrint(i, 0, *statStr)

			} else if i == MAP_HEIGHT+3 {
				goncurses.StdScr().MovePrint(i, 0, msgStr.Front().Value)
			} else if i == MAP_HEIGHT+2 {
				goncurses.StdScr().MovePrint(i, 0, msgStr.Front().Next().Value)
			} else {
				goncurses.StdScr().MoveAddChar(i, j, goncurses.Char(SPACE_CHAR))
				// goncurses.StdScr().MovePrint(i, j, '©')
			}
		}
	}
}
