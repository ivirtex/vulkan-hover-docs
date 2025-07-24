# vkCreateIndirectCommandsLayoutEXT(3) Manual Page

## Name

vkCreateIndirectCommandsLayoutEXT - Create an indirect command layout object



## [](#_c_specification)C Specification

Indirect command layouts for `VK_EXT_device_generated_commands` are created by:

```c++
// Provided by VK_EXT_device_generated_commands
VkResult vkCreateIndirectCommandsLayoutEXT(
    VkDevice                                    device,
    const VkIndirectCommandsLayoutCreateInfoEXT* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkIndirectCommandsLayoutEXT*                pIndirectCommandsLayout);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the indirect command layout.
- `pCreateInfo` is a pointer to a [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html) structure containing parameters affecting creation of the indirect command layout.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pIndirectCommandsLayout` is a pointer to a `VkIndirectCommandsLayoutEXT` handle in which the resulting indirect command layout is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateIndirectCommandsLayoutEXT-deviceGeneratedCommands-11089)VUID-vkCreateIndirectCommandsLayoutEXT-deviceGeneratedCommands-11089  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT`::`deviceGeneratedCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCommands) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCreateIndirectCommandsLayoutEXT-device-parameter)VUID-vkCreateIndirectCommandsLayoutEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateIndirectCommandsLayoutEXT-pCreateInfo-parameter)VUID-vkCreateIndirectCommandsLayoutEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html) structure
- [](#VUID-vkCreateIndirectCommandsLayoutEXT-pAllocator-parameter)VUID-vkCreateIndirectCommandsLayoutEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateIndirectCommandsLayoutEXT-pIndirectCommandsLayout-parameter)VUID-vkCreateIndirectCommandsLayoutEXT-pIndirectCommandsLayout-parameter  
  `pIndirectCommandsLayout` **must** be a valid pointer to a [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html) handle
- [](#VUID-vkCreateIndirectCommandsLayoutEXT-device-queuecount)VUID-vkCreateIndirectCommandsLayoutEXT-device-queuecount  
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

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkIndirectCommandsLayoutCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutCreateInfoEXT.html), [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateIndirectCommandsLayoutEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0