import z from "zod";

export const registerFormSchema = (t: (key: string) => string) =>
  z.object({
    username: z.string().min(3, t("username_min_length")),
    email: z.email(t("invalid_email")),
    password: z
      .string()
      .min(8, t("password_min_length"))
      .max(64, t("password_max_length")),
  });
