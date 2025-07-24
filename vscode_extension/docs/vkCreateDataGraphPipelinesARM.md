# vkCreateDataGraphPipelinesARM(3) Manual Page

## Name

vkCreateDataGraphPipelinesARM - Create data graph pipeline objects



## [](#_c_specification)C Specification

To create data graph pipelines, call:

```c++
// Provided by VK_ARM_data_graph
VkResult vkCreateDataGraphPipelinesARM(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    VkPipelineCache                             pipelineCache,
    uint32_t                                    createInfoCount,
    const VkDataGraphPipelineCreateInfoARM*     pCreateInfos,
    const VkAllocationCallbacks*                pAllocator,
    VkPipeline*                                 pPipelines);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the data graph pipelines.
- `deferredOperation` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or the handle of a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) [request deferral](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#deferred-host-operations-requesting) object for this command.
- `pipelineCache` is either [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), indicating that pipeline caching is disabled; or the handle of a valid [pipeline cache](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-cache) object, in which case use of that cache is enabled for the duration of the command.
- `createInfoCount` is the length of the `pCreateInfos` and `pPipelines` arrays.
- `pCreateInfos` is a pointer to an array of [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html) structures.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pPipelines` is a pointer to an array of [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handles in which the resulting data graph pipelines objects are returned.

## [](#_description)Description

The implementation will create a pipeline in each element of `pPipelines` from the corresponding element of `pCreateInfos`. If the creation of any pipeline fails, that pipeline will be set to [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html).

Valid Usage

- [](#VUID-vkCreateDataGraphPipelinesARM-dataGraph-09760)VUID-vkCreateDataGraphPipelinesARM-dataGraph-09760  
  The [`dataGraph`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dataGraph) feature **must** be enabled
- [](#VUID-vkCreateDataGraphPipelinesARM-device-09927)VUID-vkCreateDataGraphPipelinesARM-device-09927  
  `device` **must** support at least one queue family with the `VK_QUEUE_DATA_GRAPH_BIT_ARM` capability
- [](#VUID-vkCreateDataGraphPipelinesARM-deferredOperation-09761)VUID-vkCreateDataGraphPipelinesARM-deferredOperation-09761  
  `deferredOperation` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCreateDataGraphPipelinesARM-deferredOperation-09916)VUID-vkCreateDataGraphPipelinesARM-deferredOperation-09916  
  If `deferredOperation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `flags` member of elements of `pCreateInfos` **must** not include `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`
- [](#VUID-vkCreateDataGraphPipelinesARM-pNext-09928)VUID-vkCreateDataGraphPipelinesARM-pNext-09928  
  If at least one of the [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html) includes a [VkDataGraphPipelineIdentifierCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineIdentifierCreateInfoARM.html) structure in its `pNext` chain then `pipelineCache` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCreateDataGraphPipelinesARM-pipelineCache-09762)VUID-vkCreateDataGraphPipelinesARM-pipelineCache-09762  
  If `pipelineCache` was created with `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`, host access to `pipelineCache` **must** be [externally synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-threadingbehavior)

Valid Usage (Implicit)

- [](#VUID-vkCreateDataGraphPipelinesARM-device-parameter)VUID-vkCreateDataGraphPipelinesARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateDataGraphPipelinesARM-deferredOperation-parameter)VUID-vkCreateDataGraphPipelinesARM-deferredOperation-parameter  
  If `deferredOperation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `deferredOperation` **must** be a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) handle
- [](#VUID-vkCreateDataGraphPipelinesARM-pipelineCache-parameter)VUID-vkCreateDataGraphPipelinesARM-pipelineCache-parameter  
  If `pipelineCache` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipelineCache` **must** be a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) handle
- [](#VUID-vkCreateDataGraphPipelinesARM-pCreateInfos-parameter)VUID-vkCreateDataGraphPipelinesARM-pCreateInfos-parameter  
  `pCreateInfos` **must** be a valid pointer to an array of `createInfoCount` valid [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html) structures
- [](#VUID-vkCreateDataGraphPipelinesARM-pAllocator-parameter)VUID-vkCreateDataGraphPipelinesARM-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateDataGraphPipelinesARM-pPipelines-parameter)VUID-vkCreateDataGraphPipelinesARM-pPipelines-parameter  
  `pPipelines` **must** be a valid pointer to an array of `createInfoCount` [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handles
- [](#VUID-vkCreateDataGraphPipelinesARM-device-queuecount)VUID-vkCreateDataGraphPipelinesARM-device-queuecount  
  The device **must** have been created with at least `1` queue
- [](#VUID-vkCreateDataGraphPipelinesARM-createInfoCount-arraylength)VUID-vkCreateDataGraphPipelinesARM-createInfoCount-arraylength  
  `createInfoCount` **must** be greater than `0`
- [](#VUID-vkCreateDataGraphPipelinesARM-deferredOperation-parent)VUID-vkCreateDataGraphPipelinesARM-deferredOperation-parent  
  If `deferredOperation` is a valid handle, it **must** have been created, allocated, or retrieved from `device`
- [](#VUID-vkCreateDataGraphPipelinesARM-pipelineCache-parent)VUID-vkCreateDataGraphPipelinesARM-pipelineCache-parent  
  If `pipelineCache` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_PIPELINE_COMPILE_REQUIRED_EXT`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateDataGraphPipelinesARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0