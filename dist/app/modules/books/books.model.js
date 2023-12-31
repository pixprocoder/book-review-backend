"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Books = void 0;
const mongoose_1 = require("mongoose");
const BookSchema = new mongoose_1.Schema({
    title: {
        type: String,
        required: true,
        unique: true,
    },
    author: {
        type: String,
        name: String,
        required: true,
    },
    genre: {
        type: String,
        name: String,
        required: true,
    },
    publicationDate: {
        type: String,
        name: String,
        required: true,
    },
    image: {
        type: String,
        name: String,
        required: true,
    },
}, {
    timestamps: true,
    toJSON: {
        virtuals: true,
    },
});
exports.Books = (0, mongoose_1.model)("Books", BookSchema);
