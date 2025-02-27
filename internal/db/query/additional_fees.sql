-- name: GetTotalFeesByID :one
SELECT id, 
     CAST(
        (service_fee +
         vat_fee +
         payment_processing_fee +
         delivery_fee +
         vip_fee +
         specific_event_fee) AS NUMERIC
    ) AS total_fees
FROM additional_fees
WHERE id = $1
LIMIT 1;