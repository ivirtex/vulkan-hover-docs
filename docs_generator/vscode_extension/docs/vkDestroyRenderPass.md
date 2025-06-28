# vkDestroyRenderPass(3) Manual Page

## Name

vkDestroyRenderPass - Destroy a render pass object



## [](#_c_specification)C Specification

To destroy a render pass, call:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_0
void vkDestroyRenderPass(
    VkDevice                                    device,
    VkRenderPass                                renderPass,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the render pass.
- `renderPass` is the handle of the render pass to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyRenderPass-renderPass-00873)VUID-vkDestroyRenderPass-renderPass-00873  
  All submitted commands that refer to `renderPass` **must** have completed execution
- [](#VUID-vkDestroyRenderPass-renderPass-00874)VUID-vkDestroyRenderPass-renderPass-00874  
  If `VkAllocationCallbacks` were provided when `renderPass` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyRenderPass-renderPass-00875)VUID-vkDestroyRenderPass-renderPass-00875  
  If no `VkAllocationCallbacks` were provided when `renderPass` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyRenderPass-device-parameter)VUID-vkDestroyRenderPass-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyRenderPass-renderPass-parameter)VUID-vkDestroyRenderPass-renderPass-parameter  
  If `renderPass` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) handle
- [](#VUID-vkDestroyRenderPass-pAllocator-parameter)VUID-vkDestroyRenderPass-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyRenderPass-renderPass-parent)VUID-vkDestroyRenderPass-renderPass-parent  
  If `renderPass` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `renderPass` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyRenderPass)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0