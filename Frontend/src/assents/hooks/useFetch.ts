import { useState, useEffect } from "react";

export function useFetch(url: string) {
  const [data, setData] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [controller, setController] = useState<AbortController | null>(null);

  useEffect(() => {
    const abortController = new AbortController();
    setController(abortController);

    fetch(url, { signal: abortController.signal })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Error al obtener datos");
        }
        return response.json();
      })
      .then((json) => setData(json))
      .catch((error) => {
        if (error.name === "AbortError") {
          console.log("Petición cancelada");
        } else {
          setError(error.message);
        }
      })
      .finally(() => setLoading(false));

    return () => abortController.abort();
  }, [url]);

  const handleCancelRequest = () => {
    if (controller) {
      controller.abort();
      setError("Petición cancelada");
      setLoading(false);
    }
  };

  return { data, loading, error, handleCancelRequest };
}
