/*
 *  Copyright (c) 2019 Kumuluz and/or its affiliates
 *  and other contributors as indicated by the @author tags and
 *  the contributor list.
 *
 *  Licensed under the MIT License (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  https://opensource.org/licenses/MIT
 *
 *  The software is provided "AS IS", WITHOUT WARRANTY OF ANY KIND, express or
 *  implied, including but not limited to the warranties of merchantability,
 *  fitness for a particular purpose and noninfringement. in no event shall the
 *  authors or copyright holders be liable for any claim, damages or other
 *  liability, whether in an action of contract, tort or otherwise, arising from,
 *  out of or in connection with the software or the use or other dealings in the
 *  software. See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package com.kumuluz.ee.golang.samples.tutorial.java.service.api.v1.rest;

import com.kumuluz.ee.golang.samples.tutorial.java.service.persistence.models.Order;
import com.kumuluz.ee.golang.samples.tutorial.java.service.persistence.requests.OrderRequest;
import com.kumuluz.ee.golang.samples.tutorial.java.service.services.OrdersBean;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.List;

@ApplicationScoped
@Path("orders")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public class OrderResource {
	
	@Inject
	private OrdersBean ordersBean;
	
	// get all orders from customer with given id
	@GET
	@Path("customer/{customerId}")
	public Response getOrdersFromCustomer(@PathParam("customerId") long customerId) {
		List<Order> orders = ordersBean.getAllOrdersFromCustomer(customerId);
		return Response.status(Response.Status.OK).entity(orders).build();
	}
	
	// create new order
	@POST
	public Response createOrderForCustomer(OrderRequest newOrder) {
		Order order = new Order();
		order.setCustomerId(newOrder.getCustomerId());
		order.setTitle(newOrder.getTitle());
		order.setDescription(newOrder.getDescription());
		
		ordersBean.createOrder(order);
		
		return Response.status(Response.Status.CREATED).entity(order).build();
	}
}
