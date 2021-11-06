package timeago

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTimeCalculations(t *testing.T) {
	seconds := 62
	minutes, hours, days, weeks, months, years := getTimeCalculations(float64(seconds))
	fmt.Printf("%d => minutes:%d, hours:%d, days:%d, weeks:%d, months:%d, years:%d\n", seconds, minutes, hours, days, weeks, months, years)
	assert.Equal(t, 1, minutes)
	assert.Equal(t, 0, hours)
	assert.Equal(t, 0, days)
	assert.Equal(t, 0, weeks)
	assert.Equal(t, 0, months)
	assert.Equal(t, 0, years)

	result := calculateTheResult(62, ``, `en`)
	assert.Equal(t, `1 minute ago`, result)

	datetime := smallSubTime(-60 * time.Second)

	seconds = getSeconds(datetime, ``)
	assert.Equal(t, 60, seconds)

	result = Take(datetime, `en`)
	fmt.Println(datetime)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	assert.Equal(t, `1 minute ago`, result)

	datetime = smallSubTime(time.Second)
	seconds = getSeconds(datetime, ``)
	assert.Equal(t, 0, seconds)
	result = Take(datetime, `en`)
	assert.Equal(t, `0 seconds ago`, result)

	datetime = smallSubTime(time.Second * 2)
	seconds = getSeconds(datetime, ``)
	assert.Equal(t, -1, seconds)
	result = Take(datetime, `en`)
	assert.Equal(t, `0 seconds ago`, result)
}
