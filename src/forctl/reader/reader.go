/*
 * Copyright (C) 2015 - 2017 Wiky Lyu
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

package reader

import (
	"forctl/cmd"
	"forctl/io"
	"gopkg.in/readline.v1"
	"skywalker/config"
	"strings"
)

/* 对readline的简单封装 */
type Reader struct {
	rl *readline.Instance
}

func New(ccfg *config.CoreConfig, rcfg []*config.ProxyConfig) (*Reader, error) {
	/* 自动补全数据 */
	var proxies, cmds []readline.PrefixCompleterInterface
	for _, r := range rcfg {
		proxies = append(proxies, readline.PcItem(r.Name))
	}
	for k, _ := range cmd.GetCommands() {
		if k != cmd.COMMAND_HELP {
			cmds = append(cmds, readline.PcItem(k))
		}
	}
	/* 设置自动补全 */
	completer := readline.NewPrefixCompleter(
		readline.PcItem(cmd.COMMAND_STATUS, proxies...),
		readline.PcItem(cmd.COMMAND_START, proxies...),
		readline.PcItem(cmd.COMMAND_STOP, proxies...),
		readline.PcItem(cmd.COMMAND_RESTART, proxies...),
		readline.PcItem(cmd.COMMAND_INFO, proxies...),
		readline.PcItem(cmd.COMMAND_HELP, cmds...),
		readline.PcItem(cmd.COMMAND_RELOAD),
		readline.PcItem(cmd.COMMAND_QUIT),
		readline.PcItem(cmd.COMMAND_CLEARCACHE),
	)
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "\x1B[36mforce>>\x1B[0m ",
		HistoryFile:  ccfg.HistoryFile,
		AutoComplete: completer,
	})
	if err != nil {
		return nil, err
	}
	return &Reader{rl: rl}, nil
}

type Line struct {
	Cmd  *cmd.Command
	Args []string
}

func NewLine(buf string) *Line {
	var seps []string
	var c *cmd.Command

	for _, s := range strings.Split(buf, " ") {
		if len(s) > 0 {
			seps = append(seps, s)
		}
	}
	if len(seps) == 0 {
		return nil
	}
	if c = cmd.GetCommand(seps[0]); c == nil {
		io.PrintError("Unknown syntax: %s\n", seps[0])
		return nil
	}

	/* 参数个数不正确 */
	if c.Required > len(seps[1:]) ||
		(c.Optional >= 0 && c.Required+c.Optional < len(seps[1:])) {
		io.PrintError("Invalid argument for %s\n%s\n", seps[0], c.Help)
		return nil
	}

	return &Line{
		Cmd:  c,
		Args: seps[1:],
	}
}

/* 读取一行命令，去除首尾空格 */
func (r *Reader) Read() (*Line, error) {
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

func (r *Reader) Close() {
	r.rl.Close()
}
