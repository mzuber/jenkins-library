package toolrecord

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

type keydataset struct {
	Name        string // technical name from the tool
	Value       string // technical value
	DisplayName string // user friendly name - optional
	Url         string // direct URL to navigate to this key in the tool backend - optional
}

type toolrecord struct {
	RecordVersion int

	ToolName     string
	ToolInstance string

	// tool agnostic convenience aggregations
	// picks the most specific URL + concatenate the dimension names
	// for easy dashboard / xls creation
	DisplayName string
	DisplayUrl  string

	// detailed keydata - needs tool-specific parsing
	Keys []keydataset

	// place for additional context information
	Context map[string]interface{}

	// internal - not exported to the json
	workspace      string
	reportFileName string
}

func New(workspace, toolName, toolInstance string) *toolrecord {
	tr := toolrecord{}

	tr.RecordVersion = 1
	tr.ToolName = toolName
	tr.ToolInstance = toolInstance
	tr.Keys = []keydataset{}
	tr.Context = make(map[string]interface{})

	tr.workspace = workspace

	now := time.Now().UTC()
	reportFileName := filepath.Join(workspace,
		"toolrun_"+toolName+"_"+
			now.Format("20210731")+
			strings.ReplaceAll(now.Format("15:04:05"), ":", "")+
			".json")
	tr.reportFileName = reportFileName

	return &tr
}

func (tr *toolrecord) AddKeyData(keyname, keyvalue, displayname, url string) error {
	if keyname == "" {
		return errors.New("TR_ADD_KEY: empty keyname")
	}
	if keyvalue == "" {
		return fmt.Errorf("TR_ADD_KEY: empty keyvalue for %v", keyname)
	}
	keydata := keydataset{Name: keyname, Value: keyvalue, DisplayName: displayname, Url: url}
	tr.Keys = append(tr.Keys, keydata)
	return nil
}

func (tr *toolrecord) AddContext(label string, data interface{}) error {
	if label == "" {
		return errors.New("TR_ADD_CONTEXT: no label supplied")
	}
	tr.Context[label] = data
	return nil
}

func (tr *toolrecord) GetFileName() string {
	return tr.reportFileName
}

func (tr *toolrecord) Persist() error {
	if tr.workspace == "" {
		return errors.New("TR_PERSIST: empty workspace ")
	}
	if tr.ToolName == "" {
		return errors.New("TR_PERSIST: empty toolName")
	}
	if tr.ToolInstance == "" {
		return errors.New("TR_PERSIST: empty instanceName")
	}
	// convenience aggregation
	displayName := ""
	displayUrl := ""
	for _, keyset := range tr.Keys {
		// create "name1 - name2 - name3"
		subDisplayName := keyset.DisplayName
		if subDisplayName != "" {
			if displayName != "" {
				displayName = displayName + " - "
			}
			displayName = displayName + subDisplayName
		}
		subUrl := keyset.Url
		if subUrl != "" {
			displayUrl = subUrl
		}
	}
	tr.DisplayName = displayName
	tr.DisplayUrl = displayUrl

	file, _ := json.Marshal(tr)
	err := ioutil.WriteFile(tr.GetFileName(), file, 0644)
	if err != nil {
		return fmt.Errorf("TR_PERSIST: %v", err)
	}
	return nil
}