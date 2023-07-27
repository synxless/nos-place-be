package placewatcher

import (
	"context"
	"fmt"
	"math/big"
	"nosplace/abi"
	"nosplace/models"
	"nosplace/storage"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlaceIndexer struct {
	DefaultRPC string
	Storage    *storage.Storage
}

func (indexer *PlaceIndexer) Start() {
	fmt.Println("Starting Place pixel Indexer")

	state, err := indexer.Storage.GetIndexerState(context.Background(), "place_indexer")
	if err != nil {
		if err == mongo.ErrNoDocuments {
			state = &models.IndexerState{
				Indexer:          "place_indexer",
				LastIndexedBlock: 3029112,
			}
			err = indexer.Storage.CreateIndexerState(context.Background(), state)
			if err != nil {
				fmt.Println("Failed to create indexer state: ", err)
				return
			}
		}
	}

	indexedBlock := state.LastIndexedBlock
	//TODO store indexed block in database
	for {

		evmClient, err := ethclient.Dial(indexer.DefaultRPC)
		if err != nil {
			fmt.Println("Failed to connect to the Ethereum network: ", err)
			continue
		}

		currentHeight, err := evmClient.BlockNumber(context.Background())
		if err != nil {
			fmt.Println("Failed to get current block height: ", err)
			continue
		}
		if indexedBlock >= currentHeight {
			time.Sleep(2 * time.Second)
			continue
		}
		fmt.Println("Current block height: ", currentHeight)
		fmt.Println("Indexed block height: ", indexedBlock)
		block, err := evmClient.BlockByNumber(context.Background(), big.NewInt(int64(indexedBlock)))
		if err != nil {
			fmt.Println("Failed to get block: ", err)
			indexedBlock++
			continue
		}
		blockTime := time.Unix(int64(block.Header().Time), 0)

		contractAddress := common.HexToAddress("0x3D751fCA34a5da4d04D2FA5932439aC7199ceaF8")

		placeContract, err := abi.NewPlace(contractAddress, evmClient)
		if err != nil {
			fmt.Println("Failed to get contract: ", err)
			continue
		}
		leftBound, err := placeContract.LeftBound(nil)
		if err != nil {
			fmt.Println("Failed to get left bound: ", err)
			continue
		}
		upperBound, err := placeContract.UpperBound(nil)
		if err != nil {
			fmt.Println("Failed to get upper bound: ", err)
			continue
		}
		canvasHeight, err := placeContract.GetHeight(nil)
		if err != nil {
			fmt.Println("Failed to get canvas height: ", err)
			continue
		}
		canvasWidth, err := placeContract.GetWidth(nil)
		if err != nil {
			fmt.Println("Failed to get canvas width: ", err)
			continue
		}

		fmt.Println("Block time: ", blockTime)
		extractor := func(from, to uint64) error {
			fmt.Println("Extracting from ", from, " to ", to)
			iter, err := placeContract.FilterPixelSet(
				&bind.FilterOpts{
					Start:   from,
					End:     &to,
					Context: context.Background()}, nil, nil, nil)
			if err != nil {
				fmt.Println("Failed to get filter: ", err)
				return err
			}
			for iter.Next() {
				fmt.Println("Pixel set")
				fmt.Println(iter.Event.X, iter.Event.Y, iter.Event.R, iter.Event.G, iter.Event.B)

				color := fmt.Sprintf("%v,%v,%v", iter.Event.R, iter.Event.G, iter.Event.B)
				newPixel := &models.Pixel{
					X:                iter.Event.X,
					Y:                iter.Event.Y,
					Color:            color,
					LastScannedBlock: iter.Event.Raw.BlockNumber,
					Tx:               iter.Event.Raw.TxHash.String(),
					Owner:            iter.Event.Setter.String(),
					Price:            iter.Event.Value.String(),
				}
				err = indexer.Storage.UpdatePixel(context.Background(), newPixel)
				if err != nil {
					fmt.Println("Failed to update pixel: ", err)
					return err
				}
			}
			return nil
		}
		err = scanEvent(indexedBlock, currentHeight, extractor)
		if err != nil {
			fmt.Println("Failed to scan event: ", err)
			continue
		}

		state.LastIndexedBlock = currentHeight
		state.LastIndexedAt = blockTime
		state.CanvasHeight = canvasHeight
		state.CanvasWidth = canvasWidth
		state.CanvasLeftBound = leftBound
		state.CanvasUpperBound = upperBound
		indexedBlock = state.LastIndexedBlock
		err = indexer.Storage.UpdateIndexerState(context.Background(), state)
		if err != nil {
			fmt.Println("Failed to update indexer state: ", err)
			continue
		}
	}
}

func scanEvent(from, to uint64, extractor func(from, to uint64) error) error {
	startBlock := from
	var endBlock uint64
	for {
		if endBlock >= to {
			break
		}
		if to-startBlock > 1000 {
			endBlock = startBlock + 1000
		} else {
			endBlock = to
		}
		err := extractor(startBlock, endBlock)
		if err != nil {
			return err
		}
		startBlock = endBlock + 1
	}
	return nil
}

// func checkContract() error {
// 	evmClient, err := ethclient.Dial("https://tc-node.trustless.computer/")
// 	if err != nil {
// 		fmt.Println("Failed to connect to the Ethereum network: ", err)
// 		return err
// 	}
// 	txhash := common.HexToHash("0x40a3630570865b47c4f911ea68ac9ed0e11e335eae68597df655e6ab9f99c111")
// 	receipt, err := evmClient.TransactionReceipt(context.Background(), txhash)
// 	if err != nil {
// 		fmt.Println("Failed to get transaction receipt: ", err)
// 		return err
// 	}
// 	tx, _, _ := evmClient.TransactionByHash(context.Background(), txhash)
// 	if receipt.Status == 0 {
// 		return err
// 	}
// 	for _, txlog := range receipt.Logs {
// 		if strings.EqualFold(txlog.Topics[0].String(), "0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b") {
// 			// tokenInfo, err := indexer.Storage.GetTokenInfo(txlog.Address.Hex())
// 			// if err != nil {
// 			// 	log.Println("GetTokenInfo", tx.To().Hex(), err)
// 			// 	return err
// 			// }

// 			nftContract, err := abi.NewERC20(*tx.To(), evmClient)
// 			if err != nil {
// 				log.Println("NewERC20", err)
// 				return err
// 			}
// 			event, err := nftContract.ParseUpgraded(*txlog)
// 			if err != nil {
// 				log.Println("ParseUpgraded", err)
// 				return err
// 			}
// 			fmt.Println(event.Implementation.String())
// 			// tokenInfo.CurrentImplentaion = event.Implementation.String()
// 			// err = indexer.Storage.UpdateTokenImplementation(context.Background(), tokenInfo)
// 			// if err != nil {
// 			// 	log.Println("UpdateTokenImplementation", err)
// 			// 	return err
// 			// }
// 		}
// 		if strings.EqualFold(txlog.Topics[0].String(), "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0") {
// 			nftContract, err := abi.NewERC20(txlog.Address, evmClient)
// 			if err != nil {
// 				log.Println("NewERC20", err)
// 				return err
// 			}
// 			isNewContract := true
// 			if isNewContract {
// 				var newToken models.Token
// 				decimal, err := nftContract.Decimals(nil)
// 				if err != nil {
// 					log.Println("Decimals", err)
// 					return err
// 				}

// 				if decimal == 0 {
// 					return err
// 				}

// 				symbol, err := nftContract.Symbol(nil)
// 				if err != nil {
// 					log.Println("Symbol", err)
// 					return err
// 				}
// 				newToken.Symbol = symbol
// 				name, err := nftContract.Name(nil)
// 				if err != nil {
// 					log.Println("Name", err)
// 					return err
// 				}
// 				newToken.Name = name
// 				owner, err := nftContract.Owner(nil)
// 				if err != nil {
// 					log.Println("Owner", err)
// 					return err
// 				}
// 				newToken.Owner = owner.Hex()
// 				totalSupply, err := nftContract.TotalSupply(nil)
// 				if err != nil {
// 					log.Println("TotalSupply", err)
// 					return err
// 				}
// 				newToken.TotalSupply = totalSupply.String()
// 				newToken.Decimal = decimal
// 				newToken.Address = txlog.Address.Hex()
// 				fmt.Println(newToken.Address, tx.To().Hex())
// 			} else {
// 				event, err := nftContract.ParseOwnershipTransferred(*txlog)
// 				if err != nil {
// 					log.Println("ParseTransfer", err)
// 					return err
// 				}
// 				fmt.Println(event.NewOwner.Hex())
// 			}

// 		}
// 	}
// 	return nil
// }
