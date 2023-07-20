import express, { Express, Request, Response } from "express";
import cors from "cors";
import dotenv from "dotenv";
import data from "../data/data.json";

dotenv.config();

const app: Express = express();
const port = 5000;
app.use(cors());

app.get("/", (req: Request, res: Response) => {
  res.send("Express + TypeScript Server");
});
app.get("/books", (req: Request, res: Response) => {
  res.send(data);
});

app.listen(port, () => {
  console.log(`⚡️[server]: Server is running at http://localhost:${port}`);
});
