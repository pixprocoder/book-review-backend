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
Object.defineProperty(exports, "__esModule", { value: true });
exports.BooksService = void 0;
const books_model_1 = require("./books.model");
const createBooks = (payload) => __awaiter(void 0, void 0, void 0, function* () {
    // console.log(payload);
    const result = yield books_model_1.Books.create(payload);
    return result;
});
// all books
const getAllBook = () => __awaiter(void 0, void 0, void 0, function* () {
    const result = yield books_model_1.Books.find({});
    return result;
});
// single book
const getSingleBook = (payload) => __awaiter(void 0, void 0, void 0, function* () {
    const result = yield books_model_1.Books.findById(payload);
    return result;
});
// delete book
const deleteBook = (payload) => __awaiter(void 0, void 0, void 0, function* () {
    const result = yield books_model_1.Books.findByIdAndDelete(payload);
    return result;
});
exports.BooksService = {
    createBooks,
    getAllBook,
    getSingleBook,
    deleteBook,
};
