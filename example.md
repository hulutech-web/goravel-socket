### 如何使用
#### mqtt消息通信场景，本例中使用了mqtt对物联网设备的数据进行解析，同时支持向物联网设备发送消息，当后台感知到消息的时候需要与web页面进行交互，此时需要使用websocket进行消息的传递。以下是核心代码，
- 该例反应SOCKET通信原理，通过websocket建立连接，当后台收到mqtt消息后，通过websocket将消息推送给前端页面，前端页面收到消息后，通过websocket将消息发送给后台，后台收到消息后，通过mqtt将消息发送给物联网设备。
#### 初始化MQTT服务
```go
func createServerClient() mqtt.Client {
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID("server_clientid")
	opts.SetKeepAlive(30 * time.Second)
	//设置协议版本为3.1
	opts.SetProtocolVersion(uint(3))

	//opts.SetUsername("energy")
	//opts.SetPassword("energy")
	opts.SetCleanSession(true)
	opts.OnConnect = ConnectHandler(opts.ClientID)
	opts.OnConnectionLost = ConnectLostHandler(opts.ClientID)
	return mqtt.NewClient(opts)
}

```

#### 服务端：订阅了2组topic，一组为心跳数据，一组为业务数据
```go
func StartServer() {
	// 创建 MQTT 客户端
	//topics := []string{"devices/+/data", "devices/+/command"}
	fmt.Println("================服务端已开启==============")
	topics := []string{"devices/+/data", "devices/+/heartbeat"}
	for _, topic := range topics {
		var handler mqtt.MessageHandler
		switch topic {
		case "devices/+/data":
			handler = handleTask
		case "devices/+/heartbeat":
			handler = handleHeartbeat
		}
		if token := mqttClient.Subscribe(topic, byte(0), handler); token.Wait() && token.Error() != nil {
			log.Printf("Failed to subscribe to %s: %v", topic, token.Error())
		}
	}

	// 启动清理无效连接,并将其删除定时任务
	go cleanup(mqttClient)

	// 等待信号（Ctrl+C）关闭服务器
	waitForServerSignal()

	// 断开连接
	mqttClient.Disconnect(250)
	WG.Wait()
}
```

#### 消息订阅处理
```go
func handleTask(client mqtt.Client, msg mqtt.Message) {
	// 异步处理消息
	WG.Add(1)
	go func(c mqtt.Client) {
		defer WG.Done()
		// 提取设备唯一标识
		topic := msg.Topic()
		did := extractDeviceID(topic)
		messageType := extractMessageType(topic)
		reader := c.OptionsReader()
		facades.Log().Debug("服务器接收数据：Received txt->client_id:"+reader.ClientID()+" payload[", string(msg.Payload()), "]from topic:", topic, " (messageType:", messageType, ")", " (Device ID:", did, ")")
		go ReceiveMqttMessage(msg.Payload(), did)
	}(client)
}
```

#### SOCKET实时感知通信
```go
func handleTask(client mqtt.Client, msg mqtt.Message) {
	// 异步处理消息
	WG.Add(1)
	go func(c mqtt.Client) {
		defer WG.Done()
		// 提取设备唯一标识
		topic := msg.Topic()
		did := extractDeviceID(topic)
		messageType := extractMessageType(topic)
		reader := c.OptionsReader()
		facades.Log().Debug("服务器接收数据：Received txt->client_id:"+reader.ClientID()+" payload[", string(msg.Payload()), "]from topic:", topic, " (messageType:", messageType, ")", " (Device ID:", did, ")")
		go ReceiveMqttMessage(msg.Payload(), did)
	}(client)
}
```
### 业务处理：使用socket进行通信，确保信息及时性,通过websocket与web页面进行实时交互  
关键方法
```go
servers.SendMessage2System(did, "1", 200, "回传", string(jsonStr))
servers.SendMessage2System(did, "1", 2001, "回传", "<-失败，指令是：CRC")
```
```go
// 处理服务器的数据
func ReceiveMqttMessage(data []byte, did string) {
	select {
	case message := <-ReadCmdMapChan[did]:

		//解析回传数据，需要给定一个时间，如果失败，则放弃，后续实现
		chipData, err := utils.HandleChipData(string(data))
		if err != nil {
			CompletionSignal <- false
			facades.Log().Channel("single").Error("解析数据失败：", err)
			servers.SendMessage2System(did, "1", 2001, "回传", "<-失败，指令是：CRC")
			return
		}
		facades.Log().Channel("single").Info("<-：原始数据：string:", string(data), "err:", err)

		txt, err := CallMethodsOnStruct(utils.ParseExplain, message.ReadCmd.DataAddress, chipData)
		CompletionSignal <- true
		if err != nil {
			CompletionSignal <- false
			servers.SendMessage2System(did, "1", 2001, "回传", "<-失败，方法调用失败或找不到")
			return
		} else {
			txtStr := ""
			for _, v := range txt {
				txtStr += fmt.Sprintf("%v", v) + ","
			}
			InsertChipDataToSql(did, chipData, txtStr)
		}
		resData := ResData{
			Cmd:         message.ReadCmd.Parameter,
			DataAddress: message.ReadCmd.DataAddress,
			ChipData:    chipData,
			ParseData:   txt,
			ErrStr:      "",
		}
		//解析为json字符串
		marshal, _ := json.Marshal(resData)
		jsonStr := marshal
		servers.SendMessage2System(did, "1", 200, "回传", string(jsonStr))
	}
}
```
#### 其他方法
```go
servers.SendMessage2Client(inputData.ClientId, inputData.SendUserId, inputData.Code, inputData.Msg, &inputData.Data) //发送给指定客户端
servers.SendMessage2Group(systemId, inputData.SendUserId, inputData.GroupName, inputData.Code, inputData.Msg, &inputData.Data) //发送给分组
SendMessage2System(systemId, sendUserId string, code int, msg string, data string) //发送给系统
```