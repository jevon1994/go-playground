


### Release Notes

#### 1.0.0
- feat
  - 推送事件
  - 拉取事件
  - 支持 URL 回调方式
- todo
  - 支持 MQ 协议发布订阅模式


### 使用步骤

1. 引入包
```shell
go get github.com/asaskevich/EventBus
go get github.com/cloudevents/sdk-go/v2
```
2. 初始化
```go
func TestClient(t *testing.T) {
    ...
    client := NewStreamServiceClient()
    initGin(client.InitEngineofGin)
}
```

3. 实现

sender
```go
// 发送部署事件
func doPush() {
    data := make(map[string]interface{})
    marshal, _ := json.Marshal(data)
    data["test_record_id"] = "A111-122323"
    
    m := make(map[string]interface{})
    m["idc"] = "unicom"
    m["env"] = "dev"
    m["namespace"] = "unicom"
    m["service"] = "cloud-hcce"
    
    request := StreamPushRequest{
    Type:       "cloud-hcce.deployments.create",
    Url:        event.GetStreamHandlerURLofLocalService(),
    Extensions: m,
    Bytes:      marshal,
    }
    client := NewStreamGinServiceClient()
    push, _ := client.Push(context.Background(), &request)
    fmt.Println(push)
}
```
receiver
```go
type TestEventHandler struct {
}

func (th *TestEventHandler) Handle(event *cloudevents.Event) error {
	fmt.Println(event)
	return nil
}

func init() {
	ReqBus.AddHandler("cloud-hcce.deployments.create", &TestEventHandler{})
}
```

