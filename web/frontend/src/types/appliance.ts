export interface Appliance {
  id: number
  user_id?: number
  name: string
  location: string
  status: "online" | "offline"
  last_reading?: string | null
  device_code: string
  created_at?: string
  updated_at?: string
}