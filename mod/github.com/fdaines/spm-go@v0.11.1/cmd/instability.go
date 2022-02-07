package cmd

import (
	"github.com/fdaines/spm-go/cmd/impl"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
	pkg "github.com/fdaines/spm-go/utils/packages"
	"github.com/spf13/cobra"
)

var (
	instabilityCmd = &cobra.Command{
		Use:   "instability",
		Short: "Analyzes instability of packages",
		Args:  ValidateArgs,
		Run:   analyzeInstability,
	}
)

func init() {
	rootCmd.AddCommand(instabilityCmd)
}

func analyzeInstability(cmd *cobra.Command, args []string) {
	utils.ExecuteWithTimer(func() {
		utils.PrintMessage("Instability analysis started.")
		mainPackage, err := pkg.GetMainPackage()
		if err != nil {
			utils.PrintError("Error loading main package", err)
			return
		}
		var afferentMap = make(map[string][]string)
		pkgsInfo := pkg.GetBasicPackagesInfo()
		utils.PrintMessage("Gathering package metrics, please wait until the command is finished running.")
		for index, pkgInfo := range pkgsInfo {
			utils.PrintStep()
			impl.FillDependencies(pkgsInfo[index], pkgsInfo)
			for _, current := range pkgInfo.Dependencies.Internals {
				afferentMap[current] = append(afferentMap[pkgInfo.Path], current)
			}
		}
		for index, pkgInfo := range pkgsInfo {
			utils.PrintStep()
			pkgsInfo[index].Dependants = afferentMap[pkgInfo.Path]
			pkgsInfo[index].AfferentCoupling = len(pkgsInfo[index].Dependants)
			pkgsInfo[index].EfferentCoupling = pkgsInfo[index].Dependencies.InternalsCount
			pkgsInfo[index].Instability = calculateInstability(pkgsInfo[index])
		}
		utils.PrintVerboseMessage("Done.")
		printInstability(pkgsInfo)
		output.GenerateHtmlOutput(pkgsInfo, mainPackage, "instability")
	})
}

func calculateInstability(pksInfo *model.PackageInfo) float64 {
	if pksInfo.EfferentCoupling == 0 && pksInfo.AfferentCoupling == 0 {
		return 1
	}
	return utils.RoundValue(
		float64(pksInfo.EfferentCoupling) / float64(pksInfo.EfferentCoupling+pksInfo.AfferentCoupling))
}
