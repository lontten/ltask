package util

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/ltask/service"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

// ShellExec
// 执行shell命令
func ShellExec(script string, timeout int64) bool {
	err := core(script, timeout)
	if err != nil {
		// 检查超时
		if errors.Is(err, context.DeadlineExceeded) {
			service.SaveLog("命令执行超时")
		} else {
			service.SaveLog(err.Error())
		}
		return false
	} else {
		service.SaveLog("命令执行成功")
		return true
	}
}

func core(script string, timeout int64) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("执行panic: %v", e)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// Windows 使用 PowerShell
		cmd = exec.CommandContext(ctx, "powershell", "-Command", script)
	} else {
		// Linux/macOS 使用 bash/sh
		cmd = exec.CommandContext(ctx, "sh", "-c", script)
	}

	// 创建管道捕获标准输出和错误
	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	// 启动命令
	err = cmd.Start()
	if err != nil {
		return err
	}

	// 共享缓冲区与锁
	var (
		outputBuffer []string
		bufferLock   sync.Mutex
	)

	// 实时读取标准输出
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			line := scanner.Text()
			bufferLock.Lock()
			outputBuffer = append(outputBuffer, "[输出] "+line)
			bufferLock.Unlock()
		}
	}()

	// 实时读取错误输出
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			line := scanner.Text()
			bufferLock.Lock()
			outputBuffer = append(outputBuffer, "[错误] "+line)
			bufferLock.Unlock()
		}
	}()

	// 定时器：每2秒打印一次最新输出
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				bufferLock.Lock()
				if len(outputBuffer) > 0 {
					service.SaveLog(outputBuffer...)
					outputBuffer = nil // 清空缓冲区
				}
				bufferLock.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}()

	// 等待命令结束
	err = cmd.Wait()

	// 处理结果
	bufferLock.Lock()
	defer bufferLock.Unlock()

	// 检查超时
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		// 打印已捕获的输出（如果有）
		if len(outputBuffer) > 0 {
			service.SaveLog(outputBuffer...)
		}
		return ctx.Err()
	}
	// 处理其他错误
	if err != nil {
		return fmt.Errorf("命令执行出错: %v", err)
	}

	// 打印最终剩余输出
	if len(outputBuffer) > 0 {
		service.SaveLog(outputBuffer...)
	}
	return nil
}
