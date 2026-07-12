import type { UseMutationResult } from "@tanstack/react-query";
import { Paragraph } from "@/components/typography/paragraph";
import { useTranslation } from "react-i18next";

interface ErrorMutationAreaProps {
  mutation: UseMutationResult<any, Error, any, unknown>;
  pending?: string;
  success?: string;
}

export function ErrorMutationArea({
  mutation,
  pending,
  success,
}: ErrorMutationAreaProps) {
  const { t } = useTranslation();

  return (
    <div className="flex flex-col items-center justify-center w-full">
      {pending && mutation.isPending && (
        <div className="bg-blue-100 p-4 rounded border border-blue-300 dark:bg-blue-900 dark:border-blue-700">
          <Paragraph>
            {pending ?? t("verify_email_verifying_message")}
          </Paragraph>
        </div>
      )}
      {mutation.isError && mutation.error && (
        <div className="bg-red-100 p-4 rounded border border-red-300 dark:bg-red-900 dark:border-red-700">
          <Paragraph className="text-red-500">
            {t(mutation.error.message)}
          </Paragraph>
        </div>
      )}
      {success && mutation.isSuccess && (
        <div className="bg-green-100 p-4 rounded border border-green-300 dark:bg-green-900 dark:border-green-700">
          <Paragraph className="text-green-500">
            {success ?? t("verify_email_success_message")}
          </Paragraph>
        </div>
      )}
    </div>
  );
}
