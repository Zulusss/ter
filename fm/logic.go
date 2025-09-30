package main

import (
	"math"
	"math/rand"
)

func checkTile(dungeon *dungeon_t, field *map_t, player *player_t) {
	if field.playground[player.pos.y][player.pos.x] == '|' {
		generateNewLevel(dungeon, field, player)
	} else if field.playground[player.pos.y][player.pos.x] == '?' ||
		field.playground[player.pos.y][player.pos.x] == '*' ||
		field.playground[player.pos.y][player.pos.x] == '/' ||
		field.playground[player.pos.y][player.pos.x] == '$' {
		getItem(player, field)
		// if field.playground[player.pos.y][player.pos.x] != ' ' {
		// 	field.playground[player.pos.y][player.pos.x] = '^'
		// }
	}
}

func generateNewLevel(dungeon *dungeon_t, field *map_t, player *player_t) {
	dungeon.level++
	var newDungeon dungeon_t
	generate_dungeon(&newDungeon)
	newDungeon.level = dungeon.level
	// generate_player_pos(&dungeon)
	print_dungeon_generated(MAP_HEIGHT + 5)
	generate_entities(&newDungeon)
	print_entities_generated(MAP_HEIGHT + 6)
	// newDungeon.level = dungeon.level
	// newDungeon.level++
	// dungeon = nil
	*dungeon = newDungeon
	// field = nil
	// var newField map_t
	// field = &newField
	clearField(field)
	// var field map_t
	dungeon_to_map(dungeon, field)
	init_player(player, &field.player_spawn)
	print_map_created(MAP_HEIGHT + 7)
	check_result := check_connectivity(dungeon)
	print_connectivity_check_result(check_result, MAP_HEIGHT+9)
	// print_map(field)
	// firstMove(dungeon, field, player)

	// print_entity_pools(MAP_HEIGHT + 11)

}

func clearField(field *map_t) {
	for i := range field.items {
		var zeroEntity entity_t
		field.items[i] = zeroEntity
	}
	for i := range field.enemies {
		var zeroEnemy entity_t
		field.enemies[i] = zeroEnemy
	}
	field.items_cnt = 0
	field.enemies_cnt = 0
}

func getItem(player *player_t, field *map_t) {
	var item entity_t
	for i := 0; i < field.items_cnt; i++ {
		if field.items[i].pos.x == player.pos.x && field.items[i].pos.y == player.pos.y {
			switch field.items[i].typeI {
			case 10:
				if len(player.inventory.food) < 9 {
					player.inventory.food = append(player.inventory.food, 10)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			case 11:
				if len(player.inventory.food) < 9 {
					player.inventory.food = append(player.inventory.food, 11)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			case 12:
				if len(player.inventory.food) < 9 {
					player.inventory.food = append(player.inventory.food, 12)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			case 20:
				if len(player.inventory.scroll) < 9 {
					player.inventory.scroll = append(player.inventory.scroll, 20)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			case 21:
				if len(player.inventory.scroll) < 9 {
					player.inventory.scroll = append(player.inventory.scroll, 21)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			case 22:
				if len(player.inventory.scroll) < 9 {
					player.inventory.scroll = append(player.inventory.scroll, 22)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			case 30:
				if len(player.inventory.weapon) < 9 {
					player.inventory.weapon = append(player.inventory.weapon, 30)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			case 31:
				if len(player.inventory.weapon) < 9 {
					player.inventory.weapon = append(player.inventory.weapon, 31)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			case 32:
				if len(player.inventory.weapon) < 9 {
					player.inventory.weapon = append(player.inventory.weapon, 32)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			case 40:
				if len(player.inventory.potion) < 9 {
					player.inventory.potion = append(player.inventory.potion, 40)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			case 41:
				if len(player.inventory.potion) < 9 {
					player.inventory.potion = append(player.inventory.potion, 41)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			case 42:
				if len(player.inventory.potion) < 9 {
					player.inventory.potion = append(player.inventory.potion, 42)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
				}
			}
		}

	}
}

func enemyTurn(dungeon *dungeon_t, field *map_t, player *player_t) {
	for i := 0; i < dungeon.room_cnt; i++ {
		enemies_cnt := dungeon.sequence[i].entities_cnt

		for j := 0; j < enemies_cnt; j++ {
			if dungeon.sequence[i].entities[j].typeE == ENEMY {
				enemy := dungeon.sequence[i].entities[j]
				distY, distX := checkAgr(&enemy, player)
				if !enemy.stats.isChasing {
					switch enemy.symbol {
					case 'z':
						zombieMove(field, &enemy)
						dungeon.sequence[i].entities[j] = enemy
					case 's':
						snakeMove(field, &enemy)
						dungeon.sequence[i].entities[j] = enemy
					case 'g':
						ghostMove(dungeon, field, &enemy)
						dungeon.sequence[i].entities[j] = enemy
					case 'v':
						vampireMove(field, &enemy)
						dungeon.sequence[i].entities[j] = enemy
					case 'O':
						orgeMove(field, &enemy)
						dungeon.sequence[i].entities[j] = enemy
					case 'm':
						field.playground[enemy.pos.y][enemy.pos.x] = enemy.stats.extraSym
					}
				} else {
					enemyChasing(field, &enemy, player, &distY, &distX)
					// dirY := enemy.pos.y
					// dirX := enemy.pos.x
					// // test 4-side movement 4 sides
					// if distY > distX {
					// 	if dirY > player.pos.y {
					// 		dirY -= 1
					// 	} else if dirY < player.pos.y {
					// 		dirY += 1
					// 	}
					// } else {
					// 	if dirX > player.pos.x {
					// 		dirX -= 1
					// 	} else if dirX < player.pos.x {
					// 		dirX += 1
					// 	}
					// }
					// // if failed - try diagonal (corners)
					// if field.playground[dirY][dirX] != ' ' && field.playground[dirY][dirX] != '+' {
					// 	dirY = enemy.pos.y
					// 	dirX = enemy.pos.x

					// 	if dirY > player.pos.y {
					// 		dirY -= 1
					// 	} else if dirY < player.pos.y {
					// 		dirY += 1
					// 	} else if dirY == player.pos.y {
					// 		if rand.Int()%2 == 0 {
					// 			dirY += 1
					// 		} else {
					// 			dirY -= 1
					// 		}
					// 	}

					// 	if dirX > player.pos.x {
					// 		dirX -= 1
					// 	} else if dirX < player.pos.x {
					// 		dirX += 1
					// 	} else if dirX == player.pos.x {
					// 		if rand.Int()%2 == 0 {
					// 			dirX += 1
					// 		} else {
					// 			dirX -= 1
					// 		}
					// 	}
					// }
					// // if still failed - move to closes to player empty tile (invert)
					// if field.playground[dirY][dirX] != ' ' && field.playground[dirY][dirX] != '+' {
					// 	dirY = enemy.pos.y
					// 	dirX = enemy.pos.x

					// 	if distY > distX {
					// 		if dirX > player.pos.x {
					// 			dirX -= 1
					// 		} else if dirX < player.pos.x {
					// 			dirX += 1
					// 		} else if dirX == player.pos.x {
					// 			if rand.Int()%2 == 0 {
					// 				dirX += 1
					// 			} else {
					// 				dirX -= 1
					// 			}
					// 		}
					// 	} else {
					// 		if dirY > player.pos.y {
					// 			dirY -= 1
					// 		} else if dirY < player.pos.y {
					// 			dirY += 1
					// 		} else if dirY == player.pos.y {
					// 			if rand.Int()%2 == 0 {
					// 				dirY += 1
					// 			} else {
					// 				dirY -= 1
					// 			}
					// 		}
					// 	}
					// }
					if distY+distX < 2 {
						enemyAttack(&enemy, player)
						dungeon.sequence[i].entities[j] = enemy
					} else { /*if (field.playground[dirY][dirX] == ' ' || field.playground[dirY][dirX] == '+') &&
						(dirY != player.pos.y && dirX != player.pos.x) {
						field.playground[enemy.pos.y][enemy.pos.x] = ' '
						enemy.pos.y = dirY
						enemy.pos.x = dirX
						field.playground[enemy.pos.y][enemy.pos.x] = enemy.symbol */
						dungeon.sequence[i].entities[j] = enemy
					}
				}

			}
		}
	}
}

func enemyChasing(field *map_t, enemy *entity_t, player *player_t, distY *int, distX *int) {
	var empty bool

	dirY := enemy.pos.y
	dirX := enemy.pos.x
	// if dirY == enemy.pos.y && (dirX == player.pos.x+1 || dirX == player.pos.x-1) {
	// 	return
	// }
	// if dirX == enemy.pos.x && (dirY == player.pos.y+1 || dirY == player.pos.y-1) {
	// 	return
	// }
	// test 4-side movement 4 sides
	if *distY > *distX {
		if dirY > player.pos.y {
			dirY -= 1
		} else if dirY < player.pos.y {
			dirY += 1
		}
	} else {
		if dirX > player.pos.x {
			dirX -= 1
		} else if dirX < player.pos.x {
			dirX += 1
		}
	}

	if field.playground[dirY][dirX] == ' ' || field.playground[dirY][dirX] == '+' { /*
			(dirY != player.pos.y && dirX != player.pos.x) */
		field.playground[enemy.pos.y][enemy.pos.x] = ' '
		enemy.pos.y = dirY
		enemy.pos.x = dirX
		field.playground[enemy.pos.y][enemy.pos.x] = enemy.symbol
		empty = true
	} else if field.playground[dirY][dirX] == field.playground[player.pos.y][player.pos.x] {

		empty = true
	}
	//other side
	if !empty {
		dirY = enemy.pos.y
		dirX = enemy.pos.x
		if *distY > *distX {
			if dirX > player.pos.x {
				dirX -= 1
			} else if dirX < player.pos.x {
				dirX += 1
			}
		} else {
			if dirY > player.pos.y {
				dirY -= 1
			} else if dirY < player.pos.y {
				dirY += 1
			}
		}
		if field.playground[dirY][dirX] == ' ' || field.playground[dirY][dirX] == '+' { /*&&
			(dirY != player.pos.y && dirX != player.pos.x) { */
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y = dirY
			enemy.pos.x = dirX
			field.playground[enemy.pos.y][enemy.pos.x] = enemy.symbol
			empty = true
		} else if field.playground[dirY][dirX] == field.playground[player.pos.y][player.pos.x] {
			empty = true
		}
	}
	//diagonal
	// if !empty {
	// 	dirY = enemy.pos.y
	// 	dirX = enemy.pos.x
	// 	if dirY > player.pos.y {
	// 		dirY -= 1
	// 	} else if dirY < player.pos.y {
	// 		dirY += 1
	// 	}

	// 	if dirX > player.pos.x {
	// 		dirX -= 1
	// 	} else if dirX < player.pos.x {
	// 		dirX += 1
	// 	}
	// 	if (field.playground[dirY][dirX] == ' ' || field.playground[dirY][dirX] == '+') &&
	// 		(dirY != player.pos.y && dirX != player.pos.x) {
	// 		field.playground[enemy.pos.y][enemy.pos.x] = ' '
	// 		enemy.pos.y = dirY
	// 		enemy.pos.x = dirX
	// 		field.playground[enemy.pos.y][enemy.pos.x] = enemy.symbol
	// 		empty = true
	// 	} else if dirY == player.pos.y && dirX == player.pos.x {
	// 		if rand.Int()%2 == 0 {
	// 			if enemy.pos.y > player.pos.y {
	// 				dirY -= 1
	// 			} else {
	// 				dirY += 1
	// 			}
	// 			if field.playground[dirY][dirX] == ' ' || field.playground[dirY][dirX] == '+' {
	// 				field.playground[enemy.pos.y][enemy.pos.x] = ' '
	// 				enemy.pos.y = dirY
	// 				enemy.pos.x = dirX
	// 				field.playground[enemy.pos.y][enemy.pos.x] = enemy.symbol
	// 				empty = true
	// 			} else {
	// 				dirY = enemy.pos.y
	// 				if enemy.pos.x > player.pos.x {
	// 					dirX -= 1
	// 				} else {
	// 					dirX += 1
	// 				}
	// 				if field.playground[dirY][dirX] == ' ' || field.playground[dirY][dirX] == '+' {
	// 					field.playground[enemy.pos.y][enemy.pos.x] = ' '
	// 					enemy.pos.y = dirY
	// 					enemy.pos.x = dirX
	// 					field.playground[enemy.pos.y][enemy.pos.x] = enemy.symbol
	// 					empty = true
	// 				}
	// 			}
	// 		} else {
	// 			if enemy.pos.x > player.pos.x {
	// 				dirX -= 1
	// 			} else {
	// 				dirX += 1
	// 			}
	// 			if field.playground[dirY][dirX] == ' ' || field.playground[dirY][dirX] == '+' {
	// 				field.playground[enemy.pos.y][enemy.pos.x] = ' '
	// 				enemy.pos.y = dirY
	// 				enemy.pos.x = dirX
	// 				field.playground[enemy.pos.y][enemy.pos.x] = enemy.symbol
	// 				empty = true
	// 			} else {
	// 				dirX = enemy.pos.x
	// 				if enemy.pos.y > player.pos.y {
	// 					dirY -= 1
	// 				} else {
	// 					dirY += 1
	// 				}
	// 				if field.playground[dirY][dirX] == ' ' || field.playground[dirY][dirX] == '+' {
	// 					field.playground[enemy.pos.y][enemy.pos.x] = ' '
	// 					enemy.pos.y = dirY
	// 					enemy.pos.x = dirX
	// 					field.playground[enemy.pos.y][enemy.pos.x] = enemy.symbol
	// 					empty = true
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}

func enemyAttack(enemy *entity_t, player *player_t) {
	if enemy.symbol == 'O' {
		if !enemy.stats.isFirst {
			dodge := player.agility - enemy.stats.agility - rand.Int()%enemy.stats.agility
			if dodge < 1 {
				player.health -= enemy.stats.strength
			}
			enemy.stats.isFirst = true
		} else {
			enemy.stats.isFirst = false
		}
	} else {
		dodge := player.agility - enemy.stats.agility - rand.Int()%enemy.stats.agility
		if dodge < 1 {
			player.health -= enemy.stats.strength
			if enemy.symbol == 's' {
				if rand.Int()%4 == 0 {
					player.isSleeped = true
				}
			} else if enemy.symbol == 'v' {
				stealMaxHealth := rand.Int()%enemy.stats.strength + 1
				player.maxHealth -= stealMaxHealth
				enemy.stats.hpStealed += stealMaxHealth
			}
		}
	}
}

func checkAgr(enemy *entity_t, player *player_t) (int, int) {
	// distY := int(math.Abs(float64(enemy.pos.y)) - float64(player.pos.y))
	// distX := int(math.Abs(float64(enemy.pos.x)) - float64(player.pos.x))
	distY := int(math.Abs(float64(enemy.pos.y - player.pos.y)))
	distX := int(math.Abs(float64(enemy.pos.x - player.pos.x)))

	if distY < enemy.stats.aggression && distX < enemy.stats.aggression {
		enemy.stats.isChasing = true
	}
	return distY, distX
}

func zombieMove(field *map_t, enemy *entity_t) {
	if enemy.stats.lastMove == TOP {
		if field.playground[enemy.pos.y-1][enemy.pos.x] == ' ' {
			field.playground[enemy.pos.y-1][enemy.pos.x] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y -= 1
			enemy.stats.lastMove = TOP
		} else if field.playground[enemy.pos.y][enemy.pos.x+1] == ' ' {
			field.playground[enemy.pos.y][enemy.pos.x+1] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.x += 1
			enemy.stats.lastMove = RIGHT
		} else if field.playground[enemy.pos.y+1][enemy.pos.x] == ' ' {
			field.playground[enemy.pos.y+1][enemy.pos.x] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y += 1
			enemy.stats.lastMove = BOTTOM
		} else if field.playground[enemy.pos.y][enemy.pos.x-1] == ' ' {
			field.playground[enemy.pos.y][enemy.pos.x-1] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.x -= 1
			enemy.stats.lastMove = LEFT
		}

	} else if enemy.stats.lastMove == RIGHT {
		if field.playground[enemy.pos.y][enemy.pos.x+1] == ' ' {
			field.playground[enemy.pos.y][enemy.pos.x+1] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.x += 1
			enemy.stats.lastMove = RIGHT
		} else if field.playground[enemy.pos.y+1][enemy.pos.x] == ' ' {
			field.playground[enemy.pos.y+1][enemy.pos.x] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y += 1
			enemy.stats.lastMove = BOTTOM
		} else if field.playground[enemy.pos.y][enemy.pos.x-1] == ' ' {
			field.playground[enemy.pos.y][enemy.pos.x-1] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.x -= 1
			enemy.stats.lastMove = LEFT
		} else if field.playground[enemy.pos.y-1][enemy.pos.x] == ' ' {
			field.playground[enemy.pos.y-1][enemy.pos.x] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y -= 1
			enemy.stats.lastMove = TOP
		}
	} else if enemy.stats.lastMove == BOTTOM {
		if field.playground[enemy.pos.y+1][enemy.pos.x] == ' ' {
			field.playground[enemy.pos.y+1][enemy.pos.x] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y += 1
			enemy.stats.lastMove = BOTTOM
		} else if field.playground[enemy.pos.y][enemy.pos.x-1] == ' ' {
			field.playground[enemy.pos.y][enemy.pos.x-1] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.x -= 1
			enemy.stats.lastMove = LEFT
		} else if field.playground[enemy.pos.y-1][enemy.pos.x] == ' ' {
			field.playground[enemy.pos.y-1][enemy.pos.x] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y -= 1
			enemy.stats.lastMove = TOP
		} else if field.playground[enemy.pos.y][enemy.pos.x+1] == ' ' {
			field.playground[enemy.pos.y][enemy.pos.x+1] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.x += 1
			enemy.stats.lastMove = RIGHT
		}
	} else if enemy.stats.lastMove == LEFT {
		if field.playground[enemy.pos.y][enemy.pos.x-1] == ' ' {
			field.playground[enemy.pos.y][enemy.pos.x-1] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.x -= 1
			enemy.stats.lastMove = LEFT
		} else if field.playground[enemy.pos.y-1][enemy.pos.x] == ' ' {
			field.playground[enemy.pos.y-1][enemy.pos.x] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y -= 1
			enemy.stats.lastMove = TOP
		} else if field.playground[enemy.pos.y][enemy.pos.x+1] == ' ' {
			field.playground[enemy.pos.y][enemy.pos.x+1] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.x += 1
			enemy.stats.lastMove = RIGHT
		} else if field.playground[enemy.pos.y+1][enemy.pos.x] == ' ' {
			field.playground[enemy.pos.y+1][enemy.pos.x] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y += 1
			enemy.stats.lastMove = BOTTOM
		}
	}
}

func ghostMove(dungeon *dungeon_t, field *map_t, enemy *entity_t) {
	for i := 0; i < dungeon.room_cnt; i++ {
		if enemy.pos.y > dungeon.sequence[i].top_left.y && enemy.pos.x > dungeon.sequence[i].top_left.x &&
			enemy.pos.y < dungeon.sequence[i].bot_right.y && enemy.pos.x < dungeon.sequence[i].bot_right.x {
			lastPos := enemy.pos
			generate_entity_coords(dungeon.sequence[i], &enemy.pos)
			field.playground[lastPos.y][lastPos.x] = ' '
			field.playground[enemy.pos.y][enemy.pos.x] = enemy.symbol
			break
		}
	}
}

func snakeMove(field *map_t, enemy *entity_t) {
	directions := make([]bool, 4)
	directCoord := make([]position_t, 4)

	if field.playground[enemy.pos.y-1][enemy.pos.x+1] == ' ' {
		directions[TOP] = true
		directCoord[TOP].y = enemy.pos.y - 1
		directCoord[TOP].x = enemy.pos.x + 1
	}
	if field.playground[enemy.pos.y+1][enemy.pos.x+1] == ' ' {
		directions[RIGHT] = true
		directCoord[RIGHT].y = enemy.pos.y + 1
		directCoord[RIGHT].x = enemy.pos.x + 1
	}
	if field.playground[enemy.pos.y+1][enemy.pos.x-1] == ' ' {
		directions[BOTTOM] = true
		directCoord[BOTTOM].y = enemy.pos.y + 1
		directCoord[BOTTOM].x = enemy.pos.x - 1
	}
	if field.playground[enemy.pos.y-1][enemy.pos.x-1] == ' ' {
		directions[LEFT] = true
		directCoord[LEFT].y = enemy.pos.y - 1
		directCoord[LEFT].x = enemy.pos.x - 1
	}

	dir := rand.Int() % 4
	for enemy.stats.lastMove == dir {
		dir = rand.Int() % 4
	}

	if directions[dir] {
		// if directions[TOP] {
		field.playground[directCoord[dir].y][directCoord[dir].x] = enemy.symbol
		field.playground[enemy.pos.y][enemy.pos.x] = ' '
		enemy.pos.y = directCoord[dir].y
		enemy.pos.x = directCoord[dir].x
		enemy.stats.lastMove = dir
	} else {
		dir2 := rand.Int() % 4
		for dir2 == dir || enemy.stats.lastMove == dir2 {
			dir2 = rand.Int() % 4
		}
		if directions[dir2] {
			field.playground[directCoord[dir2].y][directCoord[dir2].x] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y = directCoord[dir2].y
			enemy.pos.x = directCoord[dir2].x
			enemy.stats.lastMove = dir2
		} else {
			dir3 := rand.Int() % 4
			for dir3 == dir2 || dir3 == dir || dir3 == enemy.stats.lastMove {
				dir3 = rand.Int() % 4
			}
			if directions[dir3] {
				field.playground[directCoord[dir3].y][directCoord[dir3].x] = enemy.symbol
				field.playground[enemy.pos.y][enemy.pos.x] = ' '
				enemy.pos.y = directCoord[dir3].y
				enemy.pos.x = directCoord[dir3].x
				enemy.stats.lastMove = dir3
			} else {
				if directions[enemy.stats.lastMove] {
					field.playground[directCoord[enemy.stats.lastMove].y][directCoord[enemy.stats.lastMove].x] = enemy.symbol
					field.playground[enemy.pos.y][enemy.pos.x] = ' '
					enemy.pos.y = directCoord[enemy.stats.lastMove].y
					enemy.pos.x = directCoord[enemy.stats.lastMove].x
					// enemy.stats.lastMove = enemy.stats.lastMove
				}
			}
		}
	}
	// if field.playground[enemy.pos.y-1][enemy.pos.x+1] == ' ' {
	// 	field.playground[enemy.pos.y-1][enemy.pos.x+1] = enemy.symbol
	// 	field.playground[enemy.pos.y][enemy.pos.x] = ' '
	// 	enemy.pos.y -= 1
	// 	enemy.pos.x += 1
	// 	enemy.stats.lastMove = TOP
	// }
}

func vampireMove(field *map_t, enemy *entity_t) {
	directions := make([]bool, 4)
	directCoord := make([]position_t, 4)

	if field.playground[enemy.pos.y-1][enemy.pos.x] == ' ' {
		directions[TOP] = true
		directCoord[TOP].y = enemy.pos.y - 1
		directCoord[TOP].x = enemy.pos.x
	}
	if field.playground[enemy.pos.y][enemy.pos.x+1] == ' ' {
		directions[RIGHT] = true
		directCoord[RIGHT].y = enemy.pos.y
		directCoord[RIGHT].x = enemy.pos.x + 1
	}
	if field.playground[enemy.pos.y+1][enemy.pos.x] == ' ' {
		directions[BOTTOM] = true
		directCoord[BOTTOM].y = enemy.pos.y + 1
		directCoord[BOTTOM].x = enemy.pos.x
	}
	if field.playground[enemy.pos.y][enemy.pos.x-1] == ' ' {
		directions[LEFT] = true
		directCoord[LEFT].y = enemy.pos.y
		directCoord[LEFT].x = enemy.pos.x - 1
	}
	dir := rand.Int() % 4
	if directions[dir] {
		field.playground[directCoord[dir].y][directCoord[dir].x] = enemy.symbol
		field.playground[enemy.pos.y][enemy.pos.x] = ' '
		enemy.pos.y = directCoord[dir].y
		enemy.pos.x = directCoord[dir].x
	} else {
		dir2 := rand.Int() % 4
		for dir2 == dir {
			dir2 = rand.Int() % 4
		}
		if directions[dir2] {
			field.playground[directCoord[dir2].y][directCoord[dir2].x] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y = directCoord[dir2].y
			enemy.pos.x = directCoord[dir2].x
		} else {
			dir3 := rand.Int() % 4
			for dir3 == dir2 || dir3 == dir {
				dir3 = rand.Int() % 4
			}
			if directions[dir3] {
				field.playground[directCoord[dir3].y][directCoord[dir3].x] = enemy.symbol
				field.playground[enemy.pos.y][enemy.pos.x] = ' '
				enemy.pos.y = directCoord[dir3].y
				enemy.pos.x = directCoord[dir3].x
			} else {
				dir4 := rand.Int() % 4
				for dir4 == dir3 || dir4 == dir2 || dir4 == dir {
					dir4 = rand.Int() % 4
				}
				if directions[dir4] {
					field.playground[directCoord[dir4].y][directCoord[dir4].x] = enemy.symbol
					field.playground[enemy.pos.y][enemy.pos.x] = ' '
					enemy.pos.y = directCoord[dir4].y
					enemy.pos.x = directCoord[dir4].x
				}
			}
		}
	}
}

func orgeMove(field *map_t, enemy *entity_t) {
	directions := make([]bool, 4)
	directCoord := make([]position_t, 4)

	if field.playground[enemy.pos.y-2][enemy.pos.x] == ' ' {
		directions[TOP] = true
		directCoord[TOP].y = enemy.pos.y - 2
		directCoord[TOP].x = enemy.pos.x
	}
	if field.playground[enemy.pos.y][enemy.pos.x+2] == ' ' {
		directions[RIGHT] = true
		directCoord[RIGHT].y = enemy.pos.y
		directCoord[RIGHT].x = enemy.pos.x + 2
	}
	if field.playground[enemy.pos.y+2][enemy.pos.x] == ' ' {
		directions[BOTTOM] = true
		directCoord[BOTTOM].y = enemy.pos.y + 2
		directCoord[BOTTOM].x = enemy.pos.x
	}
	if field.playground[enemy.pos.y][enemy.pos.x-2] == ' ' {
		directions[LEFT] = true
		directCoord[LEFT].y = enemy.pos.y
		directCoord[LEFT].x = enemy.pos.x - 2
	}
	dir := rand.Int() % 4
	if directions[dir] {
		field.playground[directCoord[dir].y][directCoord[dir].x] = enemy.symbol
		field.playground[enemy.pos.y][enemy.pos.x] = ' '
		enemy.pos.y = directCoord[dir].y
		enemy.pos.x = directCoord[dir].x
	} else {
		dir2 := rand.Int() % 4
		for dir2 == dir {
			dir2 = rand.Int() % 4
		}
		if directions[dir2] {
			field.playground[directCoord[dir2].y][directCoord[dir2].x] = enemy.symbol
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y = directCoord[dir2].y
			enemy.pos.x = directCoord[dir2].x
		} else {
			dir3 := rand.Int() % 4
			for dir3 == dir2 || dir3 == dir {
				dir3 = rand.Int() % 4
			}
			if directions[dir3] {
				field.playground[directCoord[dir3].y][directCoord[dir3].x] = enemy.symbol
				field.playground[enemy.pos.y][enemy.pos.x] = ' '
				enemy.pos.y = directCoord[dir3].y
				enemy.pos.x = directCoord[dir3].x
			} else {
				dir4 := rand.Int() % 4
				for dir4 == dir3 || dir4 == dir2 || dir4 == dir {
					dir4 = rand.Int() % 4
				}
				if directions[dir4] {
					field.playground[directCoord[dir4].y][directCoord[dir4].x] = enemy.symbol
					field.playground[enemy.pos.y][enemy.pos.x] = ' '
					enemy.pos.y = directCoord[dir4].y
					enemy.pos.x = directCoord[dir4].x
				}
			}
		}
	}
}
