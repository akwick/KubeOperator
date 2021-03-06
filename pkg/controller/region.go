package controller

import (
	"github.com/KubeOperator/KubeOperator/pkg/constant"
	"github.com/KubeOperator/KubeOperator/pkg/controller/log_save"
	"github.com/KubeOperator/KubeOperator/pkg/controller/page"
	"github.com/KubeOperator/KubeOperator/pkg/dto"
	"github.com/KubeOperator/KubeOperator/pkg/service"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12/context"
)

type RegionController struct {
	Ctx           context.Context
	RegionService service.RegionService
}

func NewRegionController() *RegionController {
	return &RegionController{
		RegionService: service.NewRegionService(),
	}
}

// List Region
// @Tags regions
// @Summary Show all regions
// @Description Show regions
// @Accept  json
// @Produce  json
// @Success 200 {object} page.Page
// @Security ApiKeyAuth
// @Router /regions/ [get]
func (r RegionController) Get() (page.Page, error) {

	p, _ := r.Ctx.Values().GetBool("page")
	if p {
		num, _ := r.Ctx.Values().GetInt(constant.PageNumQueryKey)
		size, _ := r.Ctx.Values().GetInt(constant.PageSizeQueryKey)
		return r.RegionService.Page(num, size)
	} else {
		var page page.Page
		items, err := r.RegionService.List()
		if err != nil {
			return page, err
		}
		page.Items = items
		page.Total = len(items)
		return page, nil
	}
}

// Get Region
// @Tags regions
// @Summary Show a Region
// @Description show a region by name
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.Region
// @Security ApiKeyAuth
// @Router /regions/{name}/ [get]
func (r RegionController) GetBy(name string) (dto.Region, error) {
	return r.RegionService.Get(name)
}

// Create Region
// @Tags regions
// @Summary Create a region
// @Description create a region
// @Accept  json
// @Produce  json
// @Param request body dto.RegionCreate true "request"
// @Success 200 {object} dto.Region
// @Security ApiKeyAuth
// @Router /regions/ [post]
func (r RegionController) Post() (*dto.Region, error) {
	var req dto.RegionCreate
	err := r.Ctx.ReadJSON(&req)
	if err != nil {
		return nil, err
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return nil, err
	}

	operator := r.Ctx.Values().GetString("operator")
	go log_save.LogSave(operator, constant.CREATE_REGION, req.Name)

	return r.RegionService.Create(req)
}

// Delete Region
// @Tags regions
// @Summary Delete a region
// @Description delete a region by name
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Router /regions/{name}/ [delete]
func (r RegionController) Delete(name string) error {
	operator := r.Ctx.Values().GetString("operator")
	go log_save.LogSave(operator, constant.DELETE_REGION, name)

	return r.RegionService.Delete(name)
}

func (r RegionController) PostBatch() error {
	var req dto.RegionOp
	err := r.Ctx.ReadJSON(&req)
	if err != nil {
		return err
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return err
	}
	err = r.RegionService.Batch(req)
	if err != nil {
		return err
	}

	operator := r.Ctx.Values().GetString("operator")
	delRegions := ""
	for _, item := range req.Items {
		delRegions += (item.Name + ",")
	}
	go log_save.LogSave(operator, constant.DELETE_REGION, delRegions)

	return err
}

func (r RegionController) PostDatacenter() (*dto.CloudRegionResponse, error) {
	var req dto.RegionDatacenterRequest
	err := r.Ctx.ReadJSON(&req)
	if err != nil {
		return nil, err
	}

	data, err := r.RegionService.ListDatacenter(req)
	if err != nil {
		return nil, err
	}
	return &dto.CloudRegionResponse{Result: data}, err
}
