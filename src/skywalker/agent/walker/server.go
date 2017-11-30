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
package walker

import (
	"fmt"
	. "skywalker/agent/base"
	"skywalker/cipher"
	"skywalker/pkg"
	"skywalker/util"
)

type (
	WalkerServerAgent struct {
		BaseAgent

		addr      string
		port      uint16
		encrypter cipher.Encrypter
		decrypter cipher.Decrypter

		cfg *WalkerServerConfig
	}

	WalkerServerConfig struct {
		addr   string
		port   uint16
		method string
	}
)

var (
	gWAConfigs = map[string]*WalkerServerConfig{}
)

func (a *WalkerServerAgent) Name() string {
	return "walker"
}

func (a *WalkerServerAgent) OnInit(name string, cfg map[string]interface{}) error {
	port := uint16(util.GetMapIntDefault(cfg, "port", 0))
	if port == 0 {
		return fmt.Errorf("invalid port %d", port)
	}
	gWAConfigs[name] = &WalkerServerConfig{
		addr:   util.GetMapString(cfg, "addr"),
		port:   port,
		method: util.GetMapString(cfg, "method"),
	}
	return nil
}

func (a *WalkerServerAgent) OnStart() error {
	a.cfg = gWAConfigs[a.BaseAgent.Name]
	return nil
}

func (a *WalkerServerAgent) GetRemoteAddress(addr string, port int) (string, int) {
	a.addr = addr
	a.port = uint16(port)
	return a.cfg.addr, int(a.cfg.port)
}

func (a *WalkerServerAgent) OnConnectResult(result int, addr string, port int) (interface{}, interface{}, error) {
	if result == pkg.CONNECT_RESULT_OK {
		data, encrypter := packConnectRequest(addr, uint16(port), a.cfg.method)
		a.encrypter = encrypter
		return nil, data, nil
	} else {
		return nil, nil, nil
	}
}

func (a *WalkerServerAgent) ReadFromServer(data []byte) (interface{}, interface{}, error) {
	return nil, nil, nil
}

func (a *WalkerServerAgent) ReadFromCA(data []byte) (interface{}, interface{}, error) {
	return nil, nil, nil
}

func (a *WalkerServerAgent) UDPSupported() bool {
	return false
}

func (a *WalkerServerAgent) RecvFromServer(data []byte) (interface{}, interface{}, error) {
	return nil, nil, nil
}

func (a *WalkerServerAgent) RecvFromCA([]byte, string, int) (interface{}, interface{}, string, int, error) {
	return nil, nil, "", 0, nil
}

func (a *WalkerServerAgent) OnClose(closed_by_client bool) {
}

func (a *WalkerServerAgent) GetInfo() []map[string]string {
	return nil
}
