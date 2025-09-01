# vkDestroySamplerYcbcrConversion(3) Manual Page

## Name

vkDestroySamplerYcbcrConversion - Destroy a created Y′CBCR conversion



## [](#_c_specification)C Specification

To destroy a sampler Y′CBCR conversion, call:

```c++
// Provided by VK_VERSION_1_1
void vkDestroySamplerYcbcrConversion(
    VkDevice                                    device,
    VkSamplerYcbcrConversion                    ycbcrConversion,
    const VkAllocationCallbacks*                pAllocator);
```

or the equivalent command

```c++
// Provided by VK_KHR_sampler_ycbcr_conversion
void vkDestroySamplerYcbcrConversionKHR(
    VkDevice                                    device,
    VkSamplerYcbcrConversion                    ycbcrConversion,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the Y′CBCR conversion.
- `ycbcrConversion` is the conversion to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkDestroySamplerYcbcrConversion-device-parameter)VUID-vkDestroySamplerYcbcrConversion-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroySamplerYcbcrConversion-ycbcrConversion-parameter)VUID-vkDestroySamplerYcbcrConversion-ycbcrConversion-parameter  
  If `ycbcrConversion` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `ycbcrConversion` **must** be a valid [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html) handle
- [](#VUID-vkDestroySamplerYcbcrConversion-pAllocator-parameter)VUID-vkDestroySamplerYcbcrConversion-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroySamplerYcbcrConversion-ycbcrConversion-parent)VUID-vkDestroySamplerYcbcrConversion-ycbcrConversion-parent  
  If `ycbcrConversion` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `ycbcrConversion` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroySamplerYcbcrConversion).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0