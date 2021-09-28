package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/pflag"
	aoj "github.com/stalin-777/accounting-of-jobs"
)

const url = "http://localhost:8088/workplaces"

func main() {

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, filename := path.Split(f.File)
			filename = fmt.Sprintf("%s:%d", filename, f.Line)
			return "", filename
		},
	})

	workplace := &aoj.Workplace{}
	service := &WorkplaceService{
		Request:  workplace,
		Response: &ServiceResponse{},
		URL:      url,
	}

	pflag.StringVarP(&service.Method, "method", "m", "POST", "Method flag. Possible values POST, PUT, DELETE, GET")
	pflag.StringVarP(&workplace.Username, "user", "u", "", "Username flag")
	pflag.StringVarP(&workplace.Hostname, "host", "h", "", "Hostname flag")
	pflag.IPVar(&workplace.IP, "ip", nil, "IP flag")
	pflag.IntVar(&workplace.ID, "id", 0, "ID flag")
	pflag.Parse()

	service.Method = strings.ToUpper(service.Method)

	var handler aoj.WorkplaceService = service

	switch service.Method {
	case "GET":

		data := [][]string{}

		if workplace.ID == 0 {

			handler.FindWorkplaces()

			resp := (service.Response.Data).(*[]aoj.Workplace)
			for _, wp := range *resp {
				data = append(data, []string{
					strconv.Itoa(wp.ID),
					wp.Username,
					wp.Hostname,
					wp.IP.String(),
				})
			}
		} else {

			service.Response.Data, _ = handler.FindWorkplace(workplace.ID)

			resp := (service.Response.Data).(*aoj.Workplace)

			data = append(data, []string{
				strconv.Itoa(resp.ID),
				resp.Username,
				resp.Hostname,
				resp.IP.String(),
			})
		}

		renderTable(data)
	case "POST":
		service.CreateWorkplace(workplace)
	case "PUT":
		service.UpdateWorkplace(workplace)
	case "DELETE":
		service.DeleteWorkplace(workplace.ID)
	default:
		fmt.Println("Указан некорректный метод. Доступные методы - POST, PUT, DELETE, GET")
	}

	fmt.Printf("нажмите ENTER чтобы выйти ...")
	b := make([]byte, 10)
	fmt.Scanln(&b)
}

func renderTable(data [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Username", "Hostname", "IP"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
