import { Field } from "@ark-ui/react/field";
import type { ComponentProps, ReactNode } from "react";
import clsx from "clsx/lite";

interface TextFieldProps extends ComponentProps<typeof Field.Input> {
  rootClassName?: string;
  label?: string;
  helper?: string;
  error?: string;
  suffix?: ReactNode;
}

export function TextField({
  rootClassName,
  className,
  label,
  helper,
  error,
  suffix,
  ...rest
}: TextFieldProps) {
  return (
    <Field.Root
      className={clsx(rootClassName, "flex flex-col w-full")}
      invalid={!!error}
    >
      {label && (
        <Field.Label className="text-gray-700 dark:text-gray-100 text-sm font-bold mb-2">
          {label}
        </Field.Label>
      )}
      <div
        className="flex gap-2 appearance-none border rounded w-full
                    text-gray-700 dark:text-gray-100 leading-tight focus:outline focus:shadow-outline"
      >
        <Field.Input
          className={clsx(
            className,
            "flex-1 appearance-none focus:outline rounded focus:outline-gray-500 py-2 pl-3",
          )}
          {...rest}
        />
        {suffix && (
          <div className="flex items-center justify-center">{suffix}</div>
        )}
      </div>
      {helper && (
        <Field.HelperText className="text-gray-500 dark:text-gray-400 text-xs mt-2">
          {helper}
        </Field.HelperText>
      )}
      {error && (
        <Field.ErrorText className="text-red-500 text-xs mt-2">
          <span>{error}</span>
        </Field.ErrorText>
      )}
    </Field.Root>
  );
}
