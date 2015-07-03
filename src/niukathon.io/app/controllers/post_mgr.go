package controllers

import (
	"github.com/revel/revel"
	"niukathon.io/app/models"
	"strings"
)

type PostMgr struct {
	App
}

type PostRet struct {
	Id       int64
	User     *models.User
	Content  string
	Imgs     []string
	CreateAt string
}

func (c *PostMgr) Index() revel.Result {
	return c.Render()
}

func (c *PostMgr) PostMsg(uid int64, content string, imgs []string) revel.Result {
	models.PostDao.Create(uid, content, strings.Join(imgs, ","))
	return c.RenderJson(map[string]string{"succcess": "ok"})
}

func (c *PostMgr) ShowAllMsg() revel.Result {
	posts, _ := models.PostDao.FindAll()
	retList := make([]*PostRet, 0, len(posts))
	for _, p := range posts {
		retList = append(retList, &PostRet{
			Id: p.Id,
			User: &models.User{
				Nickname: "我",
			},
			Content:  p.Content,
			Imgs:     strings.Split(p.Imgs, ","),
			CreateAt: p.CreateAt.Format("2006年01年02 15:04:05"),
		})
	}
	c.RenderArgs["posts"] = retList
	return c.RenderTemplate("PostMgr/index.html")
}
