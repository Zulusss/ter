package main

import (
	"container/list"
	"fmt"
	"math/rand"
	// "github.com/gbin/goncurses"
)

func init_player(player *player_t, spawn_point *entity_t) {
	var position position_t
	position.x, position.y = spawn_point.pos.x, spawn_point.pos.y
	player.pos = position
}

func initInventory(player *player_t) {
	player.inventory.weapon = make([]int, 0, 9)
	player.inventory.scroll = make([]int, 0, 9)
	player.inventory.food = make([]int, 0, 9)
	player.inventory.potion = make([]int, 0, 9)
	player.strength = rand.Int()%3 + 1
	player.agility = rand.Int()%4 + 2 + 2
	player.health = rand.Int()%10 + 3 + 7
	player.maxHealth = player.health
	var difficulty difficulty_t
	difficulty.hp = player.health
	player.diff = &difficulty
}

func useFood(player *player_t) {
	var fList list.List
	fList.Init()

	for i := range player.inventory.food {
		switch player.inventory.food[i] {
		case 10:
			str := fmt.Sprintf("%d. Apple", i+1)
			fList.PushBack(str)
		case 11:
			str := fmt.Sprintf("%d. Banana", i+1)
			fList.PushBack(str)
		case 12:
			str := fmt.Sprintf("%d. Cucumber", i+1)
			fList.PushBack(str)
		}
	}
	printFood(&fList)
	key := getInput()
	num := int(key - 48)
	if num > 0 && num <= len(player.inventory.food) {
		switch player.inventory.food[num-1] {
		case 10:
			player.health += 5
			if player.health > player.maxHealth {
				player.health = player.maxHealth
			}
			player.foodConsumed++
			player.inventory.food = append(player.inventory.food[:num-1], player.inventory.food[num:]...)
		case 11:
			player.health += 10
			if player.health > player.maxHealth {
				player.health = player.maxHealth
			}
			player.foodConsumed++
			player.inventory.food = append(player.inventory.food[:num-1], player.inventory.food[num:]...)
		case 12:
			player.health += 15
			if player.health > player.maxHealth {
				player.health = player.maxHealth
			}
			player.foodConsumed++
			player.inventory.food = append(player.inventory.food[:num-1], player.inventory.food[num:]...)
		}
	}
}

func useScroll(player *player_t) {
	var sList list.List
	sList.Init()
	for i := range player.inventory.scroll {
		switch player.inventory.scroll[i] {
		case 20:
			str := fmt.Sprintf("%d. Scroll of Strength", i+1)
			sList.PushBack(str)
		case 21:
			str := fmt.Sprintf("%d. Scroll of Agility", i+1)
			sList.PushBack(str)
		case 22:
			str := fmt.Sprintf("%d. Scroll of Max Health", i+1)
			sList.PushBack(str)
		}
	}
	printScrolls(&sList)
	key := getInput()
	num := int(key - 48)
	if num > 0 && num <= len(player.inventory.scroll) {
		switch player.inventory.scroll[num-1] {
		case 20:
			player.strength += 5
			player.scrollsRead++
			player.inventory.scroll = append(player.inventory.scroll[:num-1], player.inventory.scroll[num:]...)
		case 21:
			player.agility += 5
			player.scrollsRead++
			player.inventory.scroll = append(player.inventory.scroll[:num-1], player.inventory.scroll[num:]...)
		case 22:
			player.maxHealth += 5
			player.health += 5
			player.scrollsRead++
			player.inventory.scroll = append(player.inventory.scroll[:num-1], player.inventory.scroll[num:]...)
		}
	}
}

func usePotion(player *player_t) {
	var pList list.List
	pList.Init()
	for i := range player.inventory.potion {
		switch player.inventory.potion[i] {
		case 40:
			str := fmt.Sprintf("%d. Potion of Strength", i+1)
			pList.PushBack(str)
		case 41:
			str := fmt.Sprintf("%d. Potion of Agility", i+1)
			pList.PushBack(str)
		case 42:
			str := fmt.Sprintf("%d. Potion of Max Health", i+1)
			pList.PushBack(str)
		}
	}
	printPotions(&pList)
	key := getInput()
	num := int(key - 48)
	if num > 0 && num <= len(player.inventory.potion) {
		switch player.inventory.potion[num-1] {
		case 40:
			if player.potStrenght == 0 {
				player.strength += 5
				player.potionsConsumed++
			}
			player.potStrenght = 10
			player.inventory.potion = append(player.inventory.potion[:num-1], player.inventory.potion[num:]...)
		case 41:
			if player.potAgility == 0 {
				player.agility += 5
				player.potionsConsumed++
			}
			player.potAgility = 10
			player.inventory.potion = append(player.inventory.potion[:num-1], player.inventory.potion[num:]...)
		case 42:
			if player.potMaxHealth == 0 {
				player.maxHealth += 5
				player.potionsConsumed++
			}
			player.potMaxHealth = 10
			player.inventory.potion = append(player.inventory.potion[:num-1], player.inventory.potion[num:]...)
		}
	}
}

func checkPotion(player *player_t) {
	if player.potStrenght > 0 {
		if player.potStrenght == 1 {
			player.strength -= 5
		}
		player.potStrenght--
	}
	if player.potAgility > 0 {
		if player.potAgility == 1 {
			player.agility -= 5
		}
		player.potAgility--
	}
	if player.potMaxHealth > 0 {
		if player.potMaxHealth == 1 {
			player.maxHealth -= 5
			if player.health < 1 {
				player.health = 1
			}
		}
		player.potMaxHealth--
	}
}

func useWeapon(player *player_t, field *map_t) {
	var wList list.List
	wList.Init()
	for i := range player.inventory.weapon {
		switch player.inventory.weapon[i] {
		case 30:
			str := fmt.Sprintf("%d. Sword", i+1)
			wList.PushBack(str)
		case 31:
			str := fmt.Sprintf("%d. Axe", i+1)
			wList.PushBack(str)
		case 32:
			str := fmt.Sprintf("%d. Hammer", i+1)
			wList.PushBack(str)
		}
	}
	key := printWeapon(&wList)

	num := int(key - 48)
	if num >= 0 && num <= len(player.inventory.weapon) {
		if num == 0 {
			player.weapon = 0
		} else if player.inventory.weapon[num-1] == 30 {
			if player.weapon != 0 {
				var oldWeapon entity_t
				oldWeapon.symbol = '/'
				oldWeapon.typeE = ITEM
				oldWeapon.typeI = player.weapon
				if field.playground[player.pos.y+1][player.pos.x] == ' ' {
					oldWeapon.pos.y = player.pos.y + 1
					oldWeapon.pos.x = player.pos.x
				} else if field.playground[player.pos.y-1][player.pos.x] == ' ' {
					oldWeapon.pos.y = player.pos.y - 1
					oldWeapon.pos.x = player.pos.x
				} else if field.playground[player.pos.y][player.pos.x+1] == ' ' {
					oldWeapon.pos.y = player.pos.y
					oldWeapon.pos.x = player.pos.x + 1
				} else if field.playground[player.pos.y][player.pos.x-1] == ' ' {
					oldWeapon.pos.y = player.pos.y
					oldWeapon.pos.x = player.pos.x - 1
				}
				for i := 0; i < len(player.inventory.weapon); i++ {
					if player.inventory.weapon[i] == player.weapon {
						player.inventory.weapon = append(player.inventory.weapon[:i], player.inventory.weapon[i+1:]...)
						break
					}
				}
				field.playground[oldWeapon.pos.y][oldWeapon.pos.x] = oldWeapon.symbol
				field.items_cnt += 1
				field.items[field.items_cnt] = oldWeapon
			}
			player.weapon = 30
		} else if player.inventory.weapon[num-1] == 31 {
			if player.weapon != 0 {
				var oldWeapon entity_t
				oldWeapon.symbol = '/'
				oldWeapon.typeE = ITEM
				oldWeapon.typeI = player.weapon
				if field.playground[player.pos.y+1][player.pos.x] == ' ' {
					oldWeapon.pos.y = player.pos.y + 1
					oldWeapon.pos.x = player.pos.x
				} else if field.playground[player.pos.y-1][player.pos.x] == ' ' {
					oldWeapon.pos.y = player.pos.y - 1
					oldWeapon.pos.x = player.pos.x
				} else if field.playground[player.pos.y][player.pos.x+1] == ' ' {
					oldWeapon.pos.y = player.pos.y
					oldWeapon.pos.x = player.pos.x + 1
				} else if field.playground[player.pos.y][player.pos.x-1] == ' ' {
					oldWeapon.pos.y = player.pos.y
					oldWeapon.pos.x = player.pos.x - 1
				}
				for i := 0; i < len(player.inventory.weapon); i++ {
					if player.inventory.weapon[i] == player.weapon {
						player.inventory.weapon = append(player.inventory.weapon[:i], player.inventory.weapon[i+1:]...)
						break
					}
				}
				field.playground[oldWeapon.pos.y][oldWeapon.pos.x] = oldWeapon.symbol
				field.items_cnt += 1
				field.items[field.items_cnt] = oldWeapon
			}
			player.weapon = 31
		} else if player.inventory.weapon[num-1] == 32 {
			if player.weapon != 0 {
				var oldWeapon entity_t
				oldWeapon.symbol = '/'
				oldWeapon.typeE = ITEM
				oldWeapon.typeI = player.weapon
				if field.playground[player.pos.y+1][player.pos.x] == ' ' {
					oldWeapon.pos.y = player.pos.y + 1
					oldWeapon.pos.x = player.pos.x
				} else if field.playground[player.pos.y-1][player.pos.x] == ' ' {
					oldWeapon.pos.y = player.pos.y - 1
					oldWeapon.pos.x = player.pos.x
				} else if field.playground[player.pos.y][player.pos.x+1] == ' ' {
					oldWeapon.pos.y = player.pos.y
					oldWeapon.pos.x = player.pos.x + 1
				} else if field.playground[player.pos.y][player.pos.x-1] == ' ' {
					oldWeapon.pos.y = player.pos.y
					oldWeapon.pos.x = player.pos.x - 1
				}
				for i := 0; i < len(player.inventory.weapon); i++ {
					if player.inventory.weapon[i] == player.weapon {
						player.inventory.weapon = append(player.inventory.weapon[:i], player.inventory.weapon[i+1:]...)
						break
					}
				}
				field.playground[oldWeapon.pos.y][oldWeapon.pos.x] = oldWeapon.symbol
				field.items_cnt += 1
				field.items[field.items_cnt] = oldWeapon
			}

			player.weapon = 32
		}
	}
}

func player_movement(dungeon *dungeon_t, field *map_t, player *player_t, key int, msgStr *list.List) {
	if !player.isSleeped {
		var redraw bool
		redrawItems(field)
		if key == 'w' || key == 259 { //goncurses.KEY_UP
			player.pos.y -= MOVEMENT_STEP
			checkTile(dungeon, field, player, msgStr)

			switch field.playground[player.pos.y][player.pos.x] {
			case 'z', 'v', 'm', 'O', 'g', 's':
				var pos position_t
				pos.y, pos.x = player.pos.y, player.pos.x
				player.pos.y += MOVEMENT_STEP
				redraw = true
				enemy := findEnemy(dungeon, &pos)
				if enemy != nil {
					playerAttack(enemy, player, msgStr)
					if enemy.stats.hp < 1 {
						player.monsterKill++
						field.playground[enemy.pos.y][enemy.pos.x] = ' '
						monsterGold := rand.Int()%10 + 1 + dungeon.level*3
						player.gold += monsterGold
						enemy.typeE = 5
						player.maxHealth += enemy.stats.hpStealed
						msg := fmt.Sprintf("You kill %c, you got %d gold.", enemy.symbol, monsterGold)
						msgStr.PushFront(msg)
					}
				}
			case WALL_CHAR, OUTER_AREA_CHAR:
				redraw = true
				player.pos.y += MOVEMENT_STEP
			case 'x':
				var msg string
				msg, redraw = checkDoor(dungeon, player)
				msgStr.PushFront(msg)
				if redraw {
					player.pos.y += MOVEMENT_STEP
				}
			}
			if !redraw {
				player.turns++
				field.playground[player.pos.y+MOVEMENT_STEP][player.pos.x] = ' '
				field.playground[player.pos.y][player.pos.x] = '@'
			}
		} else if key == 's' || key == 258 { //goncurses.KEY_DOWN
			player.pos.y += MOVEMENT_STEP
			checkTile(dungeon, field, player, msgStr)

			switch field.playground[player.pos.y][player.pos.x] {
			case 'z', 'v', 'm', 'O', 'g', 's':
				var pos position_t
				pos.y, pos.x = player.pos.y, player.pos.x
				player.pos.y -= MOVEMENT_STEP
				redraw = true
				enemy := findEnemy(dungeon, &pos)
				if enemy != nil {
					playerAttack(enemy, player, msgStr)
					if enemy.stats.hp < 1 {
						player.monsterKill++
						field.playground[enemy.pos.y][enemy.pos.x] = ' '
						monsterGold := rand.Int()%10 + 1 + dungeon.level*3
						player.gold += monsterGold
						enemy.typeE = 5
						player.maxHealth += enemy.stats.hpStealed
						msg := fmt.Sprintf("You kill %c, you got %d gold", enemy.symbol, monsterGold)
						msgStr.PushFront(msg)
					}
				}
			case WALL_CHAR, OUTER_AREA_CHAR:
				redraw = true
				player.pos.y -= MOVEMENT_STEP
			case 'x':
				var msg string
				msg, redraw = checkDoor(dungeon, player)
				msgStr.PushFront(msg)
				if redraw {
					player.pos.y -= MOVEMENT_STEP
				}
			}
			if !redraw {
				player.turns++
				field.playground[player.pos.y-MOVEMENT_STEP][player.pos.x] = ' '
				field.playground[player.pos.y][player.pos.x] = '@'
			}
		} else if key == 'a' || key == 260 { //goncurses.KEY_LEFT
			player.pos.x -= MOVEMENT_STEP
			checkTile(dungeon, field, player, msgStr)

			switch field.playground[player.pos.y][player.pos.x] {
			case 'z', 'v', 'm', 'O', 'g', 's':
				var pos position_t
				pos.y, pos.x = player.pos.y, player.pos.x
				player.pos.x += MOVEMENT_STEP
				redraw = true
				enemy := findEnemy(dungeon, &pos)
				if enemy != nil {
					playerAttack(enemy, player, msgStr)
					if enemy.stats.hp < 1 {
						player.monsterKill++
						field.playground[enemy.pos.y][enemy.pos.x] = ' '
						monsterGold := rand.Int()%10 + 1 + dungeon.level*3
						player.gold += monsterGold
						enemy.typeE = 5
						player.maxHealth += enemy.stats.hpStealed
						msg := fmt.Sprintf("You kill %c, you got %d gold", enemy.symbol, monsterGold)
						msgStr.PushFront(msg)
					}
				}
			case WALL_CHAR, OUTER_AREA_CHAR:
				redraw = true
				player.pos.x += MOVEMENT_STEP
			case 'x':
				var msg string
				msg, redraw = checkDoor(dungeon, player)
				msgStr.PushFront(msg)
				if redraw {
					player.pos.x += MOVEMENT_STEP
				}
			}
			if !redraw {
				player.turns++
				field.playground[player.pos.y][player.pos.x+MOVEMENT_STEP] = ' '
				field.playground[player.pos.y][player.pos.x] = '@'
			}
		} else if key == 'd' || key == 261 { //goncurses.KEY_RIGHT
			player.pos.x += MOVEMENT_STEP
			checkTile(dungeon, field, player, msgStr)
			switch field.playground[player.pos.y][player.pos.x] {
			case 'z', 'v', 'm', 'O', 'g', 's':
				var pos position_t
				pos.y, pos.x = player.pos.y, player.pos.x
				player.pos.x -= MOVEMENT_STEP
				redraw = true
				enemy := findEnemy(dungeon, &pos)
				if enemy != nil {
					playerAttack(enemy, player, msgStr)
					if enemy.stats.hp < 1 {
						player.monsterKill++
						field.playground[enemy.pos.y][enemy.pos.x] = ' '
						monsterGold := rand.Int()%10 + 1 + dungeon.level*3
						player.gold += monsterGold
						enemy.typeE = 5
						player.maxHealth += enemy.stats.hpStealed
						msg := fmt.Sprintf("You kill %c, you got %d gold", enemy.symbol, monsterGold)
						msgStr.PushFront(msg)
					}
				}
			case WALL_CHAR, OUTER_AREA_CHAR:
				redraw = true
				player.pos.x -= MOVEMENT_STEP
			case 'x':
				var msg string
				msg, redraw = checkDoor(dungeon, player)
				msgStr.PushFront(msg)
				if redraw {
					player.pos.x -= MOVEMENT_STEP
				}
			}
			if !redraw {
				player.turns++
				field.playground[player.pos.y][player.pos.x-MOVEMENT_STEP] = ' '
				field.playground[player.pos.y][player.pos.x] = '@'
			}
		} else if key == 'j' {
			useFood(player)
		} else if key == 'e' {
			useScroll(player)
		} else if key == 'k' {
			usePotion(player)
		} else if key == 'h' {
			useWeapon(player, field)
		}
	} else {
		player.isSleeped = false
	}
	if msgStr.Len() > 2 {
		msgStr.Remove(msgStr.Back())
	}
}

func IsDead(player *player_t) bool {
	return player.health <= 0
}

func TheEnd(dungeon *dungeon_t, player *player_t) {
	recordAttempt(dungeon, player)
}

func redrawItems(field *map_t) {
	for i := 0; i < field.items_cnt; i++ {
		if field.items[i].typeE == ITEM {
			field.playground[field.items[i].pos.y][field.items[i].pos.x] = field.items[i].symbol
		}
	}
}
func playerAttack(enemy *entity_t, player *player_t, msgStr *list.List) {
	var attack int
	switch player.weapon {
	case 30:
		attack = player.strength + player.strength/10 + 1
	case 31:
		attack = player.strength + player.strength/10*2 + 2
	case 32:
		attack = player.strength + player.strength/10*3 + 3
	default:
		attack = player.strength
	}
	switch enemy.symbol {
	case 'v':
		if !enemy.stats.isFirst {
			enemy.stats.isFirst = true
			msgStr.PushFront("Vampire dodge the attack.")
		} else {
			dodge := enemy.stats.agility - player.agility - rand.Int()%player.agility

			if dodge < 1 {
				enemy.stats.hp -= attack
				player.strikesToEnemy++
				msg := fmt.Sprintf("You hit vampire by %d.", attack)
				msgStr.PushFront(msg)
			} else {
				msgStr.PushFront("Vampire dodge the attack.")
			}
		}

	default:
		dodge := enemy.stats.agility - player.agility - rand.Int()%player.agility
		if dodge < 1 {
			enemy.stats.hp -= attack
			player.strikesToEnemy++
			msg := fmt.Sprintf("You hit enemy by %d.", attack)
			msgStr.PushFront(msg)
		} else {
			msgStr.PushFront("Enemy dodge the attack.")
		}
	}
}

func findEnemy(dungeon *dungeon_t, pos *position_t) (ptrEnemy *entity_t) {
	for i := 0; i < dungeon.room_cnt; i++ {
		enemies_cnt := dungeon.sequence[i].entities_cnt

		for j := 0; j < enemies_cnt; j++ {
			if dungeon.sequence[i].entities[j].typeE == ENEMY {
				if dungeon.sequence[i].entities[j].pos.y == pos.y &&
					dungeon.sequence[i].entities[j].pos.x == pos.x {
					return &dungeon.sequence[i].entities[j]
				}
			}
		}
	}

	return
}

func checkDoor(dungeon *dungeon_t, player *player_t) (string, bool) {
	if player.pos.y == dungeon.lockedDoors[0].y && player.pos.x == dungeon.lockedDoors[0].x {
		if player.gotBlue {
			msg := "You open blue door."
			return msg, false
		} else {
			msg := "Need blue key."
			return msg, true
		}
	} else if player.pos.y == dungeon.lockedDoors[1].y && player.pos.x == dungeon.lockedDoors[1].x {
		if player.gotMagenta {
			msg := "You open magenta door."
			return msg, false
		} else {
			msg := "Need magenta key."
			return msg, true
		}
	} else if player.pos.y == dungeon.lockedDoors[2].y && player.pos.x == dungeon.lockedDoors[2].x {
		if player.gotCyan {
			msg := "You open cyan door."
			return msg, false
		} else {
			msg := "Need cyan key."
			return msg, true
		}
	}
	return "fake", false
}
