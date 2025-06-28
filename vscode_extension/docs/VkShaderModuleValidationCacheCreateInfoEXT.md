# VkShaderModuleValidationCacheCreateInfoEXT(3) Manual Page

## Name

VkShaderModuleValidationCacheCreateInfoEXT - Specify validation cache to use during shader module creation



## [](#_c_specification)C Specification

To use a [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html) to cache shader validation results, add a [VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html) structure to the `pNext` chain of the [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structure, specifying the cache object to use.

The `VkShaderModuleValidationCacheCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_validation_cache
typedef struct VkShaderModuleValidationCacheCreateInfoEXT {
    VkStructureType         sType;
    const void*             pNext;
    VkValidationCacheEXT    validationCache;
} VkShaderModuleValidationCacheCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `validationCache` is the validation cache object from which the results of prior validation attempts will be written, and to which new validation results for this [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html) will be written (if not already present). The implementation **must** not access this object outside of the duration of the command this structure is passed to.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkShaderModuleValidationCacheCreateInfoEXT-sType-sType)VUID-VkShaderModuleValidationCacheCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SHADER_MODULE_VALIDATION_CACHE_CREATE_INFO_EXT`
- [](#VUID-VkShaderModuleValidationCacheCreateInfoEXT-validationCache-parameter)VUID-VkShaderModuleValidationCacheCreateInfoEXT-validationCache-parameter  
  `validationCache` **must** be a valid [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html) handle

## [](#_see_also)See Also

[VK\_EXT\_validation\_cache](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_cache.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkShaderModuleValidationCacheCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0