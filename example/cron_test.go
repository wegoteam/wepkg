package example

import (
	"fmt"
	"github.com/wegoteam/wepkg/job/cron"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	job1, _ := cron.AddJob("0/1 * * * * ?", JobTest1)
	job2, _ := cron.AddJob("0/1 * * * * ?", JobTest2)
	time.Sleep(time.Second * 3)
	cron.DelJob(job1)
	cron.DelJob(job2)
	job3, _ := cron.AddJob("0/1 * * * * ?", JobTest1)
	time.Sleep(time.Second * 3)
	cron.DelJob(job3)
	time.Sleep(time.Second * 3)
}

func JobTest1() {
	fmt.Println("JobTest1")
}

func JobTest2() {
	fmt.Println("JobTest2")
}
