# vkCreateRayTracingPipelinesNV(3) Manual Page

## Name

vkCreateRayTracingPipelinesNV - Creates a new ray tracing pipeline object



## [](#_c_specification)C Specification

To create ray tracing pipelines, call:

```c++
// Provided by VK_NV_ray_tracing
VkResult vkCreateRayTracingPipelinesNV(
    VkDevice                                    device,
    VkPipelineCache                             pipelineCache,
    uint32_t                                    createInfoCount,
    const VkRayTracingPipelineCreateInfoNV*     pCreateInfos,
    const VkAllocationCallbacks*                pAllocator,
    VkPipeline*                                 pPipelines);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the ray tracing pipelines.
- `pipelineCache` is either [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), indicating that pipeline caching is disabled, or to enable caching, the handle of a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) object. The implementation **must** not access this object outside of the duration of this command.
- `createInfoCount` is the length of the `pCreateInfos` and `pPipelines` arrays.
- `pCreateInfos` is a pointer to an array of [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html) structures.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pPipelines` is a pointer to an array in which the resulting ray tracing pipeline objects are returned.

## [](#_description)Description

Pipelines are created and returned as described for [Multiple Pipeline Creation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-multiple).

Valid Usage

- [](#VUID-vkCreateRayTracingPipelinesNV-device-09677)VUID-vkCreateRayTracingPipelinesNV-device-09677  
  `device` **must** support at least one queue family with the `VK_QUEUE_COMPUTE_BIT` capability
- [](#VUID-vkCreateRayTracingPipelinesNV-flags-03415)VUID-vkCreateRayTracingPipelinesNV-flags-03415  
  If the `flags` member of any element of `pCreateInfos` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and the `basePipelineIndex` member of that same element is not `-1`, `basePipelineIndex` **must** be less than the index into `pCreateInfos` that corresponds to that element
- [](#VUID-vkCreateRayTracingPipelinesNV-flags-03416)VUID-vkCreateRayTracingPipelinesNV-flags-03416  
  If the `flags` member of any element of `pCreateInfos` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, the base pipeline **must** have been created with the `VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT` flag set
- [](#VUID-vkCreateRayTracingPipelinesNV-flags-03816)VUID-vkCreateRayTracingPipelinesNV-flags-03816  
  `flags` **must** not contain the `VK_PIPELINE_CREATE_DISPATCH_BASE_BIT` flag
- [](#VUID-vkCreateRayTracingPipelinesNV-pipelineCache-02903)VUID-vkCreateRayTracingPipelinesNV-pipelineCache-02903  
  If `pipelineCache` was created with `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`, host access to `pipelineCache` **must** be [externally synchronized](#fundamentals-threadingbehavior)

<!--THE END-->

- [](#VUID-vkCreateRayTracingPipelinesNV-pNext-09616)VUID-vkCreateRayTracingPipelinesNV-pNext-09616  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `pipelineCache` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCreateRayTracingPipelinesNV-pNext-09617)VUID-vkCreateRayTracingPipelinesNV-pNext-09617  
  If a [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfoKHR.html) structure with the `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR` flag set is included in the `pNext` chain of any element of `pCreateInfos`, `pipelineCache` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCreateRayTracingPipelinesNV-binaryCount-09620)VUID-vkCreateRayTracingPipelinesNV-binaryCount-09620  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT` **must** not be set in the `flags` of that element
- [](#VUID-vkCreateRayTracingPipelinesNV-binaryCount-09621)VUID-vkCreateRayTracingPipelinesNV-binaryCount-09621  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT` **must** not be set in the `flags` of that element
- [](#VUID-vkCreateRayTracingPipelinesNV-binaryCount-09622)VUID-vkCreateRayTracingPipelinesNV-binaryCount-09622  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_EXT` **must** not be set in the `flags` of that element
- [](#VUID-vkCreateRayTracingPipelinesNV-pNext-10150)VUID-vkCreateRayTracingPipelinesNV-pNext-10150  
  If a [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfoKHR.html) structure is included in the `pNext` chain of any element of `pCreateInfos`, `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR` flag **must** not be set

Valid Usage (Implicit)

- [](#VUID-vkCreateRayTracingPipelinesNV-device-parameter)VUID-vkCreateRayTracingPipelinesNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateRayTracingPipelinesNV-pipelineCache-parameter)VUID-vkCreateRayTracingPipelinesNV-pipelineCache-parameter  
  If `pipelineCache` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipelineCache` **must** be a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) handle
- [](#VUID-vkCreateRayTracingPipelinesNV-pCreateInfos-parameter)VUID-vkCreateRayTracingPipelinesNV-pCreateInfos-parameter  
  `pCreateInfos` **must** be a valid pointer to an array of `createInfoCount` valid [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html) structures
- [](#VUID-vkCreateRayTracingPipelinesNV-pAllocator-parameter)VUID-vkCreateRayTracingPipelinesNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateRayTracingPipelinesNV-pPipelines-parameter)VUID-vkCreateRayTracingPipelinesNV-pPipelines-parameter  
  `pPipelines` **must** be a valid pointer to an array of `createInfoCount` [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handles
- [](#VUID-vkCreateRayTracingPipelinesNV-createInfoCount-arraylength)VUID-vkCreateRayTracingPipelinesNV-createInfoCount-arraylength  
  `createInfoCount` **must** be greater than `0`
- [](#VUID-vkCreateRayTracingPipelinesNV-pipelineCache-parent)VUID-vkCreateRayTracingPipelinesNV-pipelineCache-parent  
  If `pipelineCache` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`
- `VK_PIPELINE_COMPILE_REQUIRED_EXT`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_INVALID_SHADER_NV`

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html), [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateRayTracingPipelinesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0