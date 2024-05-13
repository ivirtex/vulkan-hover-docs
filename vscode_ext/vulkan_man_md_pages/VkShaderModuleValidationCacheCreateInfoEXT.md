# VkShaderModuleValidationCacheCreateInfoEXT(3) Manual Page

## Name

VkShaderModuleValidationCacheCreateInfoEXT - Specify validation cache to
use during shader module creation



## <a href="#_c_specification" class="anchor"></a>C Specification

To use a [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html) to cache
shader validation results, add a
[VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html)
structure to the `pNext` chain of the
[VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateInfo.html) structure,
specifying the cache object to use.

The `VkShaderModuleValidationCacheCreateInfoEXT` struct is defined as:

``` c
// Provided by VK_EXT_validation_cache
typedef struct VkShaderModuleValidationCacheCreateInfoEXT {
    VkStructureType         sType;
    const void*             pNext;
    VkValidationCacheEXT    validationCache;
} VkShaderModuleValidationCacheCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `validationCache` is the validation cache object from which the
  results of prior validation attempts will be written, and to which new
  validation results for this [VkShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModule.html) will
  be written (if not already present).

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkShaderModuleValidationCacheCreateInfoEXT-sType-sType"
  id="VUID-VkShaderModuleValidationCacheCreateInfoEXT-sType-sType"></a>
  VUID-VkShaderModuleValidationCacheCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT`

- <a
  href="#VUID-VkShaderModuleValidationCacheCreateInfoEXT-validationCache-parameter"
  id="VUID-VkShaderModuleValidationCacheCreateInfoEXT-validationCache-parameter"></a>
  VUID-VkShaderModuleValidationCacheCreateInfoEXT-validationCache-parameter  
  `validationCache` **must** be a valid
  [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_validation_cache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_validation_cache.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShaderModuleValidationCacheCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
