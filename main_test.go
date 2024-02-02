// Package wlog @Author Bing
// @Date 2024/2/2 15:44:00
// @Desc
package wlog

import "testing"

func TestWlog(t *testing.T) {
	Info.Println("Info testing start!!!")
	Error.Println("Error testing start!!!")
	Debug.Println("Debug testing start!!!")
}
