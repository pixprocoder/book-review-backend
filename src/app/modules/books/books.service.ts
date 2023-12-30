import { Books } from "./books.model";

const createBooks = async (payload: any) => {
  // console.log(payload);
  const result = await Books.create(payload);
  return result;
};
const getAllBook = async () => {
  // console.log(payload);
  const result = await Books.find({});
  return result;
};

export const BooksService = {
  createBooks,
  getAllBook,
};
