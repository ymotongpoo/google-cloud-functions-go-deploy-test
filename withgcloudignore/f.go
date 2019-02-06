// Copyright 2019 Yoshi Yamaguchi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hello

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ymotongpoo/google-cloud-functions-go-deploy-test/withgcloudignore/dummy"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLogger() (*zap.SugaredLogger, error) {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "severity",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			TimeKey:        "timestamp",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			CallerKey:      "caller",
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	var err error
	var l *zap.Logger
	l, err = cfg.Build()
	if err != nil {
		return nil, err
	}
	return l.Sugar(), nil
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	logger, err := initLogger()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	logger.Infof("This is test log: random=%v, %v", dummy.MyRandomInt31(), r.Header.Get("User-Agent"))
	fmt.Fprintf(w, "Hello, World: %v", r.Host)
}
