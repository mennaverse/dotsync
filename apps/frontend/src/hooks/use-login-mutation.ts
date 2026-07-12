import { AuthService } from "@/api/services/auth";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export function useLoginMutation() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async ({
      login,
      password,
    }: {
      login: string;
      password: string;
    }) => await AuthService.login(login, password),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["user", "current"] });
    },
  });
}
