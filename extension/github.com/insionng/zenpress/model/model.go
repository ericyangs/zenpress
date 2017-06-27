package model

import (
	"github.com/insionng/zenpress/model"

	"qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/insionng/zenpress/model",

	"DataType":            model.DataType,
	"Database":            model.Database,
	"DatabaseConn":        model.DatabaseConn,
	"DatabaseTablePrefix": model.DatabaseTablePrefix,
	"Enforcer":            model.Enforcer,
	"HasDatabase":         model.HasDatabase,

	"AddLink":                                 model.AddLink,
	"AddOption":                               model.AddOption,
	"ConnDatabase":                            model.ConnDatabase,
	"CreateTables":                            model.CreateTables,
	"DeleteLink":                              model.DeleteLink,
	"DeleteOption":                            model.DeleteOption,
	"DeletePermission":                        model.DeletePermission,
	"DeleteRole":                              model.DeleteRole,
	"DeleteRolePermissionByPermissionId":      model.DeleteRolePermissionByPermissionId,
	"DeleteRolePermissionByRoleId":            model.DeleteRolePermissionByRoleId,
	"DeleteUser":                              model.DeleteUser,
	"DeleteUserRolesByUserId":                 model.DeleteUserRolesByUserId,
	"FindPermissionByUserIdAndPermissionName": model.FindPermissionByUserIdAndPermissionName,
	"FindPermissions":                         model.FindPermissions,
	"FindPermissionsByPid":                    model.FindPermissionsByPid,
	"FindRolePermissionByRoleId":              model.FindRolePermissionByRoleId,
	"FindRoles":                               model.FindRoles,
	"FindUserById":                            model.FindUserById,
	"FindUserByToken":                         model.FindUserByToken,
	"FindUserByUserName":                      model.FindUserByUserName,
	"FindUserRolesByUserId":                   model.FindUserRolesByUserId,
	"GetLink":                                 model.GetLink,
	"GetOption":                               model.GetOption,
	"Init":                                    model.Init,
	"Login":                                   model.Login,
	"NewDatabase":                             model.NewDatabase,
	"NewLink":                                 model.NewLink,
	"PageUser":                                model.PageUser,
	"Ping":                                    model.Ping,
	"SavePermission":                          model.SavePermission,
	"SaveRole":                                model.SaveRole,
	"SaveRolePermission":                      model.SaveRolePermission,
	"SaveUser":                                model.SaveUser,
	"SaveUserRole":                            model.SaveUserRole,
	"SetDatabase":                             model.SetDatabase,
	"UpdateLink":                              model.UpdateLink,
	"UpdateOption":                            model.UpdateOption,
	"UpdatePermission":                        model.UpdatePermission,
	"UpdateRole":                              model.UpdateRole,
	"UpdateUser":                              model.UpdateUser,

	"App":                spec.StructOf((*model.App)(nil)),
	"AppVersion":         spec.StructOf((*model.AppVersion)(nil)),
	"Comment":            spec.StructOf((*model.Comment)(nil)),
	"Commentmeta":        spec.StructOf((*model.Commentmeta)(nil)),
	"Link":               spec.StructOf((*model.Link)(nil)),
	"Model":              spec.StructOf((*model.Model)(nil)),
	"Option":             spec.StructOf((*model.Option)(nil)),
	"Permission":         spec.StructOf((*model.Permission)(nil)),
	"FindPermissionById": model.FindPermissionById,
	"Post":               spec.StructOf((*model.Post)(nil)),
	"Postmeta":           spec.StructOf((*model.Postmeta)(nil)),
	"RegistrationLog":    spec.StructOf((*model.RegistrationLog)(nil)),
	"Result":             spec.StructOf((*model.Result)(nil)),
	"Role":               spec.StructOf((*model.Role)(nil)),
	"FindRoleById":       model.FindRoleById,
	"RolePermission":     spec.StructOf((*model.RolePermission)(nil)),
	"Signup":             spec.StructOf((*model.Signup)(nil)),
	"Site":               spec.StructOf((*model.Site)(nil)),
	"Sitemeta":           spec.StructOf((*model.Sitemeta)(nil)),
	"Term":               spec.StructOf((*model.Term)(nil)),
	"TermRelationship":   spec.StructOf((*model.TermRelationship)(nil)),
	"TermTaxonomy":       spec.StructOf((*model.TermTaxonomy)(nil)),
	"Termmeta":           spec.StructOf((*model.Termmeta)(nil)),
	"User":               spec.StructOf((*model.User)(nil)),
	"UserRoleResult":     spec.StructOf((*model.UserRoleResult)(nil)),
	"Usermeta":           spec.StructOf((*model.Usermeta)(nil)),
}
