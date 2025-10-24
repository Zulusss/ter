package main

import (
	"container/list"
	// "log"

	"github.com/gbin/goncurses"
)

func main() {
	var dungeon dungeon_t
	var field map_t
	var screen screen_t
	var player player_t
	var key int
	var statStr string
	var msgStr list.List
	var difficulty int
	msgStr.Init()
	msgStr.PushFront("Welcome to rogue")
	msgStr.PushFront("w a s d or arrows to control, e h j k to open inventory, z to load last saved game")

	goncurses.Init()
	goncurses.StdScr().Keypad(true)
	goncurses.Echo(false)
	goncurses.Cursor(0)
	goncurses.StartColor()
	goncurses.UseDefaultColors()
	goncurses.InitPair(goncurses.C_RED, goncurses.C_RED, goncurses.C_BLACK)
	goncurses.InitPair(goncurses.C_GREEN, goncurses.C_GREEN, goncurses.C_BLACK)
	goncurses.InitPair(goncurses.C_YELLOW, goncurses.C_YELLOW, goncurses.C_BLACK)

	generate_dungeon(&dungeon)
	print_dungeon_generated(MAP_HEIGHT + 5)
	generate_entities(&dungeon, difficulty)
	print_entities_generated(MAP_HEIGHT + 6)
	dungeon_to_map(&dungeon, &field)
	init_player(&player, &field.player_spawn)
	// difficulty.hp = player.health

	initInventory(&player)

	print_map_created(MAP_HEIGHT + 7)
	check_result := check_connectivity(&dungeon)
	print_connectivity_check_result(check_result, MAP_HEIGHT+9)
	print_entity_pools(MAP_HEIGHT + 11)
	print_map(&field)

	key = int(goncurses.StdScr().GetChar())
	for key != ESC {

		player_movement(&dungeon, &field, &player, key, &msgStr)

		checkPotion(&player)
		enemyTurn(&dungeon, &field, &player, &msgStr)
		render_screen(&field, &player, &screen, &dungeon)
		createStatString(&dungeon, &player, &statStr)

		screen_print(&screen, &statStr, &msgStr)
		dead := IsDead(&player)
		if dead {
			TheEnd(&dungeon, &player)
			printRecords()
			for key != ESC {
				key = int(goncurses.StdScr().GetChar())
			}
		} else {
			key = int(goncurses.StdScr().GetChar())
			if key == 'z' {
				savedDungeon, savedField, savedPlayer, err := LoadGame()
				if err != nil {
					msgStr.PushFront("No saved games.")
				} else {
					dungeon = *savedDungeon
					field = *savedField
					player = *savedPlayer
					msgStr.PushFront("Game loaded.")
					print_map(&field)
					key = int(goncurses.StdScr().GetChar())
				}
			}
		}
	}
	goncurses.End()
}
