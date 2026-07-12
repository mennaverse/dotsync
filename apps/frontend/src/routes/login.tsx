import { ErrorMutationArea } from "@/components/form/error-mutation-area";
import { Page } from "@/components/page";
import { Paragraph } from "@/components/typography/paragraph";
import { useAppForm } from "@/hooks/form";
import { useLoginMutation } from "@/hooks/use-login-mutation";
import { loginFormSchema } from "@/schemas/login";
import { createFileRoute, Link, useNavigate } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";
import { useDocumentTitle } from "usehooks-ts";
import z from "zod";

const searchSchema = z.object({
  registered: z.boolean().optional(),
});

export const Route = createFileRoute("/login")({
  component: RouteComponent,
  validateSearch: searchSchema,
});

function RouteComponent() {
  const { t } = useTranslation();
  const { registered } = Route.useSearch();
  const navigate = useNavigate();
  const loginMutation = useLoginMutation();

  useDocumentTitle(t("login_doctitle"));

  const form = useAppForm({
    defaultValues: {
      login: "",
      password: "",
    },
    validators: {
      onChange: loginFormSchema(t),
    },
    onSubmit: async ({ value }) => {
      const { login, password } = value;
      await loginMutation.mutateAsync(
        { login, password },
        {
          onSuccess: async () => {
            await navigate({ to: "/home" });
          },
        },
      );
    },
  });

  return (
    <Page className="flex flex-col items-center justify-center w-full my-auto gap-8 p-8">
      {registered && (
        <div className="w-full max-w-lg flex justify-center gap-2 bg-green-300 dark:bg-green-900 p-4 rounded-md border border-green-700">
          <Paragraph className="text-green-700 dark:text-green-300">
            {t("login_registered_success")}
          </Paragraph>
        </div>
      )}
      <ErrorMutationArea mutation={loginMutation} />
      <form
        className="w-full max-w-lg flex flex-col justify-center"
        onSubmit={(e) => {
          e.preventDefault();
          e.stopPropagation();
          form.handleSubmit();
        }}
      >
        <h1 className="text-2xl font-bold mb-4 text-teal-800 dark:text-teal-400 mb-4">
          {t("login_title")}
        </h1>
        <form.AppField
          name="login"
          children={(field) => (
            <field.TextField label={t("login_login_label")} />
          )}
        />
        <form.AppField
          name="password"
          children={(field) => (
            <field.PasswordField label={t("login_password_label")} />
          )}
        />
        <form.AppForm>
          <form.SubmitButton
            label={t("login_submit_label")}
            labelSubmitting={t("login_submit_loading")}
          />
        </form.AppForm>
      </form>
      <div className="w-full max-w-lg flex justify-center gap-2">
        <Paragraph>{t("login_register_description")}</Paragraph>
        <Link to="/register" className="underline">
          <Paragraph className="underline text-teal-700 dark:text-teal-400">
            {t("login_register_button")}
          </Paragraph>
        </Link>
      </div>
    </Page>
  );
}
