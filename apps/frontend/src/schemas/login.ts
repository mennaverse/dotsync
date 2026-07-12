import z from "zod";

export const loginFormSchema = (t: (key: string) => string) =>
  z.object({
    login: z.string().min(1, { message: t("validation_login_required") }),
    password: z
      .string()
      .min(6, { message: t("validation_password_min_length") }),
  });
