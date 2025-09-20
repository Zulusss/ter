package main

import (
	"math"
	"math/rand"
	"strings"

	"github.com/gbin/goncurses"
)

var py, px int // @ coords
var sy, sx int // > coords
var race string
var att int
var hp int
var mana int
var stealth int
var t_placed bool
var p_placed bool
var r_placed int
var dlvl int
var turns int
var lvl_turns int
var m_defeated int
var state string // conf, pois, blee etc

type monsters struct {
	y      int
	x      int
	lvl    int
	typeGo int
	awake  bool
}

var monster [20]monsters

func DungeonDraw(rows int, cols int, field [][]byte, obj [][]byte) int {

	goncurses.StdScr().Move(1, 0)
	goncurses.StdScr().ClearToBottom()

	for y := 0; y <= rows; y++ {
		for x := 0; x <= cols; x++ {
			if y == rows {
				goncurses.StdScr().MoveAddChar(y, x, goncurses.Char(' '))
			} else if field[y][x] == ' ' {
				if obj[y][x] == '>' && (lvl_turns > 420+dlvl ||
					((y > py-5 && x > px-5) && (y < py+5 && x < px+5))) {
					goncurses.StdScr().AttrOn(goncurses.A_BOLD)
					goncurses.StdScr().MoveAddChar(y, x, '>')
					goncurses.StdScr().AttrOff(goncurses.A_BOLD)
				} else if obj[y][x] == '^' && ((y > py-5 && x > px-5) &&
					(y < py+5 && x < px+5)) {
					goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_YELLOW))
					goncurses.StdScr().MoveAddChar(y, x, '^')
					goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_YELLOW))
				} else {
					goncurses.StdScr().MoveAddChar(y, x, ' ')
				}
			} else if field[y][x] == '%' && y != 0 {
				goncurses.StdScr().AttrOn(goncurses.A_DIM | goncurses.ColorPair(goncurses.C_WHITE))
				goncurses.StdScr().MoveAddChar(y, x, '%')
				goncurses.StdScr().AttrOff(goncurses.A_DIM | goncurses.ColorPair(goncurses.C_WHITE))
			} else if field[y][x] == '#' {
				goncurses.StdScr().MoveAddChar(y, x, '#')
			} else if field[y][x] == '~' {
				goncurses.StdScr().AttrOn(goncurses.ColorPair(goncurses.C_RED))
				goncurses.StdScr().MoveAddChar(y, x, '~')
				goncurses.StdScr().AttrOff(goncurses.ColorPair(goncurses.C_RED))
				goncurses.StdScr().Standend()

			} else if field[y][x] == 't' && ((y > py-5 && x > px-5) &&
				(y < py+5 && x < px+5)) {
				for m := 0; m < 10+dlvl/2; m++ {
					if monster[m].y == y && monster[m].x == x {
						if monster[m].lvl < dlvl/2+2 {
							goncurses.StdScr().AttrOn(goncurses.ColorPair(goncurses.C_RED))
							goncurses.StdScr().MoveAddChar(y, x, goncurses.Char(monster[m].typeGo))
						} else if monster[m].lvl < dlvl+2 {
							goncurses.StdScr().AttrOn(goncurses.ColorPair(goncurses.C_YELLOW))
							goncurses.StdScr().MoveAddChar(y, x, goncurses.Char(monster[m].typeGo))

						} else {
							goncurses.StdScr().MoveAddChar(y, x, goncurses.Char(monster[m].typeGo))
						}
					}
					goncurses.StdScr().Standend()
				}
			}
		}
	}

	goncurses.StdScr().AttrOn(goncurses.A_BOLD)
	goncurses.StdScr().MoveAddChar(py, px, '@')
	goncurses.StdScr().AttrOff(goncurses.A_BOLD)

	return 0
}

func Rip(rows int, cols int, killer int) int {
	var c int

	for {
		goncurses.StdScr().Clear()
		goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_RED))
		goncurses.StdScr().MovePrintf(rows/2-3, cols/2-10, "You were captured by %c\n\n\n", killer)
		goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_RED))
		goncurses.StdScr().Printf("\tLevel reached: %d\n\tMonsers defeated: %d\n\tTurns: %d\n\n\tAttack: %d\n\tMana: %d\n\n\n\tPress 'n' to start new game or 'ESC' to exit.", dlvl, m_defeated, turns, att, mana)
		c = int(goncurses.StdScr().GetChar())

		if c == 'n' || c == 27 {
			return c
		}
	}
}

func CheckTrap(rows int, cols int, obj [][]byte) int {
	if obj[py][px] == '^' {
		goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_YELLOW))
		goncurses.StdScr().MovePrint(rows, cols-10, "[Trap]")
		goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_YELLOW))
		if (strings.Compare(race, "Halfling") == 0) && (rand.Int()%2 == 0) {
			goncurses.StdScr().MovePrint(0, 0, " You are Halfling")
		} else if rand.Int()%2 == 0 {
			goncurses.StdScr().MovePrint(0, 0, " You have stepped into a trap... dart hits you!")
			hp -= dlvl/2 + 1
			if hp < 1 {
				return '^'
			}
		} else {
			goncurses.StdScr().MovePrint(0, 0, " You've stepped into a trap... you are confused!")
			state = "conf"
			// state[0] = 'c'
			// state[1] = 'o'
			// state[2] = 'n'
			// state[3] = 'f'
			// state[4] = '\000'
		}
	}
	return 0
}

func MonsterTurn(cols int, field [][]byte) int {
	var dist_y, dist_x int

	for m := 0; m < 10+dlvl/2; m++ {
		if monster[m].lvl < 1 {
			continue
		}
		dist_y = int(math.Abs(float64(monster[m].y) - float64(py)))
		// dist_y = int(math.Abs(float64(dist_y)))
		dist_x = int(math.Abs(float64(monster[m].x) - float64(px)))

		if (rand.Int() % (dlvl + 1)) == 0 {
			goncurses.StdScr().MovePrint(0, 0, " You are Halfilng again!")
		} else if dist_y < dlvl+3-stealth && dist_x < dlvl+3-stealth {
			monster[m].awake = true
		}

		if !monster[m].awake {
			continue
		}
		var dir_y, dir_x int
		dir_y = monster[m].y
		dir_x = monster[m].x

		// test 4-side movement 4 sides
		if dist_y > dist_x {
			if dir_y > py {
				dir_y -= 1
			} else {
				dir_y += 1
			}
		} else {
			if dir_x > px {
				dir_x -= 1
			} else {
				dir_x += 1
			}
		}
		// if failed - try diagonal (corners)
		if field[dir_y][dir_x] == '#' || field[dir_y][dir_x] == '%' {
			dir_y = monster[m].y
			dir_x = monster[m].x

			if dir_y > py {
				dir_y -= 1
			} else {
				dir_y += 1
			}

			if dir_x > px {
				dir_x -= 1
			} else {
				dir_x += 1
			}
		}
		// if still failed - move to closes to player empty tile (invert)
		if field[dir_y][dir_x] == '#' || field[dir_y][dir_x] == '%' {
			dir_y = monster[m].y
			dir_x = monster[m].x

			if dist_y > dist_x {
				if dir_x > px {
					dir_x -= 1
				} else {
					dir_x += 1
				}
			} else {
				if dir_y > py {
					dir_y -= 1
				} else {
					dir_y = +1
				}
			}
		}

		if dist_y < 2 && dist_x < 2 {
			dmg := ((monster[m].typeGo - 96) + dlvl) / 2
			if rand.Int()%2 != 0 {
				if (rand.Int()%3 != 0 && (strings.Compare(race, "Halfling") == 0)) || ((rand.Int()%5 != 0) && (strings.Compare(race, "Elf") == 0)) {
					goncurses.StdScr().MovePrint(0, 0, " You dodge the attack.")
				} else {
					hp -= dmg
					if dmg > dlvl/2+1 {
						goncurses.StdScr().MovePrintf(0, 0, " The '%c' hits you hard.", monster[m].typeGo)
					} else {
						goncurses.StdScr().MovePrintf(0, 0, " The '%c' hits you.", monster[m].typeGo)
					}
					if hp < 1 {
						return monster[m].typeGo
					}
				}
			} else {
				if rand.Int()%2 == 0 {
					goncurses.StdScr().MovePrintf(0, 0, " The '%c' missed you.", monster[m].typeGo)
				}
			}

		} else if field[dir_y][dir_x] == ' ' && (dir_y != py && dir_x != px) {
			field[monster[m].y][monster[m].x] = ' '
			monster[m].y = dir_y
			monster[m].x = dir_x
			field[monster[m].y][monster[m].x] = 't'

		}
	}
	return 0
}

func battle(cols int, field [][]byte, dir_y int, dir_x int) int {
	for m := 0; m < 10+dlvl/2; m++ {
		if dir_y == monster[m].y && dir_x == monster[m].x {
			var sleeped bool
			//wake up
			if !monster[m].awake {
				monster[m].awake = true
				sleeped = true
			}
			// var dmg int
			dmg := att
			//hit
			if strings.Compare(race, "Halfling") == 0 && att > 1 {
				dmg--
			}
			monster[m].lvl -= dmg
			// if was asleep and we didn't defeat it yet
			if sleeped && monster[m].lvl > 0 {
				wasAlmostDead := false
				goncurses.StdScr().MovePrintf(0, cols/2, ">> You hit sleeping '%c' hard.   ", monster[m].typeGo)
				if monster[m].lvl < (dlvl*9)/8 {
					wasAlmostDead = true
				}
				monster[m].lvl -= dmg
				// give monster a chance to survive critical hit
				if !wasAlmostDead && monster[m].lvl < 1 && rand.Int()%(dlvl+1) != 0 {
					monster[m].lvl = 1
				}
			} else {
				goncurses.StdScr().MovePrintf(0, cols/2, ">> You hit '%c'.   ", monster[m].typeGo)
			}
			// monster defeated
			if monster[m].lvl < 1 {
				m_defeated++
				// gain player lvl
				// 1st lvl hardcoded
				if dlvl == 1 && (rand.Int()%3 == 0 || m_defeated > 2) && m_defeated < 4 {
					att++
					mana++
					hp += rand.Int()%10 + 1
				} else if (rand.Int() % (((monster[m].typeGo - 96) + dlvl) / 2)) != 0 { //gain stats
					hp += rand.Int()%10 + 1
					if rand.Int()%2 == 0 {
						mana++
					} else {
						att++
					}
					if strings.Compare(race, "Human") == 0 && rand.Int()%2 == 0 && dlvl > 1 {
						hp++
					}
				}
				if m_defeated > 9 && m_defeated%10 == 0 {
					att++
				}
				goncurses.StdScr().MovePrintf(0, cols/2, ">> You defeat '%c'.   ", monster[m].typeGo)
				// wipe monster from DB
				field[dir_y][dir_x] = ' '
				monster[m].y = 0
				monster[m].x = 0

			}
		}
	}
	return 0
}

func PlAction(c int, rows int, cols int, field [][]byte, obj [][]byte) int {
	dir_y := py
	dir_x := px
	// remap macro for movement
	if c == 'w' || c == 'k' {
		c = goncurses.KEY_UP
	} else if c == 's' || c == 'j' {
		c = goncurses.KEY_DOWN
	} else if c == 'a' || c == 'h' {
		c = goncurses.KEY_LEFT
	} else if c == 'd' || c == 'l' {
		c = goncurses.KEY_RIGHT
	}

	//confusion
	if strings.Compare(state, "conf") == 0 {
		if c == goncurses.KEY_UP || c == goncurses.KEY_DOWN || c == goncurses.KEY_LEFT || c == goncurses.KEY_RIGHT {
			rng := rand.Int() % 4
			if rng == 0 {
				c = goncurses.KEY_UP
			} else if rng == 1 {
				c = goncurses.KEY_DOWN
			} else if rng == 2 {
				c = goncurses.KEY_LEFT
			} else if rng == 3 {
				c = goncurses.KEY_RIGHT
			}
		}

	}

	if c == goncurses.KEY_UP {
		dir_y--
	} else if c == goncurses.KEY_DOWN {
		dir_y++
	} else if c == goncurses.KEY_LEFT {
		dir_x--
	} else if c == goncurses.KEY_RIGHT {
		dir_x++
	} else if (c == '>' || c == 'r' || c == '\n' || c == goncurses.KEY_ENTER || c == ' ') && obj[py][px] == '>' {
		t_placed = false
		p_placed = false
		r_placed = 0
		return 1 //go down
	} else if mana > 0 && (c == '1' || c == 'q' || c == 't') { //teleport
		if dlvl < 13 {
			mana--
			if strings.Compare(race, "Elf") == 0 && rand.Int()%2 == 0 {
				mana++
			}
			py = rand.Int() % rows
			px = rand.Int() % cols
			for {
				if field[py][px] != ' ' && obj[py][px] == ' ' {
					break
				}
				py = rand.Int() % rows
				px = rand.Int() % cols
			}
		}
		return 2 //tp
	} else if mana > 1 && (c == '2' || c == 'e' || c == 'y') { //heal
		mana -= 2
		if strings.Compare(race, "Orc") == 0 && rand.Int()%2 == 0 {
			mana++
		}
		hp += rand.Int()%dlvl + 5
		if hp > dlvl*10 {
			hp = dlvl * 10
		}
		goncurses.StdScr().MovePrint(0, cols/2, ">> You heal yourself.")
		//also heal conf
		state = "\000\000\000\000\000"
		goncurses.StdScr().MovePrint(rows, cols-20, "    ")
		return 0
	} else if hp > 1 && (c == '3' || c == 'r' || c == 'u') { //dig
		c = int(goncurses.StdScr().GetChar())
		hp -= 1
		if strings.Compare(race, "Dwarf") == 0 && rand.Int()%2 == 0 {
			hp++
		}

		if c == goncurses.KEY_UP || c == 'w' || c == 'k' {
			dir_y--
		} else if c == goncurses.KEY_DOWN || c == 's' || c == 'j' {
			dir_y++
		} else if c == goncurses.KEY_LEFT || c == 'a' || c == 'h' {
			dir_x--
		} else if c == goncurses.KEY_RIGHT || c == 'd' || c == 'l' {
			dir_x++
		}

		if field[dir_y][dir_x] == '#' {
			field[dir_y][dir_x] = ' '
		}
		return 0
	}
	//win
	if field[dir_y][dir_x] == '~' {
		return 3
	}
	//move
	if field[dir_y][dir_x] == ' ' {
		py = dir_y
		px = dir_x
	} else if field[dir_y][dir_x] == 't' {
		battle(cols, field, dir_y, dir_x)
	}
	return 0
}

func SpawnT(rows int, cols int, field [][]byte) int {
	if !t_placed {
		var my, mx int
		boss := 0
		for m := 0; m < 10+dlvl/2; m++ {
			my = rand.Int() % rows
			mx = rand.Int() % cols
			for {
				if field[my][mx] == ' ' || (my == py && mx == px) {
					break
				}
				my = rand.Int() % rows
				mx = rand.Int() % cols
			}
			monster[m].y = my
			monster[m].x = mx
			//lvl
			monster[m].lvl = rand.Int()%dlvl + 2
			if dlvl == 1 && rand.Int()%5 == 0 {
				monster[m].lvl = 1
			}
			if rand.Int()%2 == 0 {
				monster[m].lvl = dlvl + 2
			}
			//type
			if (dlvl == 13 || dlvl == 14) && boss != 9 {
				monster[m].typeGo = 'Z'
				monster[m].lvl = 666
				monster[m].awake = true
				boss++
			} else {
				monster[m].typeGo = rand.Int()%(dlvl+1) + 97
				if rand.Int()%2 == 0 && dlvl != 1 {
					monster[m].typeGo += 1
				}
				monster[m].awake = false
			}
			field[monster[m].y][monster[m].x] = 't'
		}
		t_placed = true
	}
	return 0
}

func SpawnP(rows int, cols int, field [][]byte, obj [][]byte) int {
	if !p_placed {
		sy = rand.Int() % rows
		sx = rand.Int() % cols
		for {
			if field[sy][sx] == ' ' {
				break
			}
			// if field[sy][sx] != '%' && field[sy][sx] != '#' {
			// 	break
			// }
			sy = rand.Int() % rows
			sx = rand.Int() % cols
		}
		obj[sy][sx] = '>'
		var dist_y, dist_x int
		py = rand.Int() % rows
		px = rand.Int() % cols
		//distance from stairs
		dist_y = int(math.Abs(float64(py) - float64(sy)))
		dist_x = int(math.Abs(float64(px) - float64(sx)))
		for {
			if field[py][px] == ' ' && obj[py][px] == ' ' && (dist_y > 7+dlvl/2 || dist_x > 7+dlvl/2) {
				break
			}
			py = rand.Int() % rows
			px = rand.Int() % cols
			//distance from stairs
			dist_y = int(math.Abs(float64(py) - float64(sy)))
			dist_x = int(math.Abs(float64(px) - float64(sx)))
		}
		p_placed = true
	}
	return 0
}

func SpawnObj(rows int, cols int, field [][]byte, obj [][]byte) int {
	if lvl_turns == 0 || turns == 0 {
		for y := 0; y <= rows; y++ {
			for x := 0; x <= cols; x++ {
				obj[y][x] = ' '
			}
		}
		//stair
		final_lvl := 13 + rand.Int()%2
		if dlvl != final_lvl {
			sy = rand.Int() % rows
			sx = rand.Int() % cols
			for {
				if field[sy][sx] != ' ' {
					break
				}
				// if field[sy][sx] != '%' && field[sy][sx] != '#' {
				// 	break
				// }
				sy = rand.Int() % rows
				sx = rand.Int() % cols
			}
			obj[sy][sx] = '>'
		} else { //lava
			var ly, lx int
			ly = rand.Int() % rows
			lx = rand.Int() % cols
			for {
				if field[ly][lx] != ' ' {
					break
				}
				ly = rand.Int() % rows
				lx = rand.Int() % cols
			}
			field[ly][lx] = '~'
			for i := 0; i < 33; i++ {
				if field[ly][lx] == '%' {
					i = 32
				}
				if rand.Int()%2 == 0 {
					field[ly-1][lx] = '~'
				} else {
					field[ly][lx+1] = '~'
				}
			}

		}
	}
	// each turn there is a chance to spawn trap (the faster you run away - the better)
	if rand.Int()%((18-dlvl)/2) == 0 {
		var y, x int
		y = rand.Int() % rows
		x = rand.Int() % cols
		for {
			if field[y][x] == ' ' && obj[y][x] == ' ' {
				break
			}
			y = rand.Int() % rows
			x = rand.Int() % cols
		}
		obj[y][x] = '^'
	}
	return 0
}

func DungeonGen(rows int, cols int, field [][]byte) int {
	if r_placed == 0 {
		var ry, rx int             // room coords
		var r_size_y, r_size_x int // room size
		var r_center_y, r_center_x int
		var r_old_center_y, r_old_center_x int
		room_num := rand.Int()%5 + 5
		var collision bool
		// fill dungeon with walls and borders
		for y := 0; y <= rows; y++ {
			for x := 0; x <= cols; x++ {
				//borders
				if y == 0 || y == 1 || y == rows-1 || x == 0 || x == cols || y == rows {
					field[y][x] = '%'
				} else {
					field[y][x] = '#'
				}
			}
		}
		for r_placed < room_num {
			try_counter := 0 //number of tries for prototiping
			for {
				if try_counter > 100 {
					ry = 3
					rx = 3
					r_size_y = 3
					r_size_x = 3
					break
				}
				collision = false
				ry = rand.Int()%(rows-4) + 1
				rx = rand.Int()%(cols-4) + 1
				r_size_y = rand.Int()%5 + 4
				r_size_x = rand.Int()%10 + 8
				try_counter++
				//check for collision
				for y := ry; y <= ry+r_size_y && y < rows; y++ {
					for x := rx; x < rx+r_size_x && x < cols; x++ {
						if field[y][x] == '%' || field[y][x] == ' ' || field[y+1][x] == ' ' || field[y-1][x] == ' ' || field[y][x+1] == ' ' || field[y][x-1] == ' ' {
							collision = true
							y = ry + r_size_y + 1 //exit upper loop
							break                 //exit current loop
						}
					}
				}
				if !collision {
					break
				}

			}
			// fill DB map with rooms
			for y := ry; y <= ry+r_size_y && y < rows; y++ {
				for x := rx; x <= rx+r_size_x && x < cols; x++ {
					if field[y][x] == '%' {
						y = ry + r_size_y + 1
						break
					} else {
						field[y][x] = ' '
					}
				}
			}
			r_placed++
			//corridors
			if r_placed > 1 {
				r_old_center_y = r_center_y
				r_old_center_x = r_center_x
			}
			r_center_y = ry + (r_size_y / 2)
			r_center_x = rx + (r_size_x / 2)

			if r_center_y >= rows {
				r_center_y = rows - 3
			}
			if r_center_x >= cols {
				r_center_x = cols - 3
			}
			if r_placed > 1 {
				var path_y, path_x int
				for path_y = r_old_center_y; path_y != r_center_y && path_y < rows; {
					if field[path_y][r_old_center_x] != '%' {
						field[path_y][r_old_center_x] = ' '
					}
					if r_old_center_y < r_center_y {
						path_y++
					} else if r_old_center_y > r_center_y {
						path_y--
					}
				}

				for path_x = r_old_center_x; path_x != r_center_x && path_x < cols; {
					if field[path_y][path_x] != '%' {
						field[path_y][path_x] = ' '
					}
					if r_old_center_x < r_center_x {
						path_x++
					} else if r_old_center_x > r_center_x {
						path_x--
					}
				}
			}
		}

	}
	return 0
}

func CreateChar(c int) int {
	att = 1
	mana = 1
	stealth = 0
	t_placed = false
	p_placed = false
	r_placed = 0
	dlvl = 1
	lvl_turns = 0
	m_defeated = 0
	state = "\000\000\000\000\000"

	if c == 'n' {
		if strings.Compare(race, "Human") == 0 {
			c = '1'
		} else if strings.Compare(race, "Dwarf") == 0 {
			c = '2'
		} else if strings.Compare(race, "Elf") == 0 {
			c = '3'
		} else if strings.Compare(race, "Halfling") == 0 {
			c = '4'
		} else if strings.Compare(race, "Orc") == 0 {
			c = '5'
		}
	}

	switch c {
	case '1':
		hp = 10 + rand.Int()%2
		stealth = 0
		race = "Human"
	case '2':
		hp = 10 + rand.Int()%3 + 2
		att += 2
		stealth -= 2
		race = "Dwarf"
	case '3':
		hp = 10 + rand.Int()%3
		att += 2
		stealth = 1
		race = "Elf"
	case '4':
		hp = 10 - rand.Int()%2
		stealth = 2
		race = "Halflling"
	case '5':
		hp = 10 - rand.Int()%2 + 1
		att += 1
		stealth = -1
		race = "Orc"
	}
	return 0
}

func GameLoop(c int, rows int, cols int, field [][]byte, obj [][]byte) int {
	actionResult := 0
	killer := 0

	goncurses.StdScr().Move(0, 0)
	goncurses.StdScr().ClearToEOL()

	if turns == 0 || c == 'n' {
		CreateChar(c)
	}
	DungeonGen(rows, cols, field)
	SpawnObj(rows, cols, field, obj)
	SpawnP(rows, cols, field, obj)
	SpawnT(rows, cols, field)

	if turns > 0 {
		if c != 0 { //to prevent need of double push a button '3' at the beginning
			actionResult = PlAction(c, rows, cols, field, obj) //+battle
			killer = MonsterTurn(cols, field)
		}
		if hp > 0 {
			killer = CheckTrap(rows, cols, obj)
		}
		if hp < 1 {
			c = Rip(rows, cols, killer)
			turns = 0
			return c
		}
	}
	//new level
	if actionResult == 1 {
		dlvl++
		hp += dlvl
		lvl_turns = 0
		DungeonGen(rows, cols, field)
		SpawnObj(rows, cols, field, obj)
		SpawnP(rows, cols, field, obj)
		SpawnT(rows, cols, field)
		DungeonDraw(rows, cols, field, obj)
	} else if actionResult == 2 { //teleport
		if dlvl < 13 {
			goncurses.StdScr().MovePrint(0, 0, " You teleported away.")
		} else {
			goncurses.StdScr().MovePrint(0, 0, " Anti-magic field prevents teleport.")
		}
	} else if actionResult == 3 { //win
		for {
			goncurses.StdScr().Clear()
			goncurses.StdScr().Print("\n\n\n\n \t\t\t\tThe Ring is destroyed!\n\n\tEagles came and brought you away from the fury of the night!\n\n\n")
			goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_CYAN))
			goncurses.StdScr().Print("\t\t\t\t YOU WIN!\n\n")
			goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_CYAN))
			goncurses.StdScr().Printf("\tLevel reached: %d\n\tMonsters defeated: %d\n\tTurns: %d\n\n\tAttack: %d\n\tMana: %d\n\n\n\tPress 'n' to start new game or 'ESC' to exit.", dlvl, m_defeated, turns, att, mana)
			c = int(goncurses.StdScr().GetChar())
			if c == 'n' || c == 27 {
				return c
			}
		}
	}
	DungeonDraw(rows, cols, field, obj)
	//UI line
	goncurses.StdScr().MovePrintf(rows, 0, " %s    HP: %d   Att: %d   Mana: %d \t\t Dlvl: %d", race, hp, att, mana, dlvl)
	//process state and update UI line
	if strings.Compare(state, "conf") == 0 {
		if rand.Int()%3 != 0 {
			state = "\000\000\000\000\000"
			goncurses.StdScr().MovePrint(rows, cols-20, "    ")
		} else {
			goncurses.StdScr().MovePrint(rows, cols-20, "conf")
		}
	}
	//player input
	c = int(goncurses.StdScr().GetChar())
	//exit game (ESC)
	if c == 27 {
		return c
	} else {
		turns++ //turns count
		lvl_turns++
		if (turns%50-(dlvl*2)) == 0 && hp > 1 {
			hp--
		}
	}
	return c
}

func IntroUI() int {
	c := 0
	goncurses.StdScr().Print("\n")
	goncurses.StdScr().AttrOn(goncurses.A_BOLD)
	goncurses.StdScr().Print("\t\t\t\t    Rogue prototype")
	goncurses.StdScr().AttrOff(goncurses.A_BOLD)
	goncurses.StdScr().Print("\n\n\tBring the ring to the ")
	goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_RED))
	goncurses.StdScr().Print("River of Flame")
	goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_RED))
	goncurses.StdScr().Print(" and try to get rid of it...")
	goncurses.StdScr().Print("\n\n\n\tContols:\n\n\tArrows/wasd/hjkl - move and attack\n\tSpace/Enter/> - go to the next level \n\t1/q/t - teleport (mana) \n\t2/e/y - heal (mana) \n\t3/r/u - dig (HP) \n\t'n' - start a new game\n\t'ESC' - exit game\n\n\n\tYou level up HP and Attack by defeating monsters.\n\tYou become hungry (HP) after some time.\n\n\tChoose race wisely:\n")
	goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_CYAN))
	goncurses.StdScr().Print("\t1) Human 2) Dwarf 3) Elf 4) Halfling 5) Orc\n\t   -- press '?' to see races details --")
	//  mid    sturdy    dexy   stealth-dodge  reverse
	goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_CYAN))
	// goncurses.StdScr().Printf("\ngetchar will be now used first\n")
	c = int(goncurses.StdScr().GetChar())
	goncurses.StdScr().Printf("getchar = %c\n", c)
	if c == '?' {
		goncurses.StdScr().Clear()
		goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_CYAN))
		goncurses.StdScr().Print("\n\t\t\t\tRaces:\n\n")
		goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_CYAN))
		goncurses.StdScr().AttrOn(goncurses.A_BOLD)
		goncurses.StdScr().Print("\tHuman\n")
		goncurses.StdScr().AttrOff(goncurses.A_BOLD)
		goncurses.StdScr().Print("\tlearns faster\n")

		goncurses.StdScr().AttrOn(goncurses.ColorPair(goncurses.C_YELLOW))
		goncurses.StdScr().Print("\tDwarf\n")
		goncurses.StdScr().AttrOff(goncurses.ColorPair(goncurses.C_YELLOW))
		goncurses.StdScr().Print("\t+HP, + Att, -Stealth\n\tbonus to Digging\n")

		goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_GREEN))
		goncurses.StdScr().Print("\tElf\n")
		goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_GREEN))
		goncurses.StdScr().Print("\t+HP, +Att, +Stealth\n\tbonus to Teleportation\n")

		goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_RED))
		goncurses.StdScr().Print("\tHalfling\n")
		goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_RED))
		goncurses.StdScr().Print("\t-HP, -Att, +Stealth\n\tcan dodge and avoid traps sometimes\n")

		goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_MAGENTA))
		goncurses.StdScr().Print("\tOrc\n")
		goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_MAGENTA))
		goncurses.StdScr().Print("\t+HP, +Att, -Stealth\n\tbonus to Healing\n\n\n")

		goncurses.StdScr().Print("\tSo...Which race you choose?\n")
		goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_CYAN))
		goncurses.StdScr().Print("\t1) Human  2) Dwarf  3) Elf  4) Halfling  5) Orc\n")
		goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_CYAN))

		c = int(goncurses.StdScr().GetChar())

	}

	return c

}

func main() {
	var c int //input
	var rows, cols int

	// initscr(); //init curses
	goncurses.Init()
	goncurses.StartColor()
	goncurses.UseDefaultColors()

	// init_pair(BLACK, COLOR_BLACK, COLOR_WHITE);
	// init_pair(RED, COLOR_RED, COLOR_BLACK);
	// init_pair(GREEN, COLOR_GREEN, COLOR_BLACK);
	// init_pair(YELLOW, COLOR_YELLOW, COLOR_BLACK);
	// init_pair(BLUE, COLOR_BLUE, COLOR_BLACK);
	// init_pair(MAGENTA, COLOR_MAGENTA, COLOR_BLACK);
	// init_pair(CYAN, COLOR_CYAN, COLOR_BLACK);
	// init_pair(WHITE, COLOR_WHITE, COLOR_WHITE);

	goncurses.InitPair(goncurses.C_BLACK, goncurses.C_BLACK, goncurses.C_WHITE)
	goncurses.InitPair(goncurses.C_RED, goncurses.C_RED, goncurses.C_BLACK)
	goncurses.InitPair(goncurses.C_GREEN, goncurses.C_GREEN, goncurses.C_BLACK)
	goncurses.InitPair(goncurses.C_YELLOW, goncurses.C_YELLOW, goncurses.C_BLACK)
	goncurses.InitPair(goncurses.C_BLUE, goncurses.C_BLUE, goncurses.C_BLACK)
	goncurses.InitPair(goncurses.C_MAGENTA, goncurses.C_MAGENTA, goncurses.C_BLACK)
	goncurses.InitPair(goncurses.C_CYAN, goncurses.C_CYAN, goncurses.C_BLACK)
	goncurses.InitPair(goncurses.C_WHITE, goncurses.C_WHITE, goncurses.C_BLACK)

	goncurses.StdScr().Keypad(true)

	// goncurses.NoEcho()
	goncurses.Echo(false)
	// goncurses.CursSet(0)
	goncurses.Cursor(0)

	// adaptive way of the screen size:
	rows, cols = goncurses.StdScr().MaxYX()
	// but this particular game designed to be played at tiny term:
	// rows = 23
	// cols = 80

	// var field[rows][cols] slice
	field := make([][]byte, rows)
	obj := make([][]byte, rows)
	for i := range field {
		field[i] = make([]byte, cols)
		obj[i] = make([]byte, cols)
	}
	// goncurses.StdScr().Print("\tNow")
	// c = int(goncurses.StdScr().GetChar())
	// goncurses.StdScr().Printf("getchar from main = %c", c)

	c = IntroUI()
	// goncurses.StdScr().Print("\tAfter Intro")

	if c == 27 {
		goncurses.End()
		return
	}

	for {
		c = GameLoop(c, rows-1, cols-1, field, obj)
		if c == 27 {
			goncurses.End()
			return
		}
	}

}
