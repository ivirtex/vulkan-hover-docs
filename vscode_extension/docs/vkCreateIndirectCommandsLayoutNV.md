# vkCreateIndirectCommandsLayoutNV(3) Manual Page

## Name

vkCreateIndirectCommandsLayoutNV - Create an indirect command layout object



## [](#_c_specification)C Specification

Indirect command layouts for `VK_NV_device_generated_commands` are created by:

```c++
// Provided by VK_NV_device_generated_commands
VkResult vkCreateIndirectCommandsLayoutNV(
    VkDevice                                    device,
    const VkIndirectCommandsLayoutCreateInfoNV* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkIndirectCommandsLayoutNV*                 pIndirectCommandsLayout);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the indirect command layout.
- `pCreateInfo` is a pointer to a [VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoNV.html) structure containing parameters affecting creation of the indirect command layout.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pIndirectCommandsLayout` is a pointer to a `VkIndirectCommandsLayoutNV` handle in which the resulting indirect command layout is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateIndirectCommandsLayoutNV-deviceGeneratedCommands-02929)VUID-vkCreateIndirectCommandsLayoutNV-deviceGeneratedCommands-02929  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV`::`deviceGeneratedCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCommandsNV) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCreateIndirectCommandsLayoutNV-device-parameter)VUID-vkCreateIndirectCommandsLayoutNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateIndirectCommandsLayoutNV-pCreateInfo-parameter)VUID-vkCreateIndirectCommandsLayoutNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoNV.html) structure
- [](#VUID-vkCreateIndirectCommandsLayoutNV-pAllocator-parameter)VUID-vkCreateIndirectCommandsLayoutNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateIndirectCommandsLayoutNV-pIndirectCommandsLayout-parameter)VUID-vkCreateIndirectCommandsLayoutNV-pIndirectCommandsLayout-parameter  
  `pIndirectCommandsLayout` **must** be a valid pointer to a [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutNV.html) handle
- [](#VUID-vkCreateIndirectCommandsLayoutNV-device-queuecount)VUID-vkCreateIndirectCommandsLayoutNV-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoNV.html), [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateIndirectCommandsLayoutNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0