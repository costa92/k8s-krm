package api

import (
	"github.com/costa92/krm/internal/api/controller/v1/user"
	"github.com/costa92/krm/pkg/server"
	"net/http"
)

func getUserResource() *server.GroupResource {
	var groupResourceApis []*server.GroupResourceApi
	userController := user.NewUser()
	groupResourceApis = append(groupResourceApis, &server.GroupResourceApi{
		Method:  http.MethodGet,
		Path:    "/list",
		Handler: userController.List,
	})

	userApi := &server.GroupResource{
		Version: "v1",
		Group:   "user",
		Apis:    groupResourceApis,
	}

	return userApi
}

func getUserV2Resource() *server.GroupResource {
	var groupResourceApis []*server.GroupResourceApi
	userController := user.NewUser()
	groupResourceApis = append(groupResourceApis, &server.GroupResourceApi{
		Method:  http.MethodGet,
		Path:    "/list",
		Handler: userController.List,
	}, &server.GroupResourceApi{
		Method:  http.MethodGet,
		Path:    "/list2",
		Handler: userController.List,
	})

	userApi := &server.GroupResource{
		Version: "v2",
		Group:   "user",
		Apis:    groupResourceApis,
	}

	return userApi
}

func GetGroupResource() []*server.GroupResource {
	var groupResources []*server.GroupResource
	groupResources = append(
		groupResources,
		getUserResource(),
		getUserV2Resource(),
	)
	return groupResources
}
