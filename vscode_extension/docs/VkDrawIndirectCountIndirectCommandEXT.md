# VkDrawIndirectCountIndirectCommandEXT(3) Manual Page

## Name

VkDrawIndirectCountIndirectCommandEXT - Structure specifying input data for a single draw-type command token



## [](#_c_specification)C Specification

The `VkDrawIndirectCountIndirectCommandEXT` structure specifies the input data for all draw-type tokens.

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkDrawIndirectCountIndirectCommandEXT {
    VkDeviceAddress    bufferAddress;
    uint32_t           stride;
    uint32_t           commandCount;
} VkDrawIndirectCountIndirectCommandEXT;
```

## [](#_members)Members

- `bufferAddress` specifies a physical address of the [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) used for draw commands.
- `stride` is the byte size stride for the command arguments
- `commandCount` is the number of commands to execute

## [](#_description)Description

The corresponding indirect draw structure data will be read from the buffer address.

Valid Usage

- [](#VUID-VkDrawIndirectCountIndirectCommandEXT-None-11122)VUID-VkDrawIndirectCountIndirectCommandEXT-None-11122  
  The buffer’s usage flag from which the address was acquired **must** have the `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set

Valid Usage (Implicit)

- [](#VUID-VkDrawIndirectCountIndirectCommandEXT-bufferAddress-parameter)VUID-VkDrawIndirectCountIndirectCommandEXT-bufferAddress-parameter  
  `bufferAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDrawIndirectCountIndirectCommandEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0