import z from "zod";
import { apiWrapper } from "../apiWrapper";

export const voucherSchema = z.object({
  name: z.string().min(1, "Crew Name is required"),
  id: z.string().min(1, "Crew ID is required"),
  flightNumber: z.string().min(1, "Flight Number is required"),
  date: z.string().min(1, "Flight Date is required"),
  aircraft: z.string().min(1, "Aircraft Type is required"),
});

export type VoucherFormData = z.infer<typeof voucherSchema>;

interface GeneratedResponse {
  seats: string[];
}

interface CheckResponse {
  exists: boolean;
}

export interface Voucher {
  id: number;
  crew_name: string;
  crew_id: string;
  flight_number: string;
  flight_date: string;
  aircraft_type: string;
  seat1: string;
  seat2: string;
  seat3: string;
  created_at: string;
}

export const generateVoucher = async (form: VoucherFormData) => {
  try {
    const res = await apiWrapper("generate", "POST", form);
    return res.data as GeneratedResponse;
  } catch (error) {
    throw error;
  }
};

export const checkVoucher = async (flightNumber: string, date: string) => {
  try {
    const res = await apiWrapper("check", "POST", { flightNumber, date });
    return res.data as CheckResponse;
  } catch (error) {
    throw error;
  }
};

export const getVouchers = async () => {
  try {
    const res = await apiWrapper("vouchers", "GET");
    return res.data as Voucher[];
  } catch (error) {
    throw error;
  }
};
