import { SubmitButton } from "@/components/form/submit-button";
import { createFormHook, createFormHookContexts } from "@tanstack/react-form";
import { lazy } from "react";

const TextFormField = lazy(() =>
  import("@/components/form/form-field").then((mod) => ({
    default: mod.FormField,
  })),
);

const PasswordFormField = lazy(() =>
  import("@/components/form/password-form-field").then((mod) => ({
    default: mod.PasswordFormField,
  })),
);

export const { fieldContext, formContext, useFieldContext, useFormContext } =
  createFormHookContexts();

export const { useAppForm } = createFormHook({
  formComponents: {
    SubmitButton,
  },
  fieldComponents: {
    TextField: TextFormField,
    PasswordField: PasswordFormField,
  },
  fieldContext,
  formContext,
});
