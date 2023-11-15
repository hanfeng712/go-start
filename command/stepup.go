package main

type SetupFunc func() error
type Setup []SetupFunc

func (s Setup) apply() {
	for _, fn := range s {
		if err := fn(); err != nil {
			// TODO 缺少日志
		}
	}
}

func SetupDB() error {
	return nil
}

func SetupJwt() error {
	return nil
}

func SetupLogger() error {
	return nil
}
