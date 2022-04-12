package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Storage struct {
	name    string
	connect *sql.DB
}

func NewStorage() Storage {
	return Storage{
		name: time.Now().Format("exchange_2006_01_02__15_04_05.db"),
	}
}

func (storage *Storage) Open() {
	storage.connect, _ = sql.Open("sqlite3", ":memory:")

	storage.CreateBuysTable()
	storage.CreateSellsTable()
}

func (storage *Storage) Close() {
	storage.connect.Close()
}

func (storage *Storage) CreateBuysTable() sql.Result {
	query := `
		CREATE TABLE IF NOT EXISTS buys (
			id integer primary key AUTOINCREMENT,
			symbol VARCHAR(255),
			coins FLOAT,
			exchange_rate FLOAT,
			created_at DATETIME
		);
	`
	result, _ := storage.connect.Exec(query)

	fakeQuery := `
		CREATE TABLE IF NOT EXISTS fake_buys (
			id integer primary key AUTOINCREMENT,
			symbol VARCHAR(255),
			coins FLOAT,
			exchange_rate FLOAT,
			created_at DATETIME
		);
	`
	storage.connect.Exec(fakeQuery)

	return result
}

func (storage *Storage) CreateSellsTable() sql.Result {
	query := `
		CREATE TABLE IF NOT EXISTS sells (
			id integer primary key AUTOINCREMENT,
			symbol VARCHAR(255),
			coins FLOAT,
			exchange_rate FLOAT,
			revenue FLOAT,
			buy_id INT,
			created_at DATETIME
		);
	`
	result, _ := storage.connect.Exec(query)

	fakeQuery := `
		CREATE TABLE IF NOT EXISTS fake_sells (
			id integer primary key AUTOINCREMENT,
			symbol VARCHAR(255),
			coins FLOAT,
			exchange_rate FLOAT,
			revenue FLOAT,
			buy_id INT,
			created_at DATETIME
		);
	`
	storage.connect.Exec(fakeQuery)

	return result
}

func (storage *Storage) AddBuy(symbol string, coinsCount float64, exchangeRate float64, createdAt string) sql.Result {
	//createdAt := time.Now().Format("2006-01-02 15:04:05")
	query := `
		INSERT INTO buys (symbol, coins, exchange_rate, created_at) VALUES ($1, $2, $3, $4);
	`

	result, _ := storage.connect.Exec(query, symbol, coinsCount, exchangeRate, createdAt)

	return result
}

func (storage *Storage) AddFakeBuy(symbol string, coinsCount float64, exchangeRate float64, createdAt string) sql.Result {
	//createdAt := time.Now().Format("2006-01-02 15:04:05")
	query := `
		INSERT INTO fake_buys (symbol, coins, exchange_rate, created_at) VALUES ($1, $2, $3, $4);
	`

	result, _ := storage.connect.Exec(query, symbol, coinsCount, exchangeRate, createdAt)

	return result
}

func (storage *Storage) AddSell(
	symbol string,
	coinsCount float64,
	exchangeRate float64,
	revenue float64,
	buyId int,
	createdAt string,
) sql.Result {
	//createdAt := time.Now().Format("2006-01-02 15:04:05")
	query := `
		INSERT INTO sells (symbol, coins, exchange_rate, revenue, buy_id, created_at) VALUES ($1, $2, $3, $4, $5, $6);
	`

	result, _ := storage.connect.Exec(query, symbol, exchangeRate, coinsCount, revenue, buyId, createdAt)
	return result
}

func (storage *Storage) AddFakeSell(
	symbol string,
	coinsCount float64,
	exchangeRate float64,
	revenue float64,
	buyId int,
	createdAt string,
) sql.Result {
	//createdAt := time.Now().Format("2006-01-02 15:04:05")
	query := `
		INSERT INTO fake_sells (symbol, coins, exchange_rate, revenue, buy_id, created_at) VALUES ($1, $2, $3, $4, $5, $6);
	`

	result, _ := storage.connect.Exec(query, symbol, exchangeRate, coinsCount, revenue, buyId, createdAt)
	return result
}

type Buy struct {
	id           int
	symbol       string
	coins        float64
	exchangeRate float64
	createdAt    string
}

func (storage *Storage) FindUnsoldBuys(
	symbol string,
	exchangeRate float64,
	upperPercentage float64,
	lowerPercentage float64,
	createdAt string,
) []Buy {
	query := `
		SELECT b.*
		FROM buys AS b 
        LEFT JOIN sells AS s 
        	ON s.buy_id = b.id 
        WHERE s.id IS NULL 
            AND b.symbol = $1 
            AND (
                	(
                	    ((b.exchange_rate + ((b.exchange_rate * $2) / 100)) <= $3) OR 
                		((b.exchange_rate - ((b.exchange_rate * $4) / 100)) >= $3)
                	)
                )
	`

	unsoldBuys := []Buy{}

	rows, _ := storage.connect.Query(query, symbol, upperPercentage, exchangeRate, lowerPercentage)
	defer rows.Close()

	for rows.Next() {
		buy := Buy{}
		rows.Scan(&buy.id, &buy.symbol, &buy.coins, &buy.exchangeRate, &buy.createdAt)
		unsoldBuys = append(unsoldBuys, buy)
	}

	return unsoldBuys
}

func (storage *Storage) FindFakeUnsoldBuys(
	symbol string,
	exchangeRate float64,
	upperPercentage float64,
	lowerPercentage float64,
	createdAt string,
) []Buy {
	query := `
		SELECT b.*
		FROM fake_buys AS b 
        LEFT JOIN fake_sells AS s 
        	ON s.buy_id = b.id 
        WHERE s.id IS NULL 
            AND b.symbol = $1 
            AND (
                	(
                	    ((b.exchange_rate + ((b.exchange_rate * $2) / 100)) <= $3) OR 
                		((b.exchange_rate - ((b.exchange_rate * $4) / 100)) >= $3)
                	)
                )
	`

	unsoldBuys := []Buy{}

	rows, _ := storage.connect.Query(query, symbol, upperPercentage, exchangeRate, lowerPercentage)
	defer rows.Close()

	for rows.Next() {
		buy := Buy{}
		rows.Scan(&buy.id, &buy.symbol, &buy.coins, &buy.exchangeRate, &buy.createdAt)
		unsoldBuys = append(unsoldBuys, buy)
	}

	return unsoldBuys
}

func (storage *Storage) FindFirstSellZombies(symbol string, exchangeRate float64, createdAt string, minutes int, sellPercentage float64) []Buy {
	query := `
		SELECT b.*
		FROM buys AS b 
        LEFT JOIN sells AS s 
        	ON s.buy_id = b.id 
        WHERE s.id IS NULL 
            AND b.symbol = $1 
            AND b.created_at < $2
			AND (
			    	(b.exchange_rate + ((b.exchange_rate * $3) / 100)) <= $4
				)
	`

	unsoldBuys := []Buy{}
	candleTime := ConvertDateStringToTime(createdAt)
	zombieDuration := GetCurrentMinusTime(candleTime, minutes)

	rows, _ := (*storage).connect.Query(query, symbol, zombieDuration.Format("2006-01-02 15:04:05"), sellPercentage, exchangeRate)
	defer rows.Close()

	for rows.Next() {
		buy := Buy{}
		rows.Scan(&buy.id, &buy.symbol, &buy.coins, &buy.exchangeRate, &buy.createdAt)
		unsoldBuys = append(unsoldBuys, buy)
	}

	return unsoldBuys
}

func (storage *Storage) FindFakeFirstSellZombies(symbol string, exchangeRate float64, createdAt string, minutes int, sellPercentage float64) []Buy {
	query := `
		SELECT b.*
		FROM fake_buys AS b 
        LEFT JOIN fake_sells AS s 
        	ON s.buy_id = b.id 
        WHERE s.id IS NULL 
            AND b.symbol = $1 
            AND b.created_at < $2
			AND (
			    	(b.exchange_rate + ((b.exchange_rate * $3) / 100)) <= $4
				)
	`

	unsoldBuys := []Buy{}
	candleTime := ConvertDateStringToTime(createdAt)
	zombieDuration := GetCurrentMinusTime(candleTime, minutes)

	rows, _ := (*storage).connect.Query(query, symbol, zombieDuration.Format("2006-01-02 15:04:05"), sellPercentage, exchangeRate)
	defer rows.Close()

	for rows.Next() {
		buy := Buy{}
		rows.Scan(&buy.id, &buy.symbol, &buy.coins, &buy.exchangeRate, &buy.createdAt)
		unsoldBuys = append(unsoldBuys, buy)
	}

	return unsoldBuys
}

func (storage *Storage) FindExitZombies(symbol string, createdAt string, minutes int) []Buy {
	query := `
		SELECT b.*
		FROM buys AS b 
        LEFT JOIN sells AS s 
        	ON s.buy_id = b.id 
        WHERE s.id IS NULL 
            AND b.symbol = $1 
            AND b.created_at < $2
	`

	unsoldBuys := []Buy{}
	candleTime := ConvertDateStringToTime(createdAt)
	zombieDuration := GetCurrentMinusTime(candleTime, minutes)

	rows, _ := (*storage).connect.Query(query, symbol, zombieDuration.Format("2006-01-02 15:04:05"))
	defer rows.Close()

	for rows.Next() {
		buy := Buy{}
		rows.Scan(&buy.id, &buy.symbol, &buy.coins, &buy.exchangeRate, &buy.createdAt)
		unsoldBuys = append(unsoldBuys, buy)
	}

	return unsoldBuys
}

func (storage *Storage) FindFakeExitZombies(symbol string, createdAt string, minutes int) []Buy {
	query := `
		SELECT b.*
		FROM fake_buys AS b 
        LEFT JOIN fake_sells AS s 
        	ON s.buy_id = b.id 
        WHERE s.id IS NULL 
            AND b.symbol = $1 
            AND b.created_at < $2
	`

	unsoldBuys := []Buy{}
	candleTime := ConvertDateStringToTime(createdAt)
	zombieDuration := GetCurrentMinusTime(candleTime, minutes)

	rows, _ := (*storage).connect.Query(query, symbol, zombieDuration.Format("2006-01-02 15:04:05"))
	defer rows.Close()

	for rows.Next() {
		buy := Buy{}
		rows.Scan(&buy.id, &buy.symbol, &buy.coins, &buy.exchangeRate, &buy.createdAt)
		unsoldBuys = append(unsoldBuys, buy)
	}

	return unsoldBuys
}

func (storage *Storage) CanBuyInGivenPeriod(symbol string, createdAt string, period int) bool {
	var count int
	query := `
		SELECT COUNT(s.id)
		FROM sells AS s 
        WHERE s.symbol = $1 AND s.created_at > $2
	`

	candleTime := ConvertDateStringToTime(createdAt)
	canNotBuyDuration := GetCurrentMinusTime(candleTime, period)
	(*storage).connect.QueryRow(query, symbol, canNotBuyDuration).Scan(&count)

	return count == 0
}

func (storage *Storage) CountUnsoldBuys(symbol string) int {
	var count int
	query := `
		SELECT COUNT(b.id)
		FROM buys AS b 
        LEFT JOIN sells AS s 
        	ON s.buy_id = b.id 
        WHERE b.symbol = $1 AND s.id IS NULL
	`
	(*storage).connect.QueryRow(query, symbol).Scan(&count)

	return count
}

func (storage *Storage) CountFakeUnsoldBuys(symbol string) int {
	var count int
	query := `
		SELECT COUNT(b.id)
		FROM fake_buys AS b 
        LEFT JOIN fake_sells AS s 
        	ON s.buy_id = b.id 
        WHERE b.symbol = $1 AND s.id IS NULL
	`
	(*storage).connect.QueryRow(query, symbol).Scan(&count)

	return count
}

func (storage *Storage) getBuyBySymbol(symbol string) Buy {
	symbolBuy := Buy{}
	query := `
		SELECT *
		FROM buys
		WHERE symbol = $1
		ORDER BY created_at DESC
	`

	(*storage).connect.QueryRow(query, symbol).Scan(&symbolBuy.id, &symbolBuy.symbol, &symbolBuy.coins, &symbolBuy.exchangeRate, &symbolBuy.createdAt)

	return symbolBuy
}

type revenue struct {
	value float64
}

func (storage *Storage) GetTotalRevenue() float64 {
	rev := revenue{}
	query := `
		SELECT (SUM(revenue) - COUNT(id) * 100) AS rev 
		FROM sells 
		GROUP BY symbol
	`
	row := (*storage).connect.QueryRow(query)
	row.Scan(&rev.value)

	return rev.value
}

func (storage *Storage) GetFakeTotalRevenue(symbol string) float64 {
	rev := revenue{}
	query := `
		SELECT (SUM(revenue) - COUNT(id) * 100) AS rev 
		FROM fake_sells 
		where symbol = $1
		GROUP BY symbol
	`
	row := (*storage).connect.QueryRow(query, symbol)
	row.Scan(&rev.value)

	return rev.value
}

func (storage *Storage) CalculateRevenueFromStartTime(symbol, startTime string) float64 {
	rev := revenue{}
	query := `
		SELECT (SUM(revenue) - COUNT(id) * 100) AS rev 
		FROM sells 
		WHERE symbol = $1 AND created_at > $2
		GROUP BY symbol
	`
	row := (*storage).connect.QueryRow(query, symbol, startTime)
	row.Scan(&rev.value)

	return rev.value
}

func (storage *Storage) CleanFakeBuySellTables(symbol string) {
	queryBuys := `
		DELETE
		FROM fake_buys 
		WHERE symbol = $1
	`
	(*storage).connect.Exec(queryBuys, symbol)

	querySells := `
		DELETE
		FROM fake_sells 
		WHERE symbol = $1
	`
	(*storage).connect.Exec(querySells, symbol)
}

type buysCount struct {
	value int
}

func (storage *Storage) GetBuysCount() int {
	count := buysCount{}
	query := `
		SELECT COUNT(id) AS c 
		FROM buys 
	`
	row := (*storage).connect.QueryRow(query)
	row.Scan(&count.value)

	return count.value
}

func (storage *Storage) GetSuccessBuysCount() int {
	count := buysCount{}
	query := `
		SELECT COUNT(b.id) AS c
		FROM buys AS b
		INNER JOIN sells s ON b.id = s.buy_id
		WHERE s.revenue > 100;
	`
	row := (*storage).connect.QueryRow(query)
	row.Scan(&count.value)

	return count.value
}

// Exchange manager
type ExchangeManager struct {
	config  BotConfig
	storage Storage
}

func NewExchangeManager(config BotConfig) ExchangeManager {
	storage := NewStorage()
	storage.Open()

	return ExchangeManager{
		config:  config,
		storage: storage,
	}
}

func (em *ExchangeManager) Close() {
	em.storage.Close()
}

func (em *ExchangeManager) Buy(symbol string, exchangeRate float64, createdAt string) {
	//createdAt := time.Now().Format("2006-01-02 15:04:05")
	coinsCount := TOTAL_MONEY_AMOUNT / exchangeRate

	em.storage.AddBuy(symbol, coinsCount, exchangeRate, createdAt)
}

func (em *ExchangeManager) Sell(symbol string, exchangeRate float64, createdAt string) {
	buy := em.storage.getBuyBySymbol(symbol)
	if buy.id != 0 {
		em.storage.AddSell(
			symbol,
			buy.coins,
			exchangeRate,
			(buy.coins * exchangeRate),
			buy.id,
			createdAt,
		)
	}
}

func (em *ExchangeManager) UpdateBuys(symbol string, exchangeRate float64, createdAt string) []UnsoldBuy {
	//createdAt := time.Now().Format("2006-01-02 15:04:05")

	normal := em.UpdateNormalBuys(symbol, exchangeRate, createdAt)
	firstSell := em.updateFirstSellZombies(symbol, exchangeRate, createdAt)
	exit := em.updateExitZombies(symbol, exchangeRate, createdAt)

	result := append(normal, firstSell...)

	return append(result, exit...)
}

type UnsoldBuy struct {
	Symbol       string
	ExchangeRate float64
	Revenue      float64
}

func (em *ExchangeManager) UpdateNormalBuys(symbol string, exchangeRate float64, createdAt string) []UnsoldBuy {
	reportUnsoldBuys := []UnsoldBuy{}
	//createdAt := time.Now().Format("2006-01-02 15:04:05")
	unsoldBuys := em.GetUnsoldBuys(symbol, exchangeRate, createdAt)

	for _, buy := range unsoldBuys {
		calcedRevenue := em.calcRevenue(buy.coins, exchangeRate)
		reportUnsoldBuys = append(reportUnsoldBuys, UnsoldBuy{
			Symbol:       symbol,
			ExchangeRate: exchangeRate,
			Revenue:      calcedRevenue,
		})

		PrintNeuralData(symbol, buy.createdAt, createdAt, calcedRevenue)

		fmt.Println(fmt.Sprintf("COIN: %s, SELL: %s, Volume, %f, EXCHANGE_RATE: %f, Revenue: %f", symbol, createdAt, 0.0, exchangeRate, calcedRevenue-100))
		em.storage.AddSell(
			symbol,
			buy.coins,
			exchangeRate,
			calcedRevenue,
			buy.id,
			createdAt,
		)
	}

	return reportUnsoldBuys
}

func (em *ExchangeManager) updateFirstSellZombies(symbol string, exchangeRate float64, createdAt string) []UnsoldBuy {
	reportUnsoldBuys := []UnsoldBuy{}
	firstSellZombies := em.getFirstSellZombies(
		symbol,
		exchangeRate,
		createdAt,
		em.config.UnsoldFirstSellDurationMinutes,
		em.config.UnsoldFirstSellPercentage,
	)

	for _, expiredBuy := range firstSellZombies {
		calcedRevenue := expiredBuy.coins * exchangeRate
		reportUnsoldBuys = append(reportUnsoldBuys, UnsoldBuy{
			Symbol:       symbol,
			ExchangeRate: exchangeRate,
			Revenue:      calcedRevenue,
		})

		PrintNeuralData(symbol, expiredBuy.createdAt, createdAt, calcedRevenue)

		fmt.Println(fmt.Sprintf("COIN: %s, SELL: %s, Volume: %f, EXCHANGE_RATE: %f, Revenue: %f", symbol, createdAt, 0.0, exchangeRate, expiredBuy.coins*exchangeRate-100))
		em.storage.AddSell(
			symbol,
			expiredBuy.coins,
			exchangeRate,
			expiredBuy.coins*exchangeRate,
			expiredBuy.id,
			createdAt,
		)
	}

	return reportUnsoldBuys
}

func (em *ExchangeManager) getFirstSellZombies(symbol string, exchangeRate float64, createdAt string, minutes int, sellPercentage float64) []Buy {
	return em.storage.FindFirstSellZombies(symbol, exchangeRate, createdAt, minutes, sellPercentage)
}

func (em *ExchangeManager) updateExitZombies(symbol string, exchangeRate float64, createdAt string) []UnsoldBuy {
	reportUnsoldBuys := []UnsoldBuy{}
	exitZombies := em.getExitZombies(symbol, createdAt, em.config.UnsoldFinalSellDurationMinutes)
	for _, expiredBuy := range exitZombies {
		calcedRevenue := expiredBuy.coins * exchangeRate
		reportUnsoldBuys = append(reportUnsoldBuys, UnsoldBuy{
			Symbol:       symbol,
			ExchangeRate: exchangeRate,
			Revenue:      calcedRevenue,
		})

		PrintNeuralData(symbol, expiredBuy.createdAt, createdAt, calcedRevenue)

		fmt.Println(fmt.Sprintf("COIN: %s, SELL: %s, Volume: %f, EXCHANGE_RATE: %f, Revenue: %f", symbol, createdAt, 0.0, exchangeRate, expiredBuy.coins*exchangeRate-100))
		em.storage.AddSell(
			symbol,
			expiredBuy.coins,
			exchangeRate,
			expiredBuy.coins*exchangeRate,
			expiredBuy.id,
			createdAt,
		)
	}

	return reportUnsoldBuys
}

func (em *ExchangeManager) UpdateAllExitSymbols(symbol string, exchangeRate float64, createdAt string) []UnsoldBuy {
	reportUnsoldBuys := []UnsoldBuy{}
	exitZombies := em.getExitZombies(symbol, createdAt, 0)

	for _, expiredBuy := range exitZombies {
		calcedRevenue := expiredBuy.coins * exchangeRate
		reportUnsoldBuys = append(reportUnsoldBuys, UnsoldBuy{
			Symbol:       symbol,
			ExchangeRate: exchangeRate,
			Revenue:      calcedRevenue,
		})

		PrintNeuralData(symbol, expiredBuy.createdAt, createdAt, calcedRevenue)

		fmt.Println(fmt.Sprintf("COIN: %s, SELL: %s, Volume: %f, EXCHANGE_RATE: %f, Revenue: %f", symbol, createdAt, 0.0, exchangeRate, calcedRevenue-100))
		em.storage.AddSell(
			symbol,
			expiredBuy.coins,
			exchangeRate,
			calcedRevenue,
			expiredBuy.id,
			createdAt,
		)
	}

	return reportUnsoldBuys
}

func (manager *ExchangeManager) getExitZombies(symbol string, createdAt string, minutes int) []Buy {
	return (*manager).storage.FindExitZombies(symbol, createdAt, minutes)
}

func (em *ExchangeManager) GetUnsoldBuys(symbol string, exchangeRate float64, createdAt string) []Buy {
	return em.storage.FindUnsoldBuys(
		symbol,
		exchangeRate,
		em.config.HighSellPercentage,
		em.config.LowSellPercentage,
		createdAt,
	)
}

func (em *ExchangeManager) CanBuyInGivenPeriod(symbol string) bool {
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	return em.storage.CanBuyInGivenPeriod(symbol, createdAt, em.config.StopBuyAfterSellPeriodMinutes)
}

func (em *ExchangeManager) CountUnsoldBuys(symbol string) int {
	return em.storage.CountUnsoldBuys(symbol)
}

func (manager *ExchangeManager) GetTotalRevenue() float64 {
	return (*manager).storage.GetTotalRevenue()
}

func (manager *ExchangeManager) GetBuysCount() int {
	return (*manager).storage.GetBuysCount()
}

func (manager *ExchangeManager) GetSuccessBuysCount() int {
	return (*manager).storage.GetSuccessBuysCount()
}

func (em *ExchangeManager) calcRevenue(coinsCount float64, exchangeRate float64) float64 {
	prevRevenue := coinsCount * exchangeRate
	if prevRevenue == TOTAL_MONEY_AMOUNT {
		return TOTAL_MONEY_AMOUNT
	}

	if prevRevenue > TOTAL_MONEY_AMOUNT {
		plus := (TOTAL_MONEY_AMOUNT * em.config.HighSellPercentage) / 100

		return TOTAL_MONEY_AMOUNT + plus
	}

	minus := (TOTAL_MONEY_AMOUNT * em.config.LowSellPercentage) / 100

	return TOTAL_MONEY_AMOUNT - minus
}
