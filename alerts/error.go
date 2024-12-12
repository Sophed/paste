package alerts

import "github.com/sophed/lg"

func Error(err interface{}) {
	lg.Warn(err)
}
