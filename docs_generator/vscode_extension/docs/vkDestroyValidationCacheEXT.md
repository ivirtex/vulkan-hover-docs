# vkDestroyValidationCacheEXT(3) Manual Page

## Name

vkDestroyValidationCacheEXT - Destroy a validation cache object



## [](#_c_specification)C Specification

To destroy a validation cache, call:

```c++
// Provided by VK_EXT_validation_cache
void vkDestroyValidationCacheEXT(
    VkDevice                                    device,
    VkValidationCacheEXT                        validationCache,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the validation cache object.
- `validationCache` is the handle of the validation cache to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyValidationCacheEXT-validationCache-01537)VUID-vkDestroyValidationCacheEXT-validationCache-01537  
  If `VkAllocationCallbacks` were provided when `validationCache` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyValidationCacheEXT-validationCache-01538)VUID-vkDestroyValidationCacheEXT-validationCache-01538  
  If no `VkAllocationCallbacks` were provided when `validationCache` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyValidationCacheEXT-device-parameter)VUID-vkDestroyValidationCacheEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyValidationCacheEXT-validationCache-parameter)VUID-vkDestroyValidationCacheEXT-validationCache-parameter  
  If `validationCache` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `validationCache` **must** be a valid [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html) handle
- [](#VUID-vkDestroyValidationCacheEXT-pAllocator-parameter)VUID-vkDestroyValidationCacheEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyValidationCacheEXT-validationCache-parent)VUID-vkDestroyValidationCacheEXT-validationCache-parent  
  If `validationCache` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `validationCache` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_EXT\_validation\_cache](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_cache.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyValidationCacheEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0