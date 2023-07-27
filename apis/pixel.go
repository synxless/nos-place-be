package apis

import (
	"nosplace/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

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
		"pixels": formatPixelsToBytes(pixelList, x, y, size),
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

func formatPixelsToBytes(pixels []models.Pixel, x, y, size int64) []int {
	var result []byte
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
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			if pixelMap[i] != nil && pixelMap[i][j] != nil {
				result = append(result, pixelMap[i][j]...)
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
