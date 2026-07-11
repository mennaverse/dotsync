import { AuthService } from "@/api/services/auth";
import { useQuery } from "@tanstack/react-query";

export function useCurrentUser() {
  return useQuery({
    retry: false,
    queryKey: ["user", "current"],
    queryFn: async () => {
      return (await AuthService.getCurrentUser()).data;
    },
  });
}
