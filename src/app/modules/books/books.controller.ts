import { Request, Response } from "express";
import { BooksService } from "./books.service";

const createBooks = async (req: Request, res: Response) => {
  const result = await BooksService.createBooks(req.body);

  return res.send(result);
};

export const BooksController = {
  createBooks,
};
