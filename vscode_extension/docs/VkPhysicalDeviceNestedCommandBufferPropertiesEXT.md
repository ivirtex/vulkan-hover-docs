# VkPhysicalDeviceNestedCommandBufferPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceNestedCommandBufferPropertiesEXT - Structure describing the nested command buffer limits of an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceNestedCommandBufferPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_nested_command_buffer
typedef struct VkPhysicalDeviceNestedCommandBufferPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxCommandBufferNestingLevel;
} VkPhysicalDeviceNestedCommandBufferPropertiesEXT;
```

## [](#_members)Members

The members of the `VkPhysicalDeviceNestedCommandBufferPropertiesEXT` structure describe the following features:

## [](#_description)Description

- []()`maxCommandBufferNestingLevel` indicates the maximum nesting level of calls to [vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdExecuteCommands.html) from [Secondary Command Buffers](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary). A `maxCommandBufferNestingLevel` of `UINT32_MAX` means there is no limit to the nesting level.

If the `VkPhysicalDeviceNestedCommandBufferPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceNestedCommandBufferPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceNestedCommandBufferPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_NESTED_COMMAND_BUFFER_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_nested\_command\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_nested_command_buffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceNestedCommandBufferPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0