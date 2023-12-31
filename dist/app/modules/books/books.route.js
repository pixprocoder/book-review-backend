"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.BooksRoute = void 0;
const express_1 = __importDefault(require("express"));
const books_controller_1 = require("./books.controller");
const router = express_1.default.Router();
router.post("/create-book", books_controller_1.BooksController.createBooks);
router.get("/:id", books_controller_1.BooksController.getSingleBook);
router.get("/", books_controller_1.BooksController.getAllBook);
router.delete("/:id", books_controller_1.BooksController.deleteBook);
exports.BooksRoute = router;
