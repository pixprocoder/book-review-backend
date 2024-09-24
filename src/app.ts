import cors from "cors";
import express, { Application } from "express";
import routes from "./app/routes";
import cookieParser from "cookie-parser";
import multer from 'multer';
import path from 'path';

const app: Application = express();

// Parser and Middlewares
app.use(cors({ origin: "*" }));
app.use(cookieParser());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// all routes
app.use("/api/v1", routes);

// Set storage engine
const storage = multer.diskStorage({
    destination: './uploads/',
    filename: (req, file, cb) => {
        cb(null, file.fieldname + '-' + Date.now() + path.extname(file.originalname));
    },
});

// Initialize upload
const upload = multer({
    storage: storage,
    limits: { fileSize: 5000000 }, // Limit file size to 5MB
}).single('image');

// Upload route
app.post('/upload', (req, res) => {
    upload(req, res, (err) => {
        if (err) {
            console.log(err)

            res.status(500).send('File upload failed.');
        } else {
            if (req.file === undefined) {
                res.status(400).send('No file selected.');
            } else {
                res.json({
                    success: true,
                    filePath: `/uploads/${req.file.filename}`,
                });
            }
        }
    });
});

export default app;
