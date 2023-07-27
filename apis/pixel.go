package apis

import (
	"context"
	"log"
	"nosplace/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var canvasMap = make(map[int64]map[int64][]byte)

func (rt *Router) preloadCanvas() {
	for {
		t := time.Now()
		state, err := rt.Storage.GetIndexerState(context.Background(), "place_indexer")
		if err != nil {
			log.Printf("Failed to get indexer state: %s", err.Error())
			continue
		}
		if state.CanvasHeight == 0 || state.CanvasWidth == 0 {
			time.Sleep(2 * time.Second)
			continue
		}
		size := state.CanvasHeight
		if state.CanvasWidth > state.CanvasHeight {
			size = state.CanvasWidth
		}
		pixelList, err := rt.Storage.GetPixels(context.Background(), int64(state.CanvasLeftBound), int64(state.CanvasUpperBound), size)
		if err != nil {
			log.Printf("Failed to get pixels: %s", err.Error())
			continue
		}
		canvasMap = makePixelsMap(pixelList)
		log.Printf("Preloaded canvas in %s", time.Since(t).String())
		time.Sleep(2 * time.Second)
	}
}

func (rt *Router) GetPixels(c *gin.Context) {
	xStr := c.Query("x")
	yStr := c.Query("y")
	sizeStr := c.Query("size")

	x, err := strconv.ParseInt(xStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid x",
		})
		return
	}
	y, err := strconv.ParseInt(yStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid y",
		})
		return
	}
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid size",
		})
		return
	}
	pixelList, err := rt.Storage.GetPixels(c, x, y, size)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"pixels": formatPixelsToBytes(makePixelsMap(pixelList), x, y, size),
	})
}

func (rt *Router) GetPixelsCached(c *gin.Context) {
	xStr := c.Query("x")
	yStr := c.Query("y")
	sizeStr := c.Query("size")

	x, err := strconv.ParseInt(xStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid x",
		})
		return
	}
	y, err := strconv.ParseInt(yStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid y",
		})
		return
	}
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid size",
		})
		return
	}
	c.JSON(200, gin.H{
		"pixels": formatPixelsToBytes(canvasMap, x, y, size),
	})
}

func (rt *Router) GetLatest(c *gin.Context) {
	from := c.Query("from")

	fromInt, err := strconv.ParseInt(from, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid from",
		})
		return
	}

	pixelList, err := rt.Storage.GetLatestPixels(c, fromInt)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"pixels": pixelList,
	})
}

func formatPixelsToBytes(pixelsMap map[int64]map[int64][]byte, x, y, size int64) []int {
	var result []byte
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			if pixelsMap[i] != nil && pixelsMap[i][j] != nil {
				result = append(result, pixelsMap[i][j]...)
			} else {
				result = append(result, 255, 255, 255)
			}
		}
	}
	var result2 []int
	for _, v := range result {
		result2 = append(result2, int(v))
	}
	return result2
}

func makePixelsMap(pixels []models.Pixel) map[int64]map[int64][]byte {
	var pixelMap = make(map[int64]map[int64][]byte)

	for _, pixel := range pixels {
		if _, ok := pixelMap[pixel.X]; !ok {
			pixelMap[pixel.X] = make(map[int64][]byte)
		}
		colors := strings.Split(pixel.Color, ",")
		r, _ := strconv.ParseUint(colors[0], 10, 8)
		g, _ := strconv.ParseUint(colors[1], 10, 8)
		b, _ := strconv.ParseUint(colors[2], 10, 8)
		pixelMap[pixel.X][pixel.Y] = []byte{byte(r), byte(g), byte(b)}
	}
	return pixelMap
}
