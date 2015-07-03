package models

var TopicDao = &_Topic{}

type _Topic struct{}

type Topic struct {
	Name string `xorm:"varchar(16)"`
	Pid  int64  `xorm:"post_id"`
}

func (t *_Topic) Create(name string, pid int64) (err error) {
	topic := new(Topic)
	topic.Name = name
	topic.Pid = pid
	_, err = Engine.Insert(topic)
	return
}

func (t *_Topic) FindByPid(pid int64) (topics []*Topic, err error) {
	var ret []*Topic
	err = Engine.Find(ret, &Topic{Pid: pid})
	if err != nil {
		return
	}

	topics = ret
	return
}

func (t *_Topic) FindByName(name string) (topics []*Topic, err error) {
	var ret []*Topic
	err = Engine.Find(ret, &Topic{Name: name})
	if err != nil {
		return
	}

	topics = ret
	return
}

func (t *_Topic) FindContainsName(name string) (topics []*Topic, err error) {
	var ret []*Topic
	err = Engine.Where("like %?%", name).Find(&ret)
	if err != nil {
		return
	}

	topics = ret
	return
}
