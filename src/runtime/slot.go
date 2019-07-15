package runtime

type Slot uint32
type SlotTable []Slot
type SlotStack struct {
	position int
	table SlotTable
}

func newSlotTable(length uint16) SlotTable{
	return make([]Slot,length)
}

func (self *SlotTable ) setInt(index uint16,num uint32){
	(*self)[index] = Slot(num)
}

type Object interface {
}

func (self *SlotTable ) setObjectRef(index uint16,ref *Object){
	//(*self)[index] = Slot(ref)
}

func (self *SlotTable ) setLong(index uint16,num uint64){
	(*self)[index] = Slot(num)
	(*self)[index+1] = Slot(num>>32)
}

func (self *SlotTable ) setValue(index uint16,table2 *SlotTable,t2index uint16){
	(*self)[index] = (*table2)[t2index]
}
func (self *SlotTable ) setValue2(index uint16,table2 *SlotTable,t2index uint16){
	(*self)[index] = (*table2)[t2index]
	(*self)[index+1] = (*table2)[t2index+1]
}

func (self *Slot) getInt() uint32{
	return uint32(*self)
}

/*

func (self *Slot) setLong(bits uint64){
	ptr := self
	*ptr = bits[:4]
	ptr += 1  ;


}

func (self *Slot) setDouble(num double){
	var *ptr = self
	bits := math.Float64bits(num)
	*ptr = bits[:4]

}*/