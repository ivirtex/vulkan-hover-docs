# VkValidationCacheCreateInfoEXT(3) Manual Page

## Name

VkValidationCacheCreateInfoEXT - Structure specifying parameters of a
newly created validation cache



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkValidationCacheCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_validation_cache
typedef struct VkValidationCacheCreateInfoEXT {
    VkStructureType                    sType;
    const void*                        pNext;
    VkValidationCacheCreateFlagsEXT    flags;
    size_t                             initialDataSize;
    const void*                        pInitialData;
} VkValidationCacheCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `initialDataSize` is the number of bytes in `pInitialData`. If
  `initialDataSize` is zero, the validation cache will initially be
  empty.

- `pInitialData` is a pointer to previously retrieved validation cache
  data. If the validation cache data is incompatible (as defined below)
  with the device, the validation cache will be initially empty. If
  `initialDataSize` is zero, `pInitialData` is ignored.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkValidationCacheCreateInfoEXT-initialDataSize-01534"
  id="VUID-VkValidationCacheCreateInfoEXT-initialDataSize-01534"></a>
  VUID-VkValidationCacheCreateInfoEXT-initialDataSize-01534  
  If `initialDataSize` is not `0`, it **must** be equal to the size of
  `pInitialData`, as returned by `vkGetValidationCacheDataEXT` when
  `pInitialData` was originally retrieved

- <a href="#VUID-VkValidationCacheCreateInfoEXT-initialDataSize-01535"
  id="VUID-VkValidationCacheCreateInfoEXT-initialDataSize-01535"></a>
  VUID-VkValidationCacheCreateInfoEXT-initialDataSize-01535  
  If `initialDataSize` is not `0`, `pInitialData` **must** have been
  retrieved from a previous call to `vkGetValidationCacheDataEXT`

Valid Usage (Implicit)

- <a href="#VUID-VkValidationCacheCreateInfoEXT-sType-sType"
  id="VUID-VkValidationCacheCreateInfoEXT-sType-sType"></a>
  VUID-VkValidationCacheCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT`

- <a href="#VUID-VkValidationCacheCreateInfoEXT-pNext-pNext"
  id="VUID-VkValidationCacheCreateInfoEXT-pNext-pNext"></a>
  VUID-VkValidationCacheCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkValidationCacheCreateInfoEXT-flags-zerobitmask"
  id="VUID-VkValidationCacheCreateInfoEXT-flags-zerobitmask"></a>
  VUID-VkValidationCacheCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-VkValidationCacheCreateInfoEXT-pInitialData-parameter"
  id="VUID-VkValidationCacheCreateInfoEXT-pInitialData-parameter"></a>
  VUID-VkValidationCacheCreateInfoEXT-pInitialData-parameter  
  If `initialDataSize` is not `0`, `pInitialData` **must** be a valid
  pointer to an array of `initialDataSize` bytes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_validation_cache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_validation_cache.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkValidationCacheCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheCreateFlagsEXT.html),
[vkCreateValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateValidationCacheEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkValidationCacheCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
