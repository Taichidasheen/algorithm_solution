package pool

import "sync"

type Task interface {
	Task()
}

//Pool 协程池
type Pool struct {
	size int
	tasks chan Task
	wg sync.WaitGroup
}

//New 初始化协程池大小
func New(size int) *Pool {
	return &Pool{
		size: size,
		tasks: make(chan Task),//注意：chan需要进行初始化
	}
}

//Add 添加任务
func (p *Pool) Add(task Task) {
	p.tasks <- task
}

//Run 执行任务
func (p *Pool) Run() {
	for i:=0;i<p.size;i++ {
		p.wg.Add(1)
		go func() {
			for task := range p.tasks {
				task.Task()
			}
			p.wg.Done()
		}()
	}
}

//Shutdown 结束运行
func (p *Pool) Shutdown() {
	close(p.tasks)
	p.wg.Wait()
}


