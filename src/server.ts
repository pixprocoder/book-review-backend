/* eslint-disable no-console */
import { Server } from "http";
import mongoose from "mongoose";
import config from "./config/index";
import app from "./app";

const PORT = 3000;

process.on("uncaughtException", (error) => {
  console.error(error);
  process.exit(1);
});

let server: Server;

async function bootstrap() {
  try {
    await mongoose.connect(config.database_url as string);
    console.log(`🛢   Database is connected successfully`);

    server = app.listen(PORT, () => {
      console.log(`Application  listening on port ${PORT}`);
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
