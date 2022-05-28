package main

type CoinBot struct {
	symbol                            string
	config                            BotConfig
	dataSource                        *DataSource
	buyIndicators                     []BuyTechnicalIndicator
	sellIndicators                    []string
	bitcoinSuperTrendIndicator        BitcoinSuperTrendIndicator
	btcPriceGrowthPercentageIndicator PriceGrowthIndicator
	hasReachedBtcPercentage           bool
}

func NewCoinBot(symbol string, config BotConfig, dataSource *DataSource) CoinBot {
	bot := CoinBot{
		symbol:                  symbol,
		config:                  config,
		dataSource:              dataSource,
		hasReachedBtcPercentage: false,
	}
	bot.initIndicators()

	return bot
}

func (bot *CoinBot) HasBuySignal() bool {
	// Так можно отключить супертренд по биткойну, смотри ниже
	//btcCandles := bot.dataSource.GetCandlesFor(BITCOIN_SYMBOL)
	//if !bot.bitcoinSuperTrendIndicator.HasBuySignal(btcCandles) {
	//	return false
	//}

	//if !bot.hasReachedBtcPercentage && !bot.btcPriceGrowthPercentageIndicator.HasBuySignal(btcCandles) {
	//	return false
	//}

	bot.SetHasReached()
	for _, indicator := range bot.buyIndicators {
		candles := bot.dataSource.GetCandlesFor(bot.symbol)
		if !indicator.HasBuySignal(candles) {
			return false
		}
	}

	return true
}

func (bot *CoinBot) SetHasReached() {
	bot.hasReachedBtcPercentage = true
}

func (bot *CoinBot) ResetHasReached() {
	bot.hasReachedBtcPercentage = false
}

func (bot *CoinBot) initIndicators() {
	bot.buyIndicators = []BuyTechnicalIndicator{}

	//superTrendIndicator := NewSuperTrendIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &superTrendIndicator)

	//averageVolumeIndicator := NewAverageVolumeIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &averageVolumeIndicator)

	//priceFallIndicator := NewPriceFallIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &priceFallIndicator)

	//flatLineIndicator := NewFlatLineIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &flatLineIndicator)

	//twoLineIndicator := NewTwoLineIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &twoLineIndicator)

	//medianVolumeIndicator := NewMedianVolumeIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &medianVolumeIndicator)

	//candleBodyHeightIndicator := NewCandleBodyHeightIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &candleBodyHeightIndicator)

	//adxIndicator := NewAdxIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &adxIndicator)

	//bot.bitcoinSuperTrendIndicator = NewBitcoinSuperTrendIndicator(bot.config)

	//Bitcoin price growth indicator
	//bot.btcPriceGrowthPercentageIndicator = NewPriceGrowthIndicator(bot.config)

	wholeDayTotalVolumeIndicator := NewWholeDayTotalVolumeIndicator(bot.config)
	bot.buyIndicators = append(bot.buyIndicators, &wholeDayTotalVolumeIndicator)

	halfVolumeIndicator := NewHalfVolumeIndicator(bot.config)
	bot.buyIndicators = append(bot.buyIndicators, &halfVolumeIndicator)

	//flatLineSearchIndicator := NewFlatLineSearchIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &flatLineSearchIndicator)

	//tripleGrowthIndicator := NewTripleGrowthIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &tripleGrowthIndicator)

	pastMaxPriceIndicator := NewPastMaxPriceIndicator(bot.config)
	bot.buyIndicators = append(bot.buyIndicators, &pastMaxPriceIndicator)

	//smoothGrowthIndicator := NewSmoothGrowthIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &smoothGrowthIndicator)

	//eachVolumeMinValueIndicator := NewEachVolumeMinValueIndicator(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &eachVolumeMinValueIndicator)
}

// Coin bot factory
type CoinBotFactory struct {
	bots       map[string]CoinBot
	dataSource *DataSource
}

func NewCoinBotFactory(dataSource *DataSource) CoinBotFactory {
	return CoinBotFactory{
		bots:       map[string]CoinBot{},
		dataSource: dataSource,
	}
}

func (factory *CoinBotFactory) FactoryCoinBot(symbol string, config BotConfig) CoinBot {
	if _, ok := factory.bots[symbol]; !ok {
		factory.bots[symbol] = NewCoinBot(symbol, config, factory.dataSource)
	}

	return factory.bots[symbol]
}
