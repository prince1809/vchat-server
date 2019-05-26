package jobs

type Schedulers struct {
	Stop    chan bool
	Stopped chan bool
}
