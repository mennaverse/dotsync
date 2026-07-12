import { ErrorMutationArea } from "@/components/form/error-mutation-area";
import { Page } from "@/components/page";
import { Heading } from "@/components/typography/heading";
import { useVerifyEmailMutation } from "@/hooks/use-verify-email-mutation";
import { createFileRoute } from "@tanstack/react-router";
import { useEffect } from "react";
import { useTranslation } from "react-i18next";
import z from "zod";

const searchSchema = z.object({
  token: z.string().optional(),
});

export const Route = createFileRoute("/verify-email")({
  component: RouteComponent,
  validateSearch: searchSchema,
});

function RouteComponent() {
  const { t } = useTranslation();
  const { token: searchToken } = Route.useSearch();
  const verifyEmailMutation = useVerifyEmailMutation();

  useEffect(() => {
    if (searchToken) {
      verifyEmail(searchToken);
    }
  }, [searchToken]);

  const verifyEmail = async (token: string) => {
    await verifyEmailMutation.mutateAsync(token);
  };

  return (
    <Page contained>
      <div className="flex flex-col items-center justify-center w-full my-auto gap-8 p-8">
        <Heading className="text-2xl font-bold mb-4 text-teal-800 dark:text-teal-400 mb-4">
          {t("verify_email_title")}
        </Heading>
        <ErrorMutationArea
          mutation={verifyEmailMutation}
          pending={t("verify_email_verifying_message")}
          success={t("verify_email_success_message")}
        />
      </div>
    </Page>
  );
}
