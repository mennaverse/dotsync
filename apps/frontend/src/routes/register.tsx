import { Page } from "@/components/page";
import { ErrorMutationArea } from "@/components/form/error-mutation-area";
import { Paragraph } from "@/components/typography/paragraph";
import { useAppForm } from "@/hooks/form";
import { useRegisterMutation } from "@/hooks/use-register-mutation";
import { registerFormSchema } from "@/schemas/register";
import { createFileRoute, Link, useRouter } from "@tanstack/react-router";
import { useTranslation } from "react-i18next";
import { useDocumentTitle } from "usehooks-ts";

export const Route = createFileRoute("/register")({
  component: RouteComponent,
});

function RouteComponent() {
  const { t } = useTranslation();
  const router = useRouter();
  const registerMutation = useRegisterMutation();

  useDocumentTitle(t("register_doctitle"));

  const form = useAppForm({
    defaultValues: {
      username: "",
      email: "",
      password: "",
    },
    validators: {
      onChange: registerFormSchema(t),
    },
    onSubmit: async ({ value }) => {
      await registerMutation.mutateAsync(value, {
        onSuccess: async () => {
          await router.navigate({
            to: "/login",
            search: { registered: true },
          });
        },
      });
    },
  });

  return (
    <Page className="flex flex-col items-center justify-center w-full my-auto gap-8 p-8">
      <ErrorMutationArea mutation={registerMutation} />
      <form
        className="w-full max-w-lg flex flex-col justify-center"
        onSubmit={(e) => {
          e.preventDefault();
          e.stopPropagation();
          form.handleSubmit();
        }}
      >
        <h1 className="text-2xl font-bold mb-4 text-teal-800 dark:text-teal-400 mb-4">
          {t("register_title")}
        </h1>
        <div className="flex flex-col gap-2">
          <form.AppField
            name="username"
            children={(field) => (
              <field.TextField
                label={t("register_username_label")}
                autoComplete="username"
              />
            )}
          />
          <form.AppField
            name="email"
            children={(field) => (
              <field.TextField
                label={t("register_email_label")}
                autoComplete="email"
              />
            )}
          />
          <form.AppField
            name="password"
            children={(field) => (
              <field.PasswordField label={t("register_password_label")} />
            )}
          />
          <form.AppForm>
            <form.SubmitButton
              label={t("register_submit_label")}
              labelSubmitting={t("register_submit_loading")}
            />
          </form.AppForm>
        </div>
      </form>
      <div className="w-full max-w-lg flex justify-center gap-2">
        <Paragraph>{t("register_login_description")}</Paragraph>
        <Link to="/login" className="underline">
          <Paragraph className="underline text-teal-700 dark:text-teal-400">
            {t("register_login_button")}
          </Paragraph>
        </Link>
      </div>
    </Page>
  );
}
