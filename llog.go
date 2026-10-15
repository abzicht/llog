/**
 * Copyright (C) 2026 Abzicht (https://github.com/abzicht)
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

/*
 * Log lib for mid-sized codebases.
 * Simple to use, offers logging via global and local variables.
 */

package llog

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type LogLevel int

const (
	LFatal = LogLevel(-1)
	LPanic = LogLevel(0)
	LError = LogLevel(1)
	LWarn  = LogLevel(2)
	LInfo  = LogLevel(3)
	LDebug = LogLevel(4)
)

const (
	LSymFatal = "FATAL"
	LSymPanic = "PANIC"
	LSymError = "ERROR"
	LSymWarn  = "WARN"
	LSymInfo  = "INFO"
	LSymDebug = "DEBUG"
)

type Log struct {
	level     LogLevel
	prefix    string
	writer    io.Writer
	writerErr io.Writer
}

var logDefault Log = Log{level: LInfo, prefix: "LLOG", writer: os.Stderr, writerErr: os.Stderr}

var log *Log = &logDefault

func (l *Log) SetLevel(level LogLevel) {
	if level > LDebug {
		level = LDebug
	}
	if level < LFatal {
		level = LFatal
	}
	l.level = level
}

func (l *Log) SetLevelStr(level string) {
	l.SetLevel(str2level(level))
}

func (l *Log) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *Log) SetWriter(w io.Writer) {
	l.writer = w
}

func (l *Log) SetWriterErr(w io.Writer) {
	l.writerErr = w
}

func level2str(level LogLevel) string {
	switch level {
	case LFatal:
		return LSymFatal
	case LPanic:
		return LSymPanic
	case LError:
		return LSymError
	case LWarn:
		return LSymWarn
	case LInfo:
		return LSymInfo
	case LDebug:
		return LSymDebug
	}
	if level < LDebug {
		return LSymDebug
	}
	return LSymFatal
}

func str2level(level string) LogLevel {
	switch strings.ToUpper(strings.TrimSpace(level)) {
	case LSymFatal:
		return LFatal
	case LSymPanic:
		return LPanic
	case LSymError:
		return LError
	case LSymWarn:
		return LWarn
	case LSymInfo:
		return LInfo
	case LSymDebug:
		return LDebug
	}
	return LDebug
}

func (l *Log) tag(level LogLevel) string {
	return fmt.Sprintf("%s[%s]: ", l.prefix, level2str(level))
}

// Return log string
func (l *Log) logs(level LogLevel, v ...any) string {
	return l.tag(level) + fmt.Sprint(v...)
}

// Return formatted log string
func (l *Log) logsf(level LogLevel, format string, v ...any) string {
	return l.tag(level) + fmt.Sprintf(format, v...)
}

// Log string to file, if level requirement is met
func (l *Log) log(level LogLevel, v ...any) {
	if l.level >= level {
		writer := l.writer
		if level <= LError {
			writer = l.writerErr
		}
		fmt.Fprint(writer, l.logs(level, v...))
	}
}

// Log formatted string to file, if level requirement is met
func (l *Log) logf(level LogLevel, format string, v ...any) {
	if l.level >= level {
		writer := l.writer
		if level <= LError {
			writer = l.writerErr
		}
		fmt.Fprint(writer, l.logsf(level, format, v...))
	}
}

// Log to writer and then panic with the same message
func (l *Log) Panic(v ...any) {
	panic(l.logs(LPanic, v...))
}

func (l *Log) Fatal(v ...any) {
	l.log(LFatal, v...)
	os.Exit(1)
}

func (l *Log) Error(v ...any) {
	l.log(LError, v...)
}

func (l *Log) Warn(v ...any) {
	l.log(LWarn, v...)
}

func (l *Log) Info(v ...any) {
	l.log(LInfo, v...)
}

func (l *Log) Debug(v ...any) {
	l.log(LDebug, v...)
}

// Log to writer and then panic with the same message
func (l *Log) Panicf(format string, v ...any) {
	panic(l.logsf(LPanic, format, v...))
}

func (l *Log) Fatalf(format string, v ...any) {
	l.logf(LFatal, format, v...)
	os.Exit(1)
}

func (l *Log) Errorf(format string, v ...any) {
	l.logf(LError, format, v...)
}

func (l *Log) Warnf(format string, v ...any) {
	l.logf(LWarn, format, v...)
}

func (l *Log) Infof(format string, v ...any) {
	l.logf(LInfo, format, v...)
}

func (l *Log) Debugf(format string, v ...any) {
	l.logf(LDebug, format, v...)
}

// vvv Functions that use the global log variable vvv

func SetLogger(l *Log) {
	log = l
}

func SetLevel(level LogLevel) {
	log.SetLevel(level)
}

func SetLevelStr(level string) {
	log.SetLevelStr(level)
}

func GetLevel() LogLevel {
	return log.level
}

func SetPrefix(prefix string) {
	log.prefix = prefix
}

func SetWriter(w io.Writer) {
	log.SetWriter(w)
}

// not formatted

func Panic(v ...any) {
	log.Panic(v...)
}

func Fatal(v ...any) {
	log.Fatal(v...)
}

func Error(v ...any) {
	log.Error(v...)
}

func Warn(v ...any) {
	log.Warn(v...)
}

func Info(v ...any) {
	log.Info(v...)
}

func Debug(v ...any) {
	log.Debug(v...)
}

// using format string

func Panicf(format string, v ...any) {
	log.Panicf(format, v...)
}

func Fatalf(format string, v ...any) {
	log.Fatalf(format, v...)
}

func Errorf(format string, v ...any) {
	log.Errorf(format, v...)
}

func Warnf(format string, v ...any) {
	log.Warnf(format, v...)
}

func Infof(format string, v ...any) {
	log.Infof(format, v...)
}

func Debugf(format string, v ...any) {
	log.Debugf(format, v...)
}
