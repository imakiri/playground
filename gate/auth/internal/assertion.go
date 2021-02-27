package internal

const Type_Assertion_ID = "Type_Assertion_ID"

type Assertion_ID string

func (Assertion_ID) Type() string {
	return Type_Assertion_ID
}

func (e Assertion_ID) Data() interface{} {
	return e
}

const Type_Assertion_Rand = "Type_Assertion_Rand"

type Assertion_Rand string

func (Assertion_Rand) Type() string {
	return Type_Assertion_Rand
}

func (e Assertion_Rand) Data() interface{} {
	return e
}
