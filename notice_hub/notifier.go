package notice_hub

type Notifier interface {
	Notify(string)
}

type NoneNotifier struct {
}

func (n *NoneNotifier) Notify(msg string) {
	// Nothing to do
}
