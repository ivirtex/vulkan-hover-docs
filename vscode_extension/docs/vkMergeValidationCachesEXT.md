# vkMergeValidationCachesEXT(3) Manual Page

## Name

vkMergeValidationCachesEXT - Combine the data stores of validation caches



## [](#_c_specification)C Specification

Validation cache objects **can** be merged using the command:

```c++
// Provided by VK_EXT_validation_cache
VkResult vkMergeValidationCachesEXT(
    VkDevice                                    device,
    VkValidationCacheEXT                        dstCache,
    uint32_t                                    srcCacheCount,
    const VkValidationCacheEXT*                 pSrcCaches);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the validation cache objects.
- `dstCache` is the handle of the validation cache to merge results into.
- `srcCacheCount` is the length of the `pSrcCaches` array.
- `pSrcCaches` is a pointer to an array of validation cache handles, which will be merged into `dstCache`. The previous contents of `dstCache` are included after the merge.

## [](#_description)Description

Note

The details of the merge operation are implementation-dependent, but implementations **should** merge the contents of the specified validation caches and prune duplicate entries.

Valid Usage

- [](#VUID-vkMergeValidationCachesEXT-dstCache-01536)VUID-vkMergeValidationCachesEXT-dstCache-01536  
  `dstCache` **must** not appear in the list of source caches

Valid Usage (Implicit)

- [](#VUID-vkMergeValidationCachesEXT-device-parameter)VUID-vkMergeValidationCachesEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkMergeValidationCachesEXT-dstCache-parameter)VUID-vkMergeValidationCachesEXT-dstCache-parameter  
  `dstCache` **must** be a valid [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html) handle
- [](#VUID-vkMergeValidationCachesEXT-pSrcCaches-parameter)VUID-vkMergeValidationCachesEXT-pSrcCaches-parameter  
  `pSrcCaches` **must** be a valid pointer to an array of `srcCacheCount` valid [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html) handles
- [](#VUID-vkMergeValidationCachesEXT-srcCacheCount-arraylength)VUID-vkMergeValidationCachesEXT-srcCacheCount-arraylength  
  `srcCacheCount` **must** be greater than `0`
- [](#VUID-vkMergeValidationCachesEXT-dstCache-parent)VUID-vkMergeValidationCachesEXT-dstCache-parent  
  `dstCache` **must** have been created, allocated, or retrieved from `device`
- [](#VUID-vkMergeValidationCachesEXT-pSrcCaches-parent)VUID-vkMergeValidationCachesEXT-pSrcCaches-parent  
  Each element of `pSrcCaches` **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `dstCache` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_validation\_cache](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_cache.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkMergeValidationCachesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0