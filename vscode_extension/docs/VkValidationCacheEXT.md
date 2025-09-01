# VkValidationCacheEXT(3) Manual Page

## Name

VkValidationCacheEXT - Opaque handle to a validation cache object



## [](#_c_specification)C Specification

Validation cache objects allow the result of internal validation to be reused, both within a single application run and between multiple runs. Reuse within a single run is achieved by passing the same validation cache object when creating supported Vulkan objects. Reuse across runs of an application is achieved by retrieving validation cache contents in one run of an application, saving the contents, and using them to preinitialize a validation cache on a subsequent run. The contents of the validation cache objects are managed by the validation layers. Applications **can** manage the host memory consumed by a validation cache object and control the amount of data retrieved from a validation cache object.

Validation cache objects are represented by `VkValidationCacheEXT` handles:

```c++
// Provided by VK_EXT_validation_cache
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkValidationCacheEXT)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_EXT\_validation\_cache](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_cache.html), [VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html), [vkCreateValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateValidationCacheEXT.html), [vkDestroyValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyValidationCacheEXT.html), [vkGetValidationCacheDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetValidationCacheDataEXT.html), [vkMergeValidationCachesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMergeValidationCachesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkValidationCacheEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0