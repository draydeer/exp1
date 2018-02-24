package ads

import (
	"strconv"
	"strings"
)

type AdsNode interface {
	GetIndex(index string) (interface{}, bool)
	GetValue() interface{}
	HasIndex(index string) bool
	IsContainer() bool
}

// Ads list

type AdsList struct {
	Val []interface{} `json`
}

func (adsNode AdsList) GetIndex(index string) (interface{}, bool) {
	var i, err = strconv.Atoi(index)

	if err == nil && i >= 0 && len(adsNode.Val) > i {
		return adsNode.Val[i], true
	}

	return nil, false
}

func (adsNode AdsList) GetValue() interface{} {
	return adsNode.Val
}

func (adsNode AdsList) HasIndex(index string) bool {
	var i, err = strconv.Atoi(index)

	return err != nil && i >= 0 && len(adsNode.Val) > i
}

func (adsNode AdsList) IsContainer() bool {
	return true
}

func NewAdsList(l []interface{}) AdsNode {
	var node = AdsList{make([]interface{}, 0)}

	for _, v := range l {
		switch v.(type) {
		case map[string]interface{}:
			node.Val = append(node.Val, NewAdsMap(v.(map[string]interface{})))

			break;

		case []interface{}:
			node.Val = append(node.Val, NewAdsList(v.([]interface{})))

			break;

		default:
			node.Val = append(node.Val, AdsPrimitive{v})
		}
	}

	return node
}

func NewAdsListWithMapper(l []interface{}, mapper func(interface{}) (interface{}, error)) (AdsNode, error) {
	node := AdsList{make([]interface{}, 0)}

	for _, v := range l {
		v, err := mapper(v)

		if err != nil {
			return nil, err
		}

		switch v.(type) {
		case map[string]interface{}:
			val, err := NewAdsMapWithMapper(v.(map[string]interface{}), mapper)

			if err != nil {
				return nil, err
			}

			node.Val = append(node.Val, val)

			break;

		case []interface{}:
			val, err := NewAdsListWithMapper(v.([]interface{}), mapper)

			if err != nil {
				return nil, err
			}

			node.Val = append(node.Val, val)

			break;

		default:
			node.Val = append(node.Val, AdsPrimitive{v})
		}
	}

	return node, nil
}

// Ads map

type AdsMap struct {
	Val map[string]interface{} `json:"-"`
}

func (adsNode AdsMap) GetIndex(index string) (interface{}, bool) {
	var v, present = adsNode.Val[index]

	if present {
		return v, true
	}

	return nil, false
}

func (adsNode AdsMap) GetValue() interface{} {
	return adsNode.Val
}

func (adsNode AdsMap) HasIndex(index string) bool {
	var _, present = adsNode.Val[index]

	return present
}

func (adsNode AdsMap) IsContainer() bool {
	return true
}

func NewAdsMap(m map[string]interface{}) AdsNode {
	var node = AdsMap{make(map[string]interface{})}

	for k, v := range m {
		switch v.(type) {
		case map[string]interface{}:
			node.Val[k] = NewAdsMap(v.(map[string]interface{}))

			break;

		case []interface{}:
			node.Val[k] = NewAdsList(v.([]interface{}))

			break;

		default:
			node.Val[k] = AdsPrimitive{v}
		}
	}

	return node
}

func NewAdsMapWithMapper(m map[string]interface{}, mapper func(interface{}) (interface{}, error)) (AdsNode, error) {
	var node = AdsMap{make(map[string]interface{})}

	for k, v := range m {
		v, err := mapper(v)

		if err != nil {
			return nil, err
		}

		switch v.(type) {
		case map[string]interface{}:
			val, err := NewAdsMapWithMapper(v.(map[string]interface{}), mapper)

			if err != nil {
				return nil, err
			}

			node.Val[k] = val

			break;

		case []interface{}:
			val, err := NewAdsListWithMapper(v.([]interface{}), mapper)

			if err != nil {
				return nil, err
			}

			node.Val[k] = val

			break;

		default:
			node.Val[k] = AdsPrimitive{v}
		}
	}

	return node, nil
}

// Ads primitive

type AdsPrimitive struct {
	Val interface {}
}

func (adsNode AdsPrimitive) GetIndex(index string) (interface{}, bool) {
	return nil, false
}

func (adsNode AdsPrimitive) GetValue() interface{} {
	return adsNode.Val
}

func (adsNode AdsPrimitive) HasIndex(index string) bool {
	return false
}

func (adsNode AdsPrimitive) IsContainer() bool {
	return false
}

func NewAdsPrimitive(val interface{}) AdsNode {
	return AdsPrimitive{val}
}

func NewAdsPrimitiveWithMapper(val interface{}, mapper func(interface{}) (interface{}, error)) (AdsNode, error) {
	val, err := mapper(val)

	if err != nil {
		return nil, err
	}

	return AdsPrimitive{val}, nil
}

//

func Create(val interface{}) AdsNode {
	switch val.(type) {
	case map[string]interface{}:
		return NewAdsMap(val.(map[string]interface{}))

	case []interface{}:
		return NewAdsList(val.([]interface{}))
	}

	return AdsPrimitive{val}
}

func CreateWithMapper(val interface{}, mapper func(interface{}) (interface{}, error)) (AdsNode, error) {
	switch val.(type) {
	case map[string]interface{}:
		return NewAdsMapWithMapper(val.(map[string]interface{}), mapper)

	case []interface{}:
		return NewAdsListWithMapper(val.([]interface{}), mapper)
	}

	return NewAdsPrimitiveWithMapper(val, mapper)
}

func GetKey(ads AdsNode, key string, def interface{}) interface{} {
	if ads == nil {
		panic("ads is nil")
	}

	if ! ads.IsContainer() {
		return def
	}

	for _, v := range strings.Split(key, ".") {
		if len(v) == 0 {
			return def
		}

		var val, isPresent = ads.GetIndex(v)

		if isPresent {
			ads = val.(AdsNode)
		} else {
			return def
		}
	}

	var val = ads.GetValue()

	return val
}

func GetPath(ads AdsNode, path []string, def interface{}) interface{} {
	if ads == nil {
		panic("ads is nil")
	}

	if ! ads.IsContainer() {
		return def
	}

	for _, v := range path {
		if len(v) == 0 {
			return def
		}

		var val, isPresent = ads.GetIndex(v)

		if isPresent {
			ads = val.(AdsNode)
		} else {
			return def
		}
	}

	var val = ads.GetValue()

	return val
}

//func Set(ads AdsNode, key string, val interface{}) AdsNode {
//	if ! ads.IsContainer() {
//		return ads
//	}
//
//	for _, v := range strings.Split(key, ".") {
//		var val, present = ads.GetIndex(v)
//
//		if present {
//			ads = val.(AdsNode)
//		} else {
//			return def
//		}
//	}
//
//	return ads.GetValue()
//}
