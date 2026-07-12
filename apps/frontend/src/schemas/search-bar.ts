import { z } from "zod";

export const searchBarFormSchema = (t: (key: string) => string) =>
  z.object({
    search: z.string().min(1, { message: t("search_term_required") }),
  });
