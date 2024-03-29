
# Backend API Documentation
## main.go
### Functions
#### httpHandler()
- Defines the API requests to be used in the project, setting the url expected, the function associated with the url, and the http request type. Returns handlers using CORS, allowing the transfer of data between two domains and circumventing security features implemented in modern browsers.
#### main()
- Calls InitialMigration() and UpdateStocks() and listens for requests from frontend.  


## dns.go
### Functions
#### InitialMigration()
- Create the database variable using gorm. Automatically creates tables for the Stock, UserStocks, and Credentials structs. Connection to the database is established elsewhere in the file. 


## angular_live.go
### Functions
#### getOrigin()
- Creates an url struct using localhost:4200 and returns a pointer to this url struct. 


## credentials.go
### Types
#### type Credentials
- Like the rest of our databases, credentials uses gorm.Model and contains just username, password and funds.
### Functions
#### login()
- When login() is called, it verifies that the username and password combination exits in the database. If it does, then 
it returns and http status code 200. If it does not, it will return code 401 for unauthorized
#### signup()
- This functions will create a new user account with a username and password. If the username is not take, it will return a 200 code. 
If the username is taken, then the code will return a 401 for unauthorized.
#### deleteCredentials()
- This will check if the username give is in the database, if it is, it will delete and return a 200 code. If the username
is not in the database, then the code returned is 409.


## stock.go
### Types
#### type Stock
- Uses gorm.Model. Contains a ticker and a price. 
#### type Query
- Contains a ticker and a start & end day, month, and year, to be used in the QueryStocks() function. 
### Functions
#### GetStockPrice()
- Returns price associated with stock ticker.
#### GetStocks()
- Returns all data pertaining to all stocks in json format.
#### GetStock()
- Returns all stock data pertaining to the ticker used as input in json format.
#### UpdateStocks()
- Updates stock-price data in the database for each stock stored in the database. Does not require json input.
#### QueryStocks()
- Returns daily stock price(s) and date(s) for a specific stock, over a start and end period. Requires json input data in the form of the Query struct. Returns this information in json format. 


## userstocks.go
### Types
#### type UserStocks
- Uses gorm.Model. Defines a username, ticker, and shares variables, and the json representations of each. In this context, it is meant to represent a purchase order. 
#### type StockTickerShares
- Used to store ticker, shares, price, and change variables pertaining to user's owned stock.  
### Functions
#### PurchaseStock()
- Receives input according to the UserStocks struct. Uses the username to access the funds of the user in the credentials table. Calculates the cost of the purchase order, considers if the user has enough funds. If so, the purchase order is completed, updating or creating an entry in the database containing the username, how many shares the user owns, and ticker of the owned stock. The users funds are updated as well. 
#### SellStock()
- Receives input according to the UserStocks struct. Follows the same format as the previous function, however, we check if the user owns shares of the stock they are attempting to sell. Once this is done, the sale is complete and the users funds are updated. The owned shares of the user are updated or deleted in the database. 
#### GetStocksOwned()
- Returns stock information regarding the stock the user owns. Uses the username as an input. Uses the StockTickerShares struct.
#### GetUserStocksArray()
- returns an array of stocks owned by the user. Uses the UserStocks struct. 
#### ResetAccount()
- Takes username as input. Removes records of ownership of any stocks, and from the database. Also removes logs associated with the account. User funds is reset to default value. 


## portfoliohistory.go
### Types
#### type PortfolioHistory
- Struct for storing username, ticker, shares, ordertype, and ordervalue
#### type PortfolioInfo
- Struct for storing portfolio value and change in portfolio value. 
### Functions
#### CreateLog()
- Inputs are ordertype string, UserStocks object, and ordervalue. Creates a log in the database recording a purchase/sell. 
#### DeleteAllLogs()
- Deletes all records in the database pertaining to the user. 
#### GetLogs()
- Returns all logs pertaining to the user in JSON format. 
#### GetUserPortfolioInfo()
- Encodes information into the PorfolioInfo struct and returns it in JSON format. Uses functions GetUserPortfolioValue().
#### GetUserPortfolioValue()
- Takes username as input. Sums the value of each stock per share that the user owns. Adds the user's funds to this amount and returns the total value. 
#### GetIndividualStockChange()
- Takes username, ticker, and shares as input. Gets the average purchase price or cost for each stock the user owns. Compares these values to current value of stock. Returns the change as a percent. 




