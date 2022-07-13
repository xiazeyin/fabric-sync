// Code generated by counterfeiter. DO NOT EDIT.
package resmgmt

import (
	reqContext "context"
	"sync"

	"github.com/xiazeyin/fabric-protos-go-gm/common"
	"github.com/xiazeyin/fabric-sdk-go-gm/pkg/common/providers/fab"
	"github.com/xiazeyin/fabric-sdk-go-gm/pkg/fab/resource"
)

type MockLifecycleResource struct {
	InstallStub        func(reqCtx reqContext.Context, installPkg []byte, targets []fab.ProposalProcessor, opts ...resource.Opt) ([]*resource.LifecycleInstallProposalResponse, error)
	installMutex       sync.RWMutex
	installArgsForCall []struct {
		reqCtx     reqContext.Context
		installPkg []byte
		targets    []fab.ProposalProcessor
		opts       []resource.Opt
	}
	installReturns struct {
		result1 []*resource.LifecycleInstallProposalResponse
		result2 error
	}
	installReturnsOnCall map[int]struct {
		result1 []*resource.LifecycleInstallProposalResponse
		result2 error
	}
	GetInstalledPackageStub        func(reqCtx reqContext.Context, packageID string, target fab.ProposalProcessor, opts ...resource.Opt) ([]byte, error)
	getInstalledPackageMutex       sync.RWMutex
	getInstalledPackageArgsForCall []struct {
		reqCtx    reqContext.Context
		packageID string
		target    fab.ProposalProcessor
		opts      []resource.Opt
	}
	getInstalledPackageReturns struct {
		result1 []byte
		result2 error
	}
	getInstalledPackageReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	QueryInstalledStub        func(reqCtx reqContext.Context, target fab.ProposalProcessor, opts ...resource.Opt) (*resource.LifecycleQueryInstalledCCResponse, error)
	queryInstalledMutex       sync.RWMutex
	queryInstalledArgsForCall []struct {
		reqCtx reqContext.Context
		target fab.ProposalProcessor
		opts   []resource.Opt
	}
	queryInstalledReturns struct {
		result1 *resource.LifecycleQueryInstalledCCResponse
		result2 error
	}
	queryInstalledReturnsOnCall map[int]struct {
		result1 *resource.LifecycleQueryInstalledCCResponse
		result2 error
	}
	QueryApprovedStub        func(reqCtx reqContext.Context, channelID string, req *resource.QueryApprovedChaincodeRequest, target fab.ProposalProcessor, opts ...resource.Opt) (*resource.LifecycleQueryApprovedCCResponse, error)
	queryApprovedMutex       sync.RWMutex
	queryApprovedArgsForCall []struct {
		reqCtx    reqContext.Context
		channelID string
		req       *resource.QueryApprovedChaincodeRequest
		target    fab.ProposalProcessor
		opts      []resource.Opt
	}
	queryApprovedReturns struct {
		result1 *resource.LifecycleQueryApprovedCCResponse
		result2 error
	}
	queryApprovedReturnsOnCall map[int]struct {
		result1 *resource.LifecycleQueryApprovedCCResponse
		result2 error
	}
	CreateApproveProposalStub        func(txh fab.TransactionHeader, req *resource.ApproveChaincodeRequest) (*fab.TransactionProposal, error)
	createApproveProposalMutex       sync.RWMutex
	createApproveProposalArgsForCall []struct {
		txh fab.TransactionHeader
		req *resource.ApproveChaincodeRequest
	}
	createApproveProposalReturns struct {
		result1 *fab.TransactionProposal
		result2 error
	}
	createApproveProposalReturnsOnCall map[int]struct {
		result1 *fab.TransactionProposal
		result2 error
	}
	CreateCheckCommitReadinessProposalStub        func(txh fab.TransactionHeader, req *resource.CheckChaincodeCommitReadinessRequest) (*fab.TransactionProposal, error)
	createCheckCommitReadinessProposalMutex       sync.RWMutex
	createCheckCommitReadinessProposalArgsForCall []struct {
		txh fab.TransactionHeader
		req *resource.CheckChaincodeCommitReadinessRequest
	}
	createCheckCommitReadinessProposalReturns struct {
		result1 *fab.TransactionProposal
		result2 error
	}
	createCheckCommitReadinessProposalReturnsOnCall map[int]struct {
		result1 *fab.TransactionProposal
		result2 error
	}
	CreateCommitProposalStub        func(txh fab.TransactionHeader, req *resource.CommitChaincodeRequest) (*fab.TransactionProposal, error)
	createCommitProposalMutex       sync.RWMutex
	createCommitProposalArgsForCall []struct {
		txh fab.TransactionHeader
		req *resource.CommitChaincodeRequest
	}
	createCommitProposalReturns struct {
		result1 *fab.TransactionProposal
		result2 error
	}
	createCommitProposalReturnsOnCall map[int]struct {
		result1 *fab.TransactionProposal
		result2 error
	}
	CreateQueryCommittedProposalStub        func(txh fab.TransactionHeader, req *resource.QueryCommittedChaincodesRequest) (*fab.TransactionProposal, error)
	createQueryCommittedProposalMutex       sync.RWMutex
	createQueryCommittedProposalArgsForCall []struct {
		txh fab.TransactionHeader
		req *resource.QueryCommittedChaincodesRequest
	}
	createQueryCommittedProposalReturns struct {
		result1 *fab.TransactionProposal
		result2 error
	}
	createQueryCommittedProposalReturnsOnCall map[int]struct {
		result1 *fab.TransactionProposal
		result2 error
	}
	UnmarshalApplicationPolicyStub        func(policyBytes []byte) (*common.SignaturePolicyEnvelope, string, error)
	unmarshalApplicationPolicyMutex       sync.RWMutex
	unmarshalApplicationPolicyArgsForCall []struct {
		policyBytes []byte
	}
	unmarshalApplicationPolicyReturns struct {
		result1 *common.SignaturePolicyEnvelope
		result2 string
		result3 error
	}
	unmarshalApplicationPolicyReturnsOnCall map[int]struct {
		result1 *common.SignaturePolicyEnvelope
		result2 string
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MockLifecycleResource) Install(reqCtx reqContext.Context, installPkg []byte, targets []fab.ProposalProcessor, opts ...resource.Opt) ([]*resource.LifecycleInstallProposalResponse, error) {
	var installPkgCopy []byte
	if installPkg != nil {
		installPkgCopy = make([]byte, len(installPkg))
		copy(installPkgCopy, installPkg)
	}
	var targetsCopy []fab.ProposalProcessor
	if targets != nil {
		targetsCopy = make([]fab.ProposalProcessor, len(targets))
		copy(targetsCopy, targets)
	}
	fake.installMutex.Lock()
	ret, specificReturn := fake.installReturnsOnCall[len(fake.installArgsForCall)]
	fake.installArgsForCall = append(fake.installArgsForCall, struct {
		reqCtx     reqContext.Context
		installPkg []byte
		targets    []fab.ProposalProcessor
		opts       []resource.Opt
	}{reqCtx, installPkgCopy, targetsCopy, opts})
	fake.recordInvocation("Install", []interface{}{reqCtx, installPkgCopy, targetsCopy, opts})
	fake.installMutex.Unlock()
	if fake.InstallStub != nil {
		return fake.InstallStub(reqCtx, installPkg, targets, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.installReturns.result1, fake.installReturns.result2
}

func (fake *MockLifecycleResource) InstallCallCount() int {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return len(fake.installArgsForCall)
}

func (fake *MockLifecycleResource) InstallArgsForCall(i int) (reqContext.Context, []byte, []fab.ProposalProcessor, []resource.Opt) {
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	return fake.installArgsForCall[i].reqCtx, fake.installArgsForCall[i].installPkg, fake.installArgsForCall[i].targets, fake.installArgsForCall[i].opts
}

func (fake *MockLifecycleResource) InstallReturns(result1 []*resource.LifecycleInstallProposalResponse, result2 error) {
	fake.InstallStub = nil
	fake.installReturns = struct {
		result1 []*resource.LifecycleInstallProposalResponse
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) InstallReturnsOnCall(i int, result1 []*resource.LifecycleInstallProposalResponse, result2 error) {
	fake.InstallStub = nil
	if fake.installReturnsOnCall == nil {
		fake.installReturnsOnCall = make(map[int]struct {
			result1 []*resource.LifecycleInstallProposalResponse
			result2 error
		})
	}
	fake.installReturnsOnCall[i] = struct {
		result1 []*resource.LifecycleInstallProposalResponse
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) GetInstalledPackage(reqCtx reqContext.Context, packageID string, target fab.ProposalProcessor, opts ...resource.Opt) ([]byte, error) {
	fake.getInstalledPackageMutex.Lock()
	ret, specificReturn := fake.getInstalledPackageReturnsOnCall[len(fake.getInstalledPackageArgsForCall)]
	fake.getInstalledPackageArgsForCall = append(fake.getInstalledPackageArgsForCall, struct {
		reqCtx    reqContext.Context
		packageID string
		target    fab.ProposalProcessor
		opts      []resource.Opt
	}{reqCtx, packageID, target, opts})
	fake.recordInvocation("GetInstalledPackage", []interface{}{reqCtx, packageID, target, opts})
	fake.getInstalledPackageMutex.Unlock()
	if fake.GetInstalledPackageStub != nil {
		return fake.GetInstalledPackageStub(reqCtx, packageID, target, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getInstalledPackageReturns.result1, fake.getInstalledPackageReturns.result2
}

func (fake *MockLifecycleResource) GetInstalledPackageCallCount() int {
	fake.getInstalledPackageMutex.RLock()
	defer fake.getInstalledPackageMutex.RUnlock()
	return len(fake.getInstalledPackageArgsForCall)
}

func (fake *MockLifecycleResource) GetInstalledPackageArgsForCall(i int) (reqContext.Context, string, fab.ProposalProcessor, []resource.Opt) {
	fake.getInstalledPackageMutex.RLock()
	defer fake.getInstalledPackageMutex.RUnlock()
	return fake.getInstalledPackageArgsForCall[i].reqCtx, fake.getInstalledPackageArgsForCall[i].packageID, fake.getInstalledPackageArgsForCall[i].target, fake.getInstalledPackageArgsForCall[i].opts
}

func (fake *MockLifecycleResource) GetInstalledPackageReturns(result1 []byte, result2 error) {
	fake.GetInstalledPackageStub = nil
	fake.getInstalledPackageReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) GetInstalledPackageReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.GetInstalledPackageStub = nil
	if fake.getInstalledPackageReturnsOnCall == nil {
		fake.getInstalledPackageReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getInstalledPackageReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) QueryInstalled(reqCtx reqContext.Context, target fab.ProposalProcessor, opts ...resource.Opt) (*resource.LifecycleQueryInstalledCCResponse, error) {
	fake.queryInstalledMutex.Lock()
	ret, specificReturn := fake.queryInstalledReturnsOnCall[len(fake.queryInstalledArgsForCall)]
	fake.queryInstalledArgsForCall = append(fake.queryInstalledArgsForCall, struct {
		reqCtx reqContext.Context
		target fab.ProposalProcessor
		opts   []resource.Opt
	}{reqCtx, target, opts})
	fake.recordInvocation("QueryInstalled", []interface{}{reqCtx, target, opts})
	fake.queryInstalledMutex.Unlock()
	if fake.QueryInstalledStub != nil {
		return fake.QueryInstalledStub(reqCtx, target, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.queryInstalledReturns.result1, fake.queryInstalledReturns.result2
}

func (fake *MockLifecycleResource) QueryInstalledCallCount() int {
	fake.queryInstalledMutex.RLock()
	defer fake.queryInstalledMutex.RUnlock()
	return len(fake.queryInstalledArgsForCall)
}

func (fake *MockLifecycleResource) QueryInstalledArgsForCall(i int) (reqContext.Context, fab.ProposalProcessor, []resource.Opt) {
	fake.queryInstalledMutex.RLock()
	defer fake.queryInstalledMutex.RUnlock()
	return fake.queryInstalledArgsForCall[i].reqCtx, fake.queryInstalledArgsForCall[i].target, fake.queryInstalledArgsForCall[i].opts
}

func (fake *MockLifecycleResource) QueryInstalledReturns(result1 *resource.LifecycleQueryInstalledCCResponse, result2 error) {
	fake.QueryInstalledStub = nil
	fake.queryInstalledReturns = struct {
		result1 *resource.LifecycleQueryInstalledCCResponse
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) QueryInstalledReturnsOnCall(i int, result1 *resource.LifecycleQueryInstalledCCResponse, result2 error) {
	fake.QueryInstalledStub = nil
	if fake.queryInstalledReturnsOnCall == nil {
		fake.queryInstalledReturnsOnCall = make(map[int]struct {
			result1 *resource.LifecycleQueryInstalledCCResponse
			result2 error
		})
	}
	fake.queryInstalledReturnsOnCall[i] = struct {
		result1 *resource.LifecycleQueryInstalledCCResponse
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) QueryApproved(reqCtx reqContext.Context, channelID string, req *resource.QueryApprovedChaincodeRequest, target fab.ProposalProcessor, opts ...resource.Opt) (*resource.LifecycleQueryApprovedCCResponse, error) {
	fake.queryApprovedMutex.Lock()
	ret, specificReturn := fake.queryApprovedReturnsOnCall[len(fake.queryApprovedArgsForCall)]
	fake.queryApprovedArgsForCall = append(fake.queryApprovedArgsForCall, struct {
		reqCtx    reqContext.Context
		channelID string
		req       *resource.QueryApprovedChaincodeRequest
		target    fab.ProposalProcessor
		opts      []resource.Opt
	}{reqCtx, channelID, req, target, opts})
	fake.recordInvocation("QueryApproved", []interface{}{reqCtx, channelID, req, target, opts})
	fake.queryApprovedMutex.Unlock()
	if fake.QueryApprovedStub != nil {
		return fake.QueryApprovedStub(reqCtx, channelID, req, target, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.queryApprovedReturns.result1, fake.queryApprovedReturns.result2
}

func (fake *MockLifecycleResource) QueryApprovedCallCount() int {
	fake.queryApprovedMutex.RLock()
	defer fake.queryApprovedMutex.RUnlock()
	return len(fake.queryApprovedArgsForCall)
}

func (fake *MockLifecycleResource) QueryApprovedArgsForCall(i int) (reqContext.Context, string, *resource.QueryApprovedChaincodeRequest, fab.ProposalProcessor, []resource.Opt) {
	fake.queryApprovedMutex.RLock()
	defer fake.queryApprovedMutex.RUnlock()
	return fake.queryApprovedArgsForCall[i].reqCtx, fake.queryApprovedArgsForCall[i].channelID, fake.queryApprovedArgsForCall[i].req, fake.queryApprovedArgsForCall[i].target, fake.queryApprovedArgsForCall[i].opts
}

func (fake *MockLifecycleResource) QueryApprovedReturns(result1 *resource.LifecycleQueryApprovedCCResponse, result2 error) {
	fake.QueryApprovedStub = nil
	fake.queryApprovedReturns = struct {
		result1 *resource.LifecycleQueryApprovedCCResponse
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) QueryApprovedReturnsOnCall(i int, result1 *resource.LifecycleQueryApprovedCCResponse, result2 error) {
	fake.QueryApprovedStub = nil
	if fake.queryApprovedReturnsOnCall == nil {
		fake.queryApprovedReturnsOnCall = make(map[int]struct {
			result1 *resource.LifecycleQueryApprovedCCResponse
			result2 error
		})
	}
	fake.queryApprovedReturnsOnCall[i] = struct {
		result1 *resource.LifecycleQueryApprovedCCResponse
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) CreateApproveProposal(txh fab.TransactionHeader, req *resource.ApproveChaincodeRequest) (*fab.TransactionProposal, error) {
	fake.createApproveProposalMutex.Lock()
	ret, specificReturn := fake.createApproveProposalReturnsOnCall[len(fake.createApproveProposalArgsForCall)]
	fake.createApproveProposalArgsForCall = append(fake.createApproveProposalArgsForCall, struct {
		txh fab.TransactionHeader
		req *resource.ApproveChaincodeRequest
	}{txh, req})
	fake.recordInvocation("CreateApproveProposal", []interface{}{txh, req})
	fake.createApproveProposalMutex.Unlock()
	if fake.CreateApproveProposalStub != nil {
		return fake.CreateApproveProposalStub(txh, req)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createApproveProposalReturns.result1, fake.createApproveProposalReturns.result2
}

func (fake *MockLifecycleResource) CreateApproveProposalCallCount() int {
	fake.createApproveProposalMutex.RLock()
	defer fake.createApproveProposalMutex.RUnlock()
	return len(fake.createApproveProposalArgsForCall)
}

func (fake *MockLifecycleResource) CreateApproveProposalArgsForCall(i int) (fab.TransactionHeader, *resource.ApproveChaincodeRequest) {
	fake.createApproveProposalMutex.RLock()
	defer fake.createApproveProposalMutex.RUnlock()
	return fake.createApproveProposalArgsForCall[i].txh, fake.createApproveProposalArgsForCall[i].req
}

func (fake *MockLifecycleResource) CreateApproveProposalReturns(result1 *fab.TransactionProposal, result2 error) {
	fake.CreateApproveProposalStub = nil
	fake.createApproveProposalReturns = struct {
		result1 *fab.TransactionProposal
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) CreateApproveProposalReturnsOnCall(i int, result1 *fab.TransactionProposal, result2 error) {
	fake.CreateApproveProposalStub = nil
	if fake.createApproveProposalReturnsOnCall == nil {
		fake.createApproveProposalReturnsOnCall = make(map[int]struct {
			result1 *fab.TransactionProposal
			result2 error
		})
	}
	fake.createApproveProposalReturnsOnCall[i] = struct {
		result1 *fab.TransactionProposal
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) CreateCheckCommitReadinessProposal(txh fab.TransactionHeader, req *resource.CheckChaincodeCommitReadinessRequest) (*fab.TransactionProposal, error) {
	fake.createCheckCommitReadinessProposalMutex.Lock()
	ret, specificReturn := fake.createCheckCommitReadinessProposalReturnsOnCall[len(fake.createCheckCommitReadinessProposalArgsForCall)]
	fake.createCheckCommitReadinessProposalArgsForCall = append(fake.createCheckCommitReadinessProposalArgsForCall, struct {
		txh fab.TransactionHeader
		req *resource.CheckChaincodeCommitReadinessRequest
	}{txh, req})
	fake.recordInvocation("CreateCheckCommitReadinessProposal", []interface{}{txh, req})
	fake.createCheckCommitReadinessProposalMutex.Unlock()
	if fake.CreateCheckCommitReadinessProposalStub != nil {
		return fake.CreateCheckCommitReadinessProposalStub(txh, req)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createCheckCommitReadinessProposalReturns.result1, fake.createCheckCommitReadinessProposalReturns.result2
}

func (fake *MockLifecycleResource) CreateCheckCommitReadinessProposalCallCount() int {
	fake.createCheckCommitReadinessProposalMutex.RLock()
	defer fake.createCheckCommitReadinessProposalMutex.RUnlock()
	return len(fake.createCheckCommitReadinessProposalArgsForCall)
}

func (fake *MockLifecycleResource) CreateCheckCommitReadinessProposalArgsForCall(i int) (fab.TransactionHeader, *resource.CheckChaincodeCommitReadinessRequest) {
	fake.createCheckCommitReadinessProposalMutex.RLock()
	defer fake.createCheckCommitReadinessProposalMutex.RUnlock()
	return fake.createCheckCommitReadinessProposalArgsForCall[i].txh, fake.createCheckCommitReadinessProposalArgsForCall[i].req
}

func (fake *MockLifecycleResource) CreateCheckCommitReadinessProposalReturns(result1 *fab.TransactionProposal, result2 error) {
	fake.CreateCheckCommitReadinessProposalStub = nil
	fake.createCheckCommitReadinessProposalReturns = struct {
		result1 *fab.TransactionProposal
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) CreateCheckCommitReadinessProposalReturnsOnCall(i int, result1 *fab.TransactionProposal, result2 error) {
	fake.CreateCheckCommitReadinessProposalStub = nil
	if fake.createCheckCommitReadinessProposalReturnsOnCall == nil {
		fake.createCheckCommitReadinessProposalReturnsOnCall = make(map[int]struct {
			result1 *fab.TransactionProposal
			result2 error
		})
	}
	fake.createCheckCommitReadinessProposalReturnsOnCall[i] = struct {
		result1 *fab.TransactionProposal
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) CreateCommitProposal(txh fab.TransactionHeader, req *resource.CommitChaincodeRequest) (*fab.TransactionProposal, error) {
	fake.createCommitProposalMutex.Lock()
	ret, specificReturn := fake.createCommitProposalReturnsOnCall[len(fake.createCommitProposalArgsForCall)]
	fake.createCommitProposalArgsForCall = append(fake.createCommitProposalArgsForCall, struct {
		txh fab.TransactionHeader
		req *resource.CommitChaincodeRequest
	}{txh, req})
	fake.recordInvocation("CreateCommitProposal", []interface{}{txh, req})
	fake.createCommitProposalMutex.Unlock()
	if fake.CreateCommitProposalStub != nil {
		return fake.CreateCommitProposalStub(txh, req)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createCommitProposalReturns.result1, fake.createCommitProposalReturns.result2
}

func (fake *MockLifecycleResource) CreateCommitProposalCallCount() int {
	fake.createCommitProposalMutex.RLock()
	defer fake.createCommitProposalMutex.RUnlock()
	return len(fake.createCommitProposalArgsForCall)
}

func (fake *MockLifecycleResource) CreateCommitProposalArgsForCall(i int) (fab.TransactionHeader, *resource.CommitChaincodeRequest) {
	fake.createCommitProposalMutex.RLock()
	defer fake.createCommitProposalMutex.RUnlock()
	return fake.createCommitProposalArgsForCall[i].txh, fake.createCommitProposalArgsForCall[i].req
}

func (fake *MockLifecycleResource) CreateCommitProposalReturns(result1 *fab.TransactionProposal, result2 error) {
	fake.CreateCommitProposalStub = nil
	fake.createCommitProposalReturns = struct {
		result1 *fab.TransactionProposal
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) CreateCommitProposalReturnsOnCall(i int, result1 *fab.TransactionProposal, result2 error) {
	fake.CreateCommitProposalStub = nil
	if fake.createCommitProposalReturnsOnCall == nil {
		fake.createCommitProposalReturnsOnCall = make(map[int]struct {
			result1 *fab.TransactionProposal
			result2 error
		})
	}
	fake.createCommitProposalReturnsOnCall[i] = struct {
		result1 *fab.TransactionProposal
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) CreateQueryCommittedProposal(txh fab.TransactionHeader, req *resource.QueryCommittedChaincodesRequest) (*fab.TransactionProposal, error) {
	fake.createQueryCommittedProposalMutex.Lock()
	ret, specificReturn := fake.createQueryCommittedProposalReturnsOnCall[len(fake.createQueryCommittedProposalArgsForCall)]
	fake.createQueryCommittedProposalArgsForCall = append(fake.createQueryCommittedProposalArgsForCall, struct {
		txh fab.TransactionHeader
		req *resource.QueryCommittedChaincodesRequest
	}{txh, req})
	fake.recordInvocation("CreateQueryCommittedProposal", []interface{}{txh, req})
	fake.createQueryCommittedProposalMutex.Unlock()
	if fake.CreateQueryCommittedProposalStub != nil {
		return fake.CreateQueryCommittedProposalStub(txh, req)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createQueryCommittedProposalReturns.result1, fake.createQueryCommittedProposalReturns.result2
}

func (fake *MockLifecycleResource) CreateQueryCommittedProposalCallCount() int {
	fake.createQueryCommittedProposalMutex.RLock()
	defer fake.createQueryCommittedProposalMutex.RUnlock()
	return len(fake.createQueryCommittedProposalArgsForCall)
}

func (fake *MockLifecycleResource) CreateQueryCommittedProposalArgsForCall(i int) (fab.TransactionHeader, *resource.QueryCommittedChaincodesRequest) {
	fake.createQueryCommittedProposalMutex.RLock()
	defer fake.createQueryCommittedProposalMutex.RUnlock()
	return fake.createQueryCommittedProposalArgsForCall[i].txh, fake.createQueryCommittedProposalArgsForCall[i].req
}

func (fake *MockLifecycleResource) CreateQueryCommittedProposalReturns(result1 *fab.TransactionProposal, result2 error) {
	fake.CreateQueryCommittedProposalStub = nil
	fake.createQueryCommittedProposalReturns = struct {
		result1 *fab.TransactionProposal
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) CreateQueryCommittedProposalReturnsOnCall(i int, result1 *fab.TransactionProposal, result2 error) {
	fake.CreateQueryCommittedProposalStub = nil
	if fake.createQueryCommittedProposalReturnsOnCall == nil {
		fake.createQueryCommittedProposalReturnsOnCall = make(map[int]struct {
			result1 *fab.TransactionProposal
			result2 error
		})
	}
	fake.createQueryCommittedProposalReturnsOnCall[i] = struct {
		result1 *fab.TransactionProposal
		result2 error
	}{result1, result2}
}

func (fake *MockLifecycleResource) UnmarshalApplicationPolicy(policyBytes []byte) (*common.SignaturePolicyEnvelope, string, error) {
	var policyBytesCopy []byte
	if policyBytes != nil {
		policyBytesCopy = make([]byte, len(policyBytes))
		copy(policyBytesCopy, policyBytes)
	}
	fake.unmarshalApplicationPolicyMutex.Lock()
	ret, specificReturn := fake.unmarshalApplicationPolicyReturnsOnCall[len(fake.unmarshalApplicationPolicyArgsForCall)]
	fake.unmarshalApplicationPolicyArgsForCall = append(fake.unmarshalApplicationPolicyArgsForCall, struct {
		policyBytes []byte
	}{policyBytesCopy})
	fake.recordInvocation("UnmarshalApplicationPolicy", []interface{}{policyBytesCopy})
	fake.unmarshalApplicationPolicyMutex.Unlock()
	if fake.UnmarshalApplicationPolicyStub != nil {
		return fake.UnmarshalApplicationPolicyStub(policyBytes)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.unmarshalApplicationPolicyReturns.result1, fake.unmarshalApplicationPolicyReturns.result2, fake.unmarshalApplicationPolicyReturns.result3
}

func (fake *MockLifecycleResource) UnmarshalApplicationPolicyCallCount() int {
	fake.unmarshalApplicationPolicyMutex.RLock()
	defer fake.unmarshalApplicationPolicyMutex.RUnlock()
	return len(fake.unmarshalApplicationPolicyArgsForCall)
}

func (fake *MockLifecycleResource) UnmarshalApplicationPolicyArgsForCall(i int) []byte {
	fake.unmarshalApplicationPolicyMutex.RLock()
	defer fake.unmarshalApplicationPolicyMutex.RUnlock()
	return fake.unmarshalApplicationPolicyArgsForCall[i].policyBytes
}

func (fake *MockLifecycleResource) UnmarshalApplicationPolicyReturns(result1 *common.SignaturePolicyEnvelope, result2 string, result3 error) {
	fake.UnmarshalApplicationPolicyStub = nil
	fake.unmarshalApplicationPolicyReturns = struct {
		result1 *common.SignaturePolicyEnvelope
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *MockLifecycleResource) UnmarshalApplicationPolicyReturnsOnCall(i int, result1 *common.SignaturePolicyEnvelope, result2 string, result3 error) {
	fake.UnmarshalApplicationPolicyStub = nil
	if fake.unmarshalApplicationPolicyReturnsOnCall == nil {
		fake.unmarshalApplicationPolicyReturnsOnCall = make(map[int]struct {
			result1 *common.SignaturePolicyEnvelope
			result2 string
			result3 error
		})
	}
	fake.unmarshalApplicationPolicyReturnsOnCall[i] = struct {
		result1 *common.SignaturePolicyEnvelope
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *MockLifecycleResource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.installMutex.RLock()
	defer fake.installMutex.RUnlock()
	fake.getInstalledPackageMutex.RLock()
	defer fake.getInstalledPackageMutex.RUnlock()
	fake.queryInstalledMutex.RLock()
	defer fake.queryInstalledMutex.RUnlock()
	fake.queryApprovedMutex.RLock()
	defer fake.queryApprovedMutex.RUnlock()
	fake.createApproveProposalMutex.RLock()
	defer fake.createApproveProposalMutex.RUnlock()
	fake.createCheckCommitReadinessProposalMutex.RLock()
	defer fake.createCheckCommitReadinessProposalMutex.RUnlock()
	fake.createCommitProposalMutex.RLock()
	defer fake.createCommitProposalMutex.RUnlock()
	fake.createQueryCommittedProposalMutex.RLock()
	defer fake.createQueryCommittedProposalMutex.RUnlock()
	fake.unmarshalApplicationPolicyMutex.RLock()
	defer fake.unmarshalApplicationPolicyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MockLifecycleResource) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}