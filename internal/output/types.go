package output

import "eol-checker/internal/api"

type OutputInterface interface {
	Write(checks []api.EolCheck)
}
