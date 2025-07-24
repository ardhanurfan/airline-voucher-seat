"use client";

import Dropdown from "@/components/Dropdown";
import FieldInput from "@/components/FieldInput";
import { Pagination } from "@/components/Pagination";
import React, { useState, useMemo } from "react";
import useSWR from "swr";
import { getAircraftTypes } from "../../../lib/api/aircraft";
import { getVouchers } from "../../../lib/api/voucher";

const LIMIT_PAGE = 10;

export default function Vouchers() {
  const { data: aircraftTypesData, error: aircraftError } = useSWR(
    "aircraft-types",
    getAircraftTypes
  );

  const { data: vouchersData, error: vouchersError } = useSWR(
    "vouchers",
    getVouchers
  );

  const [page, setPage] = useState(1);
  const [filter, setFilter] = useState({
    flight_number: "",
    crew_name: "",
    aircraft_type: "",
  });

  const aircraftTypes = useMemo(() => {
    if (aircraftTypesData) {
      return ["All", ...aircraftTypesData];
    }
    return [];
  }, [aircraftTypesData]);

  const filteredVouchers = useMemo(() => {
    if (!vouchersData) return [];

    return vouchersData.filter((voucher) => {
      return (
        (filter.flight_number
          ? voucher.flight_number
              .toLowerCase()
              .includes(filter.flight_number.toLowerCase())
          : true) &&
        (filter.crew_name
          ? voucher.crew_name
              .toLowerCase()
              .includes(filter.crew_name.toLowerCase())
          : true) &&
        (filter.aircraft_type
          ? filter.aircraft_type === "All"
            ? true
            : voucher.aircraft_type
                .toLowerCase()
                .includes(filter.aircraft_type.toLowerCase())
          : true)
      );
    });
  }, [vouchersData, filter]);

  const totalPages = useMemo(() => {
    setPage(1);
    return Math.ceil(filteredVouchers.length / LIMIT_PAGE);
  }, [filteredVouchers]);

  const handlePageChange = (newPage: number) => {
    setPage(newPage);
  };

  const handleFilterChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;
    setFilter((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const startIndex = (page - 1) * LIMIT_PAGE;
  const endIndex = page * LIMIT_PAGE;
  const currentData = filteredVouchers.slice(startIndex, endIndex);

  if (aircraftError || vouchersError) {
    return (
      <div className="w-full min-h-screen flex justify-center items-center">
        Error fetching data
      </div>
    );
  }

  return (
    <div className="min-h-screen py-20 bg-gray-50">
      <div className="p-8 bg-white shadow-lg rounded-2xl space-y-4 max-w-5xl mx-auto">
        <div className="flex space-x-4 mb-6">
          <FieldInput
            type="text"
            name="flight_number"
            placeholder="Flight Number"
            value={filter.flight_number}
            onChange={handleFilterChange}
          />
          <FieldInput
            type="text"
            name="crew_name"
            placeholder="Crew Name"
            value={filter.crew_name}
            onChange={handleFilterChange}
          />
          <Dropdown
            name="aircraft_type"
            value={filter.aircraft_type}
            onChange={handleFilterChange}
            options={aircraftTypes || []}
            placeholder={"Select an Aircraft"}
          />
        </div>

        <table className="min-w-full table-auto bg-white rounded-lg overflow-hidden">
          <thead className="bg-primary text-white">
            <tr>
              <th className="px-6 py-3 text-left">ID</th>
              <th className="px-6 py-3 text-left">Crew Name</th>
              <th className="px-6 py-3 text-left">Aircraft</th>
              <th className="px-6 py-3 text-left">Flight Number</th>
              <th className="px-6 py-3 text-left">Seats</th>
              <th className="px-6 py-3 text-left">Created At</th>
            </tr>
          </thead>
          <tbody>
            {currentData.map((voucher) => (
              <tr key={voucher.id} className="hover:bg-gray-50">
                <td className="px-6 py-3 border-b">{voucher.id}</td>
                <td className="px-6 py-3 border-b">{voucher.crew_name}</td>
                <td className="px-6 py-3 border-b">{voucher.aircraft_type}</td>
                <td className="px-6 py-3 border-b">{voucher.flight_number}</td>
                <td className="px-6 py-3 border-b">
                  {voucher.seat1}, {voucher.seat2}, {voucher.seat3}
                </td>
                <td className="px-6 py-3 border-b">
                  {new Date(voucher.created_at).toLocaleString()}
                </td>
              </tr>
            ))}
          </tbody>
        </table>

        <Pagination
          currentPage={page}
          totalPages={totalPages}
          onPageChange={handlePageChange}
        />
      </div>
    </div>
  );
}
