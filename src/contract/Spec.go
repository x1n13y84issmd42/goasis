package contract

// OperationIterator is an iterable channel to receive operations from.
type OperationIterator chan Operation

// Spec is an interface to access specification data.
type Spec interface {
	ProjectInfo

	Operations() OperationIterator
}

//TODO: make spec a set of operations.
