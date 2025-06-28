# vkCreateValidationCacheEXT(3) Manual Page

## Name

vkCreateValidationCacheEXT - Creates a new validation cache



## [](#_c_specification)C Specification

To create validation cache objects, call:

```c++
// Provided by VK_EXT_validation_cache
VkResult vkCreateValidationCacheEXT(
    VkDevice                                    device,
    const VkValidationCacheCreateInfoEXT*       pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkValidationCacheEXT*                       pValidationCache);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the validation cache object.
- `pCreateInfo` is a pointer to a [VkValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheCreateInfoEXT.html) structure containing the initial parameters for the validation cache object.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pValidationCache` is a pointer to a [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html) handle in which the resulting validation cache object is returned.

## [](#_description)Description

Note

Applications **can** track and manage the total host memory size of a validation cache object using the `pAllocator`. Applications **can** limit the amount of data retrieved from a validation cache object in `vkGetValidationCacheDataEXT`. Implementations **should** not internally limit the total number of entries added to a validation cache object or the total host memory consumed.

Once created, a validation cache **can** be passed to the `vkCreateShaderModule` command by adding this object to the [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html) structure’s `pNext` chain. If a [VkShaderModuleValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleValidationCacheCreateInfoEXT.html) object is included in the [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html)::`pNext` chain, and its `validationCache` field is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the implementation will query it for possible reuse opportunities and update it with new content. The use of the validation cache object in these commands is internally synchronized, and the same validation cache object **can** be used in multiple threads simultaneously.

Note

Implementations **should** make every effort to limit any critical sections to the actual accesses to the cache, which is expected to be significantly shorter than the duration of the `vkCreateShaderModule` command.

Valid Usage (Implicit)

- [](#VUID-vkCreateValidationCacheEXT-device-parameter)VUID-vkCreateValidationCacheEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateValidationCacheEXT-pCreateInfo-parameter)VUID-vkCreateValidationCacheEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheCreateInfoEXT.html) structure
- [](#VUID-vkCreateValidationCacheEXT-pAllocator-parameter)VUID-vkCreateValidationCacheEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateValidationCacheEXT-pValidationCache-parameter)VUID-vkCreateValidationCacheEXT-pValidationCache-parameter  
  `pValidationCache` **must** be a valid pointer to a [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html) handle
- [](#VUID-vkCreateValidationCacheEXT-device-queuecount)VUID-vkCreateValidationCacheEXT-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_validation\_cache](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_cache.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkValidationCacheCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheCreateInfoEXT.html), [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateValidationCacheEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0