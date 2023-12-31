"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.BooksController = void 0;
const books_service_1 = require("./books.service");
const sendResponse_1 = __importDefault(require("../../../shared/sendResponse"));
const createBooks = (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const result = yield books_service_1.BooksService.createBooks(req.body);
    (0, sendResponse_1.default)(res, {
        statusCode: 200,
        success: true,
        message: "Book created successfully",
        data: result,
    });
});
const getAllBook = (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const result = yield books_service_1.BooksService.getAllBook();
    (0, sendResponse_1.default)(res, {
        statusCode: 200,
        success: true,
        message: "Book fetched successfully",
        data: result,
    });
});
// single book
const getSingleBook = (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const result = yield books_service_1.BooksService.getSingleBook(req.params.id);
    (0, sendResponse_1.default)(res, {
        statusCode: 200,
        success: true,
        message: "single book fetched successfully",
        data: result,
    });
});
// delete book
const deleteBook = (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const result = yield books_service_1.BooksService.deleteBook(req.params.id);
    (0, sendResponse_1.default)(res, {
        statusCode: 200,
        success: true,
        message: "book deleted successfully",
        data: result,
    });
});
exports.BooksController = {
    createBooks,
    getAllBook,
    getSingleBook,
    deleteBook,
};
