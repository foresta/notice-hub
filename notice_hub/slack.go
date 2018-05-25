package notice_hub

type NotifyMethod int

const (
	WebAPI NotifyMethod = iota
	Webhook
)

func Slack(nm NotifyMethod, config string) Notifier {
	switch nm {
	case WebAPI:
		return &SlackWebAPI{ConfigFilePath: config}
	case Webhook:
		return &SlackWebhook{ConfigFilePath: config}

	default:
		return &NoneNotifier{}
	}
}
