package main

import "github.com/gbin/goncurses"

func print_map(field *map_t) {
	for i := 0; i < MAP_HEIGHT; i++ {
		for j := 0; j < MAP_WIDTH; j++ {
			goncurses.StdScr().MoveAddChar(i, j, goncurses.Char(field.playground[i][j]))
		}
	}
}

func print_dungeon_generated(y int) {
	x := 0

	goncurses.StdScr().MovePrint(y, x, "DUNGEON GENERATED...")
}

func print_entities_generated(y int) {
	x := 0

	goncurses.StdScr().MovePrint(y, x, "ENTITIES GENERATED...")
}

func print_map_created(y int) {
	x := 0

	goncurses.StdScr().MovePrint(y, x, "MAP CREATED FROM GENERATED DATA...")
}

func print_connectivity_check_result(check_result int, y int) {
	x := 0

	if check_result == CONNECTED {
		goncurses.StdScr().MovePrint(y, x, "ROOM GRAPH IS CONNECTED!")
	} else {
		goncurses.StdScr().MovePrint(y, x, "ROOM GRAPH IS NOT CONNECTED!")
	}
}

func print_entity_pools(y int) {
	x := 0

	goncurses.StdScr().MovePrint(y, x, "\"ABCDEFGHIJKLMNOPQRSTUVWXYZ\" : Enemy pool.")
	goncurses.StdScr().MovePrint(y+1, x, "\"/?*$!\"                      : Item pool.")
	goncurses.StdScr().MovePrint(y+2, x, "\"@\" \"|\"                      : Player and exit.")
}
