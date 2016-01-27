// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"microservice"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Person create(Person person)")
	fmt.Fprintln(os.Stderr, "  Person read(i32 id)")
	fmt.Fprintln(os.Stderr, "  Person update(Person person)")
	fmt.Fprintln(os.Stderr, "  void destroy(i32 id)")
	fmt.Fprintln(os.Stderr, "   getAll()")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := microservice.NewPersonServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "create":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Create requires 1 args")
			flag.Usage()
		}
		arg13 := flag.Arg(1)
		mbTrans14 := thrift.NewTMemoryBufferLen(len(arg13))
		defer mbTrans14.Close()
		_, err15 := mbTrans14.WriteString(arg13)
		if err15 != nil {
			Usage()
			return
		}
		factory16 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt17 := factory16.GetProtocol(mbTrans14)
		argvalue0 := microservice.NewPerson()
		err18 := argvalue0.Read(jsProt17)
		if err18 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Create(value0))
		fmt.Print("\n")
		break
	case "read":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Read requires 1 args")
			flag.Usage()
		}
		tmp0, err19 := (strconv.Atoi(flag.Arg(1)))
		if err19 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.Read(value0))
		fmt.Print("\n")
		break
	case "update":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Update requires 1 args")
			flag.Usage()
		}
		arg20 := flag.Arg(1)
		mbTrans21 := thrift.NewTMemoryBufferLen(len(arg20))
		defer mbTrans21.Close()
		_, err22 := mbTrans21.WriteString(arg20)
		if err22 != nil {
			Usage()
			return
		}
		factory23 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt24 := factory23.GetProtocol(mbTrans21)
		argvalue0 := microservice.NewPerson()
		err25 := argvalue0.Read(jsProt24)
		if err25 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Update(value0))
		fmt.Print("\n")
		break
	case "destroy":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Destroy requires 1 args")
			flag.Usage()
		}
		tmp0, err26 := (strconv.Atoi(flag.Arg(1)))
		if err26 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.Destroy(value0))
		fmt.Print("\n")
		break
	case "getAll":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetAll requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetAll())
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
