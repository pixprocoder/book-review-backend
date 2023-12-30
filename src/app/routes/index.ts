import express from "express";
import { BooksRoute } from "../modules/books/books.route";

const router = express.Router();

const moduleRoutes = [
  {
    path: "/books",
    route: BooksRoute,
  },
];

moduleRoutes.forEach((route) => router.use(route.path, route.route));
export default router;
