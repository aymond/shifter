/*
Copyright 2019 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"mime/multipart"
	ops "shifter/ops"

	"github.com/gin-gonic/gin"
)

// HTTP Server Based on gin-gonic
type Server struct {
	router *gin.Engine
}

type Response_Status_Healthz struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
}

type Response_Status_Settings struct {
	Timestamp   string `json:"timestamp"`
	RunningPort string `json:"runningPort"`
	Version     int    `json:"version"`
	Status      int    `json:"status"`
	Message     string `json:"message"`
}

type Response_Convert_Yaml2Yaml struct {
	InputType      string                  `json:"inputType"`
	UUID           string                  `json:"uuid"`
	ConvertedFiles []ops.File              `json:"convertedFiles"`
	UploadedFiles  []*multipart.FileHeader `json:"uploadedFiles"`
	Message        string                  `json:"message"`
}
