import datetime
from enum import Enum
from typing import List
from uuid import UUID
from enum import Enum
from datetime import datetime
from pydantic import BaseModel, Field
from typing import List, Optional

class Size(Enum):
    small = 'small'
    medium = 'medium'
    large = 'large'

class Status(Enum):
    created = 'created'
    cancelled = 'cancelled'
    dispatched = 'dispatched'
    delivered = 'delivered'
    progress = 'progress'

class OrderItemSchema(BaseModel):
    product: str
    size: Size
    quantity: Optional[int] = Field(default=1, ge=1, strict=True)

class CreateOrderSchema(BaseModel):
    order: List[OrderItemSchema] = Field(min_length=1)

class GetOrderSchema(CreateOrderSchema):
    id: UUID
    created: datetime
    status: Status

class GetOrdersSchema(BaseModel):
    orders: List[GetOrderSchema]
