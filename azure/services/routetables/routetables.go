/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package routetables

import (
	"context"

	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	"sigs.k8s.io/cluster-api-provider-azure/azure"
	"sigs.k8s.io/cluster-api-provider-azure/azure/services/async"
	"sigs.k8s.io/cluster-api-provider-azure/util/reconciler"
	"sigs.k8s.io/cluster-api-provider-azure/util/tele"
)

const serviceName = "routetables"

// RouteTableScope defines the scope interface for route table service.
type RouteTableScope interface {
	azure.Authorizer
	azure.AsyncStatusUpdater
	RouteTableSpecs() []azure.ResourceSpecGetter
	IsVnetManaged() bool
}

// Service provides operations on azure resources.
type Service struct {
	Scope RouteTableScope
	async.Reconciler
}

// New creates a new service.
func New(scope RouteTableScope) *Service {
	client := newClient(scope)
	return &Service{
		Scope:      scope,
		Reconciler: async.New(scope, client, client),
	}
}

// Reconcile gets/creates/updates route tables.
func (s *Service) Reconcile(ctx context.Context) error {
	ctx, log, done := tele.StartSpanWithLogger(ctx, "routetables.Service.Reconcile")
	defer done()

	ctx, cancel := context.WithTimeout(ctx, reconciler.DefaultAzureServiceReconcileTimeout)
	defer cancel()

	var resErr error

	if !s.Scope.IsVnetManaged() {
		log.V(4).Info("Skipping route tables reconcile in custom vnet mode")
	} else {
		// We go through the list of route tables to reconcile each one, independently of the result of the previous one.
		// If multiple errors occur, we return the most pressing one.
		//  Order of precedence (highest -> lowest) is: error that is not an operationNotDoneError (ie. error creating) -> operationNotDoneError (ie. creating in progress) -> no error (ie. created)
		for _, rtSpec := range s.Scope.RouteTableSpecs() {
			if _, err := s.CreateResource(ctx, rtSpec, serviceName); err != nil {
				if !azure.IsOperationNotDoneError(err) || resErr == nil {
					resErr = err
				}
			}
		}
	}
	s.Scope.UpdatePutStatus(infrav1.RouteTablesReadyCondition, serviceName, resErr)
	return resErr
}

// Delete deletes route tables.
func (s *Service) Delete(ctx context.Context) error {
	ctx, log, done := tele.StartSpanWithLogger(ctx, "routetables.Service.Delete")
	defer done()

	ctx, cancel := context.WithTimeout(ctx, reconciler.DefaultAzureServiceReconcileTimeout)
	defer cancel()

	// Only delete the route tables if their lifecycle is managed by this controller.
	// route tables are managed if and only if the vnet is managed.
	if !s.Scope.IsVnetManaged() {
		log.V(4).Info("Skipping route table deletion in custom vnet mode")
		return nil
	}

	var result error

	// We go through the list of RouteTableSpecs to delete each one, independently of the result of the previous one.
	// If multiple erros occur, we return the most pressing one
	// order of precedence is: error deleting -> deleting in progress -> deleted (no error)
	for _, rtSpec := range s.Scope.RouteTableSpecs() {
		if err := s.DeleteResource(ctx, rtSpec, serviceName); err != nil {
			if !azure.IsOperationNotDoneError(err) || result == nil {
				result = err
			}
		}
	}
	s.Scope.UpdateDeleteStatus(infrav1.RouteTablesReadyCondition, serviceName, result)
	return result
}
