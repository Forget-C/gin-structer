package model

const (
	ApproveUnknown = iota
	ApprovePassed
	ApproveRejected
	ApproveRunning
)

type UserRecord struct {
	ObjectMeta
	TimeMeta
	Name  string `json:"name" gorm:"index:,;size:64;comment:用户名称"`
	Title string `json:"title" gorm:"size:64;comment:用户职务"`
}

type ApproveRecode struct {
	ObjectMeta
	TimeMeta
	RequestUserName string `json:"request_user_name" gorm:"index:,;size:64;comment:提交人"`
	RequestUserID   string `json:"request_user_id" gorm:"size:64;comment:提交人ID"`
	ApproveUserName string `json:"approve_user_name" gorm:"index:,;size:64:comment:审批人"`
	ApproveUserID   string `json:"approve_user_id" gorm:"size:64;comment:审批人ID"`
	Desc            string `json:"desc" gorm:"size:256;comment:审批内容"`
	Status          int    `json:"status" gorm:"index:,;size:64;comment:审批状态"`
	Comment         string `json:"comment" gorm:"size:256;comment:审批意见"`
}

func (a *ApproveRecode) SetPassed() {
	a.Status = ApprovePassed
}

func (a *ApproveRecode) SetRejected() {
	a.Status = ApproveRejected
}

func (a *ApproveRecode) SetRunning() {
	a.Status = ApproveRunning
}

func (a *ApproveRecode) GetStatus() string {
	switch a.Status {
	case ApprovePassed:
		return "passed"
	case ApproveRejected:
		return "rejected"
	case ApproveRunning:
		return "running"
	default:
		return "unknown"
	}
}
