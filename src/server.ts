/* eslint-disable no-console */
import { Server } from "http";
import mongoose from "mongoose";
import config from "./config/index";
import app from "./app";

process.on("uncaughtException", (error) => {
  console.error(error);
  process.exit(1);
});

let server: Server;

async function bootstrap() {
  try {
    await mongoose.connect(config.database_url as string);
    console.log(`🛢   Database is connected successfully`);

    server = app.listen(config.port, () => {
      console.log(`Application  listening on port ${config.port}`);
    });
  } catch (err) {
    console.error("Failed to connect database", err);
  }

  process.on("unhandledRejection", (error) => {
    if (server) {
      server.close(() => {
        console.error(error);
        process.exit(1);
      });
    } else {
      process.exit(1);
    }
  });
}

bootstrap();
