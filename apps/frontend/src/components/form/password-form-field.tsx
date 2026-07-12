import { ErrorText } from "./error-text";
import { PasswordField } from "./password-field";
import { useFieldContext } from "@/hooks/form";

interface PasswordFormFieldProps {
  label?: string;
}

export function PasswordFormField({ label }: PasswordFormFieldProps) {
  const field = useFieldContext<string>();

  return (
    <div className="flex flex-col mb-4 gap-1">
      <PasswordField
        name={field.name}
        label={label}
        value={field.state.value}
        onChange={(e) => field.handleChange(e.target.value)}
        onBlur={field.handleBlur}
      />
      <ErrorText field={field} />
    </div>
  );
}
