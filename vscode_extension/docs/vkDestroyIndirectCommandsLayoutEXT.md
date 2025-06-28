# vkDestroyIndirectCommandsLayoutEXT(3) Manual Page

## Name

vkDestroyIndirectCommandsLayoutEXT - Destroy an indirect commands layout



## [](#_c_specification)C Specification

Indirect command layouts for `VK_EXT_device_generated_commands` are destroyed by:

```c++
// Provided by VK_EXT_device_generated_commands
void vkDestroyIndirectCommandsLayoutEXT(
    VkDevice                                    device,
    VkIndirectCommandsLayoutEXT                 indirectCommandsLayout,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the layout.
- `indirectCommandsLayout` is the layout to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyIndirectCommandsLayoutEXT-indirectCommandsLayout-11114)VUID-vkDestroyIndirectCommandsLayoutEXT-indirectCommandsLayout-11114  
  All submitted commands that refer to `indirectCommandsLayout` **must** have completed execution
- [](#VUID-vkDestroyIndirectCommandsLayoutEXT-indirectCommandsLayout-11115)VUID-vkDestroyIndirectCommandsLayoutEXT-indirectCommandsLayout-11115  
  If `VkAllocationCallbacks` were provided when `indirectCommandsLayout` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyIndirectCommandsLayoutEXT-indirectCommandsLayout-11116)VUID-vkDestroyIndirectCommandsLayoutEXT-indirectCommandsLayout-11116  
  If no `VkAllocationCallbacks` were provided when `indirectCommandsLayout` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyIndirectCommandsLayoutEXT-device-parameter)VUID-vkDestroyIndirectCommandsLayoutEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyIndirectCommandsLayoutEXT-indirectCommandsLayout-parameter)VUID-vkDestroyIndirectCommandsLayoutEXT-indirectCommandsLayout-parameter  
  If `indirectCommandsLayout` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `indirectCommandsLayout` **must** be a valid [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html) handle
- [](#VUID-vkDestroyIndirectCommandsLayoutEXT-pAllocator-parameter)VUID-vkDestroyIndirectCommandsLayoutEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyIndirectCommandsLayoutEXT-indirectCommandsLayout-parent)VUID-vkDestroyIndirectCommandsLayoutEXT-indirectCommandsLayout-parent  
  If `indirectCommandsLayout` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `indirectCommandsLayout` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyIndirectCommandsLayoutEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0