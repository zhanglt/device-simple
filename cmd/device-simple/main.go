// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018-2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a simple example of a device service.
package main

import (
	"github.com/edgexfoundry/device-sdk-go/pkg/startup"
	"github.com/zhanglt/device-simple"
	"github.com/zhanglt/device-simple/driver"
)

const (
	serviceName string = "device-simple"
)

func main() {
	//创建驱动协议驱动结构体（该结构体必须实现ProtocolDriver接口）
	sd := driver.SimpleDriver{}
	//sd := driver.Driver{}
	//引入sdk包
	//将设备服务名，版本与该结构体传入BootStrap启动设备服务，Bootstrap继续调用startService启动服务。
	startup.Bootstrap(serviceName, device.Version, &sd)
}
