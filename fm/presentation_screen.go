package main

import (
	"container/list"
	"os"

	"github.com/gbin/goncurses"
)

func screen_print(screen *screen_t, statStr *string, msgStr *list.List) {
	for i := 0; i < SCRN_H; i++ {
		for j := 0; j < SCRN_W; j++ {
			if i < MAP_HEIGHT && j < MAP_WIDTH {
				if screen.screen[i][j] == 'z' {
					goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_GREEN))
					goncurses.StdScr().MoveAddChar(i, j, goncurses.Char(screen.screen[i][j]))
					goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_GREEN))
				} else if screen.screen[i][j] == 'v' {
					goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_RED))
					goncurses.StdScr().MoveAddChar(i, j, goncurses.Char(screen.screen[i][j]))
					goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_RED))
				} else if screen.screen[i][j] == 'O' {
					goncurses.StdScr().AttrOn(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_YELLOW))
					goncurses.StdScr().MoveAddChar(i, j, goncurses.Char(screen.screen[i][j]))
					goncurses.StdScr().AttrOff(goncurses.A_BOLD | goncurses.ColorPair(goncurses.C_YELLOW))
				} else {
					goncurses.StdScr().MoveAddChar(i, j, goncurses.Char(screen.screen[i][j]))
				}
			} else if i == MAP_HEIGHT {
				goncurses.StdScr().MovePrint(i, 0, *statStr)

			} else if i == MAP_HEIGHT+3 {
				goncurses.StdScr().MovePrint(i, 0, msgStr.Front().Value)
			} else if i == MAP_HEIGHT+2 {
				goncurses.StdScr().MovePrint(i, 0, msgStr.Front().Next().Value)
			} else {
				goncurses.StdScr().MoveAddChar(i, j, goncurses.Char(SPACE_CHAR))
				// goncurses.StdScr().MovePrint(i, j, '©')
			}
		}
	}
}

func printRecords() {
	store, err := NewStore()
	data, _ := store.loadAttempts()
	goncurses.StdScr().Clear()
	goncurses.StdScr().MovePrint(0, 0, "Stats. Press esc to exit.")
	if err == nil {
		for i := range data.Attempts {
			goncurses.StdScr().MovePrint(i+2, 0, "Gold: ", data.Attempts[i].Gold, " Level: ", data.Attempts[i].Level, " Monsters defeated: ", data.Attempts[i].MonsterKill, " Turns: ", data.Attempts[i].Turns)
		}
	}
}

func printCongratulations() {
	goncurses.StdScr().Clear()
	goncurses.StdScr().MovePrint(0, 0, "Congratulation. You win. Are you feeling happy now?")
	goncurses.StdScr().MovePrint(1, 0, "Press ESC to exit.")
	var key int
	for key != ESC {
		key = int(goncurses.StdScr().GetChar())
	}
	goncurses.End()
	os.Exit(0)
}

func printWeapon(wList *list.List) int {
	goncurses.StdScr().Clear()
	goncurses.StdScr().MovePrint(0, 0, "Which weapon do you want to use? 0 - bare hands. Enter number 0-9:\n")

	for i := 0; i < wList.Len(); i++ {
		goncurses.StdScr().MovePrint(i+1, 0, wList.Front().Value)
		wList.MoveToBack(wList.Front())
	}
	key := goncurses.StdScr().GetChar()

	return int(key)
}

func printFood(fList *list.List) {
	goncurses.StdScr().Clear()
	goncurses.StdScr().MovePrint(0, 0, "Which food do you want to eat? Enter number 1-9:\n")
	for i := 0; i < fList.Len(); i++ {
		goncurses.StdScr().MovePrint(i+1, 0, fList.Front().Value)
		fList.MoveToBack(fList.Front())
	}
}

func printScrolls(sList *list.List) {
	goncurses.StdScr().Clear()
	goncurses.StdScr().MovePrint(0, 0, "Which scroll do you want to use? Enter number 1-9:\n")
	for i := 0; i < sList.Len(); i++ {
		goncurses.StdScr().MovePrint(i+1, 0, sList.Front().Value)
		sList.MoveToBack(sList.Front())
	}
}

func printPotions(pList *list.List) {
	goncurses.StdScr().Clear()
	goncurses.StdScr().MovePrint(0, 0, "Which potion do you want to consume? Enter number 1-9:\n")
	for i := 0; i < pList.Len(); i++ {
		goncurses.StdScr().MovePrint(i+1, 0, pList.Front().Value)
		pList.MoveToBack(pList.Front())
	}
}

func getInput() int {
	key := goncurses.StdScr().GetChar()

	return int(key)
}

func clear_screen() {
	goncurses.StdScr().Clear()

}
