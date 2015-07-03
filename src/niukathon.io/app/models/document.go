package models

var DocumentDao = &_Document{}

type _Document struct{}

type Document struct {
	Id   int64  `xorm:"pk autoincr"`
	Pid  int64  `xorm:"post_id"`
	Key  string `xorm:"notnull"`
	Name string `xorm:"notnull"`
	Ext  string `xorm:"notnull"`
}

func (d *_Document) Create(pid int64, key, name, ext string) (err error) {
	document := new(Document)
	document.Pid = pid
	document.Key = key
	document.Name = name
	document.Ext = ext
	_, err = Engine.Insert(document)
	return
}

func (m *Document) Update(pid int64) (err error) {
	_, err = Engine.Update(&Document{Pid: pid}, m)
	if err != nil {
		return
	}

	m.Pid = pid
	return
}

func (m *Document) Delete(id int64) (err error) {
	_, err = Engine.Delete(m)
	if err != nil {
		return
	}
	m = nil
	return
}

func (d *_Document) FindAllByPid(pid int64) (documents []*Document, err error) {
	return d.findAll(&Document{Pid: pid})
}

func (d *_Document) find(condi interface{}, limit, from int) (documents []*Document, err error) {
	var ret []*Document
	err = Engine.Limit(limit, (from-1)*limit).Find(ret, condi)
	if err != nil {
		return
	}

	documents = ret
	return
}

func (d *_Document) findAll(condi interface{}) (documents []*Document, err error) {
	var ret []*Document
	err = Engine.Find(ret, condi)
	if err != nil {
		return
	}

	documents = ret
	return
}
