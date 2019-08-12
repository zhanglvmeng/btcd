// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// +build !windows,!plan9

package limits

import (
	"fmt"
	"syscall"
)

const (
	fileLimitWant = 2048
	fileLimitMin  = 1024
)

// SetLimits raises some process limits to values which allow btcd and
// associated utilities to run.
<<<<<<< HEAD
// 系统资源限制的配置： 相关知识可以参考 https://www.jianshu.com/p/7e8726ebe338
func SetLimits() error {
	// 描述资源软硬限制的结构体
	var rLimit syscall.Rlimit
	// syscall.RLIMIT_NOFILE 是一个进程能打开的最大文件数，默认是8.
=======
func SetLimits() error {
	var rLimit syscall.Rlimit

>>>>>>> 增加了一些细节日志
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		return err
	}
<<<<<<< HEAD
	// rLimit.Cur 是系统soft limit ，是指内核所能支持的资源上限。
	if rLimit.Cur > fileLimitWant {
		return nil
	}
	// rLimit.Max 是指hard limit, 在资源中只是作为soft limit 的上限。
=======
	if rLimit.Cur > fileLimitWant {
		return nil
	}
>>>>>>> 增加了一些细节日志
	if rLimit.Max < fileLimitMin {
		err = fmt.Errorf("need at least %v file descriptors",
			fileLimitMin)
		return err
	}
	if rLimit.Max < fileLimitWant {
		rLimit.Cur = rLimit.Max
	} else {
		rLimit.Cur = fileLimitWant
	}
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		// try min value
		rLimit.Cur = fileLimitMin
		err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
		if err != nil {
			return err
		}
	}

	return nil
}
