# vkCreateComputePipelines(3) Manual Page

## Name

vkCreateComputePipelines - Creates a new compute pipeline object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create compute pipelines, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateComputePipelines(
    VkDevice                                    device,
    VkPipelineCache                             pipelineCache,
    uint32_t                                    createInfoCount,
    const VkComputePipelineCreateInfo*          pCreateInfos,
    const VkAllocationCallbacks*                pAllocator,
    VkPipeline*                                 pPipelines);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the compute pipelines.

- `pipelineCache` is either [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  indicating that pipeline caching is disabled; or the handle of a valid
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-cache"
  target="_blank" rel="noopener">pipeline cache</a> object, in which
  case use of that cache is enabled for the duration of the command.

- `createInfoCount` is the length of the `pCreateInfos` and `pPipelines`
  arrays.

- `pCreateInfos` is a pointer to an array of
  [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html)
  structures.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pPipelines` is a pointer to an array of [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)
  handles in which the resulting compute pipeline objects are returned.

## <a href="#_description" class="anchor"></a>Description

Pipelines are created and returned as described for <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-multiple"
target="_blank" rel="noopener">Multiple Pipeline Creation</a>.

Valid Usage

- <a href="#VUID-vkCreateComputePipelines-flags-00695"
  id="VUID-vkCreateComputePipelines-flags-00695"></a>
  VUID-vkCreateComputePipelines-flags-00695  
  If the `flags` member of any element of `pCreateInfos` contains the
  `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and the `basePipelineIndex`
  member of that same element is not `-1`, `basePipelineIndex` **must**
  be less than the index into `pCreateInfos` that corresponds to that
  element

- <a href="#VUID-vkCreateComputePipelines-flags-00696"
  id="VUID-vkCreateComputePipelines-flags-00696"></a>
  VUID-vkCreateComputePipelines-flags-00696  
  If the `flags` member of any element of `pCreateInfos` contains the
  `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, the base pipeline **must**
  have been created with the `VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT`
  flag set

- <a href="#VUID-vkCreateComputePipelines-pipelineCache-02873"
  id="VUID-vkCreateComputePipelines-pipelineCache-02873"></a>
  VUID-vkCreateComputePipelines-pipelineCache-02873  
  If `pipelineCache` was created with
  `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`, host access to
  `pipelineCache` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-threadingbehavior"
  target="_blank" rel="noopener">externally synchronized</a>

Valid Usage (Implicit)

- <a href="#VUID-vkCreateComputePipelines-device-parameter"
  id="VUID-vkCreateComputePipelines-device-parameter"></a>
  VUID-vkCreateComputePipelines-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateComputePipelines-pipelineCache-parameter"
  id="VUID-vkCreateComputePipelines-pipelineCache-parameter"></a>
  VUID-vkCreateComputePipelines-pipelineCache-parameter  
  If `pipelineCache` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pipelineCache` **must** be a valid
  [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html) handle

- <a href="#VUID-vkCreateComputePipelines-pCreateInfos-parameter"
  id="VUID-vkCreateComputePipelines-pCreateInfos-parameter"></a>
  VUID-vkCreateComputePipelines-pCreateInfos-parameter  
  `pCreateInfos` **must** be a valid pointer to an array of
  `createInfoCount` valid
  [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html)
  structures

- <a href="#VUID-vkCreateComputePipelines-pAllocator-parameter"
  id="VUID-vkCreateComputePipelines-pAllocator-parameter"></a>
  VUID-vkCreateComputePipelines-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateComputePipelines-pPipelines-parameter"
  id="VUID-vkCreateComputePipelines-pPipelines-parameter"></a>
  VUID-vkCreateComputePipelines-pPipelines-parameter  
  `pPipelines` **must** be a valid pointer to an array of
  `createInfoCount` [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handles

- <a href="#VUID-vkCreateComputePipelines-createInfoCount-arraylength"
  id="VUID-vkCreateComputePipelines-createInfoCount-arraylength"></a>
  VUID-vkCreateComputePipelines-createInfoCount-arraylength  
  `createInfoCount` **must** be greater than `0`

- <a href="#VUID-vkCreateComputePipelines-pipelineCache-parent"
  id="VUID-vkCreateComputePipelines-pipelineCache-parent"></a>
  VUID-vkCreateComputePipelines-pipelineCache-parent  
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

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateComputePipelines"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
