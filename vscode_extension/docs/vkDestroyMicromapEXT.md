# vkDestroyMicromapEXT(3) Manual Page

## Name

vkDestroyMicromapEXT - Destroy a micromap object



## [](#_c_specification)C Specification

To destroy a micromap, call:

```c++
// Provided by VK_EXT_opacity_micromap
void vkDestroyMicromapEXT(
    VkDevice                                    device,
    VkMicromapEXT                               micromap,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the micromap.
- `micromap` is the micromap to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyMicromapEXT-micromap-10382)VUID-vkDestroyMicromapEXT-micromap-10382  
  The [`micromap`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-micromap) feature **must** be enabled
- [](#VUID-vkDestroyMicromapEXT-micromap-07441)VUID-vkDestroyMicromapEXT-micromap-07441  
  All submitted commands that refer to `micromap` **must** have completed execution
- [](#VUID-vkDestroyMicromapEXT-micromap-07442)VUID-vkDestroyMicromapEXT-micromap-07442  
  If `VkAllocationCallbacks` were provided when `micromap` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyMicromapEXT-micromap-07443)VUID-vkDestroyMicromapEXT-micromap-07443  
  If no `VkAllocationCallbacks` were provided when `micromap` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyMicromapEXT-device-parameter)VUID-vkDestroyMicromapEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyMicromapEXT-micromap-parameter)VUID-vkDestroyMicromapEXT-micromap-parameter  
  If `micromap` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `micromap` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handle
- [](#VUID-vkDestroyMicromapEXT-pAllocator-parameter)VUID-vkDestroyMicromapEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyMicromapEXT-micromap-parent)VUID-vkDestroyMicromapEXT-micromap-parent  
  If `micromap` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `micromap` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyMicromapEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0