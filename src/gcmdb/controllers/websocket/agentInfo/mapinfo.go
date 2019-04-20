/*************************************************************
     FileName: src->agentInfo->mapinfo.go
         Date: 2019-01-04 20:53
       Author: 苦咖啡
        Email: voilet@qq.com
         blog: http://blog.kukafei520.net
      Version: 0.0.1
      History:
**************************************************************/

package agentInfo

import (
	"errors"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	Data map[string]string
}

var AgentMap = newSafeMap()

func newSafeMap() *SafeMap {
	sm := new(SafeMap)
	sm.Data = make(map[string]string)
	return sm

}

func (sm *SafeMap) ReadMap(key string) (string, error) {
	sm.RLock()
	value, ok := sm.Data[key]
	sm.RUnlock()
	if !ok {
		return "", errors.New("error")
	}
	return value, nil
}

func (sm *SafeMap) WriteMap(key string, value string) {
	sm.Lock()
	sm.Data[key] = value
	sm.Unlock()
}

func (sm *SafeMap) DeleteMap(key string, value string) {
	sm.Lock()
	delete(sm.Data, value)
	sm.Unlock()
}

func (sm *SafeMap) LenMap() (int) {
	sm.Lock()
	st := len(sm.Data)
	sm.Unlock()
	return st

}

func (sm *SafeMap) MapList() ([]string) {
	sm.Lock()
	st := sm.Data
	sm.Unlock()
	var s []string
	for _, v := range st {
		s = append(s, v)
	}
	return s

}
