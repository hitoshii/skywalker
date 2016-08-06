/*
 * Copyright (C) 2016 Wiky L
 *
 * This program is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.";
 */

package core

import (
	"fmt"
	"gopkg.in/readline.v1"
	"skywalker/config"
	"strings"
)

/* 对readline的简单封装 */
type Readline struct {
	rl *readline.Instance
}

func NewReadline(rcfg []*config.RelayConfig) (*Readline, error) {
	/* 自动补全数据 */
	var relays, cmds []readline.PrefixCompleterInterface
	for _, r := range rcfg {
		relays = append(relays, readline.PcItem(r.Name))
	}
	for k, _ := range gCommandMap {
		if k != COMMAND_HELP {
			cmds = append(cmds, readline.PcItem(k))
		}
	}
	completer := readline.NewPrefixCompleter(
		readline.PcItem(COMMAND_STATUS, relays...),
		readline.PcItem(COMMAND_HELP, cmds...),
	)
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "\x1B[36mforce>>\x1B[0m ",
		HistoryFile:  "/tmp/force_history",
		AutoComplete: completer,
	})
	if err != nil {
		return nil, err
	}
	return &Readline{rl: rl}, nil
}

type Line struct {
	CMD  string
	Args []string
}

/* 获取第i个参数 */
func (l *Line) Argument(i int) string {
	if len(l.Args) <= i {
		return ""
	}
	return l.Args[i]
}

func (l *Line) Arguments() []string {
	return l.Args
}

func InputError(format string, v ...interface{}) {
	format = "*** " + format
	fmt.Printf(format, v...)
}

func NewLine(buf string) *Line {
	var seps []string
	var cmd *CommandDefine

	for _, s := range strings.Split(buf, " ") {
		if len(s) == 0 {
			continue
		}
		seps = append(seps, s)
	}
	if len(seps) == 0 {
		return nil
	}
	if cmd = gCommandMap[seps[0]]; cmd == nil {
		InputError("Unknown syntax: %s\n", seps[0])
		return nil
	}

	if cmd.RequiredCount > len(seps[1:]) || cmd.RequiredCount+cmd.OptionalCount < len(seps[1:]) {
		InputError("Invalid argument for %s\n%s\n", seps[0], cmd.Help)
		return nil
	}

	return &Line{
		CMD:  seps[0],
		Args: seps[1:],
	}
}

/* 读取已行，去除首尾空格 */
func (r *Readline) Readline() (*Line, error) {
	var buf string
	var err error
	var line *Line
	for {
		buf, err = r.rl.Readline()
		if err != nil {
			return nil, err
		}
		if buf = strings.Trim(buf, " "); len(buf) == 0 {
			continue
		} else if line = NewLine(buf); line != nil {
			break
		}
	}
	return line, nil
}

func (r *Readline) Close() {
	r.rl.Close()
}