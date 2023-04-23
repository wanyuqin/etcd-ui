package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/wanyuqin/etcd-ui/backend/internal/domain/service/member"
	"github.com/wanyuqin/etcd-ui/backend/internal/handler/dto"
	"github.com/wanyuqin/etcd-ui/backend/x/xgin"
)

func CreateMember(c *gin.Context) {
	request := dto.AddMemberRequest{}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		xgin.Failed(c, err)
		return
	}
	err = member.NewMemberService().CreateMember(request.PeerURLs)
	xgin.Response(c, nil, err)
	return
}

func GetMember(c *gin.Context) {

}

func ListMember(c *gin.Context) {
	ml, err := member.NewMemberService().MemberList()
	if err != nil {
		xgin.Failed(c, err)
		return
	}
	response := dto.NewMemberListResponse(ml)
	xgin.Response(c, response, nil)
	return

}

func UpdateMember(c *gin.Context) {

}

func RemoveMember(c *gin.Context) {

}
