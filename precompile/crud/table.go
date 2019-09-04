package crud

type Table struct {
	TableName    string
	Key          string
	ValueFields  string
	Optional     string
}

func (t *Table) GetTableName() string {
	return t.TableName
}

func (t *Table) GetKey() string {
	return t.Key
}

func (t *Table) GetValueFields() string {
	return t.ValueFields
}

func (t *Table) GetOptional() string {
	return t.Optional
}

func (t *Table) SetTableName(name string) {
	t.TableName = name
}

func (t *Table) SetKey(key string) {
	t.Key = key
}

func (t *Table) SetValueFields(value string) {
	t.ValueFields = value
}

func (t *Table) SetOptional(optional string) {
	t.Optional = optional
}

func (t *Table) GetEntry() *Entry {
	return &Entry{fields:make(map[string]string)}
}

func (t *Table) GetCondition() *Condition {
	return &Condition{conditions:make(map[string]map[EnumOP]string)}
}