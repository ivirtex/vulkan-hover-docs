# vkMergeValidationCachesEXT(3) Manual Page

## Name

vkMergeValidationCachesEXT - Combine the data stores of validation
caches



## <a href="#_c_specification" class="anchor"></a>C Specification

Validation cache objects **can** be merged using the command:

``` c
// Provided by VK_EXT_validation_cache
VkResult vkMergeValidationCachesEXT(
    VkDevice                                    device,
    VkValidationCacheEXT                        dstCache,
    uint32_t                                    srcCacheCount,
    const VkValidationCacheEXT*                 pSrcCaches);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the validation cache objects.

- `dstCache` is the handle of the validation cache to merge results
  into.

- `srcCacheCount` is the length of the `pSrcCaches` array.

- `pSrcCaches` is a pointer to an array of validation cache handles,
  which will be merged into `dstCache`. The previous contents of
  `dstCache` are included after the merge.

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
<p>The details of the merge operation are implementation-dependent, but
implementations <strong>should</strong> merge the contents of the
specified validation caches and prune duplicate entries.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkMergeValidationCachesEXT-dstCache-01536"
  id="VUID-vkMergeValidationCachesEXT-dstCache-01536"></a>
  VUID-vkMergeValidationCachesEXT-dstCache-01536  
  `dstCache` **must** not appear in the list of source caches

Valid Usage (Implicit)

- <a href="#VUID-vkMergeValidationCachesEXT-device-parameter"
  id="VUID-vkMergeValidationCachesEXT-device-parameter"></a>
  VUID-vkMergeValidationCachesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkMergeValidationCachesEXT-dstCache-parameter"
  id="VUID-vkMergeValidationCachesEXT-dstCache-parameter"></a>
  VUID-vkMergeValidationCachesEXT-dstCache-parameter  
  `dstCache` **must** be a valid
  [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html) handle

- <a href="#VUID-vkMergeValidationCachesEXT-pSrcCaches-parameter"
  id="VUID-vkMergeValidationCachesEXT-pSrcCaches-parameter"></a>
  VUID-vkMergeValidationCachesEXT-pSrcCaches-parameter  
  `pSrcCaches` **must** be a valid pointer to an array of
  `srcCacheCount` valid
  [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html) handles

- <a href="#VUID-vkMergeValidationCachesEXT-srcCacheCount-arraylength"
  id="VUID-vkMergeValidationCachesEXT-srcCacheCount-arraylength"></a>
  VUID-vkMergeValidationCachesEXT-srcCacheCount-arraylength  
  `srcCacheCount` **must** be greater than `0`

- <a href="#VUID-vkMergeValidationCachesEXT-dstCache-parent"
  id="VUID-vkMergeValidationCachesEXT-dstCache-parent"></a>
  VUID-vkMergeValidationCachesEXT-dstCache-parent  
  `dstCache` **must** have been created, allocated, or retrieved from
  `device`

- <a href="#VUID-vkMergeValidationCachesEXT-pSrcCaches-parent"
  id="VUID-vkMergeValidationCachesEXT-pSrcCaches-parent"></a>
  VUID-vkMergeValidationCachesEXT-pSrcCaches-parent  
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

[VK_EXT_validation_cache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_validation_cache.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkMergeValidationCachesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
