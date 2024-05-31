package game

import (
	"github.com/Joeverson/numbria-game/utils"
)

type Zone int

const (
	Florest Zone = iota + 1
	Clearing
	River
	Boss
	Cave
	End_Map
)

type World struct {
	Width  int
	Height int
	area   [][]MapZone
}

type MapZone struct {
	name     string
	zone     Zone
	x        int
	y        int
	isEndMap bool
}

func (w World) GetNameZone(x, y int) string {
	return w.area[x][y].name
}

func (world *World) MapGenerate() {
	width, height, pixels := utils.LoadImageTemplate()

	world.Height = height
	world.Width = width

	for y := 0; y < height; y++ {
		var zone []MapZone
		for x := 0; x < width; x++ {
			zone = append(zone, MapZone{
				name:     translatePixelToZoneName(pixels[x][y]),
				zone:     translatePixelToZone(pixels[x][y]),
				x:        x,
				y:        y,
				isEndMap: pixelMapTypeCheck(End_Map, pixels[x][y]),
			})
		}
		world.area = append(world.area, zone)
	}
}

func pixelMapTypeCheck(zone Zone, pixel utils.Pixel) bool {
	switch zone {
	case Florest:
		return pixel == utils.Pixel{R: 26, G: 122, B: 62, A: 255} // green
	case Clearing:
		return pixel == utils.Pixel{R: 89, G: 193, B: 53, A: 255} // green light
	case River:
		return pixel == utils.Pixel{R: 36, G: 159, B: 222, A: 255} // blue
	case Boss:
		return pixel == utils.Pixel{R: 180, G: 32, B: 42, A: 255} // red
	case Cave:
		return pixel == utils.Pixel{R: 91, G: 49, B: 56, A: 255} // brown
	case End_Map:
		return pixel == utils.Pixel{R: 6, G: 6, B: 8, A: 255} // black
	}

	return false
}

func translatePixelToZoneName(pixel utils.Pixel) string {
	switch pixel {
	case utils.Pixel{R: 26, G: 122, B: 62, A: 255}:
		return "floresta"
	case utils.Pixel{R: 89, G: 193, B: 53, A: 255}:
		return "clareira"
	case utils.Pixel{R: 36, G: 159, B: 222, A: 255}:
		return "rio"
	case utils.Pixel{R: 180, G: 32, B: 42, A: 255}:
		return "boss"
	case utils.Pixel{R: 91, G: 49, B: 56, A: 255}:
		return "caverna"
	case utils.Pixel{R: 6, G: 6, B: 8, A: 255}:
		return "fim do mapa"
	default:
		return ""
	}
}

func translatePixelToZone(pixel utils.Pixel) Zone {
	switch pixel {
	case utils.Pixel{R: 26, G: 122, B: 52, A: 255}:
		return Florest
	case utils.Pixel{R: 89, G: 193, B: 53, A: 255}:
		return Clearing
	case utils.Pixel{R: 36, G: 159, B: 222, A: 255}:
		return River
	case utils.Pixel{R: 180, G: 32, B: 42, A: 255}:
		return Boss
	case utils.Pixel{R: 91, G: 49, B: 56, A: 255}:
		return Cave
	case utils.Pixel{R: 6, G: 6, B: 8, A: 255}:
		return End_Map
	default:
		return -1
	}
}
