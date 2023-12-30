import { Request, Response } from "express";
import { BooksService } from "./books.service";
import sendResponse from "../../../shared/sendResponse";

const createBooks = async (req: Request, res: Response) => {
  const result = await BooksService.createBooks(req.body);

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

export const BooksController = {
  createBooks,
  getAllBook,
};
