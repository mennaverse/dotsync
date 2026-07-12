import z from "zod";

export const loginFormSchema = (t: (key: string) => string) =>
  z.object({
    login: z.string().min(1, { message: t("login_required") }),
    password: z.string().min(6, { message: t("password_min_length") }),
  });
