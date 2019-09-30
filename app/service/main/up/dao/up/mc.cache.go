// Code generated by $GOPATH/src/go-common/app/tool/cache/mc. DO NOT EDIT.

/*
  Package up is a generated mc cache package.
  It is generated from:
  type _mc interface {
		//mc: -key=upCacheKey -expire=d.upExpire -encode=json
		AddCacheUp(c context.Context, mid int64, up *model.Up) (err error)
		//mc: -key=upCacheKey
		CacheUp(c context.Context, mid int64) (up *model.Up, err error)
		//mc: -key=upCacheKey
		DelCacheUp(c context.Context, mid int64) (err error)
		//mc: -key=upSwitchKey -expire=d.upExpire -encode=json
		AddCacheUpSwitch(c context.Context, mid int64, up *model.UpSwitch) (err error)
		//mc: -key=upSwitchKey
		CacheUpSwitch(c context.Context, mid int64) (res *model.UpSwitch, err error)
		//mc: -key=upSwitchKey
		DelCacheUpSwitch(c context.Context, mid int64) (err error)
		//mc: -key=upInfoActiveKey -expire=d.upExpire -encode=json
		AddCacheUpInfoActive(c context.Context, mid int64, up *model.UpInfoActiveReply) (err error)
		//mc: -key=upInfoActiveKey
		CacheUpInfoActive(c context.Context, mid int64) (res *model.UpInfoActiveReply, err error)
		//mc: -key=upInfoActiveKey
		DelCacheUpInfoActive(c context.Context, mid int64) (err error)
		// mc: -key=upInfoActiveKey -expire=d.upExpire -encode=json
		AddCacheUpsInfoActive(c context.Context, res map[int64]*model.UpInfoActiveReply) (err error)
		// mc: -key=upInfoActiveKey
		CacheUpsInfoActive(c context.Context, mids []int64) (res map[int64]*model.UpInfoActiveReply, err error)
		// mc: -key=upInfoActiveKey
		DelCacheUpsInfoActive(c context.Context, mids []int64) (err error)
	}
*/

package up

import (
	"context"
	"fmt"

	"go-common/app/service/main/up/model"
	"go-common/library/cache/memcache"
	"go-common/library/log"
	"go-common/library/stat/prom"
)

var _ _mc

// AddCacheUp Set data to mc
func (d *Dao) AddCacheUp(c context.Context, id int64, val *model.Up) (err error) {
	if val == nil {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upCacheKey(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.upExpire, Flags: memcache.FlagJSON}
	if err = conn.Set(item); err != nil {
		prom.BusinessErrCount.Incr("mc:AddCacheUp")
		log.Errorv(c, log.KV("AddCacheUp", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// CacheUp get data from mc
func (d *Dao) CacheUp(c context.Context, id int64) (res *model.Up, err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upCacheKey(id)
	reply, err := conn.Get(key)
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:CacheUp")
		log.Errorv(c, log.KV("CacheUp", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	res = &model.Up{}
	err = conn.Scan(reply, res)
	if err != nil {
		prom.BusinessErrCount.Incr("mc:CacheUp")
		log.Errorv(c, log.KV("CacheUp", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DelCacheUp delete data from mc
func (d *Dao) DelCacheUp(c context.Context, id int64) (err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upCacheKey(id)
	if err = conn.Delete(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:DelCacheUp")
		log.Errorv(c, log.KV("DelCacheUp", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheUpSwitch Set data to mc
func (d *Dao) AddCacheUpSwitch(c context.Context, id int64, val *model.UpSwitch) (err error) {
	if val == nil {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upSwitchKey(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.upExpire, Flags: memcache.FlagJSON}
	if err = conn.Set(item); err != nil {
		prom.BusinessErrCount.Incr("mc:AddCacheUpSwitch")
		log.Errorv(c, log.KV("AddCacheUpSwitch", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// CacheUpSwitch get data from mc
func (d *Dao) CacheUpSwitch(c context.Context, id int64) (res *model.UpSwitch, err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upSwitchKey(id)
	reply, err := conn.Get(key)
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:CacheUpSwitch")
		log.Errorv(c, log.KV("CacheUpSwitch", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	res = &model.UpSwitch{}
	err = conn.Scan(reply, res)
	if err != nil {
		prom.BusinessErrCount.Incr("mc:CacheUpSwitch")
		log.Errorv(c, log.KV("CacheUpSwitch", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DelCacheUpSwitch delete data from mc
func (d *Dao) DelCacheUpSwitch(c context.Context, id int64) (err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upSwitchKey(id)
	if err = conn.Delete(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:DelCacheUpSwitch")
		log.Errorv(c, log.KV("DelCacheUpSwitch", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheUpInfoActive Set data to mc
func (d *Dao) AddCacheUpInfoActive(c context.Context, id int64, val *model.UpInfoActiveReply) (err error) {
	if val == nil {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upInfoActiveKey(id)
	item := &memcache.Item{Key: key, Object: val, Expiration: d.upExpire, Flags: memcache.FlagJSON}
	if err = conn.Set(item); err != nil {
		prom.BusinessErrCount.Incr("mc:AddCacheUpInfoActive")
		log.Errorv(c, log.KV("AddCacheUpInfoActive", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// CacheUpInfoActive get data from mc
func (d *Dao) CacheUpInfoActive(c context.Context, id int64) (res *model.UpInfoActiveReply, err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upInfoActiveKey(id)
	reply, err := conn.Get(key)
	if err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:CacheUpInfoActive")
		log.Errorv(c, log.KV("CacheUpInfoActive", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	res = &model.UpInfoActiveReply{}
	err = conn.Scan(reply, res)
	if err != nil {
		prom.BusinessErrCount.Incr("mc:CacheUpInfoActive")
		log.Errorv(c, log.KV("CacheUpInfoActive", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// DelCacheUpInfoActive delete data from mc
func (d *Dao) DelCacheUpInfoActive(c context.Context, id int64) (err error) {
	conn := d.mc.Get(c)
	defer conn.Close()
	key := upInfoActiveKey(id)
	if err = conn.Delete(key); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		prom.BusinessErrCount.Incr("mc:DelCacheUpInfoActive")
		log.Errorv(c, log.KV("DelCacheUpInfoActive", fmt.Sprintf("%+v", err)), log.KV("key", key))
		return
	}
	return
}

// AddCacheUpsInfoActive Set data to mc
func (d *Dao) AddCacheUpsInfoActive(c context.Context, values map[int64]*model.UpInfoActiveReply) (err error) {
	if len(values) == 0 {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	for id, val := range values {
		key := upInfoActiveKey(id)
		item := &memcache.Item{Key: key, Object: val, Expiration: d.upExpire, Flags: memcache.FlagJSON}
		if err = conn.Set(item); err != nil {
			prom.BusinessErrCount.Incr("mc:AddCacheUpsInfoActive")
			log.Errorv(c, log.KV("AddCacheUpsInfoActive", fmt.Sprintf("%+v", err)), log.KV("key", key))
			return
		}
	}
	return
}

// CacheUpsInfoActive get data from mc
func (d *Dao) CacheUpsInfoActive(c context.Context, ids []int64) (res map[int64]*model.UpInfoActiveReply, err error) {
	l := len(ids)
	if l == 0 {
		return
	}
	keysMap := make(map[string]int64, l)
	keys := make([]string, 0, l)
	for _, id := range ids {
		key := upInfoActiveKey(id)
		keysMap[key] = id
		keys = append(keys, key)
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	replies, err := conn.GetMulti(keys)
	if err != nil {
		prom.BusinessErrCount.Incr("mc:CacheUpsInfoActive")
		log.Errorv(c, log.KV("CacheUpsInfoActive", fmt.Sprintf("%+v", err)), log.KV("keys", keys))
		return
	}
	for key, reply := range replies {
		var v *model.UpInfoActiveReply
		v = &model.UpInfoActiveReply{}
		err = conn.Scan(reply, v)
		if err != nil {
			prom.BusinessErrCount.Incr("mc:CacheUpsInfoActive")
			log.Errorv(c, log.KV("CacheUpsInfoActive", fmt.Sprintf("%+v", err)), log.KV("key", key))
			return
		}
		if res == nil {
			res = make(map[int64]*model.UpInfoActiveReply, len(keys))
		}
		res[keysMap[key]] = v
	}
	return
}

// DelCacheUpsInfoActive delete data from mc
func (d *Dao) DelCacheUpsInfoActive(c context.Context, ids []int64) (err error) {
	if len(ids) == 0 {
		return
	}
	conn := d.mc.Get(c)
	defer conn.Close()
	for _, id := range ids {
		key := upInfoActiveKey(id)
		if err = conn.Delete(key); err != nil {
			if err == memcache.ErrNotFound {
				err = nil
				continue
			}
			prom.BusinessErrCount.Incr("mc:DelCacheUpsInfoActive")
			log.Errorv(c, log.KV("DelCacheUpsInfoActive", fmt.Sprintf("%+v", err)), log.KV("key", key))
			return
		}
	}
	return
}
