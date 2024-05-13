# VkValidationCacheHeaderVersionEXT(3) Manual Page

## Name

VkValidationCacheHeaderVersionEXT - Encode validation cache version



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the second group of four bytes in the header returned
by [vkGetValidationCacheDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetValidationCacheDataEXT.html),
encoding the validation cache version, are:

``` c
// Provided by VK_EXT_validation_cache
typedef enum VkValidationCacheHeaderVersionEXT {
    VK_VALIDATION_CACHE_HEADER_VERSION_ONE_EXT = 1,
} VkValidationCacheHeaderVersionEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VALIDATION_CACHE_HEADER_VERSION_ONE_EXT` specifies version one of
  the validation cache.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_validation_cache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_validation_cache.html),
[vkCreateValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateValidationCacheEXT.html),
[vkGetValidationCacheDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetValidationCacheDataEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkValidationCacheHeaderVersionEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
