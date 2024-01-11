use reqwest;

#[tokio::main]
async fn main() -> Result<(), reqwest::Error> {
    let response = reqwest::get("https://jsonplaceholder.typicode.com/posts/1").await?;
    let body = response.text().await?;
    println!("Response body:\n{}", body);
    Ok(())
}
