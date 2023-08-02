package example

import (
	"fmt"
	"github.com/wegoteam/wepkg/log"
	"testing"
)

func TestLogConfig(t *testing.T) {
	log.SetLoggerConfig(log.LoggerConfig{
		Level:  "trace",
		Format: "text",
		Output: "console",
	})
	config := log.GetLoggerConfig()
	fmt.Printf("config=%v\n", config)
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
}

func TestLog(t *testing.T) {
	log.Trace("Something very low level.")
	log.Tracef("Something very low level. %s", "test")
	log.Traceln("Something very low level.")

	log.Debug("Useful debugging information.")
	log.Debugf("Useful debugging information. %s", "test")
	log.Debugln("Useful debugging information.")

	log.Info("Something noteworthy happened!")
	log.Infof("Something noteworthy happened! %s", "test")
	log.Infoln("Something noteworthy happened!")

	log.Notice("Something unusual happened.")
	log.Noticef("Something unusual happened. %s", "test")
	log.Noticef("Something unusual happened.")

	log.Warn("You should probably take a look at this.")
	log.Warnf("You should probably take a look at this. %s", "test")
	log.Warnln("You should probably take a look at this.")

	log.Error("Something failed but I'm not quitting.")
	log.Errorf("Something failed but I'm not quitting. %s", "test")
	log.Errorln("Something failed but I'm not quitting.")

	// Calls os.Exit(1) after logging
	log.Fatal("Bye.")
	log.Fatalf("Bye. %s", "test")
	log.Fatalln("Bye.")

	// Calls panic() after logging
	log.Panic("I'm bailing.")
	log.Panicf("I'm bailing. %s", "test")
	log.Panicln("I'm bailing.")

}
