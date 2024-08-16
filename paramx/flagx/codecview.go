package flagx

type MvFlag = Flag

const (
	MvFlag_pf MvFlag = "pf"
	//forward predicted MVs of P-frames

	MvFlag_bf MvFlag = "bf"
	//forward predicted MVs of B-frames

	MvFlag_bb MvFlag = "bb"
	//backward predicted MVs of B-frames

)

type MvTypeFlag = Flag
type MvtFlag = MvTypeFlag

const (
	MvFlag_fp MvFlag = "fp"
	MvFlag_bp MvFlag = "bp"
)
