import { AuthService } from "@/api/services/auth";
import { useMutation } from "@tanstack/react-query";

export function useRegisterMutation() {
  return useMutation({
    mutationFn: async (data: {
      username: string;
      email: string;
      password: string;
    }) => {
      const { username, email, password } = data;
      await AuthService.register(username, email, password);
    },
  });
}
