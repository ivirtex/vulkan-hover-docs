# VkIndirectCommandsStreamNV(3) Manual Page

## Name

VkIndirectCommandsStreamNV - Structure specifying input streams for generated command tokens



## [](#_c_specification)C Specification

The `VkIndirectCommandsStreamNV` structure specifies the input data for one or more tokens at processing time.

```c++
// Provided by VK_NV_device_generated_commands
typedef struct VkIndirectCommandsStreamNV {
    VkBuffer        buffer;
    VkDeviceSize    offset;
} VkIndirectCommandsStreamNV;
```

## [](#_members)Members

- `buffer` specifies the [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) storing the functional arguments for each sequence. These arguments **can** be written by the device.
- `offset` specified an offset into `buffer` where the arguments start.

## [](#_description)Description

Valid Usage

- [](#VUID-VkIndirectCommandsStreamNV-buffer-02942)VUID-VkIndirectCommandsStreamNV-buffer-02942  
  The `buffer`’s usage flag **must** have the `VK_BUFFER_USAGE_INDIRECT_BUFFER_BIT` bit set
- [](#VUID-VkIndirectCommandsStreamNV-offset-02943)VUID-VkIndirectCommandsStreamNV-offset-02943  
  The `offset` **must** be aligned to `VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV`::`minIndirectCommandsBufferOffsetAlignment`
- [](#VUID-VkIndirectCommandsStreamNV-buffer-02975)VUID-VkIndirectCommandsStreamNV-buffer-02975  
  If `buffer` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object

Valid Usage (Implicit)

- [](#VUID-VkIndirectCommandsStreamNV-buffer-parameter)VUID-VkIndirectCommandsStreamNV-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkGeneratedCommandsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsStreamNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0