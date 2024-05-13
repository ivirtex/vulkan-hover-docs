# VkValidationCacheEXT(3) Manual Page

## Name

VkValidationCacheEXT - Opaque handle to a validation cache object



## <a href="#_c_specification" class="anchor"></a>C Specification

Validation cache objects allow the result of internal validation to be
reused, both within a single application run and between multiple runs.
Reuse within a single run is achieved by passing the same validation
cache object when creating supported Vulkan objects. Reuse across runs
of an application is achieved by retrieving validation cache contents in
one run of an application, saving the contents, and using them to
preinitialize a validation cache on a subsequent run. The contents of
the validation cache objects are managed by the validation layers.
Applications **can** manage the host memory consumed by a validation
cache object and control the amount of data retrieved from a validation
cache object.

Validation cache objects are represented by `VkValidationCacheEXT`
handles:

``` c
// Provided by VK_EXT_validation_cache
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkValidationCacheEXT)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_validation_cache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_validation_cache.html),
[VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html),
[vkCreateValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateValidationCacheEXT.html),
[vkDestroyValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyValidationCacheEXT.html),
[vkGetValidationCacheDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetValidationCacheDataEXT.html),
[vkMergeValidationCachesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMergeValidationCachesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkValidationCacheEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
