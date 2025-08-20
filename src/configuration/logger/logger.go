package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*O zap utiliza buffers internos para logar de forma rápida(não escreve tudo no disco/stdout*/
/*
Output:

{"level":"info","message":"logger construction succeeded","foo":"bar"}*/


var Logger *zap.Logger //váriavel global para armazenar o log

func init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	} else{
		logConfig := zap.Config{
			OutputPaths: []string{"stdout"},
			Level: 		zap.NewAtomicLevel(zap.InfoLevel) //NewAtomicLevelAt - Inicializa o leval log
			Encoding:   "json",
			EncoderConfig: zapcore.EncoderConfig{
				LevelKey: "level",
				TimeKey: "time",
				MessageKey: "message",
				EncodeTime: zapcore.ISO8601TimeEncoder, //ISO8601TimeEncoder serializes a time.Time to an ISO8601-formatted string with millisecond precision
				EncodeLevel: zapcore.LowercaseLevelEncoder,
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}

		Logger, _ = logConfig.Build()
	}

	
}

// função para não pegar o path, para não deixar chumbado na aplicação o stdout