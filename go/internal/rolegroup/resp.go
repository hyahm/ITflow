package rolegroup

// type RespRoleGroup struct {
// 	RoleList []*ReqRoleGroup `json:"rolelist" type:"array" need:"是" information:"角色组成员"`
// 	Code     int             `json:"code" type:"array" need:"是"  information:"错误码"`
// 	Msg      string          `json:"msg,omitempty" type:"array" need:"否" information:"错误信息"`
// }

// func (rrg *RespRoleGroup) Marshal() []byte {
// 	send, _ := json.Marshal(rrg)
// 	return send
// }

// func (rrg *RespRoleGroup) ErrorE(err error) []byte {
// 	rrg.Code = 1
// 	rrg.Msg = err.Error()
// 	return rrg.Marshal()
// }

// func (rrg *RespRoleGroup) Error(msg string) []byte {
// 	rrg.Code = 1
// 	rrg.Msg = msg
// 	return rrg.Marshal()
// }
