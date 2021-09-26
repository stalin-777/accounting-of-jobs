package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/user"
	"strings"

	log "github.com/sirupsen/logrus"
	aoj "github.com/stalin-777/accounting-of-jobs"
	"github.com/stalin-777/accounting-of-jobs/postgres"
)

// Service - Service
type WorkplaceService struct {
	Request  *aoj.Workplace
	Response *ServiceResponse
	Method   string
	URL      string
}

type ServiceRequest struct {
	Data interface{}
}

type ServiceResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Data    interface{}
}

func (s *WorkplaceService) Workplace(id int) (*aoj.Workplace, error) {

	s.URL = fmt.Sprintf("%s/%v", url, id)
	s.Response.Data = &aoj.Workplace{}

	err := s.sendRequest()
	if err != nil {
		log.Fatal(err)
	}

	resp := (s.Response.Data).(*aoj.Workplace)

	return resp, nil
}

func (s *WorkplaceService) Workplaces() ([]*aoj.Workplace, error) {

	s.Response.Data = &[]aoj.Workplace{}

	err := s.sendRequest()
	if err != nil {
		log.Fatal(err)
	}

	return nil, err
}

func (s *WorkplaceService) CreateWorkplace(w *aoj.Workplace) error {

	setWorkplace(w)

	err := s.sendRequest()
	if err != nil {
		if err == postgres.ErrConstraintPgx {

			fmt.Println(err.Error())
			err = nil

			for {

				b := make([]byte, 1)
				fmt.Scanln(&b)

				switch strings.ToUpper(string(b)) {
				case "Y":

					s.Method = "PUT"
					s.Response = &ServiceResponse{}

					return s.UpdateWorkplace(w)
				case "N":
					return nil
				}
			}
		}

		log.Fatal(err)
	}
	resp := (s.Response.Data).(map[string]interface{})
	id := resp["id"].(float64)

	fmt.Printf("Запись успешно создана. ID: %v\n", id)

	return nil
}
func (s *WorkplaceService) UpdateWorkplace(w *aoj.Workplace) error {

	s.URL = fmt.Sprintf("%s/%v", url, w.ID)
	setWorkplace(w)
	fmt.Println(w)
	err := s.sendRequest()
	if err != nil {
		log.Fatal(err)
	}
	resp := (s.Response.Data).(map[string]interface{})
	id := resp["id"].(float64)

	fmt.Printf("Запись успешно обновлена. ID: %v\n", id)

	return nil
}
func (s *WorkplaceService) DeleteWorkplace(id int) error {

	s.URL = fmt.Sprintf("%s/%v", url, id)

	err := s.sendRequest()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Запись с ID=%v успешно удалена\n", id)

	return nil
}

func (s *WorkplaceService) sendRequest() (err error) {

	var content io.Reader

	if s.Request != nil {
		byteContent, err := json.Marshal(s.Request)
		if err != nil {
			return err
		}

		content = bytes.NewBuffer(byteContent)
	}

	client := &http.Client{}
	req, err := http.NewRequest(s.Method, s.URL, content)
	if err != nil {
		return err
	}

	if s.Method == "POST" || s.Method == "PUT" {
		req.Header.Add("Content-Type", "application/json")
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &s.Response); err != nil {
		return err
	}

	if s.Response.Error != "" {
		return errors.New(s.Response.Error)
	}

	return
}

func setWorkplace(w *aoj.Workplace) {

	if w.Username == "" && w.IP == nil && w.Hostname == "" {
		setDefaultWorkplace(w)
	} else {
		if w.Username == "" {
			log.Fatalln("Укажите значение флага username")
		}
		if w.Hostname == "" {
			log.Fatalln("Укажите значение флага hostname")
		}
		if w.IP == nil {
			log.Fatalln("Укажите значение флага ip")
		}
	}
}

func setDefaultWorkplace(w *aoj.Workplace) {

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	ip := getOutboundIP()

	w.Username = user.Username
	w.Hostname = hostName
	w.IP = ip
}

func getOutboundIP() net.IP {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP

	// addrs, err := net.InterfaceAddrs()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, address := range addrs {
	// 	// check the address type and if it is not a loopback the display it
	// 	if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	// 		if ipnet.IP.To4() != nil {
	// 			return ipnet.IP
	// 		}
	// 	}
	// }
	// return nil
}
