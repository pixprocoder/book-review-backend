import { Books } from "./books.model";

const createBooks = async (payload: any) => {
  const result = await Books.create(payload);
  return result;
};

export const BooksService = {
  createBooks,
};
