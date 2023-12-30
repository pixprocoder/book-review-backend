import cors from "cors";
import express, { Application } from "express";
import routes from "./app/routes";

const app: Application = express();

// Parser and Middlewares
app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// routes
app.use("/api/v1", routes);

export default app;
