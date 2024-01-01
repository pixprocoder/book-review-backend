import { Request, Response } from "express";
import { BooksService } from "./books.service";
import sendResponse from "../../../shared/sendResponse";

const createBooks = async (req: Request, res: Response) => {
  console.log(req.body);
  const result = await BooksService.createBooks(req.body.formData);

  sendResponse(res, {
    statusCode: 200,
    success: true,
    message: "Book created successfully",
    data: result,
  });
};
const getAllBook = async (req: Request, res: Response) => {
  const result = await BooksService.getAllBook();

  sendResponse(res, {
    statusCode: 200,
    success: true,
    message: "Book fetched successfully",
    data: result,
  });
};

// single book
const getSingleBook = async (req: Request, res: Response) => {
  const result = await BooksService.getSingleBook(req.params.id);
  sendResponse(res, {
    statusCode: 200,
    success: true,
    message: "single book fetched successfully",
    data: result,
  });
};
// delete book
const deleteBook = async (req: Request, res: Response) => {
  const result = await BooksService.deleteBook(req.params.id);
  sendResponse(res, {
    statusCode: 200,
    success: true,
    message: "book deleted successfully",
    data: result,
  });
};

export const BooksController = {
  createBooks,
  getAllBook,
  getSingleBook,
  deleteBook,
};
