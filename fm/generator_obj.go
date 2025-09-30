package main

type position_t struct {
	x int
	y int
}

type entity_t struct {
	typeE  int
	symbol int
	typeI  int
	pos    position_t
	stats  stats_t
}

type map_t struct {
	playground   [MAP_HEIGHT][MAP_WIDTH]int
	player_spawn entity_t
	exit         entity_t
	items        [MAX_ITEMS_TOTAL]entity_t
	items_cnt    int
	enemies      [MAX_ENEMIES_TOTAL]entity_t
	enemies_cnt  int
}

type room_t struct {
	sector       int
	grid_i       int
	grid_j       int
	top_left     position_t
	bot_right    position_t
	doors        [4]position_t
	connections  [4]*room_t
	entities     [MAX_ENTITIES_PER_ROOM]entity_t
	entities_cnt int
	isStart      bool
}

type corridor_t struct {
	typeE      int
	points     [4]position_t
	points_cnt int
}

type dungeon_t struct {
	room_cnt      int
	corridors_cnt int
	rooms         [ROOMS_PER_SIDE + 2][ROOMS_PER_SIDE + 2]room_t
	sequence      [MAX_ROOMS_NUMBER]*room_t
	corridors     [MAX_CORRIDORS_NUMBER]corridor_t
	level         int
}

type inventory_t struct {
	weapon []int
	scroll []int
	potion []int
	food   []int
}

type stats_t struct {
	strength   int
	agility    int
	hp         int
	isChasing  bool
	isFirst    bool
	aggression int
	lastMove   int
	hpStealed  int
	extraSym   int
}
