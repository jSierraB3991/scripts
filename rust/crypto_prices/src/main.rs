use serde::{Serialize, Deserialize};

fn main() {
    let mut coin = String::new();
    println!("Why Crypto coin, do you want search: ");
    let _ = std::io::stdin()
            .read_line(&mut coin)
            .expect("Error unexpected");
    let result_price = get_precio(&coin);
    match result_price {
        Ok(price) => println!("The price is: ${price}"),
        Err(error) => println!("Error to search price: {error}"),
    }
    
}

fn get_precio(coin: &str) ->  Result<String, ureq::Error> {
    let url_gecko = format!("https://api.coingecko.com/api/v3/coins/{}?localization=false", coin);
    let body: String = ureq::get(&url_gecko)
        .call()?
        .into_string()?;

    //unwrap, no recomendable
    let coin_data: CoinData = serde_json::from_str(&body).unwrap();
    Ok(coin_data.market_data.current_price.usd.to_string())
}


//For serialize and desearilize
#[derive(Serialize, Deserialize, Debug)]
struct CoinData {
    id: String,
    symbol: String,
    name: String,
    market_data: MarketData,
}

#[derive(Serialize, Deserialize, Debug)]
struct MarketData {
    current_price: Prices,
}

#[derive(Serialize, Deserialize, Debug)]
struct Prices {
    usd: f32,
    clp: f32
}