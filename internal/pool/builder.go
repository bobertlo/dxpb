package pool

import (
	"context"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"

	"github.com/dxpb/dxpb/internal/logger"
	"github.com/dxpb/dxpb/internal/spec"
)

func logFile(base string, tag string) (io.WriteCloser, error) {
	fullDirName := base + "/" + strconv.FormatInt(time.Now().Unix(), 10)
	err := os.MkdirAll(fullDirName, 0755)
	if err != nil {
		return nil, err
	}
	fullFileName := fullDirName + "/" + tag + ".txt"
	rV, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return rV, nil
}

func createOpts(p spec.Builder_build_Params, tag string) (spec.Builder_Opts, error) {
	what, err := p.NewOptions()
	if err != nil {
		return what, err
	}

	outFile, err := logFile(viper.GetString("logroot"), tag)
	if err != nil {
		return what, err
	}
	logee := spec.Logger_ServerToClient(&logger.Logger{To: outFile,
		Closer: time.AfterFunc(time.Minute, func() { _ = outFile.Close() })})
	what.SetIgnorePkgSpec(true)
	what.SetLog(logee)
	return what, nil
}

func translateJob(p spec.Builder_build_Params, in BuildJob) (spec.Builder_What, error) {
	what, err := p.NewWhat()
	if err != nil {
		return what, err
	}
	what.SetName("")
	what.SetVer("")
	what.SetArch(spec.Arch_noarch)
	return what, nil
}

func runBuilds(ctx context.Context, alias string, drone spec.Builder, busyGauge prometheus.Gauge, trigBuild <-chan BuildJob, update chan<- buildUpdate) error {
	end_builds := false
	for !end_builds {
		select {
		case what := <-trigBuild:
			log.Println("Starting build")
			busyGauge.Inc()
			_, err := drone.Build(ctx, func(p spec.Builder_build_Params) error {
				setThis, err := translateJob(p, what)
				if err != nil {
					return err
				}
				p.SetWhat(setThis)

				opts, err := createOpts(p, alias)
				if err != nil {
					return err
				}
				p.SetOptions(opts)
				return nil
			}).Struct()
			busyGauge.Dec()
			if err != nil {
				log.Println("Build erred: ", err)
				end_builds = true
			}
			log.Println("Build done")
		}
	}
	return nil
}
