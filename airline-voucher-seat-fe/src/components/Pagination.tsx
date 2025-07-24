import React from "react";
import clsx from "clsx";

export interface PaginationProps {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
  maxPageNumbers?: number;
}

export function Pagination({
  currentPage,
  totalPages,
  onPageChange,
  maxPageNumbers = 10,
}: PaginationProps) {
  if (totalPages <= 1) return null;

  const handlePrev = () => {
    if (currentPage > 1) {
      onPageChange(currentPage - 1);
    }
  };

  const handleNext = () => {
    if (currentPage < totalPages) {
      onPageChange(currentPage + 1);
    }
  };

  const pageNumbers: number[] = [];
  for (let i = 1; i <= totalPages; i++) {
    pageNumbers.push(i);
  }

  let startPage = Math.max(currentPage - Math.floor(maxPageNumbers / 2), 1);
  let endPage = startPage + maxPageNumbers - 1;
  if (endPage > totalPages) {
    endPage = totalPages;
    startPage = Math.max(endPage - maxPageNumbers + 1, 1);
  }

  const visiblePages = pageNumbers.slice(startPage - 1, endPage);

  return (
    <div className="flex items-center justify-center gap-2 py-3 mt-4">
      <button
        onClick={handlePrev}
        disabled={currentPage === 1}
        className={clsx(
          "flex justify-center items-center h-8 w-8 rounded",
          currentPage === 1
            ? "bg-gray-300 text-white cursor-not-allowed"
            : "bg-primary text-white hover:bg-primary/80"
        )}
      >
        &lt;
      </button>

      {visiblePages.map((page) => (
        <button
          key={page}
          onClick={() => onPageChange(page)}
          className={clsx(
            "flex justify-center items-center h-8 w-8 rounded",
            page === currentPage
              ? "bg-secondary text-white"
              : "hover:bg-secondary hover:text-white"
          )}
        >
          {page}
        </button>
      ))}

      <button
        onClick={handleNext}
        disabled={currentPage === totalPages}
        className={clsx(
          "flex justify-center items-center h-8 w-8 rounded",
          currentPage === totalPages
            ? "bg-gray-300 text-white cursor-not-allowed"
            : "bg-primary text-white hover:bg-primary/80"
        )}
      >
        &gt;
      </button>
    </div>
  );
}
