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
	msgStr.Init()
	msgStr.PushFront("Welcome to rogue")
	msgStr.PushFront("w a s d or arrows to control")

	goncurses.Init()
	goncurses.StdScr().Keypad(true)
	goncurses.Echo(false)
	goncurses.Cursor(0)

	generate_dungeon(&dungeon)
	// generate_player_pos(&dungeon)
	print_dungeon_generated(MAP_HEIGHT + 5)
	generate_entities(&dungeon)
	print_entities_generated(MAP_HEIGHT + 6)
	dungeon_to_map(&dungeon, &field)
	init_player(&player, &field.player_spawn)
	initInventory(&player)
	print_map_created(MAP_HEIGHT + 7)
	check_result := check_connectivity(&dungeon)
	print_connectivity_check_result(check_result, MAP_HEIGHT+9)
	print_entity_pools(MAP_HEIGHT + 11)
	print_map(&field)

	key = int(goncurses.StdScr().GetChar())
	// goncurses.StdScr().GetChar()
	for key != ESC {
		// goncurses.StdScr().MovePrint(0, 0, "DUNGEON GENERATED...")

		player_movement(&dungeon, &field, &player, key, &msgStr)
		// checkTile(&dungeon, &field, &player)

		checkPotion(&player)
		enemyTurn(&dungeon, &field, &player)
		render_screen(&field, &player, &screen, &dungeon)
		createStatString(&dungeon, &player, &statStr)

		// print_map(&field)
		screen_print(&screen, &statStr, &msgStr)
		key = int(goncurses.StdScr().GetChar())
		if key == 'z' {
			// var savedDungeon dungeon_t
			// var savedField map_t
			// var savedPlayer player_t
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
	goncurses.End()
}
