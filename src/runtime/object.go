package runtime

import "../classloader"
import "../util"

type u1 util.U1

type Object struct {

	_class *classloader.Class
	fieldValues util.SlotTable
	flag u1
}
