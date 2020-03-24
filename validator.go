package cronvalidator

import (
	"github.com/robfig/cron/v3"
	"time"
)

func Validate(spec string) error {
	parser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	_, err := parser.Parse(spec)
	return err
}

func Next(spec string, times int) ([]string, error) {
	parser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	sc, err := parser.Parse(spec)
	if err != nil {
		return nil, err
	}

	var results []string
	next := time.Now()
	for i := 0; i < times; i++ {
		next = sc.Next(next)
		results = append(results, next.String())
	}

	return results, nil
}
