package main

func dungeon_to_map(dungeon *dungeon_t, field *map_t) {
	init_map(field)
	rooms_to_map(dungeon, field)
	corridors_to_map(dungeon, field)
}

func init_map(field *map_t) {
	for i := 0; i < MAP_HEIGHT; i++ {
		for j := 0; j < MAP_WIDTH; j++ {
			field.playground[i][j] = OUTER_AREA_CHAR
		}
	}

	field.enemies_cnt = 0
	field.items_cnt = 0
}

func rooms_to_map(dungeon *dungeon_t, field *map_t) {
	for i := 0; i < dungeon.room_cnt; i++ {
		top_room_corner := dungeon.sequence[i].top_left
		bot_room_corner := dungeon.sequence[i].bot_right

		rectangle_to_map(&top_room_corner, &bot_room_corner, field)
		fill_rectangle(&top_room_corner, &bot_room_corner, field)

		for j := 0; j < dungeon.sequence[i].entities_cnt; j++ {
			cur_entity := dungeon.sequence[i].entities[j]

			switch cur_entity.typeE {
			case PLAYER:
				field.player_spawn = cur_entity
			case EXIT:
				field.exit = cur_entity
			case ENEMY:
				field.enemies[field.enemies_cnt] = cur_entity
				field.enemies_cnt++

			case ITEM:
				field.items[field.items_cnt] = cur_entity
				field.items_cnt++
			}

			field.playground[cur_entity.pos.y][cur_entity.pos.x] = int(cur_entity.symbol)
			// if field.enemies_cnt == 44 || field.items_cnt == 26 {
			// 	break
			// }
		}
	}
}

func rectangle_to_map(top *position_t, bot *position_t, field *map_t) {
	field.playground[top.y][top.x] = WALL_CHAR

	i := top.x + 1

	for ; i < bot.x; i++ {
		field.playground[top.y][i] = WALL_CHAR
	}
	field.playground[top.y][i] = WALL_CHAR

	for i := top.y + 1; i < bot.y; i++ {
		field.playground[i][top.x] = WALL_CHAR
		field.playground[i][bot.x] = WALL_CHAR
	}

	field.playground[bot.y][top.x] = WALL_CHAR
	i = top.x + 1
	for ; i < bot.x; i++ {
		field.playground[bot.y][i] = WALL_CHAR
	}
	field.playground[bot.y][i] = WALL_CHAR
}

func fill_rectangle(top *position_t, bot *position_t, field *map_t) {
	for i := top.y + 1; i < bot.y; i++ {
		for j := top.x + 1; j < bot.x; j++ {
			field.playground[i][j] = INNER_AREA_CHAR
		}
	}
}

func corridors_to_map(dungeon *dungeon_t, field *map_t) {
	for i := 0; i < dungeon.corridors_cnt; i++ {
		switch dungeon.corridors[i].typeE {
		case LEFT_TO_RIGHT_CORRIDOR:
			draw_horisontal_line(&dungeon.corridors[i].points[0], &dungeon.corridors[i].points[1], field)
			draw_vertical_line(&dungeon.corridors[i].points[1], &dungeon.corridors[i].points[2], field)
			draw_horisontal_line(&dungeon.corridors[i].points[2], &dungeon.corridors[i].points[3], field)
		case LEFT_TURN_CORRIDOR:
			draw_vertical_line(&dungeon.corridors[i].points[0], &dungeon.corridors[i].points[1], field)
			draw_horisontal_line(&dungeon.corridors[i].points[1], &dungeon.corridors[i].points[2], field)
		case RIGHT_TURN_CORRIDOR:
			draw_vertical_line(&dungeon.corridors[i].points[0], &dungeon.corridors[i].points[1], field)
			draw_horisontal_line(&dungeon.corridors[i].points[1], &dungeon.corridors[i].points[2], field)
		case TOP_TO_BOTTOM_CORRIDOR:
			draw_vertical_line(&dungeon.corridors[i].points[0], &dungeon.corridors[i].points[1], field)
			draw_horisontal_line(&dungeon.corridors[i].points[1], &dungeon.corridors[i].points[2], field)
			draw_vertical_line(&dungeon.corridors[i].points[2], &dungeon.corridors[i].points[3], field)
		}
	}
}

func draw_horisontal_line(first_dot *position_t, second_dot *position_t, field *map_t) {
	y := first_dot.y

	for x := min(first_dot.x, second_dot.x); x <= max(first_dot.x, second_dot.x); x++ {
		field.playground[y][x] = CORRIDOR_CHAR
	}
}

func draw_vertical_line(first_dot *position_t, second_dot *position_t, field *map_t) { //field = map
	x := first_dot.x

	for y := min(first_dot.y, second_dot.y); y <= max(first_dot.y, second_dot.y); y++ {
		field.playground[y][x] = CORRIDOR_CHAR
	}
}
