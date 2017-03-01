package thread_manage

type ChanThreadManager struct {
	total uint
	mc  chan uint
}

func NewChanThreadManager(total uint)*ChanThreadManager{
	mc := make(chan uint, total)
	return &ChanThreadManager{mc:mc,total:total}
}


func (this *ChanThreadManager)Get(){
	this.mc<-1
}

func (this *ChanThreadManager)Free(){
	<-this.mc
}

func (this *ChanThreadManager)Total()uint{
	return this.total
}

func (this *ChanThreadManager)Left()uint{
	return this.total-uint(len(this.mc))
}