import { Books } from "./books.model";

const createBooks = async (payload: any) => {
  console.log("service", payload);
  console.log(payload);
  const result = await Books.create(payload);

  return result;
};

// all books
const getAllBook = async () => {
  const result = await Books.find({});
  return result;
};

// single book
const getSingleBook = async (payload: string) => {
  const result = await Books.findById(payload);
  return result;
};

// delete book
const deleteBook = async (payload: string) => {
  const result = await Books.findByIdAndDelete(payload);

  return result;
};

export const BooksService = {
  createBooks,
  getAllBook,
  getSingleBook,
  deleteBook,
};
