import { Request, Response } from "express";

const createBooks = async (req: Request, res: Response) => {
  console.log(req.body, res);
};

export const BooksController = {
  createBooks,
};
