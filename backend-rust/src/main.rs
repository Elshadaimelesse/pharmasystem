use actix_cors::Cors;
use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder};
use serde::{Deserialize, Serialize};
use std::sync::Mutex;

#[derive(Serialize, Deserialize, Clone)]
pub struct DrugPriceComparison {
    pub drug_id: String,
    pub brand_name: String,
    pub generic_name: String,
    pub strength: String,
    pub pharmacy_name: String,
    pub price: f64,
    pub in_stock: bool,
    pub stock_quantity: u32,
    pub delivery_available: bool,
}

struct AppState {
    inventory: Mutex<Vec<DrugPriceComparison>>,
}

#[get("/api/drugs/compare")]
async fn compare_drug_prices(data: web::Data<AppState>) -> impl Responder {
    let items = data.inventory.lock().unwrap();
    HttpResponse::Ok().json(&*items)
}

#[post("/api/drugs/update-price")]
async fn update_pharmacy_price(
    data: web::Data<AppState>,
    item: web::Json<DrugPriceComparison>,
) -> impl Responder {
    let mut items = data.inventory.lock().unwrap();
    if let Some(pos) = items.iter().position(|x| x.drug_id == item.drug_id && x.pharmacy_name == item.pharmacy_name) {
        items[pos] = item.into_inner();
    } else {
        items.push(item.into_inner());
    }
    HttpResponse::Ok().json(serde_json::json!({"status": "updated"}))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("PharmRx Rust High-Throughput Inventory Engine starting on port 8081...");

    let initial_data = web::Data::new(AppState {
        inventory: Mutex::new(vec![
            DrugPriceComparison {
                drug_id: "drug_amox_clav".into(),
                brand_name: "Augmentin".into(),
                generic_name: "Amoxicillin and Clavulanate".into(),
                strength: "500mg/125mg".into(),
                pharmacy_name: "Apex Care Super Rx".into(),
                price: 18.50,
                in_stock: true,
                stock_quantity: 42,
                delivery_available: true,
            },
            DrugPriceComparison {
                drug_id: "drug_amox_clav".into(),
                brand_name: "Augmentin".into(),
                generic_name: "Amoxicillin and Clavulanate".into(),
                strength: "500mg/125mg".into(),
                pharmacy_name: "St. Jude Community Pharmacy".into(),
                price: 16.75,
                in_stock: true,
                stock_quantity: 28,
                delivery_available: true,
            },
        ]),
    });

    HttpServer::new(move || {
        let cors = Cors::permissive();
        App::new()
            .wrap(cors)
            .app_data(initial_data.clone())
            .service(compare_drug_prices)
            .service(update_pharmacy_price)
    })
    .bind(("0.0.0.0", 8081))?
    .run()
    .await
}
