package v1alpha2

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/emicklei/go-restful"
	"k8s.io/klog"

	"kubesphere.io/kubesphere/pkg/api"
)

const ScopeQueryUrl = "http://%s/api/topology/services"

type handler struct {
	weaveScopeHost string
}

func (h *handler) getScopeUrl() string {
	return fmt.Sprintf(ScopeQueryUrl, h.weaveScopeHost)
}

func (h *handler) getNamespaceTopology(request *restful.Request, response *restful.Response) {
	var query = url.Values{
		"namespace": []string{request.PathParameter("namespace")},
		"timestamp": request.QueryParameters("timestamp"),
	}
	var u = fmt.Sprintf("%s?%s", h.getScopeUrl(), query.Encode())

	resp, err := http.Get(u)

	if err != nil {
		klog.Errorf("query scope faile with err %v", err)
		api.HandleInternalError(response, nil, err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		klog.Errorf("read response error : %v", err)
		api.HandleInternalError(response, nil, err)
		return
	}

	// need to set header for proper response
	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(body)

	if err != nil {
		klog.Errorf("write response failed %v", err)
	}
}

func (h *handler) getNamespaceNodeTopology(request *restful.Request, response *restful.Response) {
	var query = url.Values{
		"namespace": []string{request.PathParameter("namespace")},
		"timestamp": request.QueryParameters("timestamp"),
	}
	var u = fmt.Sprintf("%s/%s?%s", h.getScopeUrl(), request.PathParameter("node_id"), query.Encode())

	resp, err := http.Get(u)

	if err != nil {
		klog.Errorf("query scope faile with err %v", err)
		api.HandleInternalError(response, nil, err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		klog.Errorf("read response error : %v", err)
		api.HandleInternalError(response, nil, err)
		return
	}

	// need to set header for proper response
	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(body)

	if err != nil {
		klog.Errorf("write response failed %v", err)
	}
}
