import { z } from "zod";

export const searchBarFormSchema = z.object({
  search: z.string().min(1, { message: "Search term is required" }),
});
