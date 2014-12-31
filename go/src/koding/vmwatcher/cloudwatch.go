package main

import (
	"fmt"
	"koding/db/models"
	"koding/db/mongodb/modelhelper"
	"time"

	"log"

	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/cloudwatch"
	"github.com/jinzhu/now"
	"labix.org/v2/mgo/bson"
)

var (
	AWS_NAMESPACE = "AWS/EC2"
	AWS_PERIOD    = 604800

	GB_TO_MB float64 = 1000

	startingToday = now.BeginningOfDay()
	sevenDaysAgo  = startingToday.Add(-7 * 24 * time.Hour)

	auth aws.Auth

	PaidPlanMultiplier float64 = 2
)

type Cloudwatch struct {
	Name  string
	Limit float64
}

func (c *Cloudwatch) GetName() string {
	return c.Name
}

func (c *Cloudwatch) GetLimit() float64 {
	return c.Limit
}

func (c *Cloudwatch) GetAndSaveData(username string) error {
	userMachines, err := modelhelper.GetMachinesForUsername(username)
	if err != nil {
		return err
	}

	var sum float64

	for _, machine := range userMachines {
		var meta = machine.Meta.(bson.M)
		var regionStr = meta["region"].(string)
		var instanceId = meta["instanceId"].(string)

		dimension := &cloudwatch.Dimension{
			Name:  "InstanceId",
			Value: instanceId,
		}

		cw, err := cloudwatch.NewCloudWatch(auth, aws.Regions[regionStr].CloudWatchServicepoint)
		if err != nil {
			return err
		}

		request := &cloudwatch.GetMetricStatisticsRequest{
			Dimensions: []cloudwatch.Dimension{*dimension},
			Statistics: []string{cloudwatch.StatisticDatapointSum},
			MetricName: c.GetName(),
			EndTime:    startingToday,
			StartTime:  sevenDaysAgo,
			Period:     AWS_PERIOD,
			Namespace:  AWS_NAMESPACE,
		}

		response, err := cw.GetMetricStatistics(request)
		if err != nil {
			return err
		}

		for _, raw := range response.GetMetricStatisticsResult.Datapoints {
			sum += raw.Sum / 1024 / 1024
		}
	}

	return storage.Save(c.Name, username, sum)
}

func (c *Cloudwatch) GetMachinesOverLimit() ([]*models.Machine, error) {
	// TODO: remove hardcoded NetworkOutLimt here
	usernames, err := storage.Range(c.Name, NetworkOutLimt)
	if err != nil {
		return nil, err
	}

	machines := []*models.Machine{}

	for _, username := range usernames {
		yes, err := exemptFromStopping(c.GetName(), username)
		if err != nil {
			return nil, err
		}

		if !yes {
			ms, err := modelhelper.GetMachinesForUsername(username)
			if err != nil {
				log.Println(err)
				continue
			}

			machines = append(machines, ms...)
		}
	}

	return machines, nil
}

func (c *Cloudwatch) IsUserOverLimit(username string) (*LimitResponse, error) {
	canStart := &LimitResponse{CanStart: true}

	value, err := storage.Get(c.GetName(), username)
	if err != nil && !isRedisRecordNil(err) {
		log.Println(err)
		return canStart, nil
	}

	yes, err := exemptFromStopping(c.GetName(), username)
	if err != nil {
		log.Println(err)
		return canStart, nil
	}

	if yes {
		return canStart, nil
	}

	planTitle, err := getPlanForUser(username)
	if err != nil {
		log.Println(err)
		return canStart, nil
	}

	var limit float64

	switch planTitle {
	case FreePlan:
		limit = c.GetLimit()
	default:
		limit = c.GetLimit() * PaidPlanMultiplier
	}

	lr := &LimitResponse{
		CanStart:     limit >= value,
		AllowedUsage: limit * GB_TO_MB,
		CurrentUsage: value * GB_TO_MB,
		Reason:       fmt.Sprintf("%s overlimit", c.GetName()),
	}

	return lr, nil
}

func (c *Cloudwatch) RemoveUsername(username string) error {
	return storage.Remove(c.GetName(), username)
}

func isRedisRecordNil(err error) bool {
	return err.Error() == "redigo: nil returned"
}
