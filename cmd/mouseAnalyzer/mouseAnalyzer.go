// mouseAnalyzer determines vesicle release events and latencies for our
// mouse NMJ 6 AZ model with two synaptic vesicles each according to the
// the original excess-calcium-binding-site model
// (Dittrich et al., Biophys. J, 2013, 104:2751-2763)
package main

import (
	"flag"
	"fmt"

	rel "github.com/haskelladdict/mbdr/releaser"
	"github.com/haskelladdict/mbdr/version"
)

// analyser info
var info = rel.AnalyzerInfo{
	Name: "mouseAnalyzer",
}

// simulation model
var model = rel.SimModel{
	VesicleIDs: []string{"1_1", "1_2", "2_1", "2_2", "3_1", "3_2", "4_1", "4_2",
		"5_1", "5_2", "6_1", "6_2"},
	SensorTemplate: "bound_vesicle_%s_%s_%d.%04d.dat",
	PulseDuration:  3e-3,
	NumPulses:      1,
}

// fusion model
var fusionModel = rel.FusionModel{
	NumSyt:       8,
	NumActiveSyt: 2,
	EnergyModel:  false,
}

// initialize simulation and fusion model parameters coming from commandline
func init() {

	flag.IntVar(&fusionModel.NumActiveSites, "n", 0, "number of sites required for activation "+
		"of deterministic model")
	flag.IntVar(&info.NumThreads, "T", 1, "number of threads. Each thread works on a "+
		"single binary output file\n\tso memory requirements multiply")

	// define synaptogamin and Y sites
	model.CaSensors = make([]rel.CaSensor, fusionModel.NumSyt+fusionModel.NumY)
	model.CaSensors[0] = rel.CaSensor{Sites: []int{8, 9, 29, 30, 31}, SiteType: rel.SytSite}
	model.CaSensors[1] = rel.CaSensor{Sites: []int{7, 32, 33, 34, 35}, SiteType: rel.SytSite}
	model.CaSensors[2] = rel.CaSensor{Sites: []int{3, 6, 36, 37, 38}, SiteType: rel.SytSite}
	model.CaSensors[3] = rel.CaSensor{Sites: []int{17, 39, 40, 41, 42}, SiteType: rel.SytSite}
	model.CaSensors[4] = rel.CaSensor{Sites: []int{15, 16, 43, 44, 45}, SiteType: rel.SytSite}
	model.CaSensors[5] = rel.CaSensor{Sites: []int{14, 46, 47, 48, 49}, SiteType: rel.SytSite}
	model.CaSensors[6] = rel.CaSensor{Sites: []int{4, 12, 24, 50, 51}, SiteType: rel.SytSite}
	model.CaSensors[7] = rel.CaSensor{Sites: []int{10, 25, 26, 27, 28}, SiteType: rel.SytSite}
}

// usage prints a brief usage information to stdout
func usage() {
	fmt.Printf("%s v%s  (C) %s Markus Dittrich\n\n", info.Name, version.Tag, version.Year)
	fmt.Printf("usage: %s [options] <binary mcell files>\n", info.Name)
	fmt.Println("\noptions:")
	flag.PrintDefaults()
}

// main entry point
func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		usage()
		return
	}

	rel.Run(&model, &fusionModel, &info, flag.Args())
}
