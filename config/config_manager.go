package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rfashwal/scs-utilities/discovery"
	"github.com/rfashwal/scs-utilities/net"

	"github.com/labstack/gommon/log"
)

type Manager struct {
	httpPort                 int
	address                  string
	serviceName              string
	eurekaService            string
	hostName                 string
	mqHost                   string
	mqCredentials            string
	registrationTicket       *discovery.RegistrationTicket
	serviceDurationInSeconds int
	ignoreLoopback           bool
	temperatureTopic         string
	actuatorTopic            string
	readingKey               string
}

func (c *Manager) Init() {

	if port, err := os.LookupEnv("HTTP_PORT"); !err {
		panic("Set HTTP_PORT and try again")
	} else {
		if port, e := strconv.Atoi(port); e != nil {
			panic("please set HTTP_PORT variable and try again")
		} else {
			c.httpPort = port
		}
	}

	if serviceName, err := os.LookupEnv("SERVICE_NAME"); !err {
		panic("please set SERVICE_NAME variable and try again")
	} else {
		c.serviceName = serviceName
	}

	if ignoreLoopback, err := os.LookupEnv("IGNORE_LOOPBACK"); !err {
		// set default behaviour in case variable not set
		c.ignoreLoopback = true
	} else {
		if result, err := strconv.ParseBool(ignoreLoopback); err != nil {
			log.Warn(fmt.Sprintf("could not convert %s to bool", ignoreLoopback))
			c.ignoreLoopback = true
		} else {
			c.ignoreLoopback = result
		}
	}

	address, err := net.GetIP(c.ignoreLoopback)
	if err != nil {
		panic(fmt.Sprintf("Could not resolve ip address: %s", err))
	}
	c.address = address

	hostname, err := os.Hostname()
	if err != nil {
		panic(fmt.Sprintf("Could not obtain hostname: %s", err))
	}
	c.hostName = hostname

	if serviceDuration, err := os.LookupEnv("SERVICE_DURATION_IN_SECONDS"); !err {
		c.serviceDurationInSeconds = 5
	} else {
		if port, e := strconv.Atoi(serviceDuration); e != nil {
			c.serviceDurationInSeconds = port
		}
	}

	if eurekaHost, err := os.LookupEnv("EUREKA_SERVICE"); !err {
		panic(fmt.Sprintf("set EUREKA_SERVICE variable and try again"))
	} else {
		c.eurekaService = eurekaHost
	}

	if mqHost, err := os.LookupEnv("MQ_HOST"); !err {
		panic(fmt.Sprintf("set MQ_HOST variable and try again"))
	} else {
		c.mqHost = mqHost
	}

	if creds, err := os.LookupEnv("MQ_CREDENTIALS"); !err {
		c.mqCredentials = "guest1:guest1"
	} else {
		c.mqCredentials = creds
	}

	if tempTopic, err := os.LookupEnv("TEMPERATURE_TOPIC"); !err {
		c.temperatureTopic = "temperature"
	} else {
		c.temperatureTopic = tempTopic
	}

	if act, err := os.LookupEnv("ACTUATOR_TOPIC"); !err {
		c.actuatorTopic = "actuator"
	} else {
		c.actuatorTopic = act
	}

	if readingKey, err := os.LookupEnv("READINGS_ROUTING_KEY"); !err {
		c.readingKey = "readings"
	} else {
		c.readingKey = readingKey
	}
	registrationTicket := discovery.BuildRegistrationTicket(c.serviceName, c.httpPort, c.serviceDurationInSeconds, c.ignoreLoopback)
	c.registrationTicket = registrationTicket
}

func (c *Manager) HttpPort() int {
	return c.httpPort
}

func (c *Manager) Address() string {
	return fmt.Sprintf("%s:%d", c.address, c.httpPort)
}

func (c *Manager) ServiceName() string {
	return c.serviceName
}

func (c *Manager) EurekaService() string {
	return c.eurekaService
}

func (c *Manager) Hostname() string {
	return c.hostName
}

func (c *Manager) MQHost() string {
	return c.mqHost
}

func (c *Manager) TemperatureTopic() string {
	return c.temperatureTopic
}

func (c *Manager) ActuatorTopic() string {
	return c.actuatorTopic
}

func (c *Manager) ReadingsRoutingKey() string {
	return c.readingKey
}

func (c *Manager) RegistrationTicket() *discovery.RegistrationTicket {
	return c.registrationTicket
}

func (c *Manager) IgnoreLoopback() bool {
	return c.ignoreLoopback
}

func (c *Manager) RabbitURL() string {
	return fmt.Sprintf("amqp://%s@%s", c.mqCredentials, c.mqHost)
}
