package utils

func Reduce[Input any, Acc any](
	inputs []Input,
	reducer func(Acc, Input) error,
	init Acc) error {
	for _, input := range inputs {
		err := reducer(init, input)
		if err != nil {
			return err
		}
	}
	return nil
}
