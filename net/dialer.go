// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package net

import (
	"errors"
	"net"
	"sync"

	"github.com/tmthrgd/shm-go"
)

type Dialer struct {
	rw   *shm.ReadWriteCloser
	name string

	mut sync.Mutex
}

func Dial(name string) (net.Conn, error) {
	rw, err := shm.OpenDuplex(name)
	if err != nil {
		return nil, err
	}

	return (&Dialer{
		rw:   rw,
		name: name,
	}).Dial("shm", name)
}

func NewDialer(rw *shm.ReadWriteCloser, name string) *Dialer {
	return &Dialer{
		rw:   rw,
		name: name,
	}
}

func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	if network != "shm" {
		return nil, errors.New("unrecognised network")
	}

	if address != d.name {
		return nil, errors.New("invalid address")
	}

	d.mut.Lock()
	return &Conn{d.rw, d.name, &d.mut}, nil
}
