package main

import (
	"context"
	"encoding/json"
	"flag"
	"math"
	"math/big"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/orakurudata/crystal-ball/configuration"
	"github.com/orakurudata/crystal-ball/contracts"
	"github.com/rs/zerolog/log"
)

const (
	BaseScore       = 100
	TimeBonus       = 50.0
	ExecutionWindow = 60
	BlockTime       = 3

	Decay = 1.0 / float64(ExecutionWindow-BlockTime)
)

func Score(delay uint64) uint64 {
	return uint64(
		math.Min(
			50.0,
			TimeBonus*(1.0-(Decay*float64(delay-BlockTime))))) + BaseScore
}

type Leaderboard []LeaderboardEntry

type LeaderboardEntry struct {
	Address      string  `json:"address"`
	Score        uint64  `json:"score"`
	ResponseTime float64 `json:"response_time"`
	Responses    uint64  `json:"responses"`
}

func main() {
	coreAddress := flag.String("core", "", "address of orakuru core")
	web3URL := flag.String("url", "", "web3 endpoint url")
	httpAddr := flag.String("http", "", "http bind address")
	log.Info().Caller().Msgf("starting leaderboard\n%v", configuration.Info())
	flag.Parse()

	if *coreAddress == "" || *web3URL == "" || *httpAddr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	client, err := ethclient.Dial(*web3URL)
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("cannot connect to web3")
	}
	core, err := contracts.NewIOrakuruCore(common.HexToAddress(*coreAddress), client)
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("cannot find core contract")
	}

	validators := NewValidators()

	regAddr, err := core.AddressRegistry(nil)
	if err != nil {
		log.Fatal().Err(err).Msg("could not get registry address")
	}
	registry, err := contracts.NewIAddressRegistry(regAddr, client)
	if err != nil {
		log.Fatal().Err(err).Msg("could not get instance of registry")
	}
	stakingAddr, err := registry.GetStakingAddr(nil)
	if err != nil {
		log.Fatal().Err(err).Msg("could not get staking address")
	}
	staking, err := contracts.NewIStaking(stakingAddr, client)
	if err != nil {
		log.Fatal().Err(err).Msg("could not get instance of staking")
	}

	registered, err := staking.FilterRegistered(nil, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("could not get registered oracles")
	}
	unregistered, err := staking.FilterUnregistered(nil, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("could not get unregistered oracles")
	}

	oracles := make(map[common.Address]bool)
	for registered.Next() {
		oracles[registered.Event.Oracle] = true
	}
	for unregistered.Next() {
		delete(oracles, unregistered.Event.Oracle)
	}

	for k := range oracles {
		validators.RegisterValidator(k)
	}

	sink := make(chan *contracts.IOrakuruCoreFulfilled, 500)
	f, err := core.WatchFulfilled(nil, sink, nil)
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("could not subscribe for new events")
	}
	errChan := f.Err()

	latestBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("could not get latest block")
	}
	for block := uint64(9121845); block <= latestBlock; block += 5001 {
		var end *uint64
		if block+5000 < latestBlock {
			_end := block + 5000
			end = &_end
		}
		requests, err := core.FilterFulfilled(&bind.FilterOpts{
			Start: block,
			End:   end,
		}, nil)
		if err != nil {
			log.Fatal().Err(err).Caller().Msg("cannot collect fulfilled requests")
		}

		for requests.Next() {
			// Empty answer = failed
			if len(requests.Event.Result) == 0 {
				continue
			}

			log.Info().Str("id", hexutil.Encode(requests.Event.RequestId[:])).Msg("processing past request")
			resp, err := core.GetResponses(nil, requests.Event.RequestId)
			if err != nil {
				log.Fatal().Err(err).Caller().Msg("cannot get responses from core contract")
			}
			req, err := core.GetRequest(nil, requests.Event.RequestId)
			if err != nil {
				log.Fatal().Err(err).Caller().Msg("could not get request from core contract")
			}
			for _, resp := range resp {
				elapsed := big.NewInt(0).Sub(resp.SubmittedAt, req.ExecutionTimestamp)
				validators.AddScore(resp.SubmittedBy, Score(elapsed.Uint64()), elapsed.Uint64())
			}
		}
	}

	go func() {
		for {
			select {
			case fulfill := <-sink:
				if len(fulfill.Result) == 0 {
					continue
				}

				log.Info().Str("id", hexutil.Encode(fulfill.RequestId[:])).Msg("processing new request")
				resp, err := core.GetResponses(nil, fulfill.RequestId)
				if err != nil {
					log.Fatal().Err(err).Caller().Msg("cannot get responses from core contract")
				}
				req, err := core.GetRequest(nil, fulfill.RequestId)
				if err != nil {
					log.Fatal().Err(err).Caller().Msg("cannot get request from core contract")
				}
				for _, resp := range resp {
					elapsed := big.NewInt(0).Sub(resp.SubmittedAt, req.ExecutionTimestamp)
					validators.AddScore(resp.SubmittedBy, Score(elapsed.Uint64()), elapsed.Uint64())
				}
			case err := <-errChan:
				log.Fatal().Err(err).Msg("subscription has returned an error")
			}
		}
	}()

	http.HandleFunc("/stats", func(writer http.ResponseWriter, request *http.Request) {
		leaderboard := validators.Collect()
		data, _ := json.Marshal(leaderboard)
		writer.Header().Add("Content-Type", "application/json")
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		_, err = writer.Write(data)
		if err != nil {
			log.Warn().Err(err).Caller().Msg("could not write response")
		}
	})

	err = http.ListenAndServe(*httpAddr, nil)
	log.Error().Err(err).Msg("http server crashed")
}
