# vkCreateExecutionGraphPipelinesAMDX(3) Manual Page

## Name

vkCreateExecutionGraphPipelinesAMDX - Creates a new execution graph pipeline object



## [](#_c_specification)C Specification

To create execution graph pipelines, call:

```c++
// Provided by VK_AMDX_shader_enqueue
VkResult vkCreateExecutionGraphPipelinesAMDX(
    VkDevice                                    device,
    VkPipelineCache                             pipelineCache,
    uint32_t                                    createInfoCount,
    const VkExecutionGraphPipelineCreateInfoAMDX* pCreateInfos,
    const VkAllocationCallbacks*                pAllocator,
    VkPipeline*                                 pPipelines);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the execution graph pipelines.
- `pipelineCache` is either [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), indicating that pipeline caching is disabled; or the handle of a valid [pipeline cache](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-cache) object, in which case use of that cache is enabled for the duration of the command. The implementation **must** not access this object outside of the duration of this command.
- `createInfoCount` is the length of the `pCreateInfos` and `pPipelines` arrays.
- `pCreateInfos` is a pointer to an array of [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html) structures.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pPipelines` is a pointer to an array of [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handles in which the resulting execution graph pipeline objects are returned.

## [](#_description)Description

Pipelines are created and returned as described for [Multiple Pipeline Creation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-multiple).

Valid Usage

- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-shaderEnqueue-09124)VUID-vkCreateExecutionGraphPipelinesAMDX-shaderEnqueue-09124  
  The [`shaderEnqueue`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderEnqueue) feature **must** be enabled
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-flags-09125)VUID-vkCreateExecutionGraphPipelinesAMDX-flags-09125  
  If the `flags` member of any element of `pCreateInfos` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and the `basePipelineIndex` member of that same element is not `-1`, `basePipelineIndex` **must** be less than the index into `pCreateInfos` that corresponds to that element
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-flags-09126)VUID-vkCreateExecutionGraphPipelinesAMDX-flags-09126  
  If the `flags` member of any element of `pCreateInfos` contains the `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, the base pipeline **must** have been created with the `VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT` flag set
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-pipelineCache-09127)VUID-vkCreateExecutionGraphPipelinesAMDX-pipelineCache-09127  
  If `pipelineCache` was created with `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`, host access to `pipelineCache` **must** be [externally synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-threadingbehavior)

<!--THE END-->

- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-pNext-09616)VUID-vkCreateExecutionGraphPipelinesAMDX-pNext-09616  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `pipelineCache` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-pNext-09617)VUID-vkCreateExecutionGraphPipelinesAMDX-pNext-09617  
  If a [VkPipelineCreateFlags2CreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlags2CreateInfoKHR.html) structure with the `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR` flag set is included in the `pNext` chain of any element of `pCreateInfos`, `pipelineCache` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-binaryCount-09620)VUID-vkCreateExecutionGraphPipelinesAMDX-binaryCount-09620  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT` **must** not be set in the `flags` of that element
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-binaryCount-09621)VUID-vkCreateExecutionGraphPipelinesAMDX-binaryCount-09621  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT` **must** not be set in the `flags` of that element
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-binaryCount-09622)VUID-vkCreateExecutionGraphPipelinesAMDX-binaryCount-09622  
  If [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)::`binaryCount` is not `0` for any element of `pCreateInfos`, `VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT_EXT` **must** not be set in the `flags` of that element

Valid Usage (Implicit)

- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-device-parameter)VUID-vkCreateExecutionGraphPipelinesAMDX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-pipelineCache-parameter)VUID-vkCreateExecutionGraphPipelinesAMDX-pipelineCache-parameter  
  If `pipelineCache` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `pipelineCache` **must** be a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) handle
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-pCreateInfos-parameter)VUID-vkCreateExecutionGraphPipelinesAMDX-pCreateInfos-parameter  
  `pCreateInfos` **must** be a valid pointer to an array of `createInfoCount` valid [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html) structures
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-pAllocator-parameter)VUID-vkCreateExecutionGraphPipelinesAMDX-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-pPipelines-parameter)VUID-vkCreateExecutionGraphPipelinesAMDX-pPipelines-parameter  
  `pPipelines` **must** be a valid pointer to an array of `createInfoCount` [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handles
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-createInfoCount-arraylength)VUID-vkCreateExecutionGraphPipelinesAMDX-createInfoCount-arraylength  
  `createInfoCount` **must** be greater than `0`
- [](#VUID-vkCreateExecutionGraphPipelinesAMDX-pipelineCache-parent)VUID-vkCreateExecutionGraphPipelinesAMDX-pipelineCache-parent  
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

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateExecutionGraphPipelinesAMDX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0