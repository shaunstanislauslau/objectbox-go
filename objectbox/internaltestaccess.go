package objectbox

// Reserved for internal unit tests only
type InternalTestAccessObjectBox struct {
	ObjectBox *ObjectBox
}

func (obx *InternalTestAccessObjectBox) RunInTxn(readOnly bool, txnFun txnFun) (err error) {
	return obx.ObjectBox.runInTxn(readOnly, txnFun)
}