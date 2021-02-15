package driver

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
)

// Temperature Sensor 定义
type TemperatureSensor struct {
	lc      logger.LoggingClient // 用来发送日志
	enabled bool                 // 是否使能

	temperature float64
}

// NewTemperatureSensor 创建一个新的TemperatureSensor
func NewTemperatureSensor(lc logger.LoggingClient) *TemperatureSensor {
	return &TemperatureSensor{
		lc:      lc,
		enabled: true,
	}
}

// Enable 用来让传感器工作
func (t *TemperatureSensor) Enable() {
	if t.enabled {
		t.lc.Warn("temperature sensor already enabled")
		return
	}
	t.enabled = true
	t.lc.Info("temperature sensor enabled")
}

// Disable 用来让传感器工作
func (t *TemperatureSensor) Disable() {
	if !t.enabled {
		t.lc.Warn("temperature sensor already disabled")
		return
	}
	t.enabled = false
	t.lc.Info("temperature sensor disabled")
}

// GetTemperature 获取当前传感器的温度数据
func (t *TemperatureSensor) GetTemperature() (string, error) {
	//out, err := runCmd()
	//if err != nil {
	//  return "", err
	//} else {
	//  return strings.TrimSuffix(out, "\n"), nil
	//}
	files, err := ioutil.ReadDir("temperature")
	if err != nil {
		t.lc.Error("获取传感器温度失败", err)
		return "", err
	}
	var sensorId string
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "28") {
			sensorId = file.Name()
		}
	}

	path := fmt.Sprintf("temperature/%s", sensorId)
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	outs := strings.Split(string(b), " ")
	res := strings.Split(outs[len(outs)-1], "=")
	return strings.TrimSuffix(res[len(res)-1], "\n"), nil
}
