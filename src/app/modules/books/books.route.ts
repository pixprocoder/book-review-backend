import express from "express";
import { BooksController } from "./books.controller";

const router = express.Router();

router.post("/create-book", BooksController.createBooks);

router.get("/:id", BooksController.getSingleBook);
router.get("/", BooksController.getAllBook);

export const BooksRoute = router;
