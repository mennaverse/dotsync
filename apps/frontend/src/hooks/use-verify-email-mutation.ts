import { AuthService } from "@/api/services/auth";
import { useMutation } from "@tanstack/react-query";

export function useVerifyEmailMutation() {
  return useMutation({
    retry: false,
    mutationFn: async (token: string) => {
      const response = await AuthService.verifyEmail(token);
      return response;
    },
  });
}
