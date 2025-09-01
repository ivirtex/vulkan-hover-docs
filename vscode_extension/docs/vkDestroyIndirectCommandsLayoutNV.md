# vkDestroyIndirectCommandsLayoutNV(3) Manual Page

## Name

vkDestroyIndirectCommandsLayoutNV - Destroy an indirect commands layout



## [](#_c_specification)C Specification

Indirect command layouts for `VK_NV_device_generated_commands` are destroyed by:

```c++
// Provided by VK_NV_device_generated_commands
void vkDestroyIndirectCommandsLayoutNV(
    VkDevice                                    device,
    VkIndirectCommandsLayoutNV                  indirectCommandsLayout,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the layout.
- `indirectCommandsLayout` is the layout to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02938)VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02938  
  All submitted commands that refer to `indirectCommandsLayout` **must** have completed execution
- [](#VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02939)VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02939  
  If `VkAllocationCallbacks` were provided when `indirectCommandsLayout` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02940)VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-02940  
  If no `VkAllocationCallbacks` were provided when `indirectCommandsLayout` was created, `pAllocator` **must** be `NULL`
- [](#VUID-vkDestroyIndirectCommandsLayoutNV-deviceGeneratedCommands-02941)VUID-vkDestroyIndirectCommandsLayoutNV-deviceGeneratedCommands-02941  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV`::`deviceGeneratedCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCommandsNV) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkDestroyIndirectCommandsLayoutNV-device-parameter)VUID-vkDestroyIndirectCommandsLayoutNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-parameter)VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-parameter  
  If `indirectCommandsLayout` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `indirectCommandsLayout` **must** be a valid [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutNV.html) handle
- [](#VUID-vkDestroyIndirectCommandsLayoutNV-pAllocator-parameter)VUID-vkDestroyIndirectCommandsLayoutNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-parent)VUID-vkDestroyIndirectCommandsLayoutNV-indirectCommandsLayout-parent  
  If `indirectCommandsLayout` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `indirectCommandsLayout` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyIndirectCommandsLayoutNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0