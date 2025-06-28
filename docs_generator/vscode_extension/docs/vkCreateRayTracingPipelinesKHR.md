# vkCreateRayTracingPipelinesKHR(3) Manual Page

## Name

vkCreateRayTracingPipelinesKHR - Creates a new ray tracing pipeline object



## [](#_c_specification)C Specification

To create ray tracing pipelines, call:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
VkResult vkCreateRayTracingPipelinesKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    VkPipelineCache                             pipelineCache,
    uint32_t                                    createInfoCount,
    const VkRayTracingPipelineCreateInfoKHR*    pCreateInfos,
    const VkAllocationCallbacks*                pAllocator,
    VkPipeline*                                 pPipelines);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the ray tracing pipelines.
- `deferredOperation` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or the handle of a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) [request deferral](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#deferred-host-operations-requesting) object for this command.
- `pipelineCache` is either [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), indicating that pipeline caching is disabled, or to enable caching, the handle of a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) object. The implementation **must** not access this object outside of the duration of this command.
- `createInfoCount` is the length of the `pCreateInfos` and `pPipelines` arrays.
- `pCreateInfos` is a pointer to an array of [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html) structures.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pPipelines` is a pointer to an array in which the resulting ray tracing pipeline objects are returned.

## [](#_description)Description

The `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS` error is returned if the implementation is unable to reuse the shader group handles provided in [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)::`pShaderGroupCaptureReplayHandle` when [VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html)::`rayTracingPipelineShaderGroupHandleCaptureReplay` is enabled.

Pipelines are created and returned as described for [Multiple Pipeline Creation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-multiple).

Valid Usage

- [](#VUID-vkCreateRayTracingPipelinesKHR-device-09677)VUID-vkCreateRayTracingPipelinesKHR-device-09677  
  `device` **must** support at least one queue family with the `VK_QUEUE_COMPUTE_BIT` capability
- [](#VUID-vkCreateRayTracingPipelinesKHR-flags-03415)VUID-vkCreateRayTracingPipelinesKHR-flags-03415  
  If the `flags` member of any element of `pCreateInfos` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and the `basePipelineIndex` member of that same element is not `-1`, `basePipelineIndex` **must** be less than the index into `pCreateInfos` that corresponds to that element
- [](#VUID-vkCreateRayTracingPipelinesKHR-flags-03416)VUID-vkCreateRayTracingPipelinesKHR-flags-03416  
  If the `flags` member of any element of `pCreateInfos` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, the base pipeline **must** have been created with the `VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT` flag set
- [](#VUID-vkCreateRayTracingPipelinesKHR-flags-03816)VUID-vkCreateRayTracingPipelinesKHR-flags-03816  
  `flags` **must** not contain the `VK_PIPELINE_CREATE_DISPATCH_BASE_BIT` flag
- [](#VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-02903)VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-02903  
  If `pipelineCache` was created with `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`, host access to `pipelineCache` **must** be [externally synchronized](#fundamentals-threadingbehavior)

<!--THE END-->

- [](#VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-03678)VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-03678  
  Any previous deferred operation that was associated with `deferredOperation` **must** be complete

<!--THE END-->

- [](#VUID-vkCreateRayTracingPipelinesKHR-pNext-09616)VUID-vkCreateRayTracingPipelinesKHR-pNext-09616  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `pipelineCache` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCreateRayTracingPipelinesKHR-pNext-09617)VUID-vkCreateRayTracingPipelinesKHR-pNext-09617  
  If a [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfoKHR.html) structure with the `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR` flag set is included in the `pNext` chain of any element of `pCreateInfos`, `pipelineCache` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCreateRayTracingPipelinesKHR-binaryCount-09620)VUID-vkCreateRayTracingPipelinesKHR-binaryCount-09620  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT` **must** not be set in the `flags` of that element
- [](#VUID-vkCreateRayTracingPipelinesKHR-binaryCount-09621)VUID-vkCreateRayTracingPipelinesKHR-binaryCount-09621  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT` **must** not be set in the `flags` of that element
- [](#VUID-vkCreateRayTracingPipelinesKHR-binaryCount-09622)VUID-vkCreateRayTracingPipelinesKHR-binaryCount-09622  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_EXT` **must** not be set in the `flags` of that element
- [](#VUID-vkCreateRayTracingPipelinesKHR-rayTracingPipeline-03586)VUID-vkCreateRayTracingPipelinesKHR-rayTracingPipeline-03586  
  The [`rayTracingPipeline`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-rayTracingPipeline) feature **must** be enabled
- [](#VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-03587)VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-03587  
  If `deferredOperation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the `flags` member of elements of `pCreateInfos` **must** not include `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`

Valid Usage (Implicit)

- [](#VUID-vkCreateRayTracingPipelinesKHR-device-parameter)VUID-vkCreateRayTracingPipelinesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-parameter)VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-parameter  
  If `deferredOperation` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `deferredOperation` **must** be a valid [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html) handle
- [](#VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-parameter)VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-parameter  
  If `pipelineCache` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipelineCache` **must** be a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) handle
- [](#VUID-vkCreateRayTracingPipelinesKHR-pCreateInfos-parameter)VUID-vkCreateRayTracingPipelinesKHR-pCreateInfos-parameter  
  `pCreateInfos` **must** be a valid pointer to an array of `createInfoCount` valid [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html) structures
- [](#VUID-vkCreateRayTracingPipelinesKHR-pAllocator-parameter)VUID-vkCreateRayTracingPipelinesKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateRayTracingPipelinesKHR-pPipelines-parameter)VUID-vkCreateRayTracingPipelinesKHR-pPipelines-parameter  
  `pPipelines` **must** be a valid pointer to an array of `createInfoCount` [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handles
- [](#VUID-vkCreateRayTracingPipelinesKHR-createInfoCount-arraylength)VUID-vkCreateRayTracingPipelinesKHR-createInfoCount-arraylength  
  `createInfoCount` **must** be greater than `0`
- [](#VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-parent)VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-parent  
  If `deferredOperation` is a valid handle, it **must** have been created, allocated, or retrieved from `device`
- [](#VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-parent)VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-parent  
  If `pipelineCache` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_OPERATION_DEFERRED_KHR`
- `VK_OPERATION_NOT_DEFERRED_KHR`
- `VK_PIPELINE_COMPILE_REQUIRED_EXT`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS`

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateRayTracingPipelinesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0