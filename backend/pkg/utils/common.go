package utils

import "time"

func DoWithAttemps(fn func() error, attemps int, delay time.Duration) (err error) {
	for attemps > 0 {
		if err = fn(); err != nil {
			attemps--
			time.Sleep(delay)
			continue
		}

		return nil
	}

	return
}
