import { Schema, model } from "mongoose";

const BookSchema = new Schema(
  {
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
  },
  {
    timestamps: true,
    toJSON: {
      virtuals: true,
    },
  }
);

export const Books = model("Books", BookSchema);
