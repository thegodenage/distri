package core

type Maybe struct {
	Value any
	Err   error
}

func MaybeWithErr(err error) Maybe {
	return Maybe{Err: err}
}

func MaybeWithVal(v any) Maybe {
	return Maybe{Value: v, Err: nil}
}
