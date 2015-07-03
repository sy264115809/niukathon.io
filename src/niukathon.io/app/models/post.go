package models

import (
	"time"
)

var PostDao = &_Post{}

type _Post struct{}

type Post struct {
	Id       int64     `xorm:"pk autoincr"`
	Uid      int64     `xorm:"user_id"`
	Content  string    `xorm:"notnull"`
	Imgs     string    `xorm`
	CreateAt time.Time `xorm:"timestamp create_at"`
}

func (p *_Post) Create(uid int64, content, imgs string) (err error) {
	post := new(Post)
	post.Uid = uid
	post.Content = content
	post.Imgs = imgs
	post.CreateAt = time.Now()
	_, err = Engine.Insert(post)
	return
}

func (m *Post) Update(content, imgs string) (err error) {
	_, err = Engine.Update(&Post{
		Content: content,
		Imgs:    imgs,
	}, m)
	if err != nil {
		return
	}

	m.Content = content
	return
}

func (m *Post) Delete(id int64) (err error) {
	_, err = Engine.Delete(m)
	if err != nil {
		return
	}
	m = nil
	return
}

func (p *_Post) FindAll() (posts []*Post, err error) {
	return p.findAll(nil)
}

func (p *_Post) FindAllByUid(uid int64) (posts []*Post, err error) {
	return p.findAll(&Post{Uid: uid})
}

func (p *_Post) FindByUid(uid int64, limit, from int) (posts []*Post, err error) {
	return p.find(&Post{Uid: uid}, limit, from)
}

func (p *_Post) FindAllById(id int64) (posts []*Post, err error) {
	return p.findAll(&Post{Id: id})
}

func (p *_Post) FindById(id int64, limit, from int) (posts []*Post, err error) {
	return p.find(&Post{Id: id}, limit, from)
}

func (p *_Post) find(condi interface{}, limit, from int) (posts []*Post, err error) {
	var ret []*Post
	if condi != nil {
		err = Engine.Desc("create_at").Limit(limit, (from-1)*limit).Find(&ret, condi)
	} else {
		err = Engine.Desc("create_at").Limit(limit, (from-1)*limit).Find(&ret)
	}
	if err != nil {
		return
	}

	posts = ret
	return
}

func (p *_Post) findAll(condi interface{}) (posts []*Post, err error) {
	var ret []*Post
	if condi != nil {
		err = Engine.Desc("create_at").Find(&ret, condi)
	} else {
		err = Engine.Desc("create_at").Find(&ret)
	}
	if err != nil {
		return
	}

	posts = ret
	return
}
