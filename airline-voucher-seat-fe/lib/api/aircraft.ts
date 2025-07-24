import { apiWrapper } from "../apiWrapper";

export const getAircraftTypes = async () => {
  try {
    const res = await apiWrapper("aircrafts", "GET");
    return res.data as string[];
  } catch (error) {
    throw error;
  }
};
