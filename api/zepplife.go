package api

import (
	"encoding/json"
	"net/http"
	"plugins/wechatMovement/zeepLife"
	"strconv"
)

/*
* @Author: zouyx
* @Date:   2024/5/26 15:54
* @Package:
 */

func CustomizedSteps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// 解析表单
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	// 获取表单参数
	account := r.FormValue("account")
	pwd := r.FormValue("pwd")
	steps := r.FormValue("steps")
	app := zeepLife.NewZeppLife(account, pwd)
	step, _ := strconv.Atoi(steps)
	err := app.SetSteps(step)
	if nil != err {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	resp := map[string]interface{}{"code": 200, "msg": "ok"}
	marshal, _ := json.Marshal(resp)
	w.Write(marshal)
}
