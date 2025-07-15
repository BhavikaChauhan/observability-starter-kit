package com.example.payment;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class PaymentController {

    @GetMapping("/health")
    public String health() {
        return "Payment Service is healthy!";
    }

    // Add more endpoints here later...
}
