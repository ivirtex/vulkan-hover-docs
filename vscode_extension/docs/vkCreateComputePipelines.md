# vkCreateComputePipelines(3) Manual Page

## Name

vkCreateComputePipelines - Creates a new compute pipeline object



## [](#_c_specification)C Specification

To create compute pipelines, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateComputePipelines(
    VkDevice                                    device,
    VkPipelineCache                             pipelineCache,
    uint32_t                                    createInfoCount,
    const VkComputePipelineCreateInfo*          pCreateInfos,
    const VkAllocationCallbacks*                pAllocator,
    VkPipeline*                                 pPipelines);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the compute pipelines.
- `pipelineCache` is either [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), indicating that pipeline caching is disabled, or to enable caching, the handle of a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) object. The implementation **must** not access this object outside of the duration of this command.
- `createInfoCount` is the length of the `pCreateInfos` and `pPipelines` arrays.
- `pCreateInfos` is a pointer to an array of [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html) structures.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pPipelines` is a pointer to an array of [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handles in which the resulting compute pipeline objects are returned.

## [](#_description)Description

Pipelines are created and returned as described for [Multiple Pipeline Creation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-multiple).

Valid Usage

- [](#VUID-vkCreateComputePipelines-device-09661)VUID-vkCreateComputePipelines-device-09661  
  `device` **must** support at least one queue family with the `VK_QUEUE_COMPUTE_BIT` capability
- [](#VUID-vkCreateComputePipelines-flags-00695)VUID-vkCreateComputePipelines-flags-00695  
  If the `flags` member of any element of `pCreateInfos` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and the `basePipelineIndex` member of that same element is not `-1`, `basePipelineIndex` **must** be less than the index into `pCreateInfos` that corresponds to that element
- [](#VUID-vkCreateComputePipelines-flags-00696)VUID-vkCreateComputePipelines-flags-00696  
  If the `flags` member of any element of `pCreateInfos` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, the base pipeline **must** have been created with the `VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT` flag set
- [](#VUID-vkCreateComputePipelines-pipelineCache-02873)VUID-vkCreateComputePipelines-pipelineCache-02873  
  If `pipelineCache` was created with `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`, host access to `pipelineCache` **must** be [externally synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-threadingbehavior)

<!--THE END-->

- [](#VUID-vkCreateComputePipelines-pNext-09616)VUID-vkCreateComputePipelines-pNext-09616  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `pipelineCache` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCreateComputePipelines-pNext-09617)VUID-vkCreateComputePipelines-pNext-09617  
  If a [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfoKHR.html) structure with the `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR` flag set is included in the `pNext` chain of any element of `pCreateInfos`, `pipelineCache` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCreateComputePipelines-binaryCount-09620)VUID-vkCreateComputePipelines-binaryCount-09620  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT` **must** not be set in the `flags` of that element
- [](#VUID-vkCreateComputePipelines-binaryCount-09621)VUID-vkCreateComputePipelines-binaryCount-09621  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT` **must** not be set in the `flags` of that element
- [](#VUID-vkCreateComputePipelines-binaryCount-09622)VUID-vkCreateComputePipelines-binaryCount-09622  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_EXT` **must** not be set in the `flags` of that element

Valid Usage (Implicit)

- [](#VUID-vkCreateComputePipelines-device-parameter)VUID-vkCreateComputePipelines-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateComputePipelines-pipelineCache-parameter)VUID-vkCreateComputePipelines-pipelineCache-parameter  
  If `pipelineCache` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipelineCache` **must** be a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) handle
- [](#VUID-vkCreateComputePipelines-pCreateInfos-parameter)VUID-vkCreateComputePipelines-pCreateInfos-parameter  
  `pCreateInfos` **must** be a valid pointer to an array of `createInfoCount` valid [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html) structures
- [](#VUID-vkCreateComputePipelines-pAllocator-parameter)VUID-vkCreateComputePipelines-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateComputePipelines-pPipelines-parameter)VUID-vkCreateComputePipelines-pPipelines-parameter  
  `pPipelines` **must** be a valid pointer to an array of `createInfoCount` [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handles
- [](#VUID-vkCreateComputePipelines-createInfoCount-arraylength)VUID-vkCreateComputePipelines-createInfoCount-arraylength  
  `createInfoCount` **must** be greater than `0`
- [](#VUID-vkCreateComputePipelines-pipelineCache-parent)VUID-vkCreateComputePipelines-pipelineCache-parent  
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

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateComputePipelines)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0