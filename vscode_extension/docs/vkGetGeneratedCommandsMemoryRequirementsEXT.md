# vkGetGeneratedCommandsMemoryRequirementsEXT(3) Manual Page

## Name

vkGetGeneratedCommandsMemoryRequirementsEXT - Retrieve the buffer allocation requirements for generated commands



## [](#_c_specification)C Specification

With `VK_EXT_device_generated_commands`, to retrieve the memory size and alignment requirements of a particular execution state call:

```c++
// Provided by VK_EXT_device_generated_commands
void vkGetGeneratedCommandsMemoryRequirementsEXT(
    VkDevice                                    device,
    const VkGeneratedCommandsMemoryRequirementsInfoEXT* pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the buffer.
- `pInfo` is a pointer to a [VkGeneratedCommandsMemoryRequirementsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoEXT.html) structure containing parameters required for the memory requirements query.
- `pMemoryRequirements` is a pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure in which the memory requirements of the buffer object are returned.

## [](#_description)Description

If the size returned is zero, the preprocessing step can be skipped for this layout.

Valid Usage (Implicit)

- [](#VUID-vkGetGeneratedCommandsMemoryRequirementsEXT-device-parameter)VUID-vkGetGeneratedCommandsMemoryRequirementsEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetGeneratedCommandsMemoryRequirementsEXT-pInfo-parameter)VUID-vkGetGeneratedCommandsMemoryRequirementsEXT-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkGeneratedCommandsMemoryRequirementsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoEXT.html) structure
- [](#VUID-vkGetGeneratedCommandsMemoryRequirementsEXT-pMemoryRequirements-parameter)VUID-vkGetGeneratedCommandsMemoryRequirementsEXT-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkGeneratedCommandsMemoryRequirementsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoEXT.html), [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetGeneratedCommandsMemoryRequirementsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0