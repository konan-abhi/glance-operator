/*
Copyright 2022.

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

package controllers

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	//+kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

// var cfg *rest.Config
// var err error
// var k8sClient client.Client
// var testEnv *envtest.Environment

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{printer.NewlineReporter{}})
}

// var _ = BeforeSuite(func() {
// 	logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter), zap.UseDevMode(true)))

// 	By("bootstrapping test environment")
// 	testEnv = &envtest.Environment{
// 		CRDDirectoryPaths:     []string{filepath.Join("..", "config", "crd", "bases")},
// 		ErrorIfCRDPathMissing: true,
// 	}

// 	cfg, err = testEnv.Start()
// 	Expect(err).NotTo(HaveOccurred())
// 	Expect(cfg).NotTo(BeNil())

// 	err = glancev1beta1.AddToScheme(scheme.Scheme)
// 	Expect(err).NotTo(HaveOccurred())

// 	//+kubebuilder:scaffold:scheme

// 	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
// 	Expect(err).NotTo(HaveOccurred())
// 	Expect(k8sClient).NotTo(BeNil())

// }, 60)

// var _ = AfterSuite(func() {
// 	By("tearing down the test environment")
// 	err := testEnv.Stop()
// 	Expect(err).NotTo(HaveOccurred())
// })