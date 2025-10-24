package main

import (
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"os"
)

func checkTile(dungeon *dungeon_t, field *map_t, player *player_t, msgStr *list.List) {
	if field.playground[player.pos.y][player.pos.x] == '|' {
		generateNewLevel(dungeon, field, player)
	} else if field.playground[player.pos.y][player.pos.x] == '?' ||
		field.playground[player.pos.y][player.pos.x] == '*' ||
		field.playground[player.pos.y][player.pos.x] == '/' ||
		field.playground[player.pos.y][player.pos.x] == '$' ||
		field.playground[player.pos.y][player.pos.x] == '!' {
		getItem(player, field, msgStr)
	}
}

func generateNewLevel(dungeon *dungeon_t, field *map_t, player *player_t) {
	dungeon.level++
	var newDungeon dungeon_t
	generate_dungeon(&newDungeon)
	newDungeon.level = dungeon.level

	print_dungeon_generated(MAP_HEIGHT + 5)
	if player.health > player.diff.hp {
		player.diff.difficulty = 1
		player.diff.hp = player.health
	} else if player.health < player.diff.hp {
		player.diff.difficulty = -1
		player.diff.hp = player.health
	}
	generate_entities(&newDungeon, player.diff.difficulty)
	print_entities_generated(MAP_HEIGHT + 6)
	*dungeon = newDungeon
	clearField(field)
	dungeon_to_map(dungeon, field)
	init_player(player, &field.player_spawn)
	print_map_created(MAP_HEIGHT + 7)
	check_result := check_connectivity(dungeon)
	print_connectivity_check_result(check_result, MAP_HEIGHT+9)
	player.gotBlue = false
	player.gotMagenta = false
	player.gotCyan = false

	Session := saveSessionFromGame(dungeon, field, player)
	Store, _ := NewStore()

	err := Store.SaveSession(Session)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if dungeon.level == 22 {
		printCongratulations()
	}

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

func getItem(player *player_t, field *map_t, msgStr *list.List) {
	var item entity_t
	for i := 0; i < field.items_cnt; i++ {
		if field.items[i].pos.x == player.pos.x && field.items[i].pos.y == player.pos.y {
			switch field.items[i].typeI {
			case 10:
				if len(player.inventory.food) < 9 {
					player.inventory.food = append(player.inventory.food, 10)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found apple.")
				} else {
					msgStr.PushFront("Too many food.")
				}
			case 11:
				if len(player.inventory.food) < 9 {
					player.inventory.food = append(player.inventory.food, 11)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found banana.")
				} else {
					msgStr.PushFront("Too many food.")
				}
			case 12:
				if len(player.inventory.food) < 9 {
					player.inventory.food = append(player.inventory.food, 12)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found cucumber.")
				} else {
					msgStr.PushFront("Too many food.")
				}
			case 20:
				if len(player.inventory.scroll) < 9 {
					player.inventory.scroll = append(player.inventory.scroll, 20)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found scroll of strength.")
				} else {
					msgStr.PushFront("Too many scrolls.")
				}
			case 21:
				if len(player.inventory.scroll) < 9 {
					player.inventory.scroll = append(player.inventory.scroll, 21)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found scroll of agility.")
				} else {
					msgStr.PushFront("Too many scrolls.")
				}
			case 22:
				if len(player.inventory.scroll) < 9 {
					player.inventory.scroll = append(player.inventory.scroll, 22)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found scroll of max health.")
				} else {
					msgStr.PushFront("Too many scrolls.")
				}
			case 30:
				if len(player.inventory.weapon) < 9 {
					player.inventory.weapon = append(player.inventory.weapon, 30)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found sword.")
				} else {
					msgStr.PushFront("Too many weapon.")
				}
			case 31:
				if len(player.inventory.weapon) < 9 {
					player.inventory.weapon = append(player.inventory.weapon, 31)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found axe.")
				} else {
					msgStr.PushFront("Too many weapon.")
				}
			case 32:
				if len(player.inventory.weapon) < 9 {
					player.inventory.weapon = append(player.inventory.weapon, 32)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found hammer.")
				} else {
					msgStr.PushFront("Too many weapon.")
				}
			case 40:
				if len(player.inventory.potion) < 9 {
					player.inventory.potion = append(player.inventory.potion, 40)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found potion of strength.")
				} else {
					msgStr.PushFront("Too many potions.")
				}
			case 41:
				if len(player.inventory.potion) < 9 {
					player.inventory.potion = append(player.inventory.potion, 41)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found potion of agility.")
				} else {
					msgStr.PushFront("Too many potions.")
				}
			case 42:
				if len(player.inventory.potion) < 9 {
					player.inventory.potion = append(player.inventory.potion, 42)
					field.items[i] = item
					field.playground[player.pos.y][player.pos.x] = ' '
					msgStr.PushFront("You found potion of max health.")
				} else {
					msgStr.PushFront("Too many potions.")
				}
			case 51:
				player.gotBlue = true
				msgStr.PushFront("Found blue key.")
				field.items[i] = item
				field.playground[player.pos.y][player.pos.x] = ' '
			case 52:
				player.gotMagenta = true
				msgStr.PushFront("Found magenta key")
				field.items[i] = item
				field.playground[player.pos.y][player.pos.x] = ' '
			case 53:
				player.gotCyan = true
				msgStr.PushFront("Found cyan key")
				field.items[i] = item
				field.playground[player.pos.y][player.pos.x] = ' '
			}
		}

	}
}

func enemyTurn(dungeon *dungeon_t, field *map_t, player *player_t, msgStr *list.List) {
	for i := 0; i < dungeon.room_cnt; i++ {
		enemies_cnt := dungeon.sequence[i].entities_cnt

		for j := 0; j < enemies_cnt; j++ {
			if dungeon.sequence[i].entities[j].typeE == ENEMY {
				enemy := dungeon.sequence[i].entities[j]
				distY, distX := checkAgr(&enemy, player, field)
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
					if distY+distX < 2 {
						enemyAttack(&enemy, player, msgStr)
						dungeon.sequence[i].entities[j] = enemy
					} else {
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

	if field.playground[dirY][dirX] == ' ' || field.playground[dirY][dirX] == '+' {
		field.playground[enemy.pos.y][enemy.pos.x] = ' '
		enemy.pos.y = dirY
		enemy.pos.x = dirX
		field.playground[enemy.pos.y][enemy.pos.x] = enemy.symbol
		empty = true
	} else if field.playground[dirY][dirX] == field.playground[player.pos.y][player.pos.x] {

		empty = true
	}
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
		if field.playground[dirY][dirX] == ' ' || field.playground[dirY][dirX] == '+' {
			field.playground[enemy.pos.y][enemy.pos.x] = ' '
			enemy.pos.y = dirY
			enemy.pos.x = dirX
			field.playground[enemy.pos.y][enemy.pos.x] = enemy.symbol
			empty = true
		} else if field.playground[dirY][dirX] == field.playground[player.pos.y][player.pos.x] {
			empty = true
		}
	}
}

func enemyAttack(enemy *entity_t, player *player_t, msgStr *list.List) {
	if enemy.symbol == 'O' {
		if !enemy.stats.isFirst {
			dodge := player.agility - enemy.stats.agility - rand.Int()%enemy.stats.agility
			if dodge < 1 {
				player.health -= enemy.stats.strength
				player.strikesToPlayer++
				msg := fmt.Sprintf("Ogre hits you by %d.", enemy.stats.strength)
				msgStr.PushFront(msg)
			} else {
				msgStr.PushFront("You dodge Ogre's attack.")
			}
			enemy.stats.isFirst = true
		} else {
			enemy.stats.isFirst = false
		}
	} else {
		dodge := player.agility - enemy.stats.agility - rand.Int()%enemy.stats.agility
		if dodge < 1 {
			player.health -= enemy.stats.strength
			player.strikesToPlayer++
			msg := fmt.Sprintf("Enemy hits you by %d.", enemy.stats.strength)
			msgStr.PushFront(msg)
			if enemy.symbol == 's' {
				if rand.Int()%4 == 0 {
					player.isSleeped = true
					msgStr.PushFront("Snake-Mage sleeps you.")
				}
			} else if enemy.symbol == 'v' {
				if rand.Int()%2 == 0 {
					stealMaxHealth := rand.Int()%enemy.stats.strength + 1
					player.maxHealth -= stealMaxHealth
					enemy.stats.hpStealed += stealMaxHealth
				}
			}
		} else {
			msgStr.PushFront("You dodge the attack.")
		}
	}
}

func checkAgr(enemy *entity_t, player *player_t, field *map_t) (int, int) {

	distY := int(math.Abs(float64(enemy.pos.y - player.pos.y)))
	distX := int(math.Abs(float64(enemy.pos.x - player.pos.x)))

	if distY < enemy.stats.aggression && distX < enemy.stats.aggression {
		enemyView := ComputeFOV(field, enemy.pos, enemy.stats.aggression)
		if enemyView[player.pos] {
			enemy.stats.isChasing = true
		}
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
				}
			}
		}
	}
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
