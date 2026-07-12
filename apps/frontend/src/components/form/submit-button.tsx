import { Button } from "../button";
import { useFormContext } from "@/hooks/form";

interface SubmitButtonProps {
  labelSubmitting: string;
  label: string;
}

export function SubmitButton({ label, labelSubmitting }: SubmitButtonProps) {
  const form = useFormContext();

  return (
    <form.Subscribe
      selector={(state) => [state.isSubmitting, state.canSubmit]}
      children={([isSubmitting, canSubmit]) => (
        <Button
          type="submit"
          className="mt-4"
          disabled={isSubmitting || !canSubmit}
        >
          {isSubmitting ? labelSubmitting : label}
        </Button>
      )}
    />
  );
}
