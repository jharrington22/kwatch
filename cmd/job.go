// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	jobName string
)

// JobCmd represents the job sub command
var JobCmd = &cobra.Command{
	Use:   "job",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Init(opt)
		if err != nil {
			log.Fatal(err)
		}

		batchAPI := k8sclient.BatchV1()

		listOptions := metav1.ListOptions{
			LabelSelector: Label,
			FieldSelector: Field,
		}

		if jobName != "" {
			fmt.Printf("Job name: %s\n", jobName)
		}

		jobs, err := batchAPI.Jobs(Namespace).List(listOptions)
		if err != nil {
			log.Fatal(err)
		}
		if jobName == "" {
			printJobs(jobs)
		} else {
			//fmt.Println("Looking for specific job")
			for _, ajob := range jobs.Items {
				fmt.Println(ajob.Name)
				if ajob.Name == jobName {
					//fmt.Println("Found specific job")
					printJob(ajob)
					//fmt.Printf("%+v\n", ajob)
					//fmt.Println("")
					//fmt.Printf("%+v\n", ajob.ObjectMeta)
					//fmt.Println("")
					//fmt.Printf("%+v\n", ajob.ObjectMeta.ResourceVersion)
					//fmt.Println("")
					//fmt.Printf("%+v\n", ajob.Status.Active)
					//var objectMeta string
					//fmt.Println(string(ajob.ObjectMeta))
					// for _, o := range ajob.ObjectMeta {
					// 	fmt.Printf("%s\n", o)
					// }
				}
			}
		}
	},
}

func init() {
	JobCmd.Flags().StringVar(&jobName, "name", "", "Job name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printJob(job v1.Job) {
	template := "%-32s%-8s\n"
	status := strconv.Itoa(int(job.Status.Active))
	fmt.Printf(template, "NAME", "STATUS")
	fmt.Printf(template, job.Name, status)
}

func printJobs(jobs *v1.JobList) {
	if len(jobs.Items) == 0 {
		log.Println("No jobs found")
		return
	}
	template := "%-32s%-8s\n"
	fmt.Printf(template, "NAME", "STATUS")
	for _, job := range jobs.Items {
		status := strconv.Itoa(int(job.Status.Active))
		fmt.Printf(template, job.Name, status)
	}
}
