package main

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/VanThen60hz/big-number-addition/mybignumber"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Cấu hình logger
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/add", func(c *gin.Context) {
		num1 := c.PostForm("number1")
		num2 := c.PostForm("number2")

		var logBuffer bytes.Buffer

		hook := &LogBufferHook{buffer: &logBuffer}
		log.AddHook(hook)

		bn := &mybignumber.MyBigNumber{}
		result := bn.Sum(num1, num2)

		log.StandardLogger().ReplaceHooks(make(log.LevelHooks))

		steps := processLogToVietnamese(logBuffer.String())

		c.JSON(http.StatusOK, gin.H{
			"result": result,
			"steps":  steps,
		})
	})

	r.Run(":8080")
}

type LogBufferHook struct {
	buffer *bytes.Buffer
}

func (h *LogBufferHook) Levels() []log.Level {
	return []log.Level{log.InfoLevel}
}

func (h *LogBufferHook) Fire(entry *log.Entry) error {
	line, err := entry.String()
	if err == nil {
		h.buffer.WriteString(line)
	}
	return err
}

func processLogToVietnamese(logContent string) string {
	var result strings.Builder
	lines := strings.Split(logContent, "\n")

	stepPattern := regexp.MustCompile(`time="[^"]*"\s+level=info\s+msg="Processing step"\s+carry_in=(\d+)\s+carry_out=(\d+)\s+current_digit=(\d+)\s+digit1=(\d+)\s+digit2=(\d+)\s+position=(\d+)\s+step=(\d+)\s+total=(\d+)`)

	for _, line := range lines {
		if matches := stepPattern.FindStringSubmatch(line); len(matches) > 0 {
			carryIn := matches[1]
			carryOut := matches[2]
			currentDigit := matches[3]
			digit1 := matches[4]
			digit2 := matches[5]
			pos := matches[6]
			step := matches[7]
			total := matches[8]

			explanation := fmt.Sprintf(
				"Bước %s: Tính toán tại vị trí %s từ phải qua\n"+
					"- Chữ số của số thứ nhất: %s\n"+
					"- Chữ số của số thứ hai: %s\n"+
					"- Số nhớ từ bước trước: %s\n"+
					"- Tổng tạm tính: %s + %s + %s = %s\n"+
					"- Ghi nhận chữ số: %s\n"+
					"- Nhớ sang bước tiếp theo: %s\n"+
					"----------------------------------------\n",
				step, pos, digit1, digit2, carryIn, digit1, digit2, carryIn, total, currentDigit, carryOut)

			result.WriteString(explanation)
		}
	}

	fmt.Println("Processed result:", result.String()) // Debug log
	return result.String()
}
