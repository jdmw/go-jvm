package runtime

import(
	"../classloader"
	// "../util"
)
type Thread struct {
	stack *Stack
}

func NewThread(deepth uint) *Thread{
	return &Thread{
		stack : NewStack(deepth),
	}
}

func (self *Thread) LoadStaticMethod(method *classloader.Method) *StackFrame{
	currentFrame := NewStackFrameByMethod(self,method)
	self.stack.Push(currentFrame)
	return currentFrame
}

func (self *Thread) LoadInstanceMethod(instant *Object,method *classloader.Method) *StackFrame{
	currentFrame := NewStackFrameByInstanceMethod(self,instant,method)
	self.stack.Push(currentFrame)
	return currentFrame
}

