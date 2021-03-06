package main

import (
	"flag"
	"os"

	vmioperator "github.com/kubevirt/vm-import-operator/pkg/operator/resources/operator"
	"github.com/spf13/pflag"
	"kubevirt.io/containerized-data-importer/tools/util"
)

var (
	csvVersion         = flag.String("csv-version", "", "The csv version")
	replacesCsvVersion = flag.String("replaces-csv-version", "", "The csv version this replaces")
	namespace          = flag.String("namespace", "", "Namespace used by csv")
	pullPolicy         = flag.String("pull-policy", "Always", "The pull policy to use on the operator image")
	operatorVersion    = flag.String("operator-version", "", "The operator version")
	operatorImage      = flag.String("operator-image", "", "The operator image name")
	controllerImage    = flag.String("controller-image", "", "The controller image name")
	virtV2vImage       = flag.String("virtv2v-image", "", "The virtv2v image name")
	dumpCRDs           = flag.Bool("dump-crds", false, "optional - dumps crd manifests to stdout")
	controllerOnly     = flag.Bool("controller-only", false, "optional - dumps manifests to stdout only for controller local deployment")
)

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true
	pflag.Parse()

	data := vmioperator.ClusterServiceVersionData{
		CsvVersion:         *csvVersion,
		ReplacesCsvVersion: *replacesCsvVersion,
		Namespace:          *namespace,
		ImagePullPolicy:    *pullPolicy,
		OperatorVersion:    *operatorVersion,
		OperatorImage:      *operatorImage,
		ControllerImage:    *controllerImage,
		VirtV2vImage:       *virtV2vImage,
	}

	// Generate resources required for deploying the controller only
	if *controllerOnly {
		err := util.MarshallObject(vmioperator.CreateResourceMapping(), os.Stdout)
		if err != nil {
			panic(err)
		}
		err = util.MarshallObject(vmioperator.CreateVMImport(), os.Stdout)
		if err != nil {
			panic(err)
		}
		err = util.MarshallObject(vmioperator.CreateServiceAccount(*namespace), os.Stdout)
		if err != nil {
			panic(err)
		}
		err = util.MarshallObject(vmioperator.CreateControllerRole(), os.Stdout)
		if err != nil {
			panic(err)
		}
		err = util.MarshallObject(vmioperator.CreateControllerRoleBinding(*namespace), os.Stdout)
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		return
	}

	csv, err := vmioperator.NewClusterServiceVersion(&data)
	if err != nil {
		panic(err)
	}
	err = util.MarshallObject(csv, os.Stdout)
	if err != nil {
		panic(err)
	}

	if *dumpCRDs {
		err = util.MarshallObject(vmioperator.CreateVMImportConfig(), os.Stdout)
		if err != nil {
			panic(err)
		}
	}
}
