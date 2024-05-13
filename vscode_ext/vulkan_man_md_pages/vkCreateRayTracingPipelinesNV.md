# vkCreateRayTracingPipelinesNV(3) Manual Page

## Name

vkCreateRayTracingPipelinesNV - Creates a new ray tracing pipeline
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create ray tracing pipelines, call:

``` c
// Provided by VK_NV_ray_tracing
VkResult vkCreateRayTracingPipelinesNV(
    VkDevice                                    device,
    VkPipelineCache                             pipelineCache,
    uint32_t                                    createInfoCount,
    const VkRayTracingPipelineCreateInfoNV*     pCreateInfos,
    const VkAllocationCallbacks*                pAllocator,
    VkPipeline*                                 pPipelines);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the ray tracing pipelines.

- `pipelineCache` is either [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  indicating that pipeline caching is disabled, or the handle of a valid
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-cache"
  target="_blank" rel="noopener">pipeline cache</a> object, in which
  case use of that cache is enabled for the duration of the command.

- `createInfoCount` is the length of the `pCreateInfos` and `pPipelines`
  arrays.

- `pCreateInfos` is a pointer to an array of
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)
  structures.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pPipelines` is a pointer to an array in which the resulting ray
  tracing pipeline objects are returned.

## <a href="#_description" class="anchor"></a>Description

Pipelines are created and returned as described for <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-multiple"
target="_blank" rel="noopener">Multiple Pipeline Creation</a>.

Valid Usage

- <a href="#VUID-vkCreateRayTracingPipelinesNV-flags-03415"
  id="VUID-vkCreateRayTracingPipelinesNV-flags-03415"></a>
  VUID-vkCreateRayTracingPipelinesNV-flags-03415  
  If the `flags` member of any element of `pCreateInfos` contains the
  `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and the `basePipelineIndex`
  member of that same element is not `-1`, `basePipelineIndex` **must**
  be less than the index into `pCreateInfos` that corresponds to that
  element

- <a href="#VUID-vkCreateRayTracingPipelinesNV-flags-03416"
  id="VUID-vkCreateRayTracingPipelinesNV-flags-03416"></a>
  VUID-vkCreateRayTracingPipelinesNV-flags-03416  
  If the `flags` member of any element of `pCreateInfos` contains the
  `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, the base pipeline **must**
  have been created with the `VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT`
  flag set

- <a href="#VUID-vkCreateRayTracingPipelinesNV-flags-03816"
  id="VUID-vkCreateRayTracingPipelinesNV-flags-03816"></a>
  VUID-vkCreateRayTracingPipelinesNV-flags-03816  
  `flags` **must** not contain the `VK_PIPELINE_CREATE_DISPATCH_BASE`
  flag

- <a href="#VUID-vkCreateRayTracingPipelinesNV-pipelineCache-02903"
  id="VUID-vkCreateRayTracingPipelinesNV-pipelineCache-02903"></a>
  VUID-vkCreateRayTracingPipelinesNV-pipelineCache-02903  
  If `pipelineCache` was created with
  `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`, host access to
  `pipelineCache` **must** be [externally
  synchronized](#fundamentals-threadingbehavior)

Valid Usage (Implicit)

- <a href="#VUID-vkCreateRayTracingPipelinesNV-device-parameter"
  id="VUID-vkCreateRayTracingPipelinesNV-device-parameter"></a>
  VUID-vkCreateRayTracingPipelinesNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateRayTracingPipelinesNV-pipelineCache-parameter"
  id="VUID-vkCreateRayTracingPipelinesNV-pipelineCache-parameter"></a>
  VUID-vkCreateRayTracingPipelinesNV-pipelineCache-parameter  
  If `pipelineCache` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pipelineCache` **must** be a valid
  [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html) handle

- <a href="#VUID-vkCreateRayTracingPipelinesNV-pCreateInfos-parameter"
  id="VUID-vkCreateRayTracingPipelinesNV-pCreateInfos-parameter"></a>
  VUID-vkCreateRayTracingPipelinesNV-pCreateInfos-parameter  
  `pCreateInfos` **must** be a valid pointer to an array of
  `createInfoCount` valid
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)
  structures

- <a href="#VUID-vkCreateRayTracingPipelinesNV-pAllocator-parameter"
  id="VUID-vkCreateRayTracingPipelinesNV-pAllocator-parameter"></a>
  VUID-vkCreateRayTracingPipelinesNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateRayTracingPipelinesNV-pPipelines-parameter"
  id="VUID-vkCreateRayTracingPipelinesNV-pPipelines-parameter"></a>
  VUID-vkCreateRayTracingPipelinesNV-pPipelines-parameter  
  `pPipelines` **must** be a valid pointer to an array of
  `createInfoCount` [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handles

- <a
  href="#VUID-vkCreateRayTracingPipelinesNV-createInfoCount-arraylength"
  id="VUID-vkCreateRayTracingPipelinesNV-createInfoCount-arraylength"></a>
  VUID-vkCreateRayTracingPipelinesNV-createInfoCount-arraylength  
  `createInfoCount` **must** be greater than `0`

- <a href="#VUID-vkCreateRayTracingPipelinesNV-pipelineCache-parent"
  id="VUID-vkCreateRayTracingPipelinesNV-pipelineCache-parent"></a>
  VUID-vkCreateRayTracingPipelinesNV-pipelineCache-parent  
  If `pipelineCache` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_PIPELINE_COMPILE_REQUIRED_EXT`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INVALID_SHADER_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html),
[VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateRayTracingPipelinesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
