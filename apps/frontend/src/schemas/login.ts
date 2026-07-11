import z from "zod";

export const loginFormSchema = z.object({
  login: z.string().min(1, { message: "Login is required" }),
  password: z
    .string()
    .min(6, { message: "Password must be at least 6 characters" }),
});
