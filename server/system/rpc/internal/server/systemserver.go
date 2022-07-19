// Code generated by goctl. DO NOT EDIT!
// Source: system.proto

package server

import (
	"context"

	"system/rpc/internal/logic"
	"system/rpc/internal/svc"
	"system/rpc/types/system"
)

type SystemServer struct {
	svcCtx *svc.ServiceContext
	system.UnimplementedSystemServer
}

func NewSystemServer(svcCtx *svc.ServiceContext) *SystemServer {
	return &SystemServer{
		svcCtx: svcCtx,
	}
}

//  user service
func (s *SystemServer) Login(ctx context.Context, in *system.LoginReq) (*system.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *SystemServer) ChangePassword(ctx context.Context, in *system.ChangePasswordReq) (*system.BaseResp, error) {
	l := logic.NewChangePasswordLogic(ctx, s.svcCtx)
	return l.ChangePassword(in)
}

func (s *SystemServer) CreateUser(ctx context.Context, in *system.RegisterReq) (*system.BaseResp, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *SystemServer) UpdateUser(ctx context.Context, in *system.UpdateUserInfoReq) (*system.BaseResp, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *SystemServer) GetUserById(ctx context.Context, in *system.IDReq) (*system.UserInfoResp, error) {
	l := logic.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *SystemServer) GetUserList(ctx context.Context, in *system.PageInfoReq) (*system.UserInfoListResp, error) {
	l := logic.NewGetUserListLogic(ctx, s.svcCtx)
	return l.GetUserList(in)
}

//  menu service
func (s *SystemServer) CreateMenu(ctx context.Context, in *system.CreateMenuReq) (*system.BaseResp, error) {
	l := logic.NewCreateMenuLogic(ctx, s.svcCtx)
	return l.CreateMenu(in)
}

func (s *SystemServer) UpdateMenu(ctx context.Context, in *system.UpdateMenuReq) (*system.BaseResp, error) {
	l := logic.NewUpdateMenuLogic(ctx, s.svcCtx)
	return l.UpdateMenu(in)
}

func (s *SystemServer) DeleteMenu(ctx context.Context, in *system.IDReq) (*system.BaseResp, error) {
	l := logic.NewDeleteMenuLogic(ctx, s.svcCtx)
	return l.DeleteMenu(in)
}

func (s *SystemServer) GetMenuById(ctx context.Context, in *system.IDReq) (*system.MenuInfo, error) {
	l := logic.NewGetMenuByIdLogic(ctx, s.svcCtx)
	return l.GetMenuById(in)
}

func (s *SystemServer) GetMenuList(ctx context.Context, in *system.PageInfoReq) (*system.MenuInfoList, error) {
	l := logic.NewGetMenuListLogic(ctx, s.svcCtx)
	return l.GetMenuList(in)
}

func (s *SystemServer) CreateMenuParam(ctx context.Context, in *system.CreateMenuParamReq) (*system.BaseResp, error) {
	l := logic.NewCreateMenuParamLogic(ctx, s.svcCtx)
	return l.CreateMenuParam(in)
}

func (s *SystemServer) UpdateMenuParam(ctx context.Context, in *system.UpdateMenuParamReq) (*system.BaseResp, error) {
	l := logic.NewUpdateMenuParamLogic(ctx, s.svcCtx)
	return l.UpdateMenuParam(in)
}

func (s *SystemServer) DeleteMenuParam(ctx context.Context, in *system.IDReq) (*system.BaseResp, error) {
	l := logic.NewDeleteMenuParamLogic(ctx, s.svcCtx)
	return l.DeleteMenuParam(in)
}

func (s *SystemServer) GetMenuParamById(ctx context.Context, in *system.IDReq) (*system.MenuParamResp, error) {
	l := logic.NewGetMenuParamByIdLogic(ctx, s.svcCtx)
	return l.GetMenuParamById(in)
}

func (s *SystemServer) GeMenuParamListById(ctx context.Context, in *system.IDReq) (*system.MenuParamListResp, error) {
	l := logic.NewGeMenuParamListByIdLogic(ctx, s.svcCtx)
	return l.GeMenuParamListById(in)
}

// menu role management
func (s *SystemServer) CreateMenuRole(ctx context.Context, in *system.CreateMenuRoleReq) (*system.BaseResp, error) {
	l := logic.NewCreateMenuRoleLogic(ctx, s.svcCtx)
	return l.CreateMenuRole(in)
}

func (s *SystemServer) UpdateMenuRole(ctx context.Context, in *system.UpdateMenuRoleReq) (*system.BaseResp, error) {
	l := logic.NewUpdateMenuRoleLogic(ctx, s.svcCtx)
	return l.UpdateMenuRole(in)
}

func (s *SystemServer) DeleteMenuRole(ctx context.Context, in *system.IDReq) (*system.BaseResp, error) {
	l := logic.NewDeleteMenuRoleLogic(ctx, s.svcCtx)
	return l.DeleteMenuRole(in)
}

func (s *SystemServer) GetMenuRoleById(ctx context.Context, in *system.IDReq) (*system.MenuRoleResp, error) {
	l := logic.NewGetMenuRoleByIdLogic(ctx, s.svcCtx)
	return l.GetMenuRoleById(in)
}

func (s *SystemServer) GetMenuRoleList(ctx context.Context, in *system.PageInfoReq) (*system.MenuRoleListResp, error) {
	l := logic.NewGetMenuRoleListLogic(ctx, s.svcCtx)
	return l.GetMenuRoleList(in)
}

//  role service
func (s *SystemServer) CreateRole(ctx context.Context, in *system.CreateRoleReq) (*system.BaseResp, error) {
	l := logic.NewCreateRoleLogic(ctx, s.svcCtx)
	return l.CreateRole(in)
}

func (s *SystemServer) UpdateRole(ctx context.Context, in *system.CreateRoleReq) (*system.BaseResp, error) {
	l := logic.NewUpdateRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(in)
}

func (s *SystemServer) DeleteRole(ctx context.Context, in *system.IDReq) (*system.BaseResp, error) {
	l := logic.NewDeleteRoleLogic(ctx, s.svcCtx)
	return l.DeleteRole(in)
}

func (s *SystemServer) GetRoleById(ctx context.Context, in *system.IDReq) (*system.RoleResp, error) {
	l := logic.NewGetRoleByIdLogic(ctx, s.svcCtx)
	return l.GetRoleById(in)
}

func (s *SystemServer) GetRoleList(ctx context.Context, in *system.PageInfoReq) (*system.RoleListResp, error) {
	l := logic.NewGetRoleListLogic(ctx, s.svcCtx)
	return l.GetRoleList(in)
}

//  casbin service
func (s *SystemServer) UpdatePolicy(ctx context.Context, in *system.UpdatePolicyReq) (*system.BaseResp, error) {
	l := logic.NewUpdatePolicyLogic(ctx, s.svcCtx)
	return l.UpdatePolicy(in)
}

func (s *SystemServer) CreatePolicy(ctx context.Context, in *system.CreatePolicyReq) (*system.BaseResp, error) {
	l := logic.NewCreatePolicyLogic(ctx, s.svcCtx)
	return l.CreatePolicy(in)
}

func (s *SystemServer) DeletePolicy(ctx context.Context, in *system.IDReq) (*system.BaseResp, error) {
	l := logic.NewDeletePolicyLogic(ctx, s.svcCtx)
	return l.DeletePolicy(in)
}

func (s *SystemServer) GetPolicyByRoleId(ctx context.Context, in *system.IDReq) (*system.UpdatePolicyReq, error) {
	l := logic.NewGetPolicyByRoleIdLogic(ctx, s.svcCtx)
	return l.GetPolicyByRoleId(in)
}
