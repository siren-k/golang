#!/usr/bin/env bash
mockgen -destination mocks_test.go -package asynckafka github.com/Shopify/sarama AsyncProducer