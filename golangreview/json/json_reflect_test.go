package json_test

// 尝试实现一个marshal和unmarshaler

type Marshaler interface{
	MarshalerJson()([]byte,error)
}


type Unmarshaler interface{
	UnmarshalerJson([]byte)error
}

type Author struct{
	Name string `json:"name,omitempty"`
	Age int32 `json:"age,string,omitempty"`
}

// 从json的格式出发，json是一个树形数据结构，根据树形的数据结构进行结构体的序列化就需要对其进行自顶向下的递归操作，递归操作的细分操作就是反射，根据反射获取json的键对应结构体的字段，根据反射将结构体的字段的值，转换成json键的值。
// 其中tag则是对其进行字段分解动作，通过第一个字段得为键，剩余的字段则作为选项进行判断。