import express from "express";
import { BooksController } from "./books.controller";

const router = express.Router();

router.post("/create-book", BooksController.createBooks);

router.get("/:id", BooksController.getSingleBook);
router.get("/", BooksController.getAllBook);
router.delete("/:id", BooksController.deleteBook);
router.patch("/post-review/:id", BooksController.postReview);

export const BooksRoute = router;
