package constant

const (
	USER_TYPE_SUPERVISOR string = "supervisor"
	USER_TYPE_USER       string = "user"
	USER_TYPE_ADMIN      string = "admin"
)

const (
	USER_STATUS_FIRESTORE_PENDING  int = -10 // migrated user needs changing password
	USER_STATUS_DEACTIVE           int = -1
	USER_STATUS_WAITING_ACTIVATION int = 0  // waiting activation by confirming email
	USER_STATUS_WAITING_APPROVAL   int = 10 // waiting admin's approval
	USER_STATUS_ACTIVE             int = 999
)
