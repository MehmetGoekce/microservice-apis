from datetime import datetime
from uuid import UUID
from http import HTTPStatus

from starlette.responses import Response
from starlette import status

from orders.app import app
from orders.api.schemas import CreateOrderSchema
order = {
    'id': 'cb9082f7-3aa3-4278-a602-cc3b125733a8',
    'status': "delivered",
    'created': datetime.today(),
    'order': [
        {
            'product': 'Kaffee',
            'size': 'medium',
            'quantity': 1
        }
    ]
}


@app.get('/orders')
def get_orders():
    return {'order': [order]}


@app.post('/orders', status_code=status.HTTP_201_CREATED)
def create_order(order_details: CreateOrderSchema):
    return order


@app.get('/orders/{order_id}')
def get_order(order_id: UUID):
    return order


@app.put('/orders/{order_id}')
def update_order(order_id: UUID, order_details: CreateOrderSchema):
    return order


@app.delete('/orders/{order_id}', status_code=status.HTTP_204_NO_CONTENT)
def delete_order(order_id: UUID):
    return Response(status_code=HTTPStatus.NO_CONTENT.value)


@app.post('/orders/{order_id}/cancel')
def cancel_order(order_id: UUID):
    return order


@app.post('/orders/{order_id}/pay')
def pay_order(order_id: UUID):
    return order
