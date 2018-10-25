package com.kumuluz.ee.golang.samples.tutorial.java.service.api.v1;

import com.kumuluz.ee.cors.annotations.CrossOrigin;
import com.kumuluz.ee.discovery.annotations.RegisterService;

import javax.ws.rs.ApplicationPath;
import javax.ws.rs.core.Application;

@ApplicationPath("v1")
@CrossOrigin
@RegisterService
public class OrderApplication extends Application {
}
