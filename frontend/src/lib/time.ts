export const formatTime = (ts?: string | Date) => {
  if (!ts) return "-"

  const date = typeof ts === "string" ? new Date(ts) : ts

  if (isNaN(date.getTime())) return "-"

  return date.toLocaleString(undefined, {
    year: "numeric",
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
    hour12: false, // 24-hour format (matches your settings)
  })
}