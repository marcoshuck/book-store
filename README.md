# Book Store
An event-driven application using `Temporal` for event mediation.

## Place Order Workflow
1. Place order
    1. Create the order
2. Process order
   1. Email customer that order has been placed
   2. Apply payment
   3. Decrement inventory
3. Fulfill order
    1. Pick and pack order
    2. Order more stock if necessary from supplier
4. Ship order
    1. Email customer that order is ready for shipment
    2. Ship order to customer
5. Notify customer
    1. Email customer that order has been shipped