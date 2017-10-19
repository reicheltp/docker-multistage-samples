package io.vertx.sample.hello;

import io.vertx.core.AbstractVerticle;

public class HelloVerticle extends AbstractVerticle {

  @Override
  public void start() {
    vertx.createHttpServer().requestHandler(req -> {
      req.response()
        .putHeader("content-type", "application/json")
        .end("{ message: 'Hallo Docker Berlin! 🐳' }");
    }).listen(80);
  }
}