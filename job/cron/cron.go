package cron

import (
	cronUtil "github.com/robfig/cron/v3"
	"github.com/wegoteam/wepkg/log"
	"sync"
)

var (
	cronjob *cronUtil.Cron
	jobMap  = make(map[int]cronUtil.EntryID)
	mutex   sync.Mutex
	isInit  bool
)

type JobFunc func()

// 引用：https://github.com/robfig/cron
// @Description: # 文件格式說明
// # ┌──分钟（0 - 59）
// # │ ┌──小时（0 - 23）
// # │ │ ┌──日（1 - 31）
// # │ │ │ ┌─月（1 - 12）
// # │ │ │ │ ┌─星期（0 - 6，表示从周日到周六）
// # │ │ │ │ │
// # *  *  *  *  * 被执行的命令
func init() {
	cronjob = cronUtil.New(cronUtil.WithParser(cronUtil.NewParser(
		cronUtil.SecondOptional | cronUtil.Minute | cronUtil.Hour | cronUtil.Dom | cronUtil.Month | cronUtil.Dow | cronUtil.Descriptor,
	)))
}

// AddJob
// @Description: 添加定时任务
// @param: seq 定时任务表达式
// @param: jobFunc 定时任务执行函数
// @return int 定时任务ID
// @return error
func AddJob(seq string, jobFunc JobFunc) (int, error) {
	mutex.Lock()
	defer mutex.Unlock()
	if !isInit {
		cronjob.Start()
		isInit = true
	}
	entryID, err := cronjob.AddFunc(seq, jobFunc)
	if err != nil {
		log.Errorf("AddJob err %v \n", err)
		return 0, err
	}
	log.Debugf("AddJob success entryID %v \n", entryID)
	id := int(entryID)
	jobMap[id] = entryID
	return id, nil
}

// DelJob
// @Description: 删除定时任务
// @param: jobID 定时任务ID
func DelJob(jobID int) {
	mutex.Lock()
	defer mutex.Unlock()
	entryID, ok := jobMap[jobID]
	if !ok {
		return
	}
	cronjob.Remove(entryID)
	delete(jobMap, jobID)
	if len(jobMap) == 0 {
		isInit = false
		cronjob.Stop()
	}
	log.Debugf("DelJob success entryID %v \n", entryID)
}
