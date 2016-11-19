package thread_manage

type ThreadManager interface {
	Get()
	Free()
	Total() uint
	Left() uint
}
