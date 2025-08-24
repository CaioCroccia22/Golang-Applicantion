package logger

import (
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*O zap utiliza buffers internos para logar de forma rápida(não escreve tudo no disco/stdout*/
/*
Output:

{"level":"info","message":"logger construction succeeded","foo":"bar"}*/

var (
	//váriavel global para armazenar a configuração do log
	Logger *zap.Logger
	// definir uma váriavel para pegar os valores de variaveis de ambiente
	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	} else {
		logConfig := zap.Config{
			OutputPaths: []string{getOutPutLogs()},           //Onde o log será escrito
			Level:       zap.NewAtomicLevelAt(getLevelLog()), //nivel de log
			//NewAtomicLevelAt - Inicializa o leval log \\Pode ser chumbado o zap.infolevel
			Encoding: "json", //Formato de saída
			EncoderConfig: zapcore.EncoderConfig{ //define como cada campo aparece no log
				LevelKey:     "level",
				TimeKey:      "time",
				MessageKey:   "message",
				EncodeTime:   zapcore.ISO8601TimeEncoder, //ISO8601TimeEncoder serializes a time.Time to an ISO8601-formatted string with millisecond precision
				EncodeLevel:  zapcore.LowercaseLevelEncoder,
				EncodeCaller: zapcore.ShortCallerEncoder,
			},
		}

		Logger, _ = logConfig.Build() //cria o logger configurad ona variavel global
	}
}

// Função para poder jogar o log
func Info(message string, tags ...zap.Field) {
	Logger.Info(message)
	Logger.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	Logger.Info(message, tags...)
	Logger.Sync()
}

// função para não pegar o path, para não deixar chumbado na aplicação o stdout do Outputpaths
func getOutPutLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT))) //Pegar a variavel de ambiente
	if output == "" {
		return "stdout"
	}
	return output
}

// Função para pegar o Level
func getLevelLog() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
