use std::error::Error;

use actix_web::web;
use actix_web::{web::{
    Data,
    Json,
}, post, get,  HttpResponse, Responder};
use sqlx::types::chrono::Utc;

use crate::{
    model::user::{UserModel, UserModelResponse},
    schema::user::CreateUserSchema,
    schema::common::FilterOptions,
    AppState,
};


#[get("")]
pub async fn generate_schedule(
    opts: web::Query<FilterOptions>,
    data: web::Data<AppState>,
) -> impl Responder {
    let limit = opts.limit.unwrap_or(10);
    let offset = (opts.page.unwrap_or(1) - 1) * limit;


    let notes: Vec<UserModel> = sqlx::query_as::<_, UserModel>(
        r#"SELECT * FROM user ORDER by id LIMIT ? OFFSET ?"#).bind(limit as i32).bind(offset as i32)
    .fetch_all(&data.db)
    .await
    .unwrap();

    let note_responses = notes
        .into_iter()
        .map(|note| filter_db_record(&note))
        .collect::<Vec<UserModelResponse>>();


        

    let json_response = serde_json::json!({
        "status": "success",
        "results": note_responses.len(),
        "users": note_responses
    });
    HttpResponse::Ok().json(json_response)
}

pub fn config(conf: &mut web::ServiceConfig) {
    //  print!("userddddconfig");
    //  cfg.service(
            //web::resource("/user").route(route)
      //    web::scope("/user")
        //      .service(create_user)
    //  );
  
      let scope = web::scope("")
      .service(generate_schedule);
      conf.service(scope);
  }
  
  fn filter_db_record(user: &UserModel) -> UserModelResponse {
    UserModelResponse {
        id: user.id.to_owned(),
        email: user.email.to_owned(),
        name: user.name.to_owned(),
        description: user.description.to_owned(),
        createdAt: user.created_at.unwrap(),
        updatedAt: user.updated_at.unwrap(),
    }
}