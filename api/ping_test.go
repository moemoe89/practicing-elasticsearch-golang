//
//  Practicing Elasticsearch
//
//  Copyright © 2016. All rights reserved.
//

package api_test

import (
	"github.com/moemoe89/go-elasticsearch-tita/config"
	"github.com/moemoe89/go-elasticsearch-tita/routers"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	log := config.InitLog()

	router := routers.GetRouter(log, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
