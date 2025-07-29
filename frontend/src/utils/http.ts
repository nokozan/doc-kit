export type AxiosLikeError = {
  message: string;
  response?: {
    data?: {
      error?: string;
    };
  };
};

export function isAxiosLikeError(err: unknown): err is AxiosLikeError {
  return (
    typeof err === "object" &&
    err !== null &&
    "message" in err &&
    typeof (err as Record<string, unknown>).message === "string"
  );
}

export function extractErrorMessage(err: unknown): string {
  if (isAxiosLikeError(err)) {
    return err.response?.data?.error || err.message;
  }
  return "An unexpected error occurred.";
}
