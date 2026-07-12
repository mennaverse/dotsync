import { AuthService } from "@/api/services/auth";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export function useLogoutMutation() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async () => await AuthService.logout(),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["user", "current"] });
    },
  });
}
