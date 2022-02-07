package runtime

import (
	"sort"
	"strings"

	"github.com/apmckinlay/gsuneido/lexer"
	"github.com/apmckinlay/gsuneido/runtime/types"
	"github.com/apmckinlay/gsuneido/util/hmap"
	"github.com/apmckinlay/gsuneido/util/ints"
	"github.com/apmckinlay/gsuneido/util/pack"
	"github.com/apmckinlay/gsuneido/util/varint"
)

// SuObject is a Suneido object
// i.e. a container with both list and named members
// Zero value is a valid empty object
// NOTE: Not thread safe
type SuObject struct {
	list     []Value
	named    hmap.Hmap
	readonly bool
	defval   Value
	CantConvert
}

// NewSuObject creates an SuObject from its arguments
func NewSuObject(args ...Value) *SuObject {
	ob := &SuObject{list: make([]Value, len(args))}
	for i, arg := range args {
		ob.list[i] = arg
	}
	return ob
}

func (ob *SuObject) Copy() *SuObject {
	return &SuObject{
		list:   append(ob.list[:0:0], ob.list...),
		named:  *ob.named.Copy(),
		defval: ob.defval,
	}
}

var _ Value = (*SuObject)(nil)

// Get returns the value associated with a key, or defval if not found
func (ob *SuObject) Get(_ *Thread, key Value) Value {
	return ob.GetDefault(key, ob.defval)
}

func (ob *SuObject) GetDefault(key Value, def Value) Value {
	val := ob.getIfPresent(key)
	if val == nil {
		//TODO handle copying object default
		return def
	}
	return val
}

func (ob *SuObject) getIfPresent(key Value) Value {
	if i, ok := key.IfInt(); ok && 0 <= i && i < len(ob.list) {
		return ob.list[i]
	}
	x := ob.named.Get(key)
	if x == nil {
		return nil
	}
	return x.(Value)
}

// Has returns true if the object contains the given key
func (ob *SuObject) Has(key Value) bool {
	return ob.getIfPresent(key) != nil
}

// ListGet returns a value from the list, panics if index out of range
func (ob *SuObject) ListGet(i int) Value {
	return ob.list[i]
}

// Put adds or updates the given key and value
// The value will be added to the list if the key is the "next"
func (ob *SuObject) Put(key Value, val Value) {
	ob.mustBeMutable()
	if i, ok := key.IfInt(); ok {
		if i == len(ob.list) {
			ob.Add(val)
			return
		} else if 0 <= i && i < len(ob.list) {
			ob.list[i] = val
			return
		}
	}
	ob.named.Put(key, val)
}

// Delete removes a key.
// If in the list, following list values are shifted over.
func (ob *SuObject) Delete(key Value) {
	ob.mustBeMutable()
	if i, ok := key.IfInt(); ok && 0 <= i && i < len(ob.list) {
		newlist := ob.list[:i+copy(ob.list[i:], ob.list[i+1:])]
		ob.list[len(ob.list)-1] = nil // aid garbage collection
		ob.list = newlist
	} else {
		ob.named.Del(key)
	}
}

// Erase removes a key.
// If in the list, following list values are NOT shifted over.
func (ob *SuObject) Erase(key Value) {
	ob.mustBeMutable()
	if i, ok := key.IfInt(); ok && 0 <= i && i < len(ob.list) {
		// migrate following list elements to named
		for j := len(ob.list) - 1; j > i; j-- {
			ob.named.Put(SuInt(j), ob.list[j])
			ob.list[j] = nil // aid garbage collection
		}
		ob.list = ob.list[:i]
	} else {
		ob.named.Del(key)
	}
}

// if (m.int_if_num(&i) && 0 <= i && i < vec.size()) {
// 	// migrate from vec to map
// 	for (int j = vec.size() - 1; j > i; --j)
// 		map[j] = vec[j];
// 	vec.erase(vec.begin() + i, vec.end());
// 	return true;
// }
// return map.erase(m);

// Clear removes all the contents of the object, making it empty (size 0)
func (ob *SuObject) Clear() {
	ob.mustBeMutable()
	ob.list = []Value{}
	ob.named = hmap.Hmap{}
}

func (ob *SuObject) RangeTo(from int, to int) Value {
	size := ob.Size()
	from = prepFrom(from, size)
	to = prepTo(from, to, size)
	return ob.rangeTo(from, to)
}

func (ob *SuObject) RangeLen(from int, n int) Value {
	size := ob.Size()
	from = prepFrom(from, size)
	n = prepLen(n, size-from)
	return ob.rangeTo(from, from+n)
}

func (ob *SuObject) rangeTo(i int, j int) *SuObject {
	list := make([]Value, j-i)
	copy(list, ob.list[i:j])
	return &SuObject{list: list}
}

func (ob *SuObject) ToObject() (*SuObject, bool) {
	return ob, true
}

func (ob *SuObject) ListSize() int {
	return len(ob.list)
}

func (ob *SuObject) NamedSize() int {
	return ob.named.Size()
}

// Size returns the number of values in the object
func (ob *SuObject) Size() int {
	return len(ob.list) + ob.named.Size()
}

// Add appends a value to the list portion
func (ob *SuObject) Add(val Value) {
	ob.mustBeMutable()
	ob.list = append(ob.list, val)
	ob.migrate()
}

// Insert inserts at the given position
// If the position is within the list, following values are move over
func (ob *SuObject) Insert(at int, val Value) {
	ob.mustBeMutable()
	if 0 <= at && at <= len(ob.list) {
		// insert into list
		ob.list = append(ob.list, nil)
		copy(ob.list[at+1:], ob.list[at:])
		ob.list[at] = val
	} else {
		ob.Put(IntVal(at), val)
	}
	ob.migrate()
}

func (ob *SuObject) mustBeMutable() {
	if ob.readonly {
		panic("can't modify readonly objects")
	}
}

func (ob *SuObject) migrate() {
	for {
		x := ob.named.Del(IntVal(len(ob.list)))
		if x == nil {
			break
		}
		ob.list = append(ob.list, x.(Value))
	}
}

func (ob *SuObject) String() string {
	buf, sep := ob.vecstr()
	iter := ob.named.Iter()
	for {
		k, v := iter()
		if k == nil {
			break
		}
		sep = entstr(buf, k, v, sep)
	}
	buf.WriteString(")")
	return buf.String()
}

func (ob *SuObject) vecstr() (*strings.Builder, string) {
	buf := strings.Builder{}
	sep := ""
	buf.WriteString("#(")
	for _, v := range ob.list {
		buf.WriteString(sep)
		sep = ", "
		buf.WriteString(v.String())
	}
	return &buf, sep
}

func entstr(buf *strings.Builder, k interface{}, v interface{}, sep string) string {
	buf.WriteString(sep)
	sep = ", "
	if ks, ok := k.(SuStr); ok && unquoted(string(ks)) {
		buf.WriteString(string(ks))
	} else {
		buf.WriteString(k.(Value).String())
	}
	buf.WriteString(":")
	if v != True {
		buf.WriteString(" ")
		buf.WriteString(v.(Value).String())
	}
	return sep
}

func unquoted(s string) bool {
	// want true/false to be quoted to avoid ambiguity
	return (s != "true" && s != "false") && lexer.IsIdentifier(s)
}

func (ob *SuObject) Show() string {
	buf, sep := ob.vecstr()
	mems := []Value{}
	iter := ob.named.Iter()
	for {
		k, _ := iter()
		if k == nil {
			break
		}
		mems = append(mems, k.(Value))
	}
	sort.Slice(mems,
		func(i, j int) bool { return mems[i].Compare(mems[j]) < 0 })
	for _, k := range mems {
		v := ob.named.Get(k)
		sep = entstr(buf, k, v, sep)
	}
	buf.WriteString(")")
	return buf.String()
}

func (ob *SuObject) Hash() uint32 {
	hash := ob.Hash2()
	if len(ob.list) > 0 {
		hash = 31*hash + ob.list[0].Hash()
	}
	if 0 < ob.named.Size() && ob.named.Size() <= 4 {
		iter := ob.named.Iter()
		for {
			k, v := iter()
			if k == nil {
				break
			}
			hash = 31*hash + k.(Value).Hash2()
			hash = 31*hash + v.(Value).Hash2()
		}
	}
	return hash
}

// Hash2 is shallow so prevents infinite recursion
func (ob *SuObject) Hash2() uint32 {
	hash := uint32(17)
	hash = 31*hash + uint32(ob.named.Size())
	hash = 31*hash + uint32(len(ob.list))
	return hash
}

func (ob *SuObject) Equal(other interface{}) bool {
	if val, ok := other.(Value); ok {
		if ob2, ok := val.ToObject(); ok {
			return soEqual(ob, ob2, newpairs())
		}
	}
	return false
}

func soEqual(x *SuObject, y *SuObject, inProgress pairs) bool {
	if x == y { // pointer comparison
		return true // same object
	}
	if x.ListSize() != y.ListSize() || x.NamedSize() != y.NamedSize() {
		return false
	}
	if inProgress.contains(x, y) {
		return true
	}
	inProgress.push(x, y) // no need to pop due to pass by value
	for i := 0; i < x.ListSize(); i++ {
		if !equals3(x.list[i], y.list[i], inProgress) {
			return false
		}
	}
	if x.NamedSize() > 0 {
		iter := x.named.Iter()
		for {
			k, v := iter()
			if k == nil {
				break
			}
			yk := y.named.Get(k)
			if yk == nil || !equals3(v.(Value), yk.(Value), inProgress) {
				return false
			}
		}
	}
	return true
}

func equals3(x Value, y Value, inProgress pairs) bool {
	if xo, ok := x.ToObject(); ok {
		if yo, ok := y.ToObject(); ok {
			return soEqual(xo, yo, inProgress)
		}
	}
	if xi, ok := x.(*SuInstance); ok {
		if yi, ok := y.(*SuInstance); ok {
			return siEqual(xi, yi, inProgress)
		}
	}
	return x.Equal(y)
}

func (SuObject) Type() types.Type {
	return types.Object
}

func (ob *SuObject) Compare(other Value) int {
	if cmp := ints.Compare(ordObject, Order(other)); cmp != 0 {
		return cmp
	}
	// now know other is an object so ToObject won't panic
	return cmp2(ob, ToObject(other), newpairs())
}

func cmp2(x *SuObject, y *SuObject, inProgress pairs) int {
	if x == y { // pointer comparison
		return 0
	}
	if inProgress.contains(x, y) {
		return 0
	}
	inProgress.push(x, y) // no need to pop due to pass by value
	for i := 0; i < x.Size() && i < y.Size(); i++ {
		if cmp := cmp3(x.ListGet(i), y.ListGet(i), inProgress); cmp != 0 {
			return cmp
		}
	}
	return ints.Compare(x.Size(), y.Size())
}

func cmp3(x Value, y Value, inProgress pairs) int {
	xo, xok := x.ToObject()
	yo, yok := y.ToObject()
	if !xok || !yok {
		return x.Compare(y)
	}
	return cmp2(xo, yo, inProgress)
}

func (*SuObject) Call(*Thread, *ArgSpec) Value {
	panic("can't call Object")
}

// ObjectMethods is initialized by the builtin package
var ObjectMethods Methods

var gnObjects = Global.Num("Objects")

func (*SuObject) Lookup(method string) Value {
	return Lookup(ObjectMethods, gnObjects, method)
}

// Slice returns a copy of the object, with the first n list elements removed
func (ob *SuObject) Slice(n int) *SuObject {
	newNamed := ob.named.Copy()
	if n > len(ob.list) {
		return &SuObject{named: *newNamed, defval: ob.defval}
	}
	newList := append(ob.list[:0:0], ob.list[n:]...)
	return &SuObject{list: newList, named: *newNamed, defval: ob.defval}
}

// ArgsIter is similar to Iter2 but it returns a nil key for list elements
func (ob *SuObject) ArgsIter() func() (Value, Value) {
	next := 0
	named := ob.named.Iter()
	return func() (Value, Value) {
		i := next
		next++
		if i < len(ob.list) {
			return nil, ob.list[i]
		}
		key, val := named()
		if key == nil {
			return nil, nil
		}
		return key.(Value), val.(Value)
	}
}

func (ob *SuObject) Iter2() func() (Value, Value) {
	next := 0
	named := ob.named.Iter()
	return func() (Value, Value) {
		i := next
		next++
		if i < len(ob.list) {
			return SuInt(i), ob.list[i]
		}
		key, val := named()
		if key == nil {
			return nil, nil
		}
		return key.(Value), val.(Value)
	}
}

func (ob *SuObject) Iter() Iter { // Values
	return &obIter{ob: ob, iter: ob.Iter2(),
		result: func(k, v Value) Value { return v }}
}

func (ob *SuObject) IterMembers() Iter {
	return &obIter{ob: ob, iter: ob.Iter2(),
		result: func(k, v Value) Value { return k }}
}

func (ob *SuObject) IterAssocs() Iter {
	return &obIter{ob: ob, iter: ob.Iter2(),
		result: func(k, v Value) Value { return NewSuObject(k, v) }}
}

type obIter struct {
	ob     *SuObject
	iter   func() (Value, Value)
	result func(Value, Value) Value
}

func (it *obIter) Next() Value {
	//TODO check for modification during iteration
	k, v := it.iter()
	if v == nil {
		return nil
	}
	return it.result(k, v)
}
func (it *obIter) Dup() Iter {
	return &obIter{ob: it.ob, iter: it.ob.Iter2(), result: it.result}
}
func (it *obIter) Infinite() bool {
	return false
}

func (ob *SuObject) Sort(t *Thread, lt Value) {
	ob.mustBeMutable()
	if lt == False {
		sort.SliceStable(ob.list, func(i, j int) bool {
			return ob.list[i].Compare(ob.list[j]) < 0
		})
	} else {
		sort.SliceStable(ob.list, func(i, j int) bool {
			return True == t.CallWithArgs(lt, ob.list[i], ob.list[j])
		})
	}
}

func (ob *SuObject) SetReadOnly() {
	if ob.readonly {
		return
	}
	ob.readonly = true
	iter := ob.Iter2()
	for k, v := iter(); k != nil; k, v = iter() {
		if x, ok := v.ToObject(); ok {
			x.SetReadOnly()
		}
	}
}

func (ob *SuObject) IsReadOnly() bool {
	return ob.readonly
}

func (ob *SuObject) SetDefault(def Value) {
	ob.mustBeMutable()
	ob.defval = def
}

// Packable ---------------------------------------------------------

var _ Packable = (*SuObject)(nil)

const packNestLimit = 20

func (ob *SuObject) PackSize(nest int) int {
	nest++
	if nest > packNestLimit {
		panic("pack: object nesting limit exceeded")
	}
	if ob.Size() == 0 {
		return 1 // just tag
	}
	ps := 1 // tag
	ps += varint.Len(uint64(ob.ListSize()))
	for _, v := range ob.list {
		ps += packSize(v, nest)
	}
	ps += varint.Len(uint64(ob.NamedSize()))
	iter := ob.named.Iter()
	for k, v := iter(); k != nil; k, v = iter() {
		ps += packSize(k.(Value), nest) + packSize(v.(Value), nest)
	}
	return ps
}

func packSize(x Value, nest int) int {
	if p, ok := x.(Packable); ok {
		n := p.PackSize(nest)
		return varint.Len(uint64(n)) + n
	}
	panic("can't pack " + x.Type().String())
}

func (ob *SuObject) Pack(buf *pack.Encoder) {
	ob.pack(buf, PackObject)
}

func (ob *SuObject) pack(buf *pack.Encoder, tag byte) {
	buf.Put1(tag)
	if ob.Size() == 0 {
		return
	}
	buf.VarUint(uint64(ob.ListSize()))
	for _, v := range ob.list {
		packValue(v, buf)
	}
	buf.VarUint(uint64(ob.NamedSize()))
	iter := ob.named.Iter()
	for k, v := iter(); k != nil; k, v = iter() {
		packValue(k.(Value), buf)
		packValue(v.(Value), buf)
	}
}

func packValue(x Value, buf *pack.Encoder) {
	n := x.(Packable).PackSize(0)
	buf.VarUint(uint64(n))
	x.(Packable).Pack(buf)
}

func UnpackObject(s string) *SuObject {
	return unpackObject(s, &SuObject{})
}

func unpackObject(s string, ob *SuObject) *SuObject {
	if len(s) <= 1 {
		return ob
	}
	buf := pack.NewDecoder(s[1:])
	var v Value
	n := int(buf.VarUint())
	for i := 0; i < n; i++ {
		v = unpackValue(buf)
		ob.Add(v)
	}
	var k Value
	n = int(buf.VarUint())
	for i := 0; i < n; i++ {
		k = unpackValue(buf)
		v = unpackValue(buf)
		ob.Put(k, v)
	}
	return ob
}

func unpackValue(buf *pack.Decoder) Value {
	size := int(buf.VarUint())
	return Unpack(buf.Get(size))
}

// old format

func UnpackObjectOld(s string) *SuObject {
	return unpackObjectOld(s, &SuObject{})
}

func unpackObjectOld(s string, ob *SuObject) *SuObject {
	if len(s) <= 1 {
		return ob
	}
	buf := pack.NewDecoder(s[1:])
	var v Value
	n := buf.Int32()
	for i := 0; i < n; i++ {
		v = unpackValueOld(buf)
		ob.Add(v)
	}
	var k Value
	n = buf.Int32()
	for i := 0; i < n; i++ {
		k = unpackValueOld(buf)
		v = unpackValueOld(buf)
		ob.Put(k, v)
	}
	return ob
}

func unpackValueOld(buf *pack.Decoder) Value {
	size := buf.Int32()
	return UnpackOld(buf.Get(size))
}
