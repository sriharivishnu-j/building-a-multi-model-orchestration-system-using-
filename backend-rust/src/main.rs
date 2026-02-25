use log::{info, LevelFilter};
use serde::Serialize;

#[derive(Serialize)]
struct Data {
    value: i32,
}

fn main() {
    env_logger::builder().filter_level(LevelFilter::Info).init();
    let data = Data { value: 42 };
    process_data(&data);
}

fn process_data(data: &Data) {
    info!("Processing data: {}", data.value);
}