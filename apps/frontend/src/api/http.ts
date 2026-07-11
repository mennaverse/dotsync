import axios from "axios";

export const http = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL ?? "/",
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

export type HttpResponse<T = any> = {
  code: number;
  message: string;
  data?: T;
};
