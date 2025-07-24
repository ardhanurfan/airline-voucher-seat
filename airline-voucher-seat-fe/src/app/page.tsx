"use client";

import { useState } from "react";
import Button from "@/components/Button";
import { ZodError } from "zod";
import FieldInput from "@/components/FieldInput";
import {
  checkVoucher,
  generateVoucher,
  VoucherFormData,
  voucherSchema,
} from "../../lib/api/voucher";
import { toast } from "react-toastify";
import Dropdown from "@/components/Dropdown";
import useSWR from "swr";
import { getAircraftTypes } from "../../lib/api/aircraft";

export default function Home() {
  const { data: aircraftTypes } = useSWR("aircraft-types", getAircraftTypes);

  const [form, setForm] = useState<VoucherFormData>({
    name: "",
    id: "",
    flightNumber: "",
    date: "",
    aircraft: "",
  });

  const [result, setResult] = useState<string[]>([]);
  const [loading, setLoading] = useState(false);
  const [errors, setErrors] = useState<Record<string, string>>({});

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => setForm({ ...form, [e.target.name]: e.target.value });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    setLoading(true);
    setResult([]);
    try {
      voucherSchema.parse(form);
      const resCheck = await checkVoucher(form.flightNumber, form.date);
      if (resCheck.exists) {
        toast.error(
          `Voucher for ${form.flightNumber} on ${form.date} already exists!`
        );
        return;
      }
      const resGenerate = await generateVoucher(form);
      setResult(resGenerate.seats);

      toast.success(`Voucher for ${form.flightNumber} successfully generated!`);

      setForm({
        name: "",
        id: "",
        flightNumber: "",
        date: "",
        aircraft: "",
      });
    } catch (error) {
      if (error instanceof ZodError) {
        const newErrors: Record<string, string> = {};
        error.issues.forEach((issue) => {
          const pathElement = issue.path[0];
          if (typeof pathElement === "string") {
            newErrors[pathElement] = issue.message;
          }
        });
        setErrors(newErrors);
        return;
      }

      if (error instanceof Error) {
        toast.error(error.message);
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen">
      <div className="fixed top-0 left-0 w-full h-1/3 bg-[url('https://images.unsplash.com/photo-1652339269213-b191895f75b0')] bg-cover bg-center"></div>
      <div className="absolute w-full px-4 py-20">
        <form
          onSubmit={handleSubmit}
          className="p-8 bg-white shadow-lg rounded-2xl space-y-4 max-w-lg mx-auto"
        >
          <h2 className="text-xl md:text-3xl font-semibold text-primary mb-4">
            Generate Voucher
          </h2>

          <FieldInput
            name="name"
            label="Crew Name"
            value={form.name}
            onChange={handleChange}
            type="text"
            placeholder="Enter Crew Name"
            errorMessage={errors.name}
          />

          <FieldInput
            name="id"
            label="Crew ID"
            value={form.id}
            onChange={handleChange}
            type="text"
            placeholder="Enter Crew ID"
            errorMessage={errors.id}
          />

          <FieldInput
            name="flightNumber"
            label="Flight Number"
            value={form.flightNumber}
            onChange={handleChange}
            type="text"
            placeholder="Enter Flight Number"
            errorMessage={errors.flightNumber}
          />

          <FieldInput
            name="date"
            label="Flight Date"
            value={form.date}
            onChange={handleChange}
            type="date"
            placeholder="Enter Flight Date"
            errorMessage={errors.date}
          />

          <Dropdown
            name="aircraft"
            label={"Aircraft"}
            value={form.aircraft}
            onChange={handleChange}
            errorMessage={errors.aircraft}
            options={aircraftTypes || []}
            placeholder={"Select an Aircraft"}
          />

          <Button type="submit" variant="secondary" disabled={loading}>
            {loading ? "Generating..." : "Generate"}
          </Button>
        </form>
        {result && result.length > 0 && (
          <div className="mt-8 px-4 py-6 bg-white shadow-lg rounded-lg max-w-lg mx-auto">
            <h2 className="text-2xl font-semibold text-primary mb-2">
              Voucher Generated
            </h2>
            <p className="text-lg">
              Seats:{" "}
              <span className="font-mono">
                {result[0]}, {result[1]}, {result[2]}
              </span>
            </p>
          </div>
        )}
      </div>
    </div>
  );
}
