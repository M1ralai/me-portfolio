package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	zap *zap.Logger
	db  *sqlx.DB
}

func NewLogger(db *sqlx.DB) *ZapLogger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05")
	config.DisableStacktrace = true
	config.DisableCaller = true

	z, _ := config.Build()

	return &ZapLogger{
		zap: z,
		db:  db,
	}
}

func (l *ZapLogger) Info(msg string, details map[string]interface{}) {
	fields := convertToZapFields(details)
	l.zap.Info(msg, fields...)

	go l.saveToDB("INFO", msg, details)
}

func (l *ZapLogger) Error(msg string, err error, details map[string]interface{}) {
	if details == nil {
		details = make(map[string]interface{})
	}
	if err != nil {
		details["error"] = err.Error()
	}

	fields := convertToZapFields(details)
	l.zap.Error(msg, fields...)

	go l.saveToDB("ERROR", msg, details)
}

// saveToDB saves log to PostgreSQL
func (l *ZapLogger) saveToDB(level, msg string, details map[string]interface{}) {
	if l.db == nil {
		return
	}

	detailsJSON, _ := json.Marshal(details)

	// CREATE, UPDATE, DELETE içeren loglar kalıcı olarak işaretlenir
	isPermanent := false
	if action, ok := details["action"].(string); ok {
		upperAction := strings.ToUpper(action)
		if strings.Contains(upperAction, "CREATE") ||
			strings.Contains(upperAction, "UPDATE") ||
			strings.Contains(upperAction, "DELETE") {
			isPermanent = true
		}
	}

	query := `INSERT INTO system_logs (level, message, details, is_permanent, created_at) VALUES ($1, $2, $3, $4, $5)`

	_, dbErr := l.db.Exec(query, level, msg, detailsJSON, isPermanent, time.Now())
	if dbErr != nil {
		fmt.Printf("LOG DB HATASI: %v\n", dbErr)
	}
}

func convertToZapFields(details map[string]any) []zap.Field {
	var fields []zap.Field
	for k, v := range details {
		fields = append(fields, zap.Any(k, v))
	}
	return fields
}
