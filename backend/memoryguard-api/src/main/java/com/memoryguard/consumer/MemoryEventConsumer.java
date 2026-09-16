package com.memoryguard.consumer;

import jakarta.enterprise.context.ApplicationScoped;
import org.eclipse.microprofile.reactive.messaging.Incoming;

@ApplicationScoped
public class MemoryEventConsumer {

    @Incoming("memory-events")
    public void consume(String event) {
        System.out.println("Evento recebido do Kafka:");
        System.out.println(event);
    }
}
