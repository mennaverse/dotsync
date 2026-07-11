import { http, type HttpResponse } from "../http";
import type { ClaimsDTO } from "../dtos/auth";

export class AuthService {
  static async register(username: string, email: string, password: string) {
    const response = await http.post("/api/auth/register", {
      username,
      email,
      password,
    });
    return response.data;
  }

  static async login(email: string, password: string) {
    const response = await http.post("/api/auth/login", { email, password });
    return response.data;
  }

  static async logout() {
    const response = await http.post("/api/auth/logout");
    return response.data;
  }

  static async logoutAll() {
    const response = await http.post("/api/auth/logout-all");
    return response.data;
  }

  static async verifyEmail(token: string) {
    const response = await http.post("/api/auth/verify-email", { token });
    return response.data;
  }

  static async resendVerificationEmail(email: string) {
    const response = await http.post("/api/auth/resend-verification-email", {
      email,
    });
    return response.data;
  }

  static async forgotPassword(email: string) {
    const response = await http.post("/api/auth/forgot-password", { email });
    return response.data;
  }

  static async resetPassword(token: string, newPassword: string) {
    const response = await http.post("/api/auth/reset-password", {
      token,
      newPassword,
    });
    return response.data;
  }

  static async validateAccessToken() {
    const response = await http.get<HttpResponse<ClaimsDTO>>(
      "/api/auth/validate-access-token",
    );
    return response.data;
  }

  static async getCurrentUser() {
    return await AuthService.validateAccessToken();
  }
}
