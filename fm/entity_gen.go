package main

import "math/rand"

// import "strings"

// var enemy_pool[ENEMY_POOL_LEN] byte
// strings
// var str string
const str_enemy = "zOsmvg"

// enemy_pool := []byte(str)
// item_pool[ITEM_POOL_LEN] = "/?*$!";
const str_item = "/?*$"

func check_unoccupied(room *room_t, pos *position_t) int {
	status := UNOCCUPIED

	for i := 0; i < room.entities_cnt && status == UNOCCUPIED; i++ {
		if room.entities[i].pos.x == pos.x && room.entities[i].pos.y == pos.y {
			status = OCCUPIED
		}
	}

	return status
}

func generate_entity_coords(room *room_t, pos *position_t) {
	top_x := room.top_left.x
	top_y := room.top_left.y
	bot_x := room.bot_right.x
	bot_y := room.bot_right.y

	pos.x = (rand.Int()%(bot_x-top_x-1) + 1) + top_x
	pos.y = (rand.Int()%(bot_y-top_y-1) + 1) + top_y

	for check_unoccupied(room, pos) == OCCUPIED {
		// if check_unoccupied(room, pos) == OCCUPIED {
		// 	break
		// }
		pos.x = (rand.Int()%(bot_x-top_x-1) + 1) + top_x
		pos.y = (rand.Int()%(bot_y-top_y-1) + 1) + top_y
	}
	// while (check_unoccupied(room, pos) == OCCUPIED);
}

func generate_entities(dungeon *dungeon_t) {
	generate_player_pos(dungeon)
	generate_exit(dungeon)
	generate_enemies(dungeon)
	generate_items(dungeon)
}

func generate_player_pos(dungeon *dungeon_t) {
	room_index := rand.Int() % dungeon.room_cnt
	spawn_room := dungeon.sequence[room_index]
	var player_pos position_t

	generate_entity_coords(spawn_room, &player_pos)

	var player entity_t

	player.typeE = PLAYER
	player.symbol = PLAYER_CHAR
	player.pos = player_pos
	cur_room_entity_cnt := dungeon.sequence[room_index].entities_cnt

	dungeon.sequence[room_index].entities[cur_room_entity_cnt] = player
	dungeon.sequence[room_index].entities_cnt++
	dungeon.sequence[room_index].isStart = true
}

func generate_exit(dungeon *dungeon_t) {
	room_index := rand.Int() % dungeon.room_cnt
	exit_room := dungeon.sequence[room_index]
	for exit_room.isStart {
		room_index = rand.Int() % dungeon.room_cnt
		exit_room = dungeon.sequence[room_index]
	}
	var exit_pos position_t

	generate_entity_coords(exit_room, &exit_pos)

	var exit entity_t
	exit.typeE = EXIT
	exit.symbol = EXIT_CHAR
	exit.pos = exit_pos
	cur_room_entity_cnt := dungeon.sequence[room_index].entities_cnt

	dungeon.sequence[room_index].entities[cur_room_entity_cnt] = exit
	dungeon.sequence[room_index].entities_cnt++
}

func generate_enemies(dungeon *dungeon_t) {
	for i := 0; i < dungeon.room_cnt; i++ {
		enemies_cnt := rand.Int() % MAX_ENEMIES_PER_ROOM

		for j := 0; j < enemies_cnt; j++ {
			var enemy entity_t

			if dungeon.sequence[i].isStart {
				break
			}
			generate_enemy(dungeon.sequence[i], &enemy)
			generateEnemyType(dungeon, &enemy)
			cur_room_entity_cnt := dungeon.sequence[i].entities_cnt

			dungeon.sequence[i].entities[cur_room_entity_cnt] = enemy
			dungeon.sequence[i].entities_cnt++
		}
	}
}

func generate_enemy(room *room_t, enemy *entity_t) {
	enemy.typeE = ENEMY
	enemy.symbol = int(str_enemy[rand.Int()%ENEMY_POOL_LEN])
	generate_entity_coords(room, &enemy.pos)

}

func generateEnemyType(dungeon *dungeon_t, enemy *entity_t) {
	switch enemy.symbol {
	case 'z':
		enemy.stats.strength = 2 + 2*dungeon.level
		enemy.stats.agility = 1 + 1*dungeon.level
		enemy.stats.hp = 3 + 3*dungeon.level
		enemy.stats.aggression = 4
	case 'O':
		enemy.stats.strength = 4 + 4*dungeon.level
		enemy.stats.agility = 1 + 1*dungeon.level
		enemy.stats.hp = 4 + 4*dungeon.level
		enemy.stats.aggression = 4
	case 's':
		enemy.stats.strength = 1 + 1*dungeon.level
		enemy.stats.agility = 4 + 4*dungeon.level
		enemy.stats.hp = 2 + 2*dungeon.level
		enemy.stats.aggression = 5
	case 'v':
		enemy.stats.strength = 2 + 2*dungeon.level
		enemy.stats.agility = 3 + 3*dungeon.level
		enemy.stats.hp = 3 + 3*dungeon.level
		enemy.stats.aggression = 5
	case 'm':
		enemy.stats.strength = 1 + 1*dungeon.level
		enemy.stats.agility = 3 + 3*dungeon.level
		enemy.stats.hp = 3 + 3*dungeon.level
		enemy.stats.aggression = 2
	case 'g':
		enemy.stats.strength = 1 + 1*dungeon.level
		enemy.stats.agility = 3 + 3*dungeon.level
		enemy.stats.hp = 1 + 1*dungeon.level
		enemy.stats.aggression = 3

	}

}

func generate_items(dungeon *dungeon_t) {
	for i := 0; i < dungeon.room_cnt; i++ {
		items_cnt := rand.Int() % MAX_ITEMS_PER_ROOM

		for j := 0; j < items_cnt; j++ {
			var item entity_t

			generate_item(dungeon.sequence[i], &item)

			cur_room_entity_cnt := dungeon.sequence[i].entities_cnt

			dungeon.sequence[i].entities[cur_room_entity_cnt] = item
			dungeon.sequence[i].entities_cnt++
		}
	}
}

func generate_item(room *room_t, item *entity_t) {
	item.typeE = ITEM
	item.symbol = int(str_item[rand.Int()%ITEM_POOL_LEN])
	generate_entity_coords(room, &item.pos)
	generateItemType(item)
}

func generateItemType(item *entity_t) {
	switch item.symbol {
	case '?':
		item.typeI = rand.Int()%3 + 10
	case '*':
		item.typeI = rand.Int()%3 + 20
	case '/':
		item.typeI = rand.Int()%3 + 30
	case '$':
		item.typeI = rand.Int()%3 + 40
	}
}
