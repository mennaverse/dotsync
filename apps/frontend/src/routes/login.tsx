import { loginFormSchema } from "@/schemas/login";
import { useForm } from "@tanstack/react-form";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/login")({
  component: RouteComponent,
});

function RouteComponent() {
  // TODO: login
  const form = useForm({
    defaultValues: {
      login: "",
      password: "",
    },
    validators: {
      onChange: loginFormSchema,
    },
  });

  return (
    <div>
      <h1>Login Page</h1>
      <p>This is the login page of our application.</p>
    </div>
  );
}
