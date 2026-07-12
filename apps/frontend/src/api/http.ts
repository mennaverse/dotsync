import axios from "axios";

export const http = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL ?? "/",
  withCredentials: true,
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
});

const t = (key: string) => key;
export const errorMessages: string[] = [
  t("http_error_user_banned"),
  t("http_error_invalid_credentials"),
  t("http_error_user_not_found"),
  t("http_error_invalid_refresh_token"),
  t("http_error_verification_token_expired"),
  t("http_error_invalid_token"),
  t("http_error_refresh_token_expired"),
  t("http_error_email_already_verified"),
  t("http_error_email_verification_disabled"),
  t("http_error_reset_password_token_expired"),
  t("http_error_reset_password_token_invalid"),
  t("http_error_invalid_refresh_token"),
  t("http_error_email_or_username_already_exists"),
  t("http_error_missing_token"),
  t("http_error_password_too_long"),
  t("http_error_internal_server_error"),
  t("http_error_unknown_error"),
];

http.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response) {
      const data = error.response.data as HttpResponse;
      return Promise.reject(
        new Error(`http_error_${data?.code ?? "unknown_error"}`),
      );
    }

    return Promise.reject(error);
  },
);

export type HttpResponse<T = any> = {
  code: string;
  message: string;
  data?: T;
};
