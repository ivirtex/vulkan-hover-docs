# vkDestroyFramebuffer(3) Manual Page

## Name

vkDestroyFramebuffer - Destroy a framebuffer object



## [](#_c_specification)C Specification

To destroy a framebuffer, call:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_0
void vkDestroyFramebuffer(
    VkDevice                                    device,
    VkFramebuffer                               framebuffer,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the framebuffer.
- `framebuffer` is the handle of the framebuffer to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyFramebuffer-framebuffer-00892)VUID-vkDestroyFramebuffer-framebuffer-00892  
  All submitted commands that refer to `framebuffer` **must** have completed execution
- [](#VUID-vkDestroyFramebuffer-framebuffer-00893)VUID-vkDestroyFramebuffer-framebuffer-00893  
  If `VkAllocationCallbacks` were provided when `framebuffer` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyFramebuffer-framebuffer-00894)VUID-vkDestroyFramebuffer-framebuffer-00894  
  If no `VkAllocationCallbacks` were provided when `framebuffer` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyFramebuffer-device-parameter)VUID-vkDestroyFramebuffer-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyFramebuffer-framebuffer-parameter)VUID-vkDestroyFramebuffer-framebuffer-parameter  
  If `framebuffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `framebuffer` **must** be a valid [VkFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebuffer.html) handle
- [](#VUID-vkDestroyFramebuffer-pAllocator-parameter)VUID-vkDestroyFramebuffer-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyFramebuffer-framebuffer-parent)VUID-vkDestroyFramebuffer-framebuffer-parent  
  If `framebuffer` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `framebuffer` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyFramebuffer).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0