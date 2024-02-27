import cors from "cors";
import express, { Application } from "express";
import routes from "./app/routes";
import cookieParser from "cookie-parser";

const app: Application = express();

// Parser and Middlewares
app.use(cors({ origin: "*" }));
app.use(cookieParser());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// all routes
app.use("/api/v1", routes);

export default app;
