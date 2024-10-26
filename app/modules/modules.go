package modules

import "app/app/modules/beef"

type Modules struct {
	Beef *beef.BeefModule
}

func Get() *Modules {
	return &Modules{
		Beef: beef.New(),
	}
}
