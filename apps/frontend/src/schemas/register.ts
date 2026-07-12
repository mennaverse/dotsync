import z from "zod";

export const registerFormSchema = (t: (key: string) => string) =>
  z.object({
    username: z.string().min(3, t("validation_username_min_length")),
    email: z.email(t("validation_invalid_email")),
    password: z
      .string()
      .min(8, t("validation_password_min_length"))
      .max(64, t("validation_password_max_length")),
  });
