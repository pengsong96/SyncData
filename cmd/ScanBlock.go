package cmd

import (
	"SyncEthData/config"
	"SyncEthData/syncData"
	"SyncEthData/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"math/big"
)

func ScanCmd() *cobra.Command {
	var blockNum int
	scanCmd := &cobra.Command{
		Use:   "scan",
		Short: "s",
		Long:  "It will sync the latest block ",
		RunE: func(cmd *cobra.Command, args []string) error {
			//syncData.SaveAllData(blockNum
			scanBlock(blockNum)
			return nil
		},
	}
	scanCmd.Flags().IntVarP(&blockNum, "blockNum", "b", 0, "input blockNum")
	return scanCmd
}

func scanBlock(blockNum int) {
	height := syncData.GetBlockHeight(config.CLIENT[0])
	distance := (height-blockNum)/len(config.CLIENT)

	for i:=0;i<len(config.CLIENT)-1;i++ {
		go getBlock(config.CLIENT[i],i,distance)
	}

	go scanNewBlock(config.CLIENT[len(config.CLIENT)-1],(len(config.CLIENT)-1)*distance)
}

func getBlock(client *ethclient.Client, i int, distance int) {
	from := distance*i
	end := from + distance
	for from<end {
		block := syncData.GetBlockByNum(client, big.NewInt(int64(from)))
		utils.TransformData(block)
		from += 1
	}
}
//同步最新区块
func scanNewBlock(client *ethclient.Client, from int)  {
	for true {
		block := syncData.GetBlockByNum(client, big.NewInt(int64(from)))
		utils.TransformData(block)
		from += 1
	}
}