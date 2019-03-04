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

import com.kumuluz.ee.discovery.annotations.DiscoverService;
import com.kumuluz.ee.golang.samples.tutorial.java.service.persistence.exceptions.JavaServiceException;
import com.kumuluz.ee.golang.samples.tutorial.java.service.persistence.models.Order;
import com.kumuluz.ee.golang.samples.tutorial.java.service.persistence.responses.CustomerResponse;
import com.kumuluz.ee.golang.samples.tutorial.java.service.services.OrdersBean;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import javax.ws.rs.*;
import javax.ws.rs.client.WebTarget;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.Optional;


@ApplicationScoped
@Path("customers")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public class CustomerResource {
	
	@Inject
	private OrdersBean ordersBean;
	
	@Inject
	@DiscoverService(value = "go-service", version = "1.0.0", environment = "dev")
	private Optional<WebTarget> serviceUrl;
	
	// get customer for given order id
	@GET
	@Path("{orderId}")
	public Response getCustomerFromOrder(@PathParam("orderId") long orderId) {
		Order order = ordersBean.getOrderById(orderId);
		
		if (!serviceUrl.isPresent()) {
			throw new JavaServiceException("Service URL not found!", 404);
		}
		
		WebTarget apiUrl = serviceUrl.get().path("v1/customers/" + order.getCustomerId());
		
		Response response = apiUrl.request().get();
		
		if (response.getStatus() == 200) {
			CustomerResponse customerResponse = response.readEntity(CustomerResponse.class);
			return Response.status(Response.Status.OK).entity(customerResponse).build();
		} else {
			throw new JavaServiceException("Service returned error status code: " + response.getStatus(), 500);
		}
		
	}
}
