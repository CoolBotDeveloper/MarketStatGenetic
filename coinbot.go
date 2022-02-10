package main

type CoinBot struct {
	symbol                     string
	config                     BotConfig
	dataSource                 *DataSource
	buyIndicators              []BuyTechnicalIndicator
	sellIndicators             []string
	bitcoinSuperTrendIndicator BitcoinSuperTrendIndicator
}

func NewCoinBot(symbol string, config BotConfig, dataSource *DataSource) CoinBot {
	bot := CoinBot{
		symbol:     symbol,
		config:     config,
		dataSource: dataSource,
	}
	bot.initIndicators()

	return bot
}

func (bot *CoinBot) HasBuySignal() bool {
	btcCandles := bot.dataSource.GetCandlesFor(BITCOIN_SYMBOL)
	if !bot.bitcoinSuperTrendIndicator.HasBuySignal(btcCandles) {
		return false
	}

	for _, indicator := range bot.buyIndicators {
		candles := bot.dataSource.GetCandlesFor(bot.symbol)
		if !indicator.HasBuySignal(candles) {
			return false
		}
	}

	return true
}

func (bot *CoinBot) initIndicators() {
	bot.buyIndicators = []BuyTechnicalIndicator{}

	superTrendIndicator := NewSuperTrendIndicator(bot.config)
	bot.buyIndicators = append(bot.buyIndicators, &superTrendIndicator)

	//superTrendIndicatorSecond := NewSuperTrendIndicatorSecond(bot.config)
	//bot.buyIndicators = append(bot.buyIndicators, &superTrendIndicatorSecond)

	averageVolumeIndicator := NewAverageVolumeIndicator(bot.config)
	bot.buyIndicators = append(bot.buyIndicators, &averageVolumeIndicator)

	adxIndicator := NewAdxIndicator(bot.config)
	bot.buyIndicators = append(bot.buyIndicators, &adxIndicator)

	bot.bitcoinSuperTrendIndicator = NewBitcoinSuperTrendIndicator(bot.config)
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
