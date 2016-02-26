/*
 * Copyright (C) 2015 Wiky L
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

package shell

import (
    "flag"
)

type Options struct {
    BindAddr string
    BindPort string
    RemoteAddr string
    RemotePort string
}

var (
    Opts *Options
)

func init() {
    Opts = parseOptions()
}

func parseOptions() *Options {
    bindAddr := flag.String("bindAddr", "127.0.0.1", "the IP address to listen")
    bindPort := flag.String("bindPort", "42312", "the port to listen")
    remoteAddr := flag.String("remoteAddr", "www.baidu.com", "")
    remotePort := flag.String("remotePort", "80", "")
    flag.Parse()

    return &Options{
        BindAddr: *bindAddr,
        BindPort: *bindPort,
        RemoteAddr: *remoteAddr,
        RemotePort: *remotePort,
    }
}
