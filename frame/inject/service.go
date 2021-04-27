package inject

//IStartService 定义IStartService接口
type IStartService interface {
	Say(message string) string
}

//StartService 注入IStartRepo
type StartService struct {
	Repo IStartRepo `inject:""`
}

//Say 实现Say方法
func (s *StartService) Say(message string) string {
	return s.Repo.Speak(message)
}
