package main

import "github.com/gbin/goncurses"

func main() {
	var dungeon dungeon_t
	var field map_t
	var screen screen_t
	var player player_t
	var key int
	var statStr string

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
		player_movement(&field, &player, key)
		checkTile(&dungeon, &field, &player)
		checkPotion(&player)
		enemyTurn(&dungeon, &field, &player)
		render_screen(&field, &player, &screen)
		createStatString(&dungeon, &player, &statStr)

		// print_map(&field)
		screen_print(&screen, &statStr)
		key = int(goncurses.StdScr().GetChar())
	}
	goncurses.End()
}
