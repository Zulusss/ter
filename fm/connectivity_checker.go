package main

func check_connectivity(rooms *dungeon_t) int {
	rc := CONNECTED
	visited := make([]int, 9)
	visited_count := depth_first_search(rooms.sequence[0], visited)

	if visited_count != rooms.room_cnt {
		rc = NOT_CONNECTED
	}

	return rc
}

func depth_first_search(cur *room_t, visited []int) int {
	visited_count := 1

	visited[cur.sector] = 1

	for i := 0; i < 4; i++ {
		if cur.connections[i] != nil && visited[cur.connections[i].sector] == 0 {
			visited_count += depth_first_search(cur.connections[i], visited)
		}
	}
	return visited_count
}
