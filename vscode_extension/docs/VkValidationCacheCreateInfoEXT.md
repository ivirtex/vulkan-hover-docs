# VkValidationCacheCreateInfoEXT(3) Manual Page

## Name

VkValidationCacheCreateInfoEXT - Structure specifying parameters of a newly created validation cache



## [](#_c_specification)C Specification

The `VkValidationCacheCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_validation_cache
typedef struct VkValidationCacheCreateInfoEXT {
    VkStructureType                    sType;
    const void*                        pNext;
    VkValidationCacheCreateFlagsEXT    flags;
    size_t                             initialDataSize;
    const void*                        pInitialData;
} VkValidationCacheCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `initialDataSize` is the number of bytes in `pInitialData`. If `initialDataSize` is zero, the validation cache will initially be empty.
- `pInitialData` is a pointer to previously retrieved validation cache data. If the validation cache data is incompatible (as defined below) with the device, the validation cache will be initially empty. If `initialDataSize` is zero, `pInitialData` is ignored.

## [](#_description)Description

Valid Usage

- [](#VUID-VkValidationCacheCreateInfoEXT-initialDataSize-01534)VUID-VkValidationCacheCreateInfoEXT-initialDataSize-01534  
  If `initialDataSize` is not `0`, it **must** be equal to the size of `pInitialData`, as returned by `vkGetValidationCacheDataEXT` when `pInitialData` was originally retrieved
- [](#VUID-VkValidationCacheCreateInfoEXT-initialDataSize-01535)VUID-VkValidationCacheCreateInfoEXT-initialDataSize-01535  
  If `initialDataSize` is not `0`, `pInitialData` **must** have been retrieved from a previous call to `vkGetValidationCacheDataEXT`

Valid Usage (Implicit)

- [](#VUID-VkValidationCacheCreateInfoEXT-sType-sType)VUID-VkValidationCacheCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VALIDATION_CACHE_CREATE_INFO_EXT`
- [](#VUID-VkValidationCacheCreateInfoEXT-pNext-pNext)VUID-VkValidationCacheCreateInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkValidationCacheCreateInfoEXT-flags-zerobitmask)VUID-VkValidationCacheCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkValidationCacheCreateInfoEXT-pInitialData-parameter)VUID-VkValidationCacheCreateInfoEXT-pInitialData-parameter  
  If `initialDataSize` is not `0`, `pInitialData` **must** be a valid pointer to an array of `initialDataSize` bytes

## [](#_see_also)See Also

[VK\_EXT\_validation\_cache](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_cache.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkValidationCacheCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheCreateFlagsEXT.html), [vkCreateValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateValidationCacheEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkValidationCacheCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0