import { Schema, model } from "mongoose";

const reviewSchema = new Schema({
  userPhotoURL: { type: String },
  userName: { type: String },
  address: { type: String },
  description: { type: String, required: true },
  rating: { type: Number, required: true },
  comment: { type: String, required: true },
});

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
    reviews: { type: [reviewSchema], default: [] },
  },
  {
    timestamps: true,
    toJSON: {
      virtuals: true,
    },
  }
);

export const Books = model("Books", BookSchema);
