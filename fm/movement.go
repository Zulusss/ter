package main

import (
	"container/list"
	"fmt"
	"math/rand"

	"github.com/gbin/goncurses"
)

func init_player(player *player_t, spawn_point *entity_t) {
	var position position_t
	position.x, position.y = spawn_point.pos.x, spawn_point.pos.y
	player.pos = position
	// spawn_point.pos = player.pos
}

func initInventory(player *player_t) {
	player.inventory.weapon = make([]int, 0, 9)
	player.inventory.scroll = make([]int, 0, 9)
	player.inventory.food = make([]int, 0, 9)
	player.inventory.potion = make([]int, 0, 9)
	player.strength = rand.Int()%3 + 1
	player.agility = rand.Int()%4 + 2
	player.health = rand.Int()%10 + 3
	player.maxHealth = player.health
}

func useFood(player *player_t) {
	goncurses.StdScr().Clear()
	goncurses.StdScr().MovePrint(0, 0, "Which food do you want to consume? Enter number 1-9:\n")
	for i := range player.inventory.food {
		if player.inventory.food[i] == 10 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Apple", i+1)
		} else if player.inventory.food[i] == 11 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Banana", i+1)
		} else if player.inventory.food[i] == 12 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Cucumber", i+1)
		}
	}
	key := goncurses.StdScr().GetChar()
	num := int(key - 48)
	if num > 0 && num <= len(player.inventory.food) {
		if player.inventory.food[num-1] == 10 {
			player.health += 5
			player.foodConsumed++
			player.inventory.food = append(player.inventory.food[:num-1], player.inventory.food[num:]...)
		} else if player.inventory.food[num-1] == 11 {
			player.health += 10
			player.foodConsumed++
			player.inventory.food = append(player.inventory.food[:num-1], player.inventory.food[num:]...)
		} else if player.inventory.food[num-1] == 12 {
			player.health += 15
			player.foodConsumed++
			player.inventory.food = append(player.inventory.food[:num-1], player.inventory.food[num:]...)
		}
	}
}

func useScroll(player *player_t) {
	goncurses.StdScr().Clear()
	goncurses.StdScr().MovePrint(0, 0, "Which scroll do you want to use? Enter number 1-9:\n")
	for i := range player.inventory.scroll {
		if player.inventory.scroll[i] == 20 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Scroll of Strength", i+1)
		} else if player.inventory.scroll[i] == 21 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Scroll of Agility", i+1)
		} else if player.inventory.scroll[i] == 22 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Scroll of Max Health", i+1)
		}
	}
	key := goncurses.StdScr().GetChar()
	num := int(key - 48)
	if num > 0 && num <= len(player.inventory.scroll) {
		if player.inventory.scroll[num-1] == 20 {
			player.strength += 5
			player.scrollsRead++
			player.inventory.scroll = append(player.inventory.scroll[:num-1], player.inventory.scroll[num:]...)
		} else if player.inventory.scroll[num-1] == 21 {
			player.agility += 5
			player.scrollsRead++
			player.inventory.scroll = append(player.inventory.scroll[:num-1], player.inventory.scroll[num:]...)
		} else if player.inventory.scroll[num-1] == 22 {
			player.maxHealth += 5
			player.scrollsRead++
			player.inventory.scroll = append(player.inventory.scroll[:num-1], player.inventory.scroll[num:]...)
		}
	}
}

func usePotion(player *player_t) {
	goncurses.StdScr().Clear()
	goncurses.StdScr().MovePrint(0, 0, "Which potion do you want to drink? Enter number 1-9:\n")
	for i := range player.inventory.potion {
		if player.inventory.potion[i] == 40 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Potion of Strength", i+1)
		} else if player.inventory.potion[i] == 41 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Potion of Agility", i+1)
		} else if player.inventory.potion[i] == 42 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Potion of Max Health", i+1)
		}
	}
	key := goncurses.StdScr().GetChar()
	num := int(key - 48)
	if num > 0 && num <= len(player.inventory.potion) {
		if player.inventory.potion[num-1] == 40 {
			if player.potStrenght == 0 {
				player.strength += 5
				player.potionsConsumed++
			}
			player.potStrenght = 10
			player.inventory.potion = append(player.inventory.potion[:num-1], player.inventory.potion[num:]...)
		} else if player.inventory.potion[num-1] == 41 {
			if player.potAgility == 0 {
				player.agility += 5
				player.potionsConsumed++
			}
			player.potAgility = 10
			player.inventory.potion = append(player.inventory.potion[:num-1], player.inventory.potion[num:]...)
		} else if player.inventory.potion[num-1] == 42 {
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
		}
		player.potMaxHealth--
	}
}

func useWeapon(player *player_t, field *map_t) {
	goncurses.StdScr().Clear()
	goncurses.StdScr().MovePrint(0, 0, "Which weapon do you want to use? 0 - bare hands. Enter number 0-9:\n")
	for i := range player.inventory.weapon {
		if player.inventory.weapon[i] == 30 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Sword", i+1)
		} else if player.inventory.weapon[i] == 31 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Axe", i+1)
		} else if player.inventory.weapon[i] == 32 {
			goncurses.StdScr().MovePrintf(i+1, 0, "%d. Hammer", i+1)
		}
	}
	key := goncurses.StdScr().GetChar()
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
			// player.inventory.weapon = append(player.inventory.weapon[:num-1], player.inventory.weapon[num:]...)
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
			// player.inventory.weapon = append(player.inventory.weapon[:num-1], player.inventory.weapon[num:]...)
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
			// player.inventory.weapon = append(player.inventory.weapon[:num-1], player.inventory.weapon[num:]...)
		}
	}
}

func player_movement(dungeon *dungeon_t, field *map_t, player *player_t, key int, msgStr *list.List) {
	if !player.isSleeped {
		var redraw bool
		redrawItems(field)
		if key == 'w' || key == goncurses.KEY_UP {
			player.pos.y -= MOVEMENT_STEP
			checkTile(dungeon, field, player, msgStr)

			if field.playground[player.pos.y][player.pos.x] == 'z' ||
				field.playground[player.pos.y][player.pos.x] == 'v' ||
				field.playground[player.pos.y][player.pos.x] == 'm' ||
				field.playground[player.pos.y][player.pos.x] == 'O' ||
				field.playground[player.pos.y][player.pos.x] == 'g' ||
				field.playground[player.pos.y][player.pos.x] == 's' {
				var pos position_t
				pos.y, pos.x = player.pos.y, player.pos.x
				player.pos.y += MOVEMENT_STEP
				redraw = true
				enemy := findEnemy(dungeon, &pos)
				if enemy != nil {
					playerAttack(enemy, player)
					if enemy.stats.hp < 1 {
						player.monsterKill++
						field.playground[enemy.pos.y][enemy.pos.x] = ' '
						monsterGold := rand.Int()%10 + 1 + dungeon.level*3
						player.gold += monsterGold
						enemy.typeE = 5
						player.maxHealth += enemy.stats.hpStealed
						msg := fmt.Sprintf("You kill %c, you got %d gold.", enemy.symbol, monsterGold)
						msgStr.PushFront(msg)
						if msgStr.Len() > 2 {
							msgStr.Remove(msgStr.Back())
						}
					}
				}
			} else if field.playground[player.pos.y][player.pos.x] == WALL_CHAR ||
				field.playground[player.pos.y][player.pos.x] == OUTER_AREA_CHAR {
				// player->pos.x -= sinf(player->view_angle) * MOVEMENT_STEP;
				redraw = true
				player.pos.y += MOVEMENT_STEP
			} else if field.playground[player.pos.y][player.pos.x] == 'x' {
				var msg string
				msg, redraw = checkDoor(dungeon, player)
				msgStr.PushFront(msg)
				if msgStr.Len() > 2 {
					msgStr.Remove(msgStr.Back())
				}
				if redraw {
					player.pos.y += MOVEMENT_STEP
				}
			}
			if !redraw {
				player.turns++
				// ch := field.playground[player.pos.y+MOVEMENT_STEP][player.pos.x]
				// if ch == '$' || ch == '*' || ch == '?' || ch == '/' || ch == '^' {
				// 	field.playground[player.pos.y][player.pos.x] = '@'
				// } else {
				field.playground[player.pos.y+MOVEMENT_STEP][player.pos.x] = ' '
				field.playground[player.pos.y][player.pos.x] = '@'
				// }
			}
		} else if key == 's' || key == goncurses.KEY_DOWN {
			// player->pos.x -= sinf(player->view_angle) * MOVEMENT_STEP;
			player.pos.y += MOVEMENT_STEP
			checkTile(dungeon, field, player, msgStr)

			if field.playground[player.pos.y][player.pos.x] == 'z' ||
				field.playground[player.pos.y][player.pos.x] == 'v' ||
				field.playground[player.pos.y][player.pos.x] == 'm' ||
				field.playground[player.pos.y][player.pos.x] == 'O' ||
				field.playground[player.pos.y][player.pos.x] == 'g' ||
				field.playground[player.pos.y][player.pos.x] == 's' {
				var pos position_t
				pos.y, pos.x = player.pos.y, player.pos.x
				player.pos.y -= MOVEMENT_STEP
				redraw = true
				enemy := findEnemy(dungeon, &pos)
				if enemy != nil {
					playerAttack(enemy, player)
					if enemy.stats.hp < 1 {
						player.monsterKill++
						field.playground[enemy.pos.y][enemy.pos.x] = ' '
						monsterGold := rand.Int()%10 + 1 + dungeon.level*3
						player.gold += monsterGold
						enemy.typeE = 5
						player.maxHealth += enemy.stats.hpStealed
						msg := fmt.Sprintf("You kill %c, you got %d gold", enemy.symbol, monsterGold)
						msgStr.PushFront(msg)
						if msgStr.Len() > 2 {
							msgStr.Remove(msgStr.Back())
						}
					}
				}
			} else if field.playground[player.pos.y][player.pos.x] == WALL_CHAR ||
				field.playground[player.pos.y][player.pos.x] == OUTER_AREA_CHAR {
				// player->pos.x += sinf(player->view_angle) * MOVEMENT_STEP;
				redraw = true
				player.pos.y -= MOVEMENT_STEP
			} else if field.playground[player.pos.y][player.pos.x] == 'x' {
				var msg string
				msg, redraw = checkDoor(dungeon, player)
				msgStr.PushFront(msg)
				if msgStr.Len() > 2 {
					msgStr.Remove(msgStr.Back())
				}
				if redraw {
					player.pos.y -= MOVEMENT_STEP
				}
			}
			if !redraw {
				player.turns++
				// ch := field.playground[player.pos.y-MOVEMENT_STEP][player.pos.x]
				// if ch == '$' || ch == '*' || ch == '?' || ch == '/' || ch == '^' {
				// 	field.playground[player.pos.y][player.pos.x] = '@'
				// } else {
				field.playground[player.pos.y-MOVEMENT_STEP][player.pos.x] = ' '
				field.playground[player.pos.y][player.pos.x] = '@'
				// }
			}
		} else if key == 'a' || key == goncurses.KEY_LEFT {
			player.pos.x -= MOVEMENT_STEP
			checkTile(dungeon, field, player, msgStr)

			if field.playground[player.pos.y][player.pos.x] == 'z' ||
				field.playground[player.pos.y][player.pos.x] == 'v' ||
				field.playground[player.pos.y][player.pos.x] == 'm' ||
				field.playground[player.pos.y][player.pos.x] == 'O' ||
				field.playground[player.pos.y][player.pos.x] == 'g' ||
				field.playground[player.pos.y][player.pos.x] == 's' {
				var pos position_t
				pos.y, pos.x = player.pos.y, player.pos.x
				player.pos.x += MOVEMENT_STEP
				redraw = true
				enemy := findEnemy(dungeon, &pos)
				if enemy != nil {
					playerAttack(enemy, player)
					if enemy.stats.hp < 1 {
						player.monsterKill++
						field.playground[enemy.pos.y][enemy.pos.x] = ' '
						monsterGold := rand.Int()%10 + 1 + dungeon.level*3
						player.gold += monsterGold
						enemy.typeE = 5
						player.maxHealth += enemy.stats.hpStealed
						msg := fmt.Sprintf("You kill %c, you got %d gold", enemy.symbol, monsterGold)
						msgStr.PushFront(msg)
						if msgStr.Len() > 2 {
							msgStr.Remove(msgStr.Back())
						}
					}
				}
			} else if field.playground[player.pos.y][player.pos.x] == WALL_CHAR ||
				field.playground[player.pos.y][player.pos.x] == OUTER_AREA_CHAR {
				redraw = true
				player.pos.x += MOVEMENT_STEP
			} else if field.playground[player.pos.y][player.pos.x] == 'x' {
				var msg string
				msg, redraw = checkDoor(dungeon, player)
				msgStr.PushFront(msg)
				if msgStr.Len() > 2 {
					msgStr.Remove(msgStr.Back())
				}
				if redraw {
					player.pos.x += MOVEMENT_STEP
				}
			}
			if !redraw {
				player.turns++
				// ch := field.playground[player.pos.y][player.pos.x+MOVEMENT_STEP]
				// if ch == '$' || ch == '*' || ch == '?' || ch == '/' || ch == '^' {
				// 	field.playground[player.pos.y][player.pos.x] = '@'
				// } else {
				field.playground[player.pos.y][player.pos.x+MOVEMENT_STEP] = ' '
				field.playground[player.pos.y][player.pos.x] = '@'
				// }
			}
		} else if key == 'd' || key == goncurses.KEY_RIGHT {
			player.pos.x += MOVEMENT_STEP
			checkTile(dungeon, field, player, msgStr)
			if field.playground[player.pos.y][player.pos.x] == 'z' ||
				field.playground[player.pos.y][player.pos.x] == 'v' ||
				field.playground[player.pos.y][player.pos.x] == 'm' ||
				field.playground[player.pos.y][player.pos.x] == 'O' ||
				field.playground[player.pos.y][player.pos.x] == 'g' ||
				field.playground[player.pos.y][player.pos.x] == 's' {
				var pos position_t
				pos.y, pos.x = player.pos.y, player.pos.x
				player.pos.x -= MOVEMENT_STEP
				redraw = true
				enemy := findEnemy(dungeon, &pos)
				if enemy != nil {
					playerAttack(enemy, player)
					if enemy.stats.hp < 1 {
						player.monsterKill++
						field.playground[enemy.pos.y][enemy.pos.x] = ' '
						monsterGold := rand.Int()%10 + 1 + dungeon.level*3
						player.gold += monsterGold
						enemy.typeE = 5
						player.maxHealth += enemy.stats.hpStealed
						msg := fmt.Sprintf("You kill %c, you got %d gold", enemy.symbol, monsterGold)
						msgStr.PushFront(msg)
						if msgStr.Len() > 2 {
							msgStr.Remove(msgStr.Back())
						}
					}
				}
			} else if field.playground[player.pos.y][player.pos.x] == WALL_CHAR ||
				field.playground[player.pos.y][player.pos.x] == OUTER_AREA_CHAR {
				redraw = true
				player.pos.x -= MOVEMENT_STEP
			} else if field.playground[player.pos.y][player.pos.x] == 'x' {
				var msg string
				msg, redraw = checkDoor(dungeon, player)
				msgStr.PushFront(msg)
				if msgStr.Len() > 2 {
					msgStr.Remove(msgStr.Back())
				}
				if redraw {
					player.pos.x -= MOVEMENT_STEP
				}
			}
			if !redraw {
				player.turns++
				// ch := field.playground[player.pos.y][player.pos.x-MOVEMENT_STEP]
				// if ch == '$' || ch == '*' || ch == '?' || ch == '/' || ch == '^' {
				// 	field.playground[player.pos.y][player.pos.x] = '@'
				// } else {
				field.playground[player.pos.y][player.pos.x-MOVEMENT_STEP] = ' '
				field.playground[player.pos.y][player.pos.x] = '@'
				// }
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

		// redrawItems(field)
		// if player.turns > 0 {
		// 	field.playground[field.player_spawn.pos.y][field.player_spawn.pos.x] = ' '
		// 	// field.player_spawn.pos.y
		// }
		// if field.playground[player.pos.y][player.pos.x] == EXIT_CHAR {
		// 	// generate_dungeon(dungeon)
		// }
	} else {
		player.isSleeped = false
	}
	if msgStr.Len() > 2 {
		msgStr.Remove(msgStr.Back())
	}
}

func redrawItems(field *map_t) {
	for i := 0; i < field.items_cnt; i++ {
		if field.items[i].typeE == ITEM {
			field.playground[field.items[i].pos.y][field.items[i].pos.x] = field.items[i].symbol
		}
	}
}
func playerAttack(enemy *entity_t, player *player_t) {
	switch enemy.symbol {
	case 'v':
		if !enemy.stats.isFirst {
			enemy.stats.isFirst = true
		} else {
			dodge := enemy.stats.agility - player.agility - rand.Int()%player.agility
			if dodge < 1 {
				enemy.stats.hp -= player.strength
				player.strikesToEnemy++
			}
		}

	default:
		dodge := enemy.stats.agility - player.agility - rand.Int()%player.agility
		if dodge < 1 {
			enemy.stats.hp -= player.strength
			player.strikesToEnemy++
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
					// break
					// enemy := dungeon.sequence[i].entities[j]
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
	return "fuck", false
}

// func firstMove(dungeon *dungeon_t, field *map_t, player *player_t) {
// 	key := int(goncurses.StdScr().GetChar())
// 	player_movement(dungeon, field, player, key)
// }
