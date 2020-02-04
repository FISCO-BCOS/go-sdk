package permission

type PermissionInfo struct {
	Address     string  `json:"address"`
	EnableNum   string  `json:"enable_num"`
	TableName   string  `json:"table_name"`
}

func (p *PermissionInfo) GetTableName() string {
	return p.TableName
}

func (p *PermissionInfo) SetTableName(name string) {
	p.TableName = name
}

func (p *PermissionInfo) GetAddress() string {
	return p.Address
}

func (p *PermissionInfo) SetAddress(addr string) {
	p.Address = addr
}

func (p *PermissionInfo) GetEnableNum() string {
	return p.EnableNum
}

func (p *PermissionInfo) SetEnableNum(enable string) {
	p.EnableNum = enable
}
