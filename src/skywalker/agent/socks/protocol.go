/*
 * Copyright (C) 2015 - 2016 Wiky L
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
package socks

import (
	"bytes"
	"encoding/binary"
	"net"
	"skywalker/util"
)

const (
	ERROR_INVALID_NMETHODS             = 1
	ERROR_INVALID_INVALID_MESSAGE_SIZE = 2
	ERROR_UNSUPPORTED_CMD              = 3
	ERROR_UNSUPPORTED_VERSION          = 4
	ERROR_UNSUPPORTED_METHOD           = 5
	ERROR_INVALID_REPLY                = 6
	ERROR_INVALID_CONFIG               = 7
)

/* 方法常量 */
const (
	METHOD_NO_AUTH_REQUIRED  = byte('\x00')
	METHOD_GSSAPI            = byte('\x01')
	METHOD_USERNAME_PASSWORD = byte('\x02')
	METHOD_NO_ACCEPTABLE     = byte('\xFF')
)

/* 地址类型 */
const (
	ATYPE_IPV4       = 1
	ATYPE_DOMAINNAME = 3
	ATYPE_IPV6       = 4
)

/* 返回结果 */
const (
	REPLY_SUCCEED                    = 0
	REPLY_GENERAL_FAILURE            = 1
	REPLY_CONNECTION_NOW_ALLOWED     = 2
	REPLY_NETWORK_UNREACHABLE        = 3
	REPLY_HOST_UNREACHABLE           = 4
	REPLY_CONNECTION_REFUSED         = 5
	REPLY_TTL_EXPIRED                = 6
	REPLY_COMMAND_NOT_SUPPORTED      = 7
	REPLY_ADDRESS_TYPE_NOT_SUPPORTED = 8
)

const (
	CMD_CONNECT       = 1
	CMD_BIND          = 2
	CMD_UDP_ASSOCIATE = 3
)

const (
	STATE_INIT    = 0 /* 初始化状态，等待客户端发送握手请求 */
	STATE_CONNECT = 1 /* 等待客户端发送链接请求 */
	STATE_TUNNEL  = 2 /* 转发数据 */
	STATE_ERROR   = 3 /* 已经出错 */
)

/*
 * +----+----------+----------+
 * |VER | NMETHODS | METHODS  |
 * +----+----------+----------+
 * | 1  |    1     | 1 to 255 |
 * +----+----------+----------+
 */
type socks5VersionRequest struct {
	version  uint8
	nmethods uint8
	methods  []byte
}

func (req *socks5VersionRequest) build() []byte {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, req.version)
	binary.Write(&buf, binary.BigEndian, req.nmethods)
	binary.Write(&buf, binary.BigEndian, req.methods)
	return buf.Bytes()
}

func (req *socks5VersionRequest) parse(data []byte) error {
	if len(data) < 3 {
		return util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "version request message is too short")
	}
	version := uint8(data[0])
	nmethods := uint8(data[1])
	if version != 5 {
		return util.NewError(ERROR_UNSUPPORTED_VERSION, "unsupported protocol version %d", version)
	} else if nmethods < 1 {
		return util.NewError(ERROR_INVALID_NMETHODS, "nmethods cannot be zero")
	} else if len(data) != 2+int(nmethods) {
		return util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "unexpected version request message size")
	}
	req.version = version
	req.nmethods = nmethods
	req.methods = []uint8(data[2:])
	return nil
}

/*
 * +----+--------+
 * |VER | METHOD |
 * +----+--------+
 * | 1  |   1    |
 * +----+--------+
 */
type socks5VersionResponse struct {
	version uint8
	method  uint8
}

func (rep *socks5VersionResponse) build() []byte {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, rep.version)
	binary.Write(&buf, binary.BigEndian, rep.method)
	return buf.Bytes()
}

/* 解析SOCKS版本请求 */
func (rep *socks5VersionResponse) parse(data []byte) error {
	if len(data) != 2 {
		return util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "unexpected version reply message size")
	}
	rep.version = data[0]
	rep.method = data[1]
	return nil
}

/*
 * +----+-----+-------+------+----------+----------+
 * |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
 * +----+-----+-------+------+----------+----------+
 * | 1  |  1  | X'00' |  1   | Variable |    2     |
 * +----+-----+-------+------+----------+----------+
 */
type socks5Request struct {
	version uint8
	cmd     uint8
	atype   uint8
	addr    string
	port    uint16
}

/* 生成连接请求 */
func (req *socks5Request) build() []byte {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, req.version)
	binary.Write(&buf, binary.BigEndian, req.cmd)
	binary.Write(&buf, binary.BigEndian, uint8(0))
	binary.Write(&buf, binary.BigEndian, req.atype)
	if req.atype == ATYPE_DOMAINNAME {
		binary.Write(&buf, binary.BigEndian, uint8(len(req.addr)))
		binary.Write(&buf, binary.BigEndian, []byte(req.addr))
	} else {
		ip := net.ParseIP(req.addr)
		if ip == nil {
			return nil
		}
		binary.Write(&buf, binary.BigEndian, []byte(ip))
	}
	binary.Write(&buf, binary.BigEndian, req.port)
	return buf.Bytes()
}

/*
 * 解析SOCKS5请求
 * 返回解析剩余数据，错误
 */
func (req *socks5Request) parse(data []byte) ([]byte, error) {
	if len(data) < 6 {
		return nil, util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "address request message size is too short")
	}
	version := uint8(data[0])
	cmd := uint8(data[1])
	atype := uint8(data[3])
	var addr string
	var port uint16
	var left []byte = nil
	if atype == ATYPE_IPV4 {
		if len(data) < 10 {
			return nil, util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "address request message size is too short")
		} else if len(data) > 10 {
			left = data[10:]
		}
		addr = net.IP(data[4:8]).String()
		data = data[8:]
	} else if atype == ATYPE_IPV6 {
		if len(data) < 22 {
			return nil, util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "address request message size is too short")
		} else if len(data) > 22 {
			left = data[22:]
		}
		addr = net.IP(data[4:20]).String()
		data = data[20:]
	} else {
		length := uint8(data[4])
		if len(data) < 7+int(length) {
			return nil, util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "address request message size is too short")
		} else if len(data) > 7+int(length) {
			left = data[7+int(length):]
		}
		addr = string(data[5:(5 + length)])
		data = data[(5 + length):]
	}
	buf := bytes.NewReader(data)
	binary.Read(buf, binary.BigEndian, &port)

	req.version = version
	req.cmd = cmd
	req.atype = atype
	req.addr = addr
	req.port = port
	return left, nil
}

/*
 * +----+-----+-------+------+----------+----------+
 * |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
 * +----+-----+-------+------+----------+----------+
 * | 1  |  1  | X'00' |  1   | Variable |    2     |
 * +----+-----+-------+------+----------+----------+
 */
type socks5Response struct {
	version uint8
	reply   uint8
	atype   uint8
	addr    string
	port    uint16
}

func (rep *socks5Response) build() []byte {
	buf := bytes.Buffer{}
	binary.Write(&buf, binary.BigEndian, rep.version)
	binary.Write(&buf, binary.BigEndian, rep.reply)
	binary.Write(&buf, binary.BigEndian, uint8(0))
	binary.Write(&buf, binary.BigEndian, rep.atype)
	if rep.atype == ATYPE_IPV4 || rep.atype == ATYPE_IPV6 {
		binary.Write(&buf, binary.BigEndian, []byte(net.ParseIP(rep.addr)))
	} else {
		binary.Write(&buf, binary.BigEndian, uint8(len(rep.addr)))
		binary.Write(&buf, binary.BigEndian, []byte(rep.addr))
	}
	binary.Write(&buf, binary.BigEndian, rep.port)
	return buf.Bytes()
}

func (rep *socks5Response) parse(data []byte) error {
	if len(data) < 10 {
		return util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "address reply message size is too short")
	}
	version := data[0]
	reply := data[1]
	atype := data[3]
	var addr string
	var port uint16
	var left []byte
	if atype == ATYPE_IPV4 {
		if len(data) != 10 {
			return util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "unexpected address request message size")
		}
		addr = net.IP(data[4:8]).String()
		left = data[8:]
	} else if atype == ATYPE_IPV6 {
		if len(data) != 22 {
			return util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "unexpected address request message size")
		}
		addr = net.IP(data[4:20]).String()
		left = data[20:]
	} else {
		length := data[4]
		if len(data) != int(length+7) {
			return util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "unexpected address request message size")
		}
		addr = string(data[5:(5 + length)])
		left = data[(5 + length):]
	}
	buf := bytes.NewReader(left)
	binary.Read(buf, binary.BigEndian, &port)
	rep.version = version
	rep.reply = reply
	rep.atype = atype
	rep.addr = addr
	rep.port = port
	return nil
}

func parseAddressReply(data []byte) (uint8, uint8, uint8, string, uint16, error) {
	if len(data) < 10 {
		return 0, 0, 0, "", 0, util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "address reply message size is too short")
	}
	ver := data[0]
	rep := data[1]
	atype := data[3]
	var address string
	var port uint16
	var left []byte
	if atype == ATYPE_IPV4 {
		if len(data) != 10 {
			return 0, 0, 0, "", 0, util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "unexpected address request message size")
		}
		address = net.IP(data[4:8]).String()
		left = data[8:]
	} else if atype == ATYPE_IPV6 {
		if len(data) != 22 {
			return 0, 0, 0, "", 0, util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "unexpected address request message size")
		}
		address = net.IP(data[4:20]).String()
		left = data[20:]
	} else {
		length := data[4]
		if len(data) != int(length+7) {
			return 0, 0, 0, "", 0, util.NewError(ERROR_INVALID_INVALID_MESSAGE_SIZE, "unexpected address request message size")
		}
		address = string(data[5:(5 + length)])
		left = data[(5 + length):]
	}
	buf := bytes.NewReader(left)
	binary.Read(buf, binary.BigEndian, &port)
	return ver, rep, atype, address, port, nil
}
