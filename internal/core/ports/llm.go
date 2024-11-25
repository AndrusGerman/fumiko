package ports

type LLM interface {
	BasicQuest(text string) string
}
