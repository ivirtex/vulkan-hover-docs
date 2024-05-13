# vkCreatePipelineCache(3) Manual Page

## Name

vkCreatePipelineCache - Creates a new pipeline cache



## <a href="#_c_specification" class="anchor"></a>C Specification

To create pipeline cache objects, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreatePipelineCache(
    VkDevice                                    device,
    const VkPipelineCacheCreateInfo*            pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkPipelineCache*                            pPipelineCache);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the pipeline cache object.

- `pCreateInfo` is a pointer to a
  [VkPipelineCacheCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateInfo.html) structure
  containing initial parameters for the pipeline cache object.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pPipelineCache` is a pointer to a
  [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html) handle in which the resulting
  pipeline cache object is returned.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications <strong>can</strong> track and manage the total host
memory size of a pipeline cache object using the
<code>pAllocator</code>. Applications <strong>can</strong> limit the
amount of data retrieved from a pipeline cache object in
<code>vkGetPipelineCacheData</code>. Implementations
<strong>should</strong> not internally limit the total number of entries
added to a pipeline cache object or the total host memory
consumed.</p></td>
</tr>
</tbody>
</table>

Once created, a pipeline cache **can** be passed to the
[vkCreateGraphicsPipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateGraphicsPipelines.html)
[vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesKHR.html),
[vkCreateRayTracingPipelinesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesNV.html), and
[vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateComputePipelines.html) commands. If
the pipeline cache passed into these commands is not
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the implementation will query it
for possible reuse opportunities and update it with new content. The use
of the pipeline cache object in these commands is internally
synchronized, and the same pipeline cache object **can** be used in
multiple threads simultaneously.

If `flags` of `pCreateInfo` includes
`VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`, all commands
that modify the returned pipeline cache object **must** be <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-threadingbehavior"
target="_blank" rel="noopener">externally synchronized</a>.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Implementations <strong>should</strong> make every effort to limit
any critical sections to the actual accesses to the cache, which is
expected to be significantly shorter than the duration of the
<code>vkCreate*Pipelines</code> commands.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkCreatePipelineCache-device-parameter"
  id="VUID-vkCreatePipelineCache-device-parameter"></a>
  VUID-vkCreatePipelineCache-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreatePipelineCache-pCreateInfo-parameter"
  id="VUID-vkCreatePipelineCache-pCreateInfo-parameter"></a>
  VUID-vkCreatePipelineCache-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkPipelineCacheCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateInfo.html) structure

- <a href="#VUID-vkCreatePipelineCache-pAllocator-parameter"
  id="VUID-vkCreatePipelineCache-pAllocator-parameter"></a>
  VUID-vkCreatePipelineCache-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreatePipelineCache-pPipelineCache-parameter"
  id="VUID-vkCreatePipelineCache-pPipelineCache-parameter"></a>
  VUID-vkCreatePipelineCache-pPipelineCache-parameter  
  `pPipelineCache` **must** be a valid pointer to a
  [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html),
[VkPipelineCacheCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreatePipelineCache"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
