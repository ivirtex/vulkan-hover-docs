# vkDestroyValidationCacheEXT(3) Manual Page

## Name

vkDestroyValidationCacheEXT - Destroy a validation cache object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a validation cache, call:

``` c
// Provided by VK_EXT_validation_cache
void vkDestroyValidationCacheEXT(
    VkDevice                                    device,
    VkValidationCacheEXT                        validationCache,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the validation cache
  object.

- `validationCache` is the handle of the validation cache to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyValidationCacheEXT-validationCache-01537"
  id="VUID-vkDestroyValidationCacheEXT-validationCache-01537"></a>
  VUID-vkDestroyValidationCacheEXT-validationCache-01537  
  If `VkAllocationCallbacks` were provided when `validationCache` was
  created, a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyValidationCacheEXT-validationCache-01538"
  id="VUID-vkDestroyValidationCacheEXT-validationCache-01538"></a>
  VUID-vkDestroyValidationCacheEXT-validationCache-01538  
  If no `VkAllocationCallbacks` were provided when `validationCache` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyValidationCacheEXT-device-parameter"
  id="VUID-vkDestroyValidationCacheEXT-device-parameter"></a>
  VUID-vkDestroyValidationCacheEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyValidationCacheEXT-validationCache-parameter"
  id="VUID-vkDestroyValidationCacheEXT-validationCache-parameter"></a>
  VUID-vkDestroyValidationCacheEXT-validationCache-parameter  
  If `validationCache` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `validationCache` **must** be a valid
  [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html) handle

- <a href="#VUID-vkDestroyValidationCacheEXT-pAllocator-parameter"
  id="VUID-vkDestroyValidationCacheEXT-pAllocator-parameter"></a>
  VUID-vkDestroyValidationCacheEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyValidationCacheEXT-validationCache-parent"
  id="VUID-vkDestroyValidationCacheEXT-validationCache-parent"></a>
  VUID-vkDestroyValidationCacheEXT-validationCache-parent  
  If `validationCache` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `validationCache` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_validation_cache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_validation_cache.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyValidationCacheEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
