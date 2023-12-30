import express from "express";
import { BooksController } from "./books.controller";

const router = express.Router();

router.post("/create-books", BooksController.createBooks);

export const BooksRoute = router;
