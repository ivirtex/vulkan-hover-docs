# vkMergePipelineCaches(3) Manual Page

## Name

vkMergePipelineCaches - Combine the data stores of pipeline caches



## <a href="#_c_specification" class="anchor"></a>C Specification

Pipeline cache objects **can** be merged using the command:

``` c
// Provided by VK_VERSION_1_0
VkResult vkMergePipelineCaches(
    VkDevice                                    device,
    VkPipelineCache                             dstCache,
    uint32_t                                    srcCacheCount,
    const VkPipelineCache*                      pSrcCaches);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the pipeline cache objects.

- `dstCache` is the handle of the pipeline cache to merge results into.

- `srcCacheCount` is the length of the `pSrcCaches` array.

- `pSrcCaches` is a pointer to an array of pipeline cache handles, which
  will be merged into `dstCache`. The previous contents of `dstCache`
  are included after the merge.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The details of the merge operation are implementation-dependent, but
implementations <strong>should</strong> merge the contents of the
specified pipelines and prune duplicate entries.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkMergePipelineCaches-dstCache-00770"
  id="VUID-vkMergePipelineCaches-dstCache-00770"></a>
  VUID-vkMergePipelineCaches-dstCache-00770  
  `dstCache` **must** not appear in the list of source caches

Valid Usage (Implicit)

- <a href="#VUID-vkMergePipelineCaches-device-parameter"
  id="VUID-vkMergePipelineCaches-device-parameter"></a>
  VUID-vkMergePipelineCaches-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkMergePipelineCaches-dstCache-parameter"
  id="VUID-vkMergePipelineCaches-dstCache-parameter"></a>
  VUID-vkMergePipelineCaches-dstCache-parameter  
  `dstCache` **must** be a valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html)
  handle

- <a href="#VUID-vkMergePipelineCaches-pSrcCaches-parameter"
  id="VUID-vkMergePipelineCaches-pSrcCaches-parameter"></a>
  VUID-vkMergePipelineCaches-pSrcCaches-parameter  
  `pSrcCaches` **must** be a valid pointer to an array of
  `srcCacheCount` valid [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html) handles

- <a href="#VUID-vkMergePipelineCaches-srcCacheCount-arraylength"
  id="VUID-vkMergePipelineCaches-srcCacheCount-arraylength"></a>
  VUID-vkMergePipelineCaches-srcCacheCount-arraylength  
  `srcCacheCount` **must** be greater than `0`

- <a href="#VUID-vkMergePipelineCaches-dstCache-parent"
  id="VUID-vkMergePipelineCaches-dstCache-parent"></a>
  VUID-vkMergePipelineCaches-dstCache-parent  
  `dstCache` **must** have been created, allocated, or retrieved from
  `device`

- <a href="#VUID-vkMergePipelineCaches-pSrcCaches-parent"
  id="VUID-vkMergePipelineCaches-pSrcCaches-parent"></a>
  VUID-vkMergePipelineCaches-pSrcCaches-parent  
  Each element of `pSrcCaches` **must** have been created, allocated, or
  retrieved from `device`

Host Synchronization

- Host access to `dstCache` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkMergePipelineCaches"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
