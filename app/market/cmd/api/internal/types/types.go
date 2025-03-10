// Code generated by goctl. DO NOT EDIT.
package types

type Nil struct {
}

type CommonResponse struct {
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
}

type RoleListReq struct {
	Param interface{} `form:"param"`
}

type GetConfigReq struct {
	AppId   int64 `form:"appId"`
	GroupId int64 `form:"groupId"`
}

type ListReq struct {
	GroupAppDistribution interface{} `form:"GroupAppDistributionData"`
}

type ListPageReq struct {
	GroupAppDistribution interface{} `form:"GroupAppDistributionData"`
	Current              int64       `form:"current"`
	Size                 int64       `form:"size"`
}

type GetTokenReq struct {
	Data interface{} `form:"data"`
}

type AppCheckReq struct {
	IdList      []string `json:"idList"`
	CheckStatus bool     `json:"checkStatus"`
}

type CancelApplyReq struct {
	IdList []int64 `json:"idList"`
}

type DeleteAppReq struct {
	AppId       int64 `form:"appId"`
	RemoveAdmin bool  `form:"removeAdmin,optional,default=false"`
}

type DetailReq struct {
	Id int64 `form:"id"`
}

type GetAllReq struct {
	Id int64 `form:"id"`
}

type GetAllAppReq struct {
	Id int64 `form:"id"`
}

type FreezeAppReq struct {
	Ids interface{} `form:"ids"`
}

type GetRedeployAppListReq struct {
	AppName string `form:"appName,optional"`
	Current int64  `form:"current"`
	Flag    int64  `form:"flag"`
	Size    int64  `form:"size"`
}

type ListReq1 struct {
	Current   int64       `form:"current"`
	MarketApp interface{} `form:"marketApp"`
	Size      int64       `form:"size"`
}

type AppListReq struct {
	Current    int64  `form:"current"`
	AppName    string `form:"appName,optional"`
	Size       int64  `form:"size"`
	TenantId   string `form:"tenantId,optional"`
	SaleStatus int64  `form:"saleStatus,optional,default=-1"`
}

type ListCheckReq struct {
	AppName string `form:"appName,optional"`
	Count   int64  `form:"count,optional,default=-1"`
	Current int64  `form:"current"`
	Size    int64  `form:"size"`
}

type ListAllReq struct {
	MarketApp interface{} `form:"marketApp"`
}

type ListAppBySortReq struct {
	Data interface{} `form:"data"`
}

type ListAppListReq struct {
	TenantCode string `form:"tenantCode"`
}

type ListByDistributeTenantIdReq struct {
	Current  int64  `form:"current"`
	Size     int64  `form:"size"`
	TenantId string `form:"tenantId"`
}

type ListByPurchaseGroupIdReq struct {
	AppName string `form:"appName,optional"`
	Current int64  `form:"current"`
	GroupId int64  `form:"groupId"`
	Size    int64  `form:"size"`
}

type ListByPurchaseTenantIdReq struct {
	AppName  string `form:"appName,optional"`
	Current  int64  `form:"current"`
	Size     int64  `form:"size"`
	TenantId string `form:"tenantId"`
}

type ReformAppReq struct {
	Ids int64 `form:"ids"`
}

type RemoveReq struct {
	Ids interface{} `form:"ids"`
}

type SetOffSaleReq struct {
	IdList []int64 `json:"idList"`
}

type SetOnSaleReq struct {
	IdList []int64 `json:"idList"`
}

type SubmitReq struct {
	MarketApp interface{} `form:"marketApp"`
}

type SubmitAllReq struct {
	Application       interface{} `form:"application"`
	MarketApplication interface{} `form:"marketApplication"`
}

type SubmitFirstReq struct {
	MarketApp   interface{} `form:"MarketApp"`
	Application interface{} `form:"application"`
}

type SubmitSecondReq struct {
	AppId int64 `json:"appId,string"`
}

type SubmitThirdReq struct {
	Application       interface{} `form:"application"`
	MarketApplication interface{} `form:"marketApplication"`
}

type TokenDetailReq struct {
	Id int64 `form:"id"`
}

type TokenListAppBySortReq struct {
	Data interface{} `form:"data"`
}

type TokenListWithoutTokenReq struct {
	MarketApp interface{} `form:"marketApp"`
}

type CancelCheckAlertReq struct {
	Id int64 `form:"id"`
}

type CheckAlertReq struct {
	Id int64 `form:"id"`
}

type ListReq2 struct {
	Current    int64       `form:"current"`
	AlertTitle interface{} `form:"alertTitle"`
	Size       int64       `form:"size"`
}

type MyAlertListReq struct {
	JsonObject interface{} `form:"jsonObject"`
}

type RemoveReq1 struct {
	Ids interface{} `form:"ids"`
}

type SendAlertReq struct {
	Dto interface{} `form:"dto"`
}

type SubmitReq1 struct {
	MarketAppAlert interface{} `form:"marketAppAlert"`
}

type ApplyReq struct {
	AppDescription string `form:"appDescription"`
	AppName        string `form:"appName"`
}

type ApplyCancelReq struct {
	Id int64 `form:"id"`
}

type ApplyhandleReq struct {
	Feedback string `form:"feedback"`
	Id       int64  `form:"id"`
	Status   int64  `form:"status"`
}

type DetailReq1 struct {
	Id int64 `form:"id"`
}

type GetApplyDetailReq struct {
	Id int64 `form:"id"`
}

type GetMyListReq struct {
	Current        int64       `form:"current"`
	MarketAppApply interface{} `form:"marketAppApply"`
	Size           int64       `form:"size"`
}

type GetSecretByKeyReq struct {
	AppKey string `form:"appKey"`
}

type ListReq3 struct {
	Current        int64       `form:"current"`
	MarketAppApply interface{} `form:"marketAppApply"`
	Size           int64       `form:"size"`
}

type RemoveReq2 struct {
	Ids interface{} `form:"ids"`
}

type SubmitReq2 struct {
	MarketAppApply interface{} `form:"marketAppApply"`
}

type TestReq struct {
	A int64 `form:"a"`
	B int64 `form:"b"`
}

type UpdateApplyReq struct {
	AppDescription string `form:"appDescription"`
	AppName        string `form:"appName"`
	Id             int64  `form:"id"`
}

type GetComponentListByAppIdReq struct {
	AppId   int64 `form:"appId"`
	Current int64 `form:"current"`
	Size    int64 `form:"size"`
}

type ListReq4 struct {
	Current int64  `form:"current"`
	Size    int64  `form:"size"`
	Name    string `form:"name,optional"`
	Flag    int64  `form:"flag,optional,default=-1"`
}

type RemoveReq3 struct {
	Ids interface{} `form:"ids"`
}

type SubmitReq3 struct {
	MarketAppComponent interface{} `form:"marketAppComponent"`
}

type GetComponentByTemplateIdReq struct {
	TemplateId int64 `form:"templateId"`
}

type GetUserTemplateReq struct {
	Current int64  `form:"current"`
	Name    string `form:"name,optional"`
	Size    int64  `form:"size"`
}

type ListReq5 struct {
	Current int64  `form:"current"`
	Size    int64  `form:"size"`
	Name    string `form:"name,optional"`
	Status  int64  `form:"status"`
}

type RemoveReq4 struct {
	Ids interface{} `form:"ids"`
}

type SetDefaultReq struct {
	Id int64 `form:"id"`
}

type SubmitReq4 struct {
	MarketAppComponentTe interface{} `form:"marketAppComponentTemplate"`
}

type DetailReq2 struct {
	Id int64 `form:"id"`
}

type DistributeAppTenantListReq struct {
	AppId    int64  `form:"appId"`
	Current  int64  `form:"current"`
	GroupId  int64  `form:"groupId,optional"`
	UnitName string `form:"unitName,optional"`
	Size     int64  `form:"size"`
}

type GetDistributedGroupIdListReq struct {
	AppId      int64  `form:"appId"`
	TenantCode string `form:"tenantCode"`
}

type GetGroupDistributeTenantAppReq struct {
	GroupId    int64  `form:"groupId"`
	TenantCode string `form:"tenantCode"`
}

type GetGroupIdListReq struct {
	AppId int64 `form:"appId"`
}

type RemoveReq5 struct {
	IdList interface{} `form:"idList"`
}

type SubmitReq5 struct {
	MarketAppGroupDistri interface{} `form:"marketAppGroupDistribution"`
}

type SubmitAllReq1 struct {
	MarketAppGroupDistri interface{} `form:"marketAppGroupDistribution"`
}

type UnitCancelDistributeReq struct {
	AppId    int64  `json:"appId,string"`
	GroupId  int64  `json:"groupId,string"`
	TenantId string `json:"tenantId"`
}

type UnitDistributeReq struct {
	MarketAppDistributio interface{} `form:"marketAppDistributionRelationDto"`
}

type GetRelationReq struct {
	AppId   int64 `form:"appId"`
	GroupId int64 `form:"groupId"`
}

type DetailReq3 struct {
	Id int64 `form:"id"`
}

type GetByAppIdReq struct {
	AppId int64 `form:"appId"`
}

type ListReq6 struct {
	Current            int64       `form:"current"`
	MarketAppKeySecret interface{} `form:"marketAppKeySecret"`
	Size               int64       `form:"size"`
}

type RemoveReq6 struct {
	IdList interface{} `form:"idList"`
}

type SubmitReq6 struct {
	MarketAppKeySecret interface{} `form:"marketAppKeySecret"`
}

type CancelReleaseNoteReq struct {
	Id int64 `form:"id"`
}

type ListReq7 struct {
	Size        int64  `form:"size"`
	Current     int64  `form:"current"`
	GroupOrUnit int64  `form:"groupOrUnit"`
	NoticeTitle string `form:"noticeTitle,optional"`
}

type MyNoticeListReq struct {
	Size    int64 `json:"size"`
	Current int64 `json:"current"`
}

type ReleaseNoteReq struct {
	Id int64 `form:"id"`
}

type RemoveReq7 struct {
	Ids string `form:"ids"`
}

type SubmitReq7 struct {
	Id                  string `json:"id,optional"`
	GroupOrUnit         int64  `json:"groupOrUnit,optional,default=-1"`
	NoticeTitle         string `json:"noticeTitle,optional"`
	NoticeContent       string `json:"noticeContent,optional"`
	NoticeUnitIds       string `json:"noticeUnitIds,optional"`
	NoticeRoleIds       string `json:"noticeRoleIds,optional"`
	NoticeReleaseStatus int64  `json:"noticeReleaseStatus,optional"`
	NoticeReleaseUnitId string `json:"noticeReleaseUnitId,optional"`
}

type DetailReq4 struct {
	Id int64 `form:"id"`
}

type GetListsReq struct {
	Current    string `form:"current"`
	Flag       string `form:"flag"`
	GroupId    string `form:"groupId"`
	Size       string `form:"size"`
	TenantCode string `form:"tenantCode"`
	UserId     string `form:"userId"`
}

type GetAppPurchaseConfigReq struct {
	AppId int64 `form:"appId"`
}

type ListReq8 struct {
	Current           int64       `form:"current"`
	MarketAppPurchase interface{} `form:"marketAppPurchase"`
	Size              int64       `form:"size"`
}

type ListAppByTenantIdReq struct {
	Current  int64  `form:"current"`
	Size     int64  `form:"size"`
	TenantId string `form:"tenantId"`
}

type ListByGroupIdVOReq struct {
	GroupId    int64  `form:"groupId"`
	SaleStatus int64  `form:"saleStatus,optional"`
	Current    int64  `form:"current"`
	AppName    string `form:"appName,optional"`
	Size       int64  `form:"size"`
}

type ListUnitAndPersonalVOReq struct {
	Current int64  `form:"current"`
	Name    string `form:"name,optional"`
	Size    int64  `form:"size"`
}

type ListVOReq struct {
	SaleStatus interface{} `form:"saleStatus"`
	Current    int64       `form:"current"`
	AppName    string      `form:"appName"`
	Size       int64       `form:"size"`
}

type PurchaseAppGroupListReq struct {
	Current int64 `form:"current"`
	AppId   int64 `form:"appId"`
	Size    int64 `form:"size"`
}

type PurchaseAppTenantListReq struct {
	Current int64 `form:"current"`
	AppId   int64 `form:"appId"`
	Size    int64 `form:"size"`
}

type PurchaseByGroupsReq struct {
	AppId    int64    `json:"appId,string"`
	GroupIds []string `json:"groupId"`
}

type RemoveReq8 struct {
	Ids interface{} `form:"ids"`
}

type MarketAppPurchasSubmitReq struct {
	AppId      int64  `json:"appId,string"`
	TenantCode string `json:"tenantId,optional"`
	GroupId    string `json:"groupId,optional"`
}

type UnitunsubscribeReq struct {
	AppIds  []string `json:"appIds"`
	GroupId int64    `json:"groupid,string"`
}

type UnsubscribeReq struct {
	AppIds []int64 `json:"appIds,string"`
}

type UpdateGroupUseStatusReq struct {
	Data interface{} `form:"data"`
}

type UpdateTenantUseStatusReq struct {
	Data interface{} `form:"data"`
}

type WithdrawGroupAuthorityReq struct {
	Data interface{} `form:"data"`
}

type WithdrawTenantAuthorityReq struct {
	Data interface{} `form:"data"`
}

type DetailReq5 struct {
	Id int64 `form:"id"`
}

type GetRoleListReq struct {
	AppId int64 `form:"appId"`
}

type GetRoleNameListReq struct {
	AppId int64 `form:"appId"`
}

type ListReq9 struct {
	AppRole interface{} `form:"appRole"`
	Current int64       `form:"current"`
	Size    int64       `form:"size"`
}

type RemoveReq9 struct {
	Ids interface{} `form:"ids"`
}

type SubmitReq9 struct {
	MarketAppRole interface{} `form:"marketAppRole"`
}

type ApptokenVerifyUserReq struct {
	AppKey    string `form:"appKey"`
	AppSecret string `form:"appSecret"`
	UserId    int64  `form:"userId"`
}

type DetailReq6 struct {
	Id int64 `form:"id"`
}

type GetAppDistributedRoleReq struct {
	AppId      int64  `form:"appId,string"`
	TenantCode string `form:"tenantCode"`
}

type GetAppIdByUserIdReq struct {
	UserId int64 `form:"userId"`
}

type GetDistributedAgencyReq struct {
	RoleId     int64  `form:"roleId"`
	TenantCode string `form:"tenantCode"`
	AppId      int64  `form:"appId"`
}

type GetDistributedJobReq struct {
	AppId      int64  `form:"appId"`
	RoleId     int64  `form:"roleId"`
	TenantCode string `form:"tenantCode"`
}

type GetDistributedPersonListByAppIdReq struct {
	AppId  int64 `form:"appId"`
	RoleId int64 `form:"roleId"`
}

type GetDistributionAgencyReq struct {
	RoleId int64 `form:"roleId"`
}

type GetDistributionJobReq struct {
	RoleId int64 `form:"roleId"`
}

type GetDistributionPersonReq struct {
	AppId      int64  `form:"appId"`
	RoleName   string `form:"roleName"`
	TenantCode string `form:"tenantCode"`
}

type GetDistributionUserReq struct {
	RoleId int64 `form:"roleId"`
}

type GetRoleByUserIdReq struct {
	UserId int64 `form:"userId"`
}

type GetUserRoleListReq struct {
	AppId      int64  `form:"appId"`
	TenantCode string `form:"tenantCode"`
	UserId     int64  `form:"userId"`
}

type IsDistributedReq struct {
	AppId int64 `form:"appId"`
}

type ListReq10 struct {
	AppRoleDistribution interface{} `form:"appRoleDistribution"`
	Current             int64       `form:"current"`
	Size                int64       `form:"size"`
}

type RemoveReq10 struct {
	Ids interface{} `form:"ids"`
}

type SubmitReq10 struct {
	MarketAppRoleDistrib interface{} `form:"marketAppRoleDistribution"`
}

type SubmitV2Req struct {
	RoleId    int64  `json:"roleId,string"`
	JobIds    string `json:"jobIds,optional"`
	AgencyIds string `json:"agencyIds,optional"`
	UserIds   string `json:"userIds,optional"`
}

type DetailReq7 struct {
	Id int64 `form:"id"`
}

type ListReq11 struct {
	Current           int64       `form:"current"`
	MarketAppRoleMenu interface{} `form:"marketAppRoleMenu"`
	Size              int64       `form:"size"`
}

type RemoveReq11 struct {
	Ids interface{} `form:"ids"`
}

type SubmitReq11 struct {
	MarketAppRoleMenu interface{} `form:"marketAppRoleMenu"`
}

type GetDefaultTemplateReq struct {
	UserId int64 `form:"userId"`
}

type GetTemplateIdByUserIdInUseReq struct {
	UserId int64 `form:"userId"`
}

type GetTemplateIdListByUserIdReq struct {
	UserId int64 `form:"userId"`
}

type ListReq12 struct {
	Current              int64       `form:"current"`
	MarketAppUserTemplat interface{} `form:"marketAppUserTemplate"`
	Size                 int64       `form:"size"`
}

type AppComponentGroupByUserReq struct {
	ComponentName string `form:"componentName"`
	TenantCode    string `form:"tenantCode"`
	AppName       string `form:"appName"`
}

type RemoveReq12 struct {
	Ids interface{} `form:"ids"`
}

type SetDefaultTemplateReq struct {
	Data interface{} `form:"data"`
}

type SetGroupTenantDefaultTemplateReq struct {
	Data interface{} `form:"data"`
}

type SubmitReq12 struct {
	MarketAppUserTemplat interface{} `form:"marketAppUserTemplate"`
}

type DetailReq8 struct {
	Id int64 `form:"id"`
}

type GetAppMenuListReq struct {
	AppId int64 `form:"appId"`
}

type GetMenuListReq struct {
	Platform int64  `form:"platform"`
	TenantId string `form:"tenantId"`
	UserId   int64  `form:"userId"`
}

type GetMobileMenuListReq struct {
	AppId  int64 `form:"appId"`
	UserId int64 `form:"userId"`
}

type GetParentMobileMenuListReq struct {
	UserId int64 `form:"userId"`
}

type GetSonMobileMenuListReq struct {
	MenuId int64 `form:"menuId"`
	UserId int64 `form:"userId"`
}

type ListReq13 struct {
	Current    int64       `form:"current"`
	MarketMenu interface{} `form:"marketMenu"`
	Size       int64       `form:"size"`
}

type RemoveReq13 struct {
	Ids interface{} `form:"ids"`
}

type SubmitReq13 struct {
	MarketMenu interface{} `form:"marketMenu"`
}

type DetailReq9 struct {
	Id int64 `form:"id"`
}

type GetUserMenuSortListReq struct {
	UserId int64 `form:"userId"`
}

type ListReq14 struct {
	Current            int64       `form:"current"`
	MarketMenuUserSort interface{} `form:"marketMenuUserSort"`
	Size               int64       `form:"size"`
}

type RemoveReq14 struct {
	Ids interface{} `form:"ids"`
}

type SortUserMenuReq struct {
	Data interface{} `form:"data"`
}

type SortUserMenuByUnitReq struct {
	Data interface{} `form:"data"`
}

type SubmitReq14 struct {
	MarketMenuUserSort interface{} `form:"marketMenuUserSort"`
}

type GetUsedAppReq struct {
	UserId int64 `form:"userId"`
}

type GetUsedAppMenuReq struct {
	UserId int64 `form:"userId"`
}

type ListReq15 struct {
	Current       int64       `form:"current"`
	MarketUsedApp interface{} `form:"marketUsedApp"`
	Size          int64       `form:"size"`
}

type RemoveReq15 struct {
	Ids interface{} `form:"ids"`
}

type RemoveByRelationReq struct {
	AppId  int64 `json:"appId,string"`
	UserId int64 `json:"userId,string"`
}

type SortUsedAppReq struct {
	AppIdList interface{} `form:"appIdList"`
}

type SubmitReq15 struct {
	UserId int64 `json:"userId,string"`
	AppId  int64 `json:"appId,string"`
}

type AddUniAppReq struct {
	Ids interface{} `form:"ids"`
}

type CancelUniAppReq struct {
	Ids interface{} `form:"ids"`
}

type GetAllAppReq1 struct {
	Platform int64 `form:"platform"`
}

type GetCommonUseAppReq struct {
	TenantId string `form:"tenantId"`
	UserId   int64  `form:"userId"`
}

type GetRemainAppReq struct {
	TenantId string `form:"tenantId"`
	UserId   int64  `form:"userId"`
}

type GetUserPortalReq struct {
	TenantId string `form:"tenantId"`
	UserId   int64  `form:"userId"`
}

type ListReq16 struct {
	Current int64 `form:"current"`
	Size    int64 `form:"size"`
}

type RemoveReq16 struct {
	Ids interface{} `form:"ids"`
}

type SaveOrUpdateAllPortalReq struct {
	Portals  interface{} `form:"portals"`
	UserId   int64       `form:"userId"`
	TenantId string      `form:"tenantId"`
}

type SubmitReq16 struct {
	Portal interface{} `form:"portal"`
}

type V2GetCommonUseAppReq struct {
	TenantId string `form:"tenantId"`
	UserId   int64  `form:"userId"`
}

type CheckRedeployReq struct {
	CheckStatus bool     `json:"checkStatus"`
	RedeployIds []string `json:"redeployIds"`
}

type IsAddRoleReq struct {
	AppId    int64  `form:"appId"`
	RoleName string `form:"roleName"`
}

type IsDeleteRoleReq struct {
	AppId    int64  `form:"appId"`
	RoleName string `form:"roleName"`
}

type ListReq17 struct {
	Current      int64       `form:"current"`
	ReDeployData interface{} `form:"reDeployData"`
	Size         int64       `form:"size"`
}

type ListAllReq2 struct {
	ReDeployData interface{} `form:"reDeployData"`
}

type RedeployReq struct {
	MarketApplication interface{} `form:"marketApplication"`
}
