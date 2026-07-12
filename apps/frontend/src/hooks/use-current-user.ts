import { AuthService } from "@/api/services/auth";
import { useQuery } from "@tanstack/react-query";

export function useCurrentUser() {
  return useQuery({
    retry: false,
    queryKey: ["user", "current"],
    queryFn: async () => {
      try {
        const { data } = await AuthService.getCurrentUser();
        return data;
      } catch (error) {
        return null;
      }
    },
  });
}
