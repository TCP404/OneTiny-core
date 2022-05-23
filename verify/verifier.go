package verify

type Verifier interface {
	AddToHead(Handler) Verifier
	AddToTail(Handler) Verifier
	Iterator() error
}

var _ Verifier = (*verifyChain)(nil)

func NewVerifyChain() Verifier {
	return &verifyChain{}
}
