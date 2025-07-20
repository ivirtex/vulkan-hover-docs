# vkCreateIndirectExecutionSetEXT(3) Manual Page

## Name

vkCreateIndirectExecutionSetEXT - Create an indirect execution set



## [](#_c_specification)C Specification

Indirect Execution Sets are created by calling:

```c++
// Provided by VK_EXT_device_generated_commands
VkResult vkCreateIndirectExecutionSetEXT(
    VkDevice                                    device,
    const VkIndirectExecutionSetCreateInfoEXT*  pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkIndirectExecutionSetEXT*                  pIndirectExecutionSet);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the indirect execution set.
- `pCreateInfo` is a pointer to a [VkIndirectExecutionSetCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetCreateInfoEXT.html) structure containing parameters affecting creation of the indirect execution set.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pIndirectExecutionSet` is a pointer to a [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html) handle in which the resulting indirect execution set is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateIndirectExecutionSetEXT-deviceGeneratedCommands-11013)VUID-vkCreateIndirectExecutionSetEXT-deviceGeneratedCommands-11013  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesEXT`::`deviceGeneratedCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCommands) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCreateIndirectExecutionSetEXT-device-parameter)VUID-vkCreateIndirectExecutionSetEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateIndirectExecutionSetEXT-pCreateInfo-parameter)VUID-vkCreateIndirectExecutionSetEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkIndirectExecutionSetCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetCreateInfoEXT.html) structure
- [](#VUID-vkCreateIndirectExecutionSetEXT-pAllocator-parameter)VUID-vkCreateIndirectExecutionSetEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateIndirectExecutionSetEXT-pIndirectExecutionSet-parameter)VUID-vkCreateIndirectExecutionSetEXT-pIndirectExecutionSet-parameter  
  `pIndirectExecutionSet` **must** be a valid pointer to a [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html) handle
- [](#VUID-vkCreateIndirectExecutionSetEXT-device-queuecount)VUID-vkCreateIndirectExecutionSetEXT-device-queuecount  
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

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkIndirectExecutionSetCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetCreateInfoEXT.html), [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateIndirectExecutionSetEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0