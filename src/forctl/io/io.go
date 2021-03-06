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

package io

import (
	"net"
	"skywalker/rpc"
	"skywalker/util"
	"strconv"
	"time"
)

/* 连接TCP服务 */
func TCPConnect(ip string, port int) (*rpc.Conn, error) {
	addr := net.JoinHostPort(ip, strconv.Itoa(port))
	if conn, err := util.TCPConnectTo(addr); err != nil {
		return nil, err
	} else {
		return rpc.NewConn(conn), nil
	}
}

/* 连接Unix套接字服务 */
func UnixConnect(filepath string) (*rpc.Conn, error) {
	if conn, err := net.DialTimeout("unix", filepath, 10*time.Second); err != nil {
		return nil, err
	} else {
		return rpc.NewConn(conn), nil
	}
}
