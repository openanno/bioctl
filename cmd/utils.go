package cmd

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	cvrt "github.com/openbiox/ligo/convert"
	cio "github.com/openbiox/ligo/io"
	clog "github.com/openbiox/ligo/log"
	"github.com/openbiox/ligo/stringo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var log = clog.Logger
var logBash = clog.LoggerBash
var logEnv = log.WithFields(logrus.Fields{
	"prefix": "Env"})
var logPrefix string
var wd string

func setGlobalFlag(cmd *cobra.Command) {
	wd, _ = os.Getwd()
	cmd.PersistentFlags().IntVarP(&(rootClis.Verbose), "verbose", "", 1, "verbose level(0:no output, 1: basic level, 2: with env info")
	cmd.PersistentFlags().StringVarP(&(rootClis.TaskID), "task-id", "k", stringo.RandString(15), "task ID (default is random).")
	cmd.PersistentFlags().StringVarP(&(rootClis.LogDir), "log-dir", "", path.Join(wd, "_log"), "log dir.")
	cmd.PersistentFlags().BoolVarP(&(rootClis.SaveLog), "save-log", "s", false, "Save log to file.")
	cmd.PersistentFlags().BoolVarP(&(rootClis.Clean), "clean", "", false, "remove log dir.")
	cmd.PersistentFlags().StringVarP(&rootClis.Out, "out", "o", "", "out specifies destination of the returned data (default to stdout or current woring directory).")

}
func initCmd(cmd *cobra.Command, args []string) {
	setLog()
	if rootClis.Verbose == 2 {
		logEnv.Infof("Prog: %s", cmd.CommandPath())
		logEnv.Infof("TaskID: %s", rootClis.TaskID)
		if rootClis.SaveLog && logPrefix != "" {
			logEnv.Infof("Log: %s.log", logPrefix)
		}
		if len(args) > 0 {
			logEnv.Infof("Args: %s", strings.Join(args, " "))
		}
		logEnv.Infof("Global: %v", cvrt.Struct2Map(rootClis))
	}
	if rootClis.Clean {
		cleanLog()
	}
}

func setLog() {
	var logCon io.Writer
	var logDir = rootClis.LogDir

	if rootClis.SaveLog {
		if logDir == "" {
			logDir = filepath.Join(os.TempDir(), "_log")
		}
		logPrefix = fmt.Sprintf("%s/%s", logDir, rootClis.TaskID)
		cio.CreateDir(logDir)
		logCon, _ = cio.Open(logPrefix + ".log")
	}
	clog.SetLogStream(log, rootClis.Verbose == 0, rootClis.SaveLog, &logCon)
}

func cleanLog() {
	rootClis.HelpFlags = false
	if err := os.RemoveAll(rootClis.LogDir); err != nil {
		log.Warn(err)
	}
}

func convertIndex(str string) (index []int) {
	indexTmp := strings.Split(str, ";")
	for i := range indexTmp {
		if strings.Contains(indexTmp[i], ":") {
			startEnd := strings.Split(indexTmp[i], ":")
			start, _ := strconv.ParseInt(startEnd[0], 10, 64)
			end, _ := strconv.ParseInt(startEnd[1], 10, 64)
			for j := start; j < end+1; j++ {
				index = append(index, int(j))
			}
		} else {
			val, _ := strconv.ParseInt(indexTmp[i], 10, 64)
			index = append(index, int(val))
		}
	}
	return index
}
