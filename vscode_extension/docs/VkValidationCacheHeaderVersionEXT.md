# VkValidationCacheHeaderVersionEXT(3) Manual Page

## Name

VkValidationCacheHeaderVersionEXT - Encode validation cache version



## [](#_c_specification)C Specification

Possible values of the second group of four bytes in the header returned by [vkGetValidationCacheDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetValidationCacheDataEXT.html), encoding the validation cache version, are:

```c++
// Provided by VK_EXT_validation_cache
typedef enum VkValidationCacheHeaderVersionEXT {
    VK_VALIDATION_CACHE_HEADER_VERSION_ONE_EXT = 1,
} VkValidationCacheHeaderVersionEXT;
```

## [](#_description)Description

- `VK_VALIDATION_CACHE_HEADER_VERSION_ONE_EXT` specifies version one of the validation cache.

## [](#_see_also)See Also

[VK\_EXT\_validation\_cache](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_cache.html), [vkCreateValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateValidationCacheEXT.html), [vkGetValidationCacheDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetValidationCacheDataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkValidationCacheHeaderVersionEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0