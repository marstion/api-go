package jsonvalue

import "encoding/json"

type J struct {
	obj interface{}
}

// json.Unmarshal(response, &b.apiBus) // json 解析到对象
// json.Marshal()                      // 对象解析到json

/*
Marshal returns the JSON encoding of v.

将任意对象v, 转换成json

参数:
	v := map[string]interface {}{"data":map[string]interface {}{"age":18, "name":"Jack"}, "msg":0}
返回:
	destByte := []byte(`{"msg": 0, "data": {"name": "Jack", "age": 18}}`)
	nil := nil
*/
func (j *J) Marshal(v *interface{}) (destByte []byte, err error) {
	/*
		obj => json, 例:

	*/
	destByte, err = json.Marshal(&v)
	return
}

// func (j J) Unmarshal2Obj(srcObj interface{}, destObj interface{}) (err error) {
// 	/*
// 		将 srcObj => []byte => destObj
// 		例:
// 			srcObj: map[string]interface {}{"data":map[string]interface {}{"age":18, "name":"Jack"}, "msg":0}
// 			destObj: obj
// 	*/
// 	tobyte, err := j.Marshal(&srcObj)
// 	if err != nil {
// 		return
// 	}
// 	err = json.Unmarshal(tobyte, &destObj)
// 	return
// }

/*
Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.

解析json编码数据, 并将结果存储到v指向的值中.
*/
func (j J) Unmarshal(srcByte []byte, v *interface{}) (err error) {
	err = json.Unmarshal(srcByte, &v)
	return
}

/*
Unmarshal parses the j.obj data and stores the result in the value pointed to by v.

解析j.obj数据, 并将结果存储到v指向的值中.
*/
func (j J) Unmarshal1Self(v interface{}) (err error) {
	byt, err := j.Marshal(&j.obj)
	if err != nil {
		return
	}
	err = json.Unmarshal(byt, &v)
	return
}

/*
Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by j.obj.

解析json编码数据, 并将结果存储到j.obj指向的值中.
*/
func (j *J) Unmarshal2Self(srcByte []byte) (err error) {
	err = j.Unmarshal(srcByte, &j.obj)
	return
}

func (j J) Get(args ...interface{}) *J {
	for _, arg := range args {
		switch arg := arg.(type) {
		case string:
			j.obj = j.obj.(map[string]interface{})[arg]
		case int:
			j.obj = j.obj.([]interface{})[arg]
		}
	}
	return &j
}

func (j J) Map() map[string]*J {
	arr := make(map[string]*J)
	switch v := j.obj.(type) {
	case map[string]interface{}:
		for k, vv := range v {
			arr[k] = &J{obj: vv}
		}
	}
	return arr
}

func (j J) Array() []*J {
	var arr []*J
	switch v := j.obj.(type) {
	case []interface{}:
		for _, vv := range v {
			arr = append(arr, &J{obj: vv})
		}
	}
	return arr
}

func (j J) String() string {
	return j.obj.(string)
}

func (j J) Float64() float64 {
	return j.obj.(float64)
}

func (j J) Float32() float32 {
	return float32(j.Float64())
}

func (j J) Integer() int {
	return int(j.Float64())
}

func (j J) Bool() bool {
	return j.obj.(bool)
}
