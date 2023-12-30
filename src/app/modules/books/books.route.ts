import express from "express";
import { BooksController } from "./books.controller";

const router = express.Router();

router.post("/create-book", BooksController.createBooks);

export const BooksRoute = router;
