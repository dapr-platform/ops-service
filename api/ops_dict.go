package api

import (
	"github.com/dapr-platform/common"
	"github.com/go-chi/chi/v5"
	"net/http"
	"ops-service/model"

	"strings"

	"time"
)

var _ = time.Now()

func InitOps_dictRoute(r chi.Router) {

	r.Get(common.BASE_CONTEXT+"/ops-dict/page", Ops_dictPageListHandler)
	r.Get(common.BASE_CONTEXT+"/ops-dict", Ops_dictListHandler)

	r.Post(common.BASE_CONTEXT+"/ops-dict", UpsertOps_dictHandler)

	r.Delete(common.BASE_CONTEXT+"/ops-dict/{id}", DeleteOps_dictHandler)

	r.Post(common.BASE_CONTEXT+"/ops-dict/batch-delete", batchDeleteOps_dictHandler)

	r.Post(common.BASE_CONTEXT+"/ops-dict/batch-upsert", batchUpsertOps_dictHandler)

}

// @Summary batch update
// @Description batch update
// @Tags 字典
// @Accept  json
// @Param entities body []map[string]any true "objects array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ops-dict/batch-upsert [post]
func batchUpsertOps_dictHandler(w http.ResponseWriter, r *http.Request) {

	var entities []model.Ops_dict
	err := common.ReadRequestBody(r, &entities)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(entities) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of entities is 0"))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Ops_dict")
	if exists {
		for _, v := range entities {
			_, err1 := beforeHook(r, v)
			if err1 != nil {
				common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
				return
			}
		}

	}
	for _, v := range entities {
		if v.ID == "" {
			v.ID = common.NanoId()
		}

		if time.Time(v.CreatedTime).IsZero() {
			v.CreatedTime = common.LocalTime(time.Now())
		}

		if time.Time(v.UpdatedTime).IsZero() {
			v.UpdatedTime = common.LocalTime(time.Now())
		}

	}

	err = common.DbBatchUpsert[model.Ops_dict](r.Context(), common.GetDaprClient(), entities, model.Ops_dictTableInfo.Name, model.Ops_dict_FIELD_NAME_id)
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}

// @Summary page query
// @Description page query, _page(from 1 begin), _page_size, _order, and others fields, status=1, name=$like.%CAM%
// @Tags 字典
// @Param _page query int true "current page"
// @Param _page_size query int true "page size"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param dict_type query string false "dict_type"
// @Param dict_code query string false "dict_code"
// @Param dict_name query string false "dict_name"
// @Param dict_value query string false "dict_value"
// @Param sort_order query string false "sort_order"
// @Param remarks query string false "remarks"
// @Produce  json
// @Success 200 {object} common.Response{data=common.Page{items=[]model.Ops_dict}} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ops-dict/page [get]
func Ops_dictPageListHandler(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("_page")
	pageSize := r.URL.Query().Get("_page_size")
	if page == "" || pageSize == "" {
		common.HttpResult(w, common.ErrParam.AppendMsg("page or pageSize is empty"))
		return
	}
	common.CommonPageQuery[model.Ops_dict](w, r, common.GetDaprClient(), "o_ops_dict", "id")

}

// @Summary query objects
// @Description query objects
// @Tags 字典
// @Param _select query string false "_select"
// @Param _order query string false "order"
// @Param id query string false "id"
// @Param created_by query string false "created_by"
// @Param created_time query string false "created_time"
// @Param updated_by query string false "updated_by"
// @Param updated_time query string false "updated_time"
// @Param dict_type query string false "dict_type"
// @Param dict_code query string false "dict_code"
// @Param dict_name query string false "dict_name"
// @Param dict_value query string false "dict_value"
// @Param sort_order query string false "sort_order"
// @Param remarks query string false "remarks"
// @Produce  json
// @Success 200 {object} common.Response{data=[]model.Ops_dict} "objects array"
// @Failure 500 {object} common.Response ""
// @Router /ops-dict [get]
func Ops_dictListHandler(w http.ResponseWriter, r *http.Request) {
	common.CommonQuery[model.Ops_dict](w, r, common.GetDaprClient(), "o_ops_dict", "id")
}

// @Summary save
// @Description save
// @Tags 字典
// @Accept       json
// @Param item body model.Ops_dict true "object"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ops_dict} "object"
// @Failure 500 {object} common.Response ""
// @Router /ops-dict [post]
func UpsertOps_dictHandler(w http.ResponseWriter, r *http.Request) {
	var val model.Ops_dict
	err := common.ReadRequestBody(r, &val)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}

	beforeHook, exists := common.GetUpsertBeforeHook("Ops_dict")
	if exists {
		v, err1 := beforeHook(r, val)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
		val = v.(model.Ops_dict)
	}
	if val.ID == "" {
		val.ID = common.NanoId()
	}

	if time.Time(val.CreatedTime).IsZero() {
		val.CreatedTime = common.LocalTime(time.Now())
	}

	if time.Time(val.UpdatedTime).IsZero() {
		val.UpdatedTime = common.LocalTime(time.Now())
	}

	err = common.DbUpsert[model.Ops_dict](r.Context(), common.GetDaprClient(), val, model.Ops_dictTableInfo.Name, "id")
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}
	common.HttpSuccess(w, common.OK.WithData(val))
}

// @Summary delete
// @Description delete
// @Tags 字典
// @Param id  path string true "实例id"
// @Produce  json
// @Success 200 {object} common.Response{data=model.Ops_dict} "object"
// @Failure 500 {object} common.Response ""
// @Router /ops-dict/{id} [delete]
func DeleteOps_dictHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	beforeHook, exists := common.GetDeleteBeforeHook("Ops_dict")
	if exists {
		_, err1 := beforeHook(r, id)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	common.CommonDelete(w, r, common.GetDaprClient(), "o_ops_dict", "id", "id")
}

// @Summary batch delete
// @Description batch delete
// @Tags 字典
// @Accept  json
// @Param ids body []string true "id array"
// @Produce  json
// @Success 200 {object} common.Response ""
// @Failure 500 {object} common.Response ""
// @Router /ops-dict/batch-delete [post]
func batchDeleteOps_dictHandler(w http.ResponseWriter, r *http.Request) {

	var ids []string
	err := common.ReadRequestBody(r, &ids)
	if err != nil {
		common.HttpResult(w, common.ErrParam.AppendMsg(err.Error()))
		return
	}
	if len(ids) == 0 {
		common.HttpResult(w, common.ErrParam.AppendMsg("len of ids is 0"))
		return
	}
	beforeHook, exists := common.GetBatchDeleteBeforeHook("Ops_dict")
	if exists {
		_, err1 := beforeHook(r, ids)
		if err1 != nil {
			common.HttpResult(w, common.ErrService.AppendMsg(err1.Error()))
			return
		}
	}
	idstr := strings.Join(ids, ",")
	err = common.DbDeleteByOps(r.Context(), common.GetDaprClient(), "o_ops_dict", []string{"id"}, []string{"in"}, []any{idstr})
	if err != nil {
		common.HttpResult(w, common.ErrService.AppendMsg(err.Error()))
		return
	}

	common.HttpResult(w, common.OK)
}
