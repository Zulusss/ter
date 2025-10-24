package main

import (
	"math"
	"math/rand"
	"sort"
)

func generate_dungeon(dungeon *dungeon_t) {
	init_dungeon(dungeon)
	generate_sectors(dungeon)
	generate_connections(dungeon)
	generate_rooms_geometry(dungeon)
	generate_corridors_geometry(dungeon)
	dungeon.level++
}

func init_dungeon(dungeon *dungeon_t) {

	for i := 0; i < ROOMS_PER_SIDE+2; i++ {
		for j := 0; j < ROOMS_PER_SIDE+2; j++ {
			dungeon.rooms[i][j].sector = UNINITIALIZED

			for k := 0; k < 4; k++ {
				dungeon.rooms[i][j].connections[k] = nil
				dungeon.rooms[i][j].doors[k].x = UNINITIALIZED
				dungeon.rooms[i][j].doors[k].y = UNINITIALIZED
			}
		}
	}
	for i := 0; i < MAX_ROOMS_NUMBER; i++ {
		dungeon.sequence[i] = nil
	}
	for i := 0; i < MAX_CORRIDORS_NUMBER; i++ {
		dungeon.corridors[i].typeE = UNINITIALIZED
	}
}

func generate_sectors(dungeon *dungeon_t) {
	for dungeon.room_cnt < 9 {
		sector := 0

		for i := 1; i < ROOMS_PER_SIDE+1; i++ {
			for j := 1; j < ROOMS_PER_SIDE+1; sector++ {
				if float64(rand.Int())/math.MaxInt64 < ROOM_CHANCE &&
					dungeon.rooms[i][j].sector == UNINITIALIZED {
					dungeon.rooms[i][j].sector = sector
					dungeon.rooms[i][j].grid_i = i
					dungeon.rooms[i][j].grid_j = j
					dungeon.sequence[dungeon.room_cnt] = &dungeon.rooms[i][j]
					dungeon.room_cnt++
				}
				j++
			}
		}
	}
	slc := make([]*room_t, 9)
	var count int
	for i := range slc {
		if dungeon.sequence[i] != nil {
			slc[count] = dungeon.sequence[i]
			count++
		}
	}
	slcNoNil := slc[:count]
	sort.Slice(slcNoNil, func(i, j int) bool /*room_placement_comparator*/ {
		return slcNoNil[i].sector < slcNoNil[j].sector
	})

	for i := range slc {
		if slc[i] != nil {
			dungeon.sequence[i] = slcNoNil[i]
		} else {
			dungeon.sequence[i] = nil
		}
	}
}

func generate_connections(dungeon *dungeon_t) {
	generate_primary_connections(dungeon)
	generate_secondary_connections(dungeon)
	disconnectRooms(dungeon)
}

func generate_primary_connections(dungeon *dungeon_t) {
	for i := 1; i < ROOMS_PER_SIDE+1; i++ {
		for j := 1; j < ROOMS_PER_SIDE+1; j++ {
			if dungeon.rooms[i][j].sector != UNINITIALIZED {
				if dungeon.rooms[i-1][j].sector != UNINITIALIZED {
					dungeon.rooms[i][j].connections[0] = &dungeon.rooms[i-1][j]

				}
				if dungeon.rooms[i][j+1].sector != UNINITIALIZED {
					dungeon.rooms[i][j].connections[1] = &dungeon.rooms[i][j+1]

				}
				if dungeon.rooms[i+1][j].sector != UNINITIALIZED {
					dungeon.rooms[i][j].connections[2] = &dungeon.rooms[i+1][j]

				}
				if dungeon.rooms[i][j-1].sector != UNINITIALIZED {
					dungeon.rooms[i][j].connections[3] = &dungeon.rooms[i][j-1]

				}
			}
		}
	}
}

func disconnectRooms(dungeon *dungeon_t) {
	y := (rand.Int() % 3) + 1
	x := (rand.Int() % 3) + 1
	room := &dungeon.rooms[y][x]
	roomConns := make([]*room_t, 0)
	for i := 0; i < 4; i++ {
		if room.connections[i] != nil {
			roomConns = append(roomConns, room.connections[i])
		}
	}
	if len(roomConns) == 2 {
		index := rand.Int() % 2
		temp := roomConns[index]
		isolateRoom(temp, room)

	} else if len(roomConns) == 3 {
		index := rand.Int() % 3
		temp := roomConns[index]
		isolateRoom(temp, room)

		index2 := rand.Int() % 3
		for index2 == index {
			index2 = rand.Int() % 3
		}

		temp = roomConns[index2]
		isolateRoom(temp, room)

	} else if len(roomConns) == 4 {
		index := rand.Int() % 4
		temp := roomConns[index]
		isolateRoom(temp, room)

		index2 := rand.Int() % 4
		for index2 == index {
			index2 = rand.Int() % 4
		}

		temp = roomConns[index2]
		isolateRoom(temp, room)

		index3 := rand.Int() % 4
		for index3 == index || index3 == index2 {
			index3 = rand.Int() % 4
		}
		temp = roomConns[index3]
		isolateRoom(temp, room)
	}
}

func isolateRoom(temp *room_t, room *room_t) {
	if temp.connections[0] != nil {
		if temp.connections[0] == room {
			temp.connections[0] = nil
			room.connections[2] = nil
		}
	}
	if temp.connections[1] != nil {
		if temp.connections[1] == room {
			temp.connections[1] = nil
			room.connections[3] = nil
		}
	}
	if temp.connections[2] != nil {
		if temp.connections[2] == room {
			temp.connections[2] = nil
			room.connections[0] = nil
		}
	}
	if temp.connections[3] != nil {
		if temp.connections[3] == room {
			temp.connections[3] = nil
			room.connections[1] = nil
		}
	}
}

func generate_secondary_connections(dungeon *dungeon_t) {
	for i := 0; i < dungeon.room_cnt-1; i++ {
		cur := dungeon.sequence[i]
		next := dungeon.sequence[i+1]

		if cur.grid_i == next.grid_i && next.connections[LEFT] == nil {
			cur.connections[RIGHT] = next
			next.connections[LEFT] = cur
		} else if cur.grid_i-next.grid_i == -1 && cur.connections[BOTTOM] == nil {
			if cur.grid_j < next.grid_j && next.connections[LEFT] == nil {
				cur.connections[BOTTOM] = next
				next.connections[LEFT] = cur
			} else if cur.grid_j > next.grid_j && next.connections[RIGHT] == nil {
				cur.connections[BOTTOM] = next
				next.connections[RIGHT] = cur
			} else if cur.grid_j > next.grid_j && cur.connections[BOTTOM] == nil &&
				i < dungeon.room_cnt-2 {
				cur.connections[BOTTOM] = dungeon.sequence[i+2]
				dungeon.sequence[i+2].connections[RIGHT] = cur
			}
		} else if cur.grid_i-next.grid_i == -2 && next.connections[TOP] == nil {
			cur.connections[BOTTOM] = next
			next.connections[TOP] = cur
		}
	}
}

func generate_rooms_geometry(dungeon *dungeon_t) {
	for i := 1; i < ROOMS_PER_SIDE+1; i++ {
		for j := 1; j < ROOMS_PER_SIDE+1; j++ {
			if dungeon.rooms[i][j].sector != UNINITIALIZED {
				generate_corners(&dungeon.rooms[i][j], (i-1)*SECTOR_HEIGHT, (j-1)*SECTOR_WIDTH)
				generate_doors(&dungeon.rooms[i][j])
			}
		}
	}
}

func generate_corners(room *room_t, offset_y int, offset_x int) {
	room.top_left.y = rand.Int()%CORNER_VERT_RANGE + offset_y + 1
	room.top_left.x = rand.Int()%CORNER_HOR_RANGE + offset_x + 1

	room.bot_right.y = SECTOR_HEIGHT - 1 - rand.Int()%CORNER_VERT_RANGE + offset_y - 1
	room.bot_right.x = SECTOR_WIDTH - 1 - rand.Int()%CORNER_HOR_RANGE + offset_x - 1
}

func generate_doors(room *room_t) {
	if room.connections[TOP] != nil {
		room.doors[TOP].y = room.top_left.y
		room.doors[TOP].x = rand.Int()%(int)(room.bot_right.x-room.top_left.x-1) + room.top_left.x + 1
	}

	if room.connections[RIGHT] != nil {
		room.doors[RIGHT].y = rand.Int()%(int)(room.bot_right.y-room.top_left.y-1) + room.top_left.y + 1
		room.doors[RIGHT].x = room.bot_right.x
	}

	if room.connections[BOTTOM] != nil {
		room.doors[BOTTOM].y = room.bot_right.y
		room.doors[BOTTOM].x = rand.Int()%(int)(room.bot_right.x-room.top_left.x-1) + room.top_left.x + 1
	}

	if room.connections[LEFT] != nil {
		room.doors[LEFT].y = rand.Int()%(int)(room.bot_right.y-room.top_left.y-1) + room.top_left.y + 1
		room.doors[LEFT].x = room.top_left.x
	}
}

func generate_corridors_geometry(dungeon *dungeon_t) {
	for i := 1; i < ROOMS_PER_SIDE+1; i++ {
		for j := 1; j < ROOMS_PER_SIDE+1; j++ {
			cur_room := &dungeon.rooms[i][j]

			if cur_room.connections[RIGHT] != nil && cur_room.connections[RIGHT].connections[LEFT] == cur_room {
				generate_left_to_right_corridor(dungeon, cur_room, cur_room.connections[RIGHT], &dungeon.corridors[dungeon.corridors_cnt])
				dungeon.corridors_cnt++
			}

			if cur_room.connections[BOTTOM] != nil {
				grid_i_diff := cur_room.grid_i - cur_room.connections[BOTTOM].grid_i
				grid_j_diff := cur_room.grid_j - cur_room.connections[BOTTOM].grid_j

				if grid_i_diff == -1 && grid_j_diff > 0 {
					generate_left_turn_corridor(dungeon, cur_room, cur_room.connections[BOTTOM], &dungeon.corridors[dungeon.corridors_cnt])
					dungeon.corridors_cnt++
				} else if grid_i_diff == -1 && grid_j_diff < 0 {
					generate_right_turn_corridor(dungeon, cur_room, cur_room.connections[BOTTOM], &dungeon.corridors[dungeon.corridors_cnt])
					dungeon.corridors_cnt++
				} else {
					generate_top_to_bottom_corridor(dungeon, cur_room, cur_room.connections[BOTTOM], &dungeon.corridors[dungeon.corridors_cnt])
					dungeon.corridors_cnt++
				}
			}
		}
	}
}

func generate_left_to_right_corridor(dungeon *dungeon_t, left_room *room_t, right_room *room_t, corridor *corridor_t) {
	corridor.typeE = LEFT_TO_RIGHT_CORRIDOR
	corridor.points_cnt = 4
	corridor.points[0] = left_room.doors[RIGHT]

	x_min := left_room.doors[RIGHT].x
	x_max := right_room.doors[LEFT].x

	for i := 1; i < ROOMS_PER_SIDE+1; i++ {
		if dungeon.rooms[i][left_room.grid_j].sector != UNINITIALIZED && i != left_room.grid_i {
			x_min = max(dungeon.rooms[i][left_room.grid_j].bot_right.x, x_min)
		}
	}
	for i := 1; i < ROOMS_PER_SIDE+1; i++ {
		if dungeon.rooms[i][right_room.grid_j].sector != UNINITIALIZED && i != right_room.grid_i {
			x_max = min(dungeon.rooms[i][right_room.grid_j].top_left.x, x_min)
		}
	}

	random_center_x := ((rand.Int() % (x_max - x_min - 1)) + 1) + (x_min)
	var second_point, third_point position_t
	second_point.x, second_point.y = random_center_x, left_room.doors[RIGHT].y
	third_point.x, third_point.y = random_center_x, right_room.doors[LEFT].y

	corridor.points[1] = second_point
	corridor.points[2] = third_point
	corridor.points[3] = right_room.doors[3]
}

func generate_left_turn_corridor(dungeon *dungeon_t, top_room *room_t, bottom_left_room *room_t, corridor *corridor_t) {
	corridor.typeE = LEFT_TURN_CORRIDOR
	corridor.points_cnt = 3
	corridor.points[0] = top_room.doors[BOTTOM]

	var second_point position_t
	second_point.x, second_point.y = top_room.doors[BOTTOM].x, bottom_left_room.doors[RIGHT].y

	corridor.points[1] = second_point
	corridor.points[2] = bottom_left_room.doors[RIGHT]
}

func generate_right_turn_corridor(dungeon *dungeon_t, top_room *room_t, bottom_right_room *room_t, corridor *corridor_t) {
	corridor.typeE = RIGHT_TURN_CORRIDOR
	corridor.points_cnt = 3
	corridor.points[0] = top_room.doors[BOTTOM]

	var second_point position_t
	second_point.x, second_point.y = top_room.doors[BOTTOM].x, bottom_right_room.doors[LEFT].y

	corridor.points[1] = second_point
	corridor.points[2] = bottom_right_room.doors[LEFT]
}

func generate_top_to_bottom_corridor(dungeon *dungeon_t, top_room *room_t, bottom_room *room_t, corridor *corridor_t) {
	corridor.typeE = TOP_TO_BOTTOM_CORRIDOR
	corridor.points_cnt = 4
	corridor.points[0] = top_room.doors[BOTTOM]

	y_min := top_room.doors[BOTTOM].y
	y_max := bottom_room.doors[TOP].y

	for j := 1; j < ROOMS_PER_SIDE+1; j++ {
		if dungeon.rooms[top_room.grid_i][j].sector != UNINITIALIZED {
			y_min = max(dungeon.rooms[top_room.grid_i][j].bot_right.y, y_min)
		}
	}
	for j := 1; j < ROOMS_PER_SIDE+1; j++ {
		if dungeon.rooms[bottom_room.grid_i][j].sector != UNINITIALIZED {
			y_max = min(dungeon.rooms[bottom_room.grid_i][j].top_left.y, y_max)
		}
	}

	random_center_y := ((rand.Int() % (y_max - y_min - 1)) + 1) + (y_min)
	var second_point, third_point position_t
	second_point.x = top_room.doors[BOTTOM].x
	second_point.y = random_center_y
	third_point.x, third_point.y = bottom_room.doors[TOP].x, random_center_y

	corridor.points[1] = second_point
	corridor.points[2] = third_point
	corridor.points[3] = bottom_room.doors[TOP]
}

func generateLockedRooms(dungeon *dungeon_t) {
	canBeLocked := make([]*room_t, 0)
	for i := range dungeon.sequence {
		if dungeon.sequence[i] != nil {
			if dungeon.sequence[i].isStart {
				continue
			}
			var count int
			for k := range dungeon.sequence[i].connections {

				if dungeon.sequence[i].connections[k] != nil {
					count++
				}
			}
			if count == 1 {
				canBeLocked = append(canBeLocked, dungeon.sequence[i])
			}
		}
	}
	if len(canBeLocked) == 1 {
		canBeLocked[0].isLocked = true
		canBeLocked[0].blueLock = true
	} else if len(canBeLocked) == 2 {
		index := rand.Int() % 2
		if index == 0 {
			canBeLocked[0].isLocked = true
			canBeLocked[0].blueLock = true
			canBeLocked[1].isLocked = true
			canBeLocked[1].magentaLock = true
		} else {
			canBeLocked[0].isLocked = true
			canBeLocked[1].blueLock = true
			canBeLocked[1].isLocked = true
			canBeLocked[0].magentaLock = true
		}
	} else if len(canBeLocked) == 3 {
		index := rand.Int() % 3
		canBeLocked[index].isLocked = true
		canBeLocked[index].blueLock = true
		for canBeLocked[index].isLocked {
			index = rand.Int() % 3
		}
		canBeLocked[index].isLocked = true
		canBeLocked[index].magentaLock = true
		for canBeLocked[index].isLocked {
			index = rand.Int() % 3
		}
		canBeLocked[index].isLocked = true
		canBeLocked[index].cyanLock = true
	} else if len(canBeLocked) > 3 {
		num := len(canBeLocked) - 1
		index := rand.Int() % num
		canBeLocked[index].isLocked = true
		canBeLocked[index].blueLock = true
		for canBeLocked[index].isLocked {
			index = rand.Int() % num
		}
		canBeLocked[index].isLocked = true
		canBeLocked[index].magentaLock = true
		for canBeLocked[index].isLocked {
			index = rand.Int() % num
		}
		canBeLocked[index].isLocked = true
		canBeLocked[index].cyanLock = true
	}
}
