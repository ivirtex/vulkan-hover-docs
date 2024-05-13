# vkCreateGraphicsPipelines(3) Manual Page

## Name

vkCreateGraphicsPipelines - Create graphics pipelines



## <a href="#_c_specification" class="anchor"></a>C Specification

To create graphics pipelines, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateGraphicsPipelines(
    VkDevice                                    device,
    VkPipelineCache                             pipelineCache,
    uint32_t                                    createInfoCount,
    const VkGraphicsPipelineCreateInfo*         pCreateInfos,
    const VkAllocationCallbacks*                pAllocator,
    VkPipeline*                                 pPipelines);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the graphics pipelines.

- `pipelineCache` is either [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  indicating that pipeline caching is disabled; or the handle of a valid
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-cache"
  target="_blank" rel="noopener">pipeline cache</a> object, in which
  case use of that cache is enabled for the duration of the command.

- `createInfoCount` is the length of the `pCreateInfos` and `pPipelines`
  arrays.

- `pCreateInfos` is a pointer to an array of
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)
  structures.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pPipelines` is a pointer to an array of [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)
  handles in which the resulting graphics pipeline objects are returned.

## <a href="#_description" class="anchor"></a>Description

The [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)
structure includes an array of
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html)
structures for each of the desired active shader stages, as well as
creation information for all relevant fixed-function stages, and a
pipeline layout.

Pipelines are created and returned as described for <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-multiple"
target="_blank" rel="noopener">Multiple Pipeline Creation</a>.

Valid Usage

- <a href="#VUID-vkCreateGraphicsPipelines-flags-00720"
  id="VUID-vkCreateGraphicsPipelines-flags-00720"></a>
  VUID-vkCreateGraphicsPipelines-flags-00720  
  If the `flags` member of any element of `pCreateInfos` contains the
  `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and the `basePipelineIndex`
  member of that same element is not `-1`, `basePipelineIndex` **must**
  be less than the index into `pCreateInfos` that corresponds to that
  element

- <a href="#VUID-vkCreateGraphicsPipelines-flags-00721"
  id="VUID-vkCreateGraphicsPipelines-flags-00721"></a>
  VUID-vkCreateGraphicsPipelines-flags-00721  
  If the `flags` member of any element of `pCreateInfos` contains the
  `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, the base pipeline **must**
  have been created with the `VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT`
  flag set

- <a href="#VUID-vkCreateGraphicsPipelines-pipelineCache-02876"
  id="VUID-vkCreateGraphicsPipelines-pipelineCache-02876"></a>
  VUID-vkCreateGraphicsPipelines-pipelineCache-02876  
  If `pipelineCache` was created with
  `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`, host access to
  `pipelineCache` **must** be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-threadingbehavior"
  target="_blank" rel="noopener">externally synchronized</a>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>An implicit cache may be provided by the implementation or a layer.
For this reason, it is still valid to set
<code>VK_PIPELINE_CREATE_FAIL_ON_PIPELINE_COMPILE_REQUIRED_BIT</code> on
<code>flags</code> for any element of <code>pCreateInfos</code> while
passing <a href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html">VK_NULL_HANDLE</a> for
<code>pipelineCache</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkCreateGraphicsPipelines-device-parameter"
  id="VUID-vkCreateGraphicsPipelines-device-parameter"></a>
  VUID-vkCreateGraphicsPipelines-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateGraphicsPipelines-pipelineCache-parameter"
  id="VUID-vkCreateGraphicsPipelines-pipelineCache-parameter"></a>
  VUID-vkCreateGraphicsPipelines-pipelineCache-parameter  
  If `pipelineCache` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pipelineCache` **must** be a valid
  [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html) handle

- <a href="#VUID-vkCreateGraphicsPipelines-pCreateInfos-parameter"
  id="VUID-vkCreateGraphicsPipelines-pCreateInfos-parameter"></a>
  VUID-vkCreateGraphicsPipelines-pCreateInfos-parameter  
  `pCreateInfos` **must** be a valid pointer to an array of
  `createInfoCount` valid
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)
  structures

- <a href="#VUID-vkCreateGraphicsPipelines-pAllocator-parameter"
  id="VUID-vkCreateGraphicsPipelines-pAllocator-parameter"></a>
  VUID-vkCreateGraphicsPipelines-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateGraphicsPipelines-pPipelines-parameter"
  id="VUID-vkCreateGraphicsPipelines-pPipelines-parameter"></a>
  VUID-vkCreateGraphicsPipelines-pPipelines-parameter  
  `pPipelines` **must** be a valid pointer to an array of
  `createInfoCount` [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handles

- <a href="#VUID-vkCreateGraphicsPipelines-createInfoCount-arraylength"
  id="VUID-vkCreateGraphicsPipelines-createInfoCount-arraylength"></a>
  VUID-vkCreateGraphicsPipelines-createInfoCount-arraylength  
  `createInfoCount` **must** be greater than `0`

- <a href="#VUID-vkCreateGraphicsPipelines-pipelineCache-parent"
  id="VUID-vkCreateGraphicsPipelines-pipelineCache-parent"></a>
  VUID-vkCreateGraphicsPipelines-pipelineCache-parent  
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
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html), [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateGraphicsPipelines"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
