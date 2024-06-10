package Numbria

import (
	"github.com/Joeverson/numbria-game/types"
	"github.com/Joeverson/numbria-game/utils"
)

type World struct {
	Width  int
	Height int
	area   [][]MapZone
}

type MapZone struct {
	name     string
	zone     types.ZoneType
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
				isEndMap: isEndMap(pixels[x][y]),
			})
		}
		world.area = append(world.area, zone)
	}
}

func isEndMap(pixel utils.Pixel) bool {
	return pixel == utils.Pixel{R: 6, G: 6, B: 8, A: 255}
}

func translatePixelToZoneName(pixel utils.Pixel) string {
	switch pixel {
	case utils.Pixel{R: 26, G: 122, B: 62, A: 255}:
		return string(types.ZoneEnum.Florest)
	case utils.Pixel{R: 89, G: 193, B: 53, A: 255}:
		return string(types.ZoneEnum.Clearing)
	case utils.Pixel{R: 36, G: 159, B: 222, A: 255}:
		return string(types.ZoneEnum.River)
	case utils.Pixel{R: 180, G: 32, B: 42, A: 255}:
		return string(types.ZoneEnum.Boss)
	case utils.Pixel{R: 91, G: 49, B: 56, A: 255}:
		return string(types.ZoneEnum.Cave)
	case utils.Pixel{R: 6, G: 6, B: 8, A: 255}:
		return string(types.ZoneEnum.End_Map)
	case utils.Pixel{R: 34, G: 28, B: 26, A: 255}:
		return string(types.ZoneEnum.Wall)
	case utils.Pixel{R: 142, G: 82, B: 82, A: 255}:
		return string(types.ZoneEnum.Way)
	case utils.Pixel{R: 188, G: 74, B: 155, A: 255}:
		return string(types.ZoneEnum.Room)
	default:
		return ""
	}
}

func translatePixelToZone(pixel utils.Pixel) types.ZoneType {
	switch pixel {
	case utils.Pixel{R: 26, G: 122, B: 52, A: 255}:
		return types.ZoneEnum.Florest
	case utils.Pixel{R: 89, G: 193, B: 53, A: 255}:
		return types.ZoneEnum.Clearing
	case utils.Pixel{R: 36, G: 159, B: 222, A: 255}:
		return types.ZoneEnum.River
	case utils.Pixel{R: 180, G: 32, B: 42, A: 255}:
		return types.ZoneEnum.Boss
	case utils.Pixel{R: 91, G: 49, B: 56, A: 255}:
		return types.ZoneEnum.Cave
	case utils.Pixel{R: 34, G: 28, B: 26, A: 255}:
		return types.ZoneEnum.Wall
	case utils.Pixel{R: 142, G: 82, B: 82, A: 255}:
		return types.ZoneEnum.Way
	case utils.Pixel{R: 188, G: 74, B: 155, A: 255}:
		return types.ZoneEnum.Room
	case utils.Pixel{R: 6, G: 6, B: 8, A: 255}:
		return types.ZoneEnum.End_Map
	default:
		return ""
	}
}
