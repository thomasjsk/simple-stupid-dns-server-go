package dns

type Message struct {
	Header   Header
	Question Question
	Answer   Answer
}

func (m *Message) ToBytes() []byte {
	header := (*m).Header.toBytes()
	question := (*m).Question.ToBytes()
	answer := (*m).Answer.ToBytes()

	data := append(header, question...)
	return append(data, answer...)
}

func NewMessage(buf []byte) *Message {
	question := RawQuestion(buf[12:]).ToQuestion()

	return &Message{
		Header:   RawHeader(buf[0:12]).toHeader(),
		Question: question,
		Answer:   AnswerDomain(question.Name).ToAnswer(),
	}
}
