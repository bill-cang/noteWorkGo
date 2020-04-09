package eunm

type PolicyType int32

const (
	Policy_MIN      PolicyType = 0
	Policy_MAX      PolicyType = 1
	Policy_MID      PolicyType = 2
	Policy_AVG      PolicyType = 3
)

type TaskBody struct {
	 Name string
	 Total int
}

func (p PolicyType) String() *TaskBody {
	switch (p) {
	case Policy_MIN: return &TaskBody{
		Name:  "MIN",
		Total: 0,
	}
	case Policy_MAX: return &TaskBody{
		Name:  "MAN",
		Total: 0,
	}
	case Policy_MID: return &TaskBody{
		Name:  "MID",
		Total: 0,
	}
	case Policy_AVG: return &TaskBody{
		Name:  "AVG",
		Total: 0,
	}
	default:         return nil
	}
}


