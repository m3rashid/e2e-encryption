import cors from 'cors';
import express from 'express';
import { handleExchangeKeys } from './handlers/keys';

const app = express();
app.use(cors({}));

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.post('/', (req, res) => res.send('Hello World!'));
app.post('/exchange-keys', handleExchangeKeys);

const PORT = process.env.PORT || 8080;
app.listen(PORT, () => console.log('Server running on port ' + PORT));
