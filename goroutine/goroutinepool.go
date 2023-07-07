package main

type Pool struct {
	worker chan func()
	sem    chan struct{} // 用来控制数量
}

func NewPool(num int) *Pool {
	return &Pool{
		worker: make(chan func()),
		sem:    make(chan struct{}, num),
	}
}

func (p *Pool) NewTask(task func()) {
	select {
	case p.worker <- task:
	case p.sem <- struct{}{}:
		go p.work(task)
	}
}

func (p *Pool) work(task func()) {
	defer func() { <-p.sem }()
	for {
		task()
		task = <-p.worker
	}
}
