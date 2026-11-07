//go:build windows

package api

import (
    "syscall"
    "unsafe"
)

func freeSpace(path string) (int64, error) {
    p, err := syscall.UTF16PtrFromString(path)
    if err != nil {
        return 0, err
    }

    var available uint64
    proc := syscall.NewLazyDLL("kernel32.dll").NewProc("GetDiskFreeSpaceExW")

    result, _, callErr := proc.Call(
        uintptr(unsafe.Pointer(p)),
        uintptr(unsafe.Pointer(&available)),
        0,
        0,
    )

    if result == 0 {
        return 0, callErr
    }

    return int64(available), nil
}
