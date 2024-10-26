package beef

type BeefModule struct {
	Ctl *BeefController
	Svc *BeefService
}

func New() *BeefModule {
	svc := newService()
	return &BeefModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
